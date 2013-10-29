package main

import (
	
	"websocket"
	"webservice"
	"rshell"
	"fmt"
	"net/http"
	"runtime"
	
	 	
)

type LineReader interface {
	ReadLine() ([]byte, bool, error)
		
}


// This example demonstrates a trivial echo server.
func main() {

	runtime.GOMAXPROCS(10000)
	rsm := rshell.CreateRsManager()
		
	boardsrv := webservice.CreateBoard()
	imgsrv := webservice.CreateImgserve()
	
	rsm.SetDefaultService(boardsrv)
	go rsm.Run()
	
	url := ":8515"
	fmt.Println("boomtown server is running at ",url)
		 
	http.Handle("/ws", websocket.Handler(rsm.GetHandler(boardsrv)))
	http.Handle("/img", websocket.Handler(rsm.GetHandler(imgsrv)))
	err := http.ListenAndServe(url, nil)	
	
		
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
	
	

	
}
