package multichannel

import (
	"fmt"
)

type MChan struct {	
	InChMap map[int] *chan []byte
	OutChMap map[int] *chan []byte
	
	In chan []byte
	Out chan []byte
}

func CreateMChan() *MChan{

	mchan := new(MChan)
	mchan.InChMap = make(map[int] *chan []byte) 
	mchan.OutChMap = make(map[int] *chan []byte)
	mchan.In = make(chan []byte,10)
	mchan.Out = make(chan []byte,10)	
	
	return mchan	
}

func getEmptyKey(m map[int] *chan []byte) int {
	
	nkey := 0
	for _,ok := m[nkey];ok; {
		nkey += 1
		_,ok = m[nkey]	
	}
	
	return nkey
}

func (mchan MChan) AddIOChannel(ich,och *chan []byte) int {
/*
 * add input channel by searching empty int key value 
 */
	key := getEmptyKey(mchan.InChMap)
	mchan.InChMap[key] = ich
	mchan.OutChMap[key] = och
	return key
}   

func (mchan MChan) deleteCh(key int){

	delete(mchan.InChMap, key)
	delete(mchan.OutChMap, key)

} 

func (mchan MChan) listen(){
	fmt.Println("mc listner: start for " , len(mchan.InChMap))
	/* 
	for{	
		for key,ch := range(mchan.InChMap) {			
			fmt.Println(key)			
			msg := <-*ch
			mchan.In <- msg		
			
		}					  		
	}
	*/		
}

func (mchan MChan) Run(){
	fmt.Println("mc_running")
	go mchan.listen()
	
	
	for	{	
		msg:= <-mchan.Out
		fmt.Println(string(msg))		
		for _,ch := range(mchan.OutChMap) {				
			*ch  <- []byte(msg)  		
		}	
	}
	
}





func (mchan MChan) Broadcast(msg string){
	
	for _,ch := range(mchan.OutChMap) {				
		*ch  <- []byte(msg)  		
	} 	

}