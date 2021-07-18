package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type CommandArgs struct {
	port   int
	webDir string
}

func parseArgs() *CommandArgs {
	var port int
	var webDir string
	flag.IntVar(&port, "port", 8080, "TCP port that the HTTP server will listen on.")
	flag.StringVar(&webDir, "web-dir", "web", "The directory from which files are served over HTTP.")
	flag.Parse()
	return &CommandArgs{
		port:   port,
		webDir: webDir,
	}
}

type handler struct {
	content string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, h.content)
}

func main() {
	args := parseArgs()

	// Read the index.html into memory.
	dat, err := ioutil.ReadFile(args.webDir + "/index.html")
	if err != nil {
		panic(err)
	}
	content := string(dat)

	// Create a handler that will serve the file.
	handler := &handler{content: content}
	http.Handle("/", handler)

	// Bind the TCP listener.
	listenAddr := fmt.Sprintf(":%d", args.port)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
