package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)
var b = aaa
var aaa int
func main() {

	// pprofMux := http.DefaultServeMux
	// http.DefaultServeMux = http.NewServeMux()
	rootHandler := func(w http.ResponseWriter, r *http.Request){
		fmt.Println(1)
		fmt.Println(r.Method)
		fmt.Fprint(w,"hello I am root in service 1")
	}
	root2Handler := func(w http.ResponseWriter, r *http.Request){
		fmt.Println("I am in 8002")
		fmt.Println(r.Method)

		fmt.Fprint(w,"hello I am root service 2")
	}
	root3Handler := func(w http.ResponseWriter, r *http.Request){
		fmt.Println("I am in 8003")
		
		// fmt.Println(r.URL.RawQuery)
		// fmt.Println(r.Header.Get("Cookie"))
		// fmt.Println(r.Cookies())
		fmt.Println("Host in Header",r.Header.Get("Host"))
		fmt.Println("Host in req",r.Host)
		fmt.Println("Host in url",r.URL.Host)
		fmt.Println("Path",r.URL.Path)
		fmt.Println("Scheme",r.URL.Scheme)
		fmt.Println("Header",r.Header)
		fmt.Fprint(w,"hello I am root service 3")
	}
	
	// using custom httpmux
	mux := http.NewServeMux()
	mux.HandleFunc("/",rootHandler)
	go http.ListenAndServe(":8001",mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/",root2Handler)
	go http.ListenAndServe(":8002",mux2)
	
	mux3 := http.NewServeMux()
	mux3.HandleFunc("/",root3Handler)
	server3 := &http.Server{Addr:":8003",Handler: mux3}
	server3.ListenAndServe()

}