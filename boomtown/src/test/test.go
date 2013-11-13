package main

import (	
	"log"
	"dot"
)

func main() {

	str := "a.b"
	p := dot.CreateParser([]byte(str))
	n := p.Parse([]byte("root"),0)
	log.Println(n.ToString())

}