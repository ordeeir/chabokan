package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const (
	DateTime = "2006-01-02 15:04:05"
)

var dip string

func main() {

	//dip = os.Getenv("DIP")

	//if dip == "" {
	//dip = os.Args[1]
	//}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, _ := os.ReadFile("index.html")
		fmt.Fprint(w, string(b))
	})

	// Anything we don't do in Go, we pass to the old platform
	http.HandleFunc("/", blog) //

	fmt.Println("Listening..............................................")

	// Start the server
	http.ListenAndServe(":4430", nil)
}

func blog(w http.ResponseWriter, r *http.Request) {

	fmt.Println("data is transfering...")

	// change the request host to match the target
	//vars := strings.Split(r.URL.Path, "/")

	//data, _ := base64.StdEncoding.DecodeString(vars[2])

	//slices := strings.Split(string(data)+"///", "/")

	//fmt.Println("---- User: " + slices[1] + " | Address: " + slices[2] + ":" + slices[3] + " (" + slices[4] + ") | Host1: " + r.Host + " | Host2: " + slices[5] + " | ")

	u, _ := url.Parse("https://srvgr3.farsino.xyz")
	pro := httputil.NewSingleHostReverseProxy(u)

	//vars[1] = vars[1] + "444"
	//r.URL.Path = strings.Join(vars, "/")
	//fmt.Println("---- r.URL.Path: " + r.URL.Path)

	pro.ServeHTTP(w, r)
}
