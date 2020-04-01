# docker file for ayolescore app
FROM golang:latest as builder
ADD . /go/src/github.com/renosyah/AyoLesPortal
WORKDIR /go/src/github.com/renosyah/AyoLesPortal
RUN go get -u github.com/golang/dep/cmd/dep
COPY . .
RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
RUN rm -rf /api
RUN rm -rf /auth
RUN rm -rf /cmd
RUN rm -rf /model
RUN rm -rf /design
RUN rm -rf /router
RUN rm -rf /vendor
RUN rm -rf /util
RUN rm .dockerignore
RUN rm .gitignore
RUN rm Dockerfile
RUN rm Gopkg.lock
RUN rm Gopkg.toml
RUN rm heroku.yml
RUN rm main.go
EXPOSE 8000
EXPOSE 80
CMD ./main --config=.server.toml
MAINTAINER syahputrareno975@gmail.com


# FROM alpine:latest  
# RUN apk --no-cache add ca-certificates
# WORKDIR /app
# COPY . .
# EXPOSE 8000
# EXPOSE 80
# CMD ./main --config=.heroku.toml
# MAINTAINER syahputrareno975@gmail.com