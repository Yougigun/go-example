package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"sync/atomic"
)

var b = aaa
var aaa int

func main() {

	// pprofMux := http.DefaultServeMux
	// http.DefaultServeMux = http.NewServeMux()
	var counts int64 = 0
	var counts2 int64 = 0
	var counts3 int64 = 0
	var counts4 int64 = 0

	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		// Cache-Control
		// no-cache: irregardless of max-age, always validate with server using with cache-related header like Etag/If-None-Match.
		// no-store: will not cache the response and the next request will not include cache-related header like If-None-Match.
		atomic.AddInt64(&counts, 1)
		fmt.Println(counts)
		w.Header().Add("Cache-Control", "max-age=10")
		w.Header().Add("Content-Type", "text/html")
		w.Header().Add("Etag", "xxx")
		// w.WriteHeader(403)
		w.Write([]byte(fmt.Sprintf("counts: %v", counts) + " <a href='test'>go to link</a>"))
	}

	// Eag and If-None-Match
	root2Handler := func(w http.ResponseWriter, r *http.Request) {
		// alway use cache work for f5
		tag:= "never-change"
		atomic.AddInt64(&counts2, 1)
		if r.Header.Get("If-None-Match") == tag {
			fmt.Println("return 304")
			w.WriteHeader(http.StatusNotModified)
			w.Write([]byte("xxxxx"))
			return
		}
		fmt.Println(counts2)
		w.Header().Add("Content-Type", "text/html")
		w.Header().Add("Etag", tag)
		w.Write([]byte(fmt.Sprintf("counts: %v", counts2) + " <a href='test'>go to link</a>"))
	}
	root3Handler := func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&counts3, 1)
		if r.Header.Get("If-None-Match") == strconv.Itoa(int(counts2)) {
			fmt.Println("return 304")
			w.WriteHeader(http.StatusNotModified)
			return
		}
		fmt.Println(counts3)
		w.Header().Add("Content-Type", "text/html")
		w.Header().Add("Etag", strconv.Itoa(int(counts3)))
		w.Write([]byte(fmt.Sprintf("counts: %v", counts) + " <a href='test'>go to link</a>"))
	}
	// Last Modified and If-Modified-Since
	root4Handler := func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&counts4, 1)
		modifiedDate := "2012.2.2"
		if r.Header.Get("If-Modified-Since") == modifiedDate {
			fmt.Println("return 304")
			w.WriteHeader(http.StatusNotModified)
			return
		}
		fmt.Println(counts4)
		w.Header().Add("Content-Type", "text/html")
		w.Header().Add("Last-Modified", modifiedDate)
		w.Write([]byte(fmt.Sprintf("counts: %v", counts4) + " <a href='test'>go to link</a>"))
	}

	// using custom httpmux
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	go http.ListenAndServe(":8001", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", root2Handler)
	go http.ListenAndServe(":8002", mux2)

	mux3 := http.NewServeMux()
	mux3.HandleFunc("/", root3Handler)
	go http.ListenAndServe(":8003", mux3)
	

	mux4 := http.NewServeMux()
	mux4.HandleFunc("/", root4Handler)
	server4 := &http.Server{Addr: ":8004", Handler: mux4}
	server4.ListenAndServe()

}
