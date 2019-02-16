package main

import (
	"context"
	"github.com/anthonydenecheau/gopocservice/config/db"
	"github.com/anthonydenecheau/gopocservice/config/middleware"
	person "github.com/anthonydenecheau/gopocservice/repository"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

// Librairies
// https://hackernoon.com/the-myth-about-golang-frameworks-and-external-libraries-93cb4b7da50f
// https://www.getrevue.co/profile/golang/issues/writing-a-go-chat-server-the-myths-about-golang-frameworks-much-more-140766
// https://qiita.com/moz450/items/bdd0eb8dff24caa5174a

// https://github.com/sepulsa/rest_echo
// https://github.com/uchonyy/echo-rest-api
// https://github.com/PacktPublishing/Echo-Essentials/tree/master/chapter8
// go get -u github.com/labstack/echo

// Functions
// https://github.com/s1s1ty/Data-Structures-and-Algorithms

// Architecture
// https://hackernoon.com/golang-clean-archithecture-efd6d7c43047
// https://github.com/hirotakan/go-cleanarchitecture-sample
// https://github.com/bxcodec/go-clean-arch
//
// Google Cloud
// https://github.com/abronan/todo-grpc/blob/master/main.go
func main() {

	// Db Connection
	person.InitPerson(&person.DbPerson{Db: db.Connect()})

	// Echo instance
	r := middleware.NewRouter()

	// Start Server
	go func() {
		r.Logger.Info("Starting Server")
		if err := r.Start(":8080"); err != nil {
			r.Logger.Info("shutting down the server")
		}
	}()

	// Graceful Shutdown
	waitForShutdown(r)
}

func waitForShutdown(r *echo.Echo) {
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.Shutdown(ctx); err != nil {
		r.Logger.Fatal(err)
	}
}
