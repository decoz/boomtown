package main

import (
	
	"bufio"
	"os"	
	"net/http"
	"websocket"
	"webkernel"
	"multichannel"
	"fmt"
	"log"
	//"strconv"
	
)

type LineReader interface {
	ReadLine() ([]byte, bool, error)
		
}

// Echo the data received on the WebSocket.
func serviceCaster(kernel webkernel.Webkernel) func(ws *websocket.Conn) {
				
	count := 0	
		
	return func(ws *websocket.Conn){
		
		count += 1
		rs := webkernel.CreateRShell(ws)
	
	
		kernel.LinkRs(rs)			
		rs.Run()
	}
	

}

 

func ReadInput(input LineReader,board *webkernel.Board) {
	log.Println("console reader start")
	
	for {
		fmt.Print("ws_v0.12:")
		cmd,_,_ := input.ReadLine()	
		board.Broadcast(cmd)		
		fmt.Println(string(cmd))		
	}	
}



// This example demonstrates a trivial echo server.
func main() {


	broadcaster := multichannel.CreateMChan()
	go broadcaster.Run()	
	
	kernel := webkernel.CreateBoard()
	kernel.DummyInit()
	fmt.Println(kernel.KernelInfo(), " is loading")
	svr := serviceCaster(kernel)

	stdin := bufio.NewReader(os.Stdin)	
	go ReadInput(stdin,kernel)	 
		
	
	url := ":8515"
	fmt.Println("server is running at ",url)
	
	http.Handle("/ws", websocket.Handler(svr))
	err := http.ListenAndServe(url, nil)	
	
		
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
	
	

	
}
