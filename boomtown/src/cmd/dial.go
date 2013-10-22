package main

import (
	"websocket"
	"fmt"
	"log"
)

func main() {
	origin := "http://localhost/"
	url := "ws://localhost:8515/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
		//fmt.Println(err)
	}
	
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
			log.Fatal(err)
		}
	fmt.Println("----")
	

	
	var msg = make([]byte, 512)
	var n int
	
	for{
				
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}  else {
			fmt.Println("server:",string(msg[0:n]))			
		}
		
	}
	
	fmt.Printf("Received: %s.\n", msg[:n])

}
