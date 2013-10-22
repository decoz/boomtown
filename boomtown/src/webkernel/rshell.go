package webkernel

import (
	"websocket"
	"fmt"	
	"log"
)

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
		
	kernel  Webkernel
			
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


func ( rs *RShell ) SetKernel(kernel Webkernel){
/*
	kernel
*/
	fmt.Println("link shell to kernel " , kernel.KernelInfo)
	rs.kernel = kernel	
	
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
	
	defer rs.kernel.UnLinkRs(rs.key) 
}




func ( rs *RShell ) Run(){	
	fmt.Println("rs_running")
	go rs.listen()
	for {
		select {		
			case indata  := <-rs.In:
				rs.ExeCmd(string(indata))
				
				/*
				if result != "" {
					rs.Ws.Write([]byte(result))
				}
				*/
				
							
			case outdata := <-rs.Out:
				rs.Ws.Write((outdata))
		}	
	}	
	
	
	
}


func ( rs *RShell ) ExeCmd(cmd string) {
/*
 *	명령을 처리하는 메소드 
 */
 	log.Print("send request ", cmd , " to " , rs.kernel.KernelInfo())
 	result  := ""
 	if rs.kernel != nil  { 	
		//result := rs.kernel.Require("list")
		result = rs.kernel.Require(cmd,rs)	
		if result == "" { result = "bad command" }		
		log.Println("request ok ")
	} else {
		log.Println("missing kernel")
	}
	
} 






