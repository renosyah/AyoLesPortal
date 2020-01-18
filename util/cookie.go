package util

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

const DefaultName = "session"

var cookies = map[string]*securecookie.SecureCookie{
	"previous": securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32),
	),
	"current": securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32),
	),
}

type CookieConfig struct {
	Name string
	Sc   map[string]*securecookie.SecureCookie
}

type CookieData struct {
	Name  string
	Value string
}

func NewCookieConfig() *CookieConfig {
	return &CookieConfig{
		Name: "teacher-cookie",
		Sc:   cookies,
	}
}

func (c *CookieConfig) Set(w http.ResponseWriter, r *http.Request, cd CookieData) error {

	value := map[string]string{
		cd.Name: cd.Value,
	}
	encoded, err := securecookie.EncodeMulti(c.Name, value, c.Sc["current"])
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:  c.Name,
		Value: encoded,
		Path:  "/",
	}
	http.SetCookie(w, cookie)

	return nil
}

func (c *CookieConfig) Get(r *http.Request, cd *CookieData) error {
	cookie, err := r.Cookie(c.Name)
	if err != nil {
		return err
	}
	value := make(map[string]string)
	err = securecookie.DecodeMulti(c.Name, cookie.Value, &value, c.Sc["current"], c.Sc["previous"])
	if err != nil {
		return err
	}

	cd.Value = value[cd.Name]

	return nil
}
func (c *CookieConfig) Delete(res http.ResponseWriter) {
	v := &http.Cookie{
		Name:   c.Name,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(res, v)
}
