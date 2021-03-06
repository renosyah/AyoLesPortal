package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/renosyah/AyoLesPortal/auth"
	"github.com/renosyah/AyoLesPortal/router"
	"github.com/renosyah/AyoLesPortal/util"
)

var (
	req     *util.PostData
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use: "app",
	PreRun: func(cmd *cobra.Command, args []string) {

		auth.Init()
		router.Init(req, viper.GetString("template.path"))

	},
	Run: func(cmd *cobra.Command, args []string) {

		r := mux.NewRouter()

		r.Handle("/", auth.EmptyMiddleware(http.HandlerFunc(router.Index)))
		r.Handle("/login", auth.EmptyMiddleware(http.HandlerFunc(router.Login)))
		r.Handle("/submit/login", auth.EmptyMiddleware(http.HandlerFunc(router.SubmitLogin)))

		r.Handle("/register", auth.EmptyMiddleware(http.HandlerFunc(router.Register)))
		r.Handle("/submit/register", auth.EmptyMiddleware(http.HandlerFunc(router.SubmitRegister)))

		r.Handle("/update/profile", auth.SessionMiddleware(http.HandlerFunc(router.UpdateProfile)))
		r.Handle("/logout", auth.EmptyMiddleware(http.HandlerFunc(router.Logout)))

		r.Handle("/dashboard", auth.SessionMiddleware(http.HandlerFunc(router.Dashboard)))

		// static file serve server
		r.PathPrefix("/data/").Handler(http.StripPrefix("/data/", http.FileServer(http.Dir(viper.GetString("dir.files")))))

		r.HandleFunc("/error", router.ErrorPage)
		r.NotFoundHandler = r.NewRoute().HandlerFunc(router.NotFound).GetHandler()

		port := viper.GetInt("app.port")
		p := os.Getenv("PORT")
		if p != "" {
			port, _ = strconv.Atoi(p)
		}

		server := &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			Handler:      r,
			ReadTimeout:  time.Duration(viper.GetInt("read_timeout")) * time.Second,
			WriteTimeout: time.Duration(viper.GetInt("write_timeout")) * time.Second,
			IdleTimeout:  time.Duration(viper.GetInt("idle_timeout")) * time.Second,
		}

		done := make(chan bool, 1)
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, os.Interrupt)

		go func() {
			<-quit
			log.Println("Server is shutting down...")

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			server.SetKeepAlivesEnabled(false)
			if err := server.Shutdown(ctx); err != nil {
				log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
			}
			close(done)
		}()

		log.Println("Server is ready to handle requests at", fmt.Sprintf(":%d", port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", fmt.Sprintf(":%d", port), err)
		}

		<-done
		log.Println("Server stopped")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is github.com/renosyah/AyoLesPortal/.server.toml)")
	cobra.OnInitialize(initConfig, initRequestPost)
}

func initRequestPost() {
	r := &util.PostData{
		URL: viper.GetString("server.url"),
	}
	req = r
}

func initConfig() {
	viper.SetConfigType("toml")
	if cfgFile != "" {

		viper.SetConfigFile(cfgFile)
	} else {

		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/AyoLesPortal")
		viper.SetConfigName(".server")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
