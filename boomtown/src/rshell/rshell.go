package rshell

import (
	"websocket"
	"fmt"	
	"log"

)

type Service interface {
	
	Request(int, []byte) [][]byte
	KernelInfo() string
			
}


type RShell struct {
/*
 *	사용자의 명령을 처리하는 클래스 
 */
 	key		int
	Id 		string 	
	Path 	string
	Ws  	*websocket.Conn
	
	In 	chan []byte
	Out chan []byte
		
	service  Service
			
}


func CreateRShell( ws *websocket.Conn ) *RShell{
	rs := new(RShell)
	rs.In = make(chan []byte)
	rs.Out = make(chan []byte)
	rs.Ws = ws
	
	//rs.ExeCmd = func(s string)string{return ""}
	return rs
}


type Cmd struct {	
	sender 	string
	cmd		string
	data	string
}


func ( rs *RShell ) SetKernel(srv Service){
/*
	kernel
*/	
	fmt.Println("link shell to kernel " , srv.KernelInfo())
	rs.service = srv	
	
}


func ( rs *RShell ) listen(){

	msg := make([]byte,1024 * 100)
	for{		
		size,_ := rs.Ws.Read(msg)
		
		if(size > 2){
			log.Println("rs_listen: ",string(msg[0:size]))				
			rs.In <- msg[0:size]
		}
	}		
	
	 
}





