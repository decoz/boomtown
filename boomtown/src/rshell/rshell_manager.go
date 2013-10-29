
package rshell

import (
	
)
/*



type RsManager struct{
	rsmap map[int] *RShell	
		
}


func CreateRsManager() *RsManager{
	rm := new(RsManager)
	rm.initrm()	
	return rm
}

func (rm *RsManager) initrm(){
	rm.rsmap = make(map[int] *RShell)
}

func (rm *RsManager) addNewRs(rs *RShell){
	newkey := getEmptyKey(rm.rsmap)
	rm.rsmap[newkey] = rs	
	rs.key = newkey
	 
}

func (rm *RsManager) removeRs(key int){
	delete(rm.rsmap,key)
}


func (rm *RsManager) Broadcast(msg []byte){
	
	for _,rs := range(rm.rsmap) {
				
		//rs.Out <- msg
		if rs!=nil { rs.Ws.Write(msg) }	
		
	}
}

func getEmptyKey(m map[int] *RShell) int {
	
	nkey := 0
	for _,ok := m[nkey];ok; {
		nkey += 1
		_,ok = m[nkey]	
	}
	
	return nkey
}
*/