package boomhttp

import (
	"net/http"
	"log"
)

func StartHttp(){

	//http.Handle("/boom",http.HandleFunc(load_client)) 
	//err := http.ListenAndServe(":8080",nil)
	log.Fatal(http.ListenAndServe(":80", http.FileServer(http.Dir("html"))))
		

}
