
package rshell

import (
	"log"
	"websocket"
	"fmt"
	"strings"
	"strconv"
	//"bytes"
)

type  RsManager struct{
	lastid int
	rsmap map[int] *RShell
	newconn	chan *websocket.Conn	
	msgin	chan RsMsg
	
	defsrv 	Service
	
}

type RsMsg  struct {
	key int
	msg []byte
}


func CreateRsManager() *RsManager{
	 
	rsm := new( RsManager )
	rsm.rsmap = make(map[int] *RShell)
	rsm.newconn = make(chan *websocket.Conn)
	rsm.msgin = make(chan RsMsg)
	rsm.lastid = 0
	
	return rsm
}

func (rsm *RsManager) Run(){
	fmt.Println("RS Manager running")
	for{log.Print("-")
		select {			
			case rsmsg :=  <- rsm.msgin:
				//log.Println("request from ",rsmsg.key," :",string(rsmsg.msg))
				key := rsmsg.key	
				rs,_ :=  rsm.rsmap[key]		
				results := rs.service.Request(key,rsmsg.msg)
				for _,result := range(results){ 
					rsm.patchResult(key,result)
				}
				
								 
		}
	}
	defer func() { log.Println("end run RSM :") }()
}



func (rsm *RsManager) patchResult(key int,result []byte){
	if len(result) < 1 { return } 
	
	head :=  string(result[0:1])
	
	switch head {
	case "%" : 
		for _,rs := range(rsm.rsmap){
			rs.Ws.Write(result[1:])		
		}		
	case "[" : 
		pos := strings.IndexRune(string(result),']')
		
		targets := string(result[1:pos])
		message := result[pos+1:]
		
		list  := strings.Split(targets,",")
		for _,kstr := range(list) {
			rkey,_ := strconv.Atoi(kstr)
			if rs,ok := rsm.rsmap[rkey];ok{
				rs.Ws.Write(message)			
			}					
		}
	default:
		rs,_ := rsm.rsmap[key]
		rs.Ws.Write(result)
	}
		
	/*	
	if str := string(result[0:1]);str  == "%" {
		for _,rs := range(rsm.rsmap){
			rs.Ws.Write(result[1:])		
		}
	} else { 
		rs,_ := rsm.rsmap[key]
		rs.Ws.Write(result)	
	}
	*/
}




func (rsm *RsManager) SetDefaultService(service Service){
/* default kernel 을 설정한다. */
	rsm.defsrv = service

}
 
func (rsm *RsManager) GetHandler(srv Service) func(ws *websocket.Conn) {
			
	return func(ws *websocket.Conn){
		log.Println("ws: new connection")
		log.Println("connection to:",srv.KernelInfo())
		
		//rsm.newconn <- ws
		rs := CreateRShell(ws)				 
		if srv == nil {
			rs.SetKernel(rsm.defsrv)
		}else {
			rs.SetKernel(srv)
		}
		
		key := rsm.lastid
		rsm.lastid++		
		rsm.rsmap[key] = rs						
		rsm.listen(key)
		
		
		
		defer func() { log.Println("handler out") }()	
	}

}

func  (rsm *RsManager) listen(key int){
	log.Println("now listen from rs:",key)
	rs,_ :=  rsm.rsmap[key]
	msg := make([]byte,1024 * 1024)
	for rs.Ws.IsServerConn() {		
		
		//size,_ := rs.Ws.Read(msg)
		err := websocket.Message.Receive(rs.Ws,&msg)
		if err != nil { return } 
		
		
		log.Println("len:",len(msg))
		rsm.msgin <- RsMsg{key,msg}		
		//if(size > 2){
						
			//log.Println("test:",bytes.Equal(msg[:size],msg))						
		
			//log.Println("rs",key,".listen ",size,"byte: ",string(msg[0:size]))
			//log.Println("rs",key,".listen ",size,"byte")				
			//rs.In <- msg[0:size]
			
			//rsm.msgin <- RsMsg{key, msg[0:size]}
		//}
	}		
	
	defer func() { log.Println("end listen from rs:",key) }()
}
 
 