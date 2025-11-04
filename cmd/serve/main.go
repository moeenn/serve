package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

const (
	DEFAULT_PORT int    = 8000
	DEFAULT_DIR  string = "."
)

func startServer(dir string, port int) error {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(dir))
	mux.Handle("/", fs)
	fmt.Printf("Starting server: http://localhost:%d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func run() error {
	port := flag.Int("port", DEFAULT_PORT, "Port to run file server on.")
	dir := flag.String("dir", DEFAULT_DIR, "Directory to serve.")
	flag.Parse()

	return startServer(*dir, *port)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
