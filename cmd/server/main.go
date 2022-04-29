package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/derekargueta/irm/pkg/irm/probes"
)

type CommandArgs struct {
	port   int
	webDir string
}

func parseArgs() *CommandArgs {
	var port int
	var webDir string
	flag.IntVar(&port, "port", 8085, "TCP port that the HTTP server will listen on.")
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
type myvals struct {
	Http10 string
	Http11 string
	Http12 string
	Http13 string
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	t, err := template.ParseFiles("./web/index.html")
	if err != nil {
		fmt.Print("iuykfsduyfsdf")
	}
	t.Execute(w, nil)
}

func singledomain(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./web/index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		http10 := (&probes.Http10probe{}).Run(r.Form.Get("website"))
		http1Result := (&probes.HTTP11Probe{}).Run(r.Form.Get("website"))
		http2Result := (&probes.HTTP2Probe{}).Run(r.Form.Get("website"))
		http3Result := (&probes.HTTP3Probe{}).Run(r.Form.Get("website"))

		// logic part of log in

		test := myvals{
			Http10: strconv.FormatBool(http10.Supported),
			Http11: strconv.FormatBool(http1Result.Supported),
			Http12: strconv.FormatBool(http2Result.Supported),
			Http13: strconv.FormatBool(http3Result.Supported),
		}
		t, _ := template.ParseFiles("./web/index.html")
		t.Execute(w, test)
		//	fmt.Println("website:", r.Form["website"])

	}
}

func main() {
	args := parseArgs()
	//Read the index.html into memory.
	// dat, err := ioutil.ReadFile(args.webDir + "/index.html")
	// if err != nil {
	// 	panic(err)
	// }
	// content := string(dat)

	// // Create a handler that will serve the file.
	// handler := &handler{content: content}
	// http.Handle("/", handler)

	// http.Handle("/", http.FileServer(http.Dir(args.webDir)))
	// http.Handle("/js", http.FileServer(http.Dir(args.webDir+"/js")))

	http.HandleFunc("/", ServeHTTP)
	http.HandleFunc("/singledomain", singledomain)
	http.HandleFunc("/about.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/about.html")
	})
	//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(args.webDir))))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/css"))))

	listenAddr := fmt.Sprintf(":%d", args.port)
	log.Fatal(http.ListenAndServe("127.0.0.1"+listenAddr, nil))
}
