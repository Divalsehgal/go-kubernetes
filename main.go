package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"crypto/tls"
	"/home/nineleaps/goProjects/go-kubernetes/helloworld.crt"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "Developer"
	}
	log.Printf("Received request for %s\n", name)
	w.Write([]byte(fmt.Sprintf("Hi there !, %s\n", name)))

}


func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", handler)
	cert,_:=tls.LoadX509KeyPair("helloworld.crt","helloworld.key")



	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		TLSConfig:&tls.Config{
			Certificates:[]tls.Certificate{cert},
		},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	// go func() {
	// 	log.Println("Starting Server")
	// 	if err := srv.ListenAndServe(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServeTLS("helloworld.crt","helloworld.key"); err != nil {
			log.Fatal("error",err)
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

//example use curl localhost:8080/?name=dival
