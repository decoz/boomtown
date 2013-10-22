package main

import (
	//"multichannel"
	"fmt"
)

func main(){
	
	
	
}



func thr(ch *chan []byte,key int){
	for{
	cmd := <- *ch 
	fmt.Println(key,":",string(cmd))
	}
	
}