package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

const (
	DateTime = "2006-01-02 15:04:05"
)

var dip string

func main() {

	dip = os.Getenv("DIP")

	if dip == "" {
		dip = os.Args[1]
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, _ := os.ReadFile("index.html")
		fmt.Fprint(w, string(b))
	})

	// Anything we don't do in Go, we pass to the old platform
	http.HandleFunc("/blog/", blog) //

	fmt.Println("Listening on localhost:443")

	// Start the server
	http.ListenAndServe(":443", nil)
}

func blog(w http.ResponseWriter, r *http.Request) {

	// change the request host to match the target
	vars := strings.Split(r.URL.Path, "/")

	data, _ := base64.StdEncoding.DecodeString(vars[2])

	slices := strings.Split(string(data)+"///", "/")

	fmt.Println("---- User: " + slices[1] + " | Address: " + slices[2] + ":" + slices[3] + " (" + slices[4] + ") | Host1: " + r.Host + " | Host2: " + slices[5] + " | ")

	u, _ := url.Parse("http://" + dip + ":" + slices[0])
	pro := httputil.NewSingleHostReverseProxy(u)

	vars[1] = vars[1] + "444"
	r.URL.Path = "/"
	//r.URL.Path = strings.Join(vars, "/")
	fmt.Println("---- r.URL.Path: " + r.URL.Path)

	pro.ServeHTTP(w, r)
}
