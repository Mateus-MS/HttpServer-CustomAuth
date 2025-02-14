package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	router := http.NewServeMux()

	startServer(router, os.Getenv("ENV"))

}

func startServer(router *http.ServeMux, env string) {
	if env == "dev" {
		fmt.Println("Starting SERVER in DEV mode")
		go func() {
			if err := http.ListenAndServe(":3000", router); err != nil {
				fmt.Println("HTTP server error:", err)
			}
		}()
		select {}
	}

	certManager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("./backend/certs"),
		HostPolicy: autocert.HostWhitelist("dommain.com"),
	}

	go func() {
		httpServer := &http.Server{
			Addr:    ":80",
			Handler: certManager.HTTPHandler(nil),
		}

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	httpsServer := &http.Server{
		Addr:      ":443",
		Handler:   router,
		TLSConfig: certManager.TLSConfig(),
	}

	if err := httpsServer.ListenAndServeTLS("", ""); err != nil {
		fmt.Println("HTTPS server error:", err)
	}

	fmt.Println("Starting SERVER")
}
