package main

import (
	"context"
	"github.com/anthonydenecheau/gopocservice/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Notes :
// https://www.jetbrains.com/help/go/install-and-set-up-product.html
// https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
// https://itnext.io/structuring-a-production-grade-rest-api-in-golang-c0229b3feedc
// https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b
// https://github.com/tsuru/tsuru/blob/master/router/router.go
//https://thenewstack.io/make-a-restful-json-api-go/

// Librairies
// https://hackernoon.com/the-myth-about-golang-frameworks-and-external-libraries-93cb4b7da50f
// https://www.getrevue.co/profile/golang/issues/writing-a-go-chat-server-the-myths-about-golang-frameworks-much-more-140766

// Projet
// https://github.com/uchonyy/echo-rest-api

// Docker :
// https://www.callicoder.com/docker-golang-image-container-example/
// https://container-solutions.com/faster-builds-in-docker-with-go-1-11/
func main() {
	//teest
	r := router.NewRouter()

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)

}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
