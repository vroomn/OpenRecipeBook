package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// Supply a port to the program, defaults to 8080
	// Set as an int to make it easier to filter out bad port inputs that would come from getting strings
	port := flag.Int("port", 8080, "Defines the port the HTTP server will listen on")
	// Series of flags to enable TLS and define special file locations to look
	tls := flag.Bool("tls", false, "Boolean flag of whether or no to look for TLS-required keys")
	// TODO: Add actual TLS default file paths
	certFile := flag.String("TLScert", "", "Filepath to a TLS cert file")
	keyFile := flag.String("TLSkey", "", "Filepath to a TLS key file")

	flag.Parse() // Needs to be done before any usage of the flag vars

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome")
	})

	portStr := strconv.Itoa(*port)
	var err error
	if !*tls {
		fmt.Println("Listening on http://localhost:" + portStr)
		err = http.ListenAndServe(":"+portStr, nil)
	} else {
		fmt.Println("Listening on https://localhost:" + portStr + " [TLS Enabled]")
		err = http.ListenAndServeTLS(":"+portStr, *certFile, *keyFile, nil)
	}
	log.Fatal(err)
}
