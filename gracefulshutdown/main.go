package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
)
var b = aaa
var aaa int
func main() {

	// pprofMux := http.DefaultServeMux
	// http.DefaultServeMux = http.NewServeMux()
	rootHandler := func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w,"hello I am root in service 1")
	}
	root2Handler := func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w,"hello I am root service 2")
	}
	root3Handler := func(w http.ResponseWriter, r *http.Request){
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
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func(){
		defer wg.Done()
		err := server3.ListenAndServe()
		fmt.Println("server3 :",err)
	}()

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		c := make(chan os.Signal,1)
		signal.Notify(c,os.Interrupt,syscall.SIGTERM)
		<-c
		err := server3.Shutdown(context.Background())
		fmt.Println("shutdown",err)
	}()

	wg.Wait()
	fmt.Println("main exit")

}