package main

import (	
	"log"
	"dot"
	"tag"
)

type Ti interface {
	print()
}

func main() {
	//test_attach()
	test_tagDB()
	
	/*
	str := "x(a,b)"	
	p := dot.CreateParser([]byte(str))
	n := p.Parse([]byte("root"),0)
	log.Println(n.ToString())
	*/

}

func test_tagDB(){
 
	tdb := tag.CreateDotTagDB()
	tdb.AddTag("test(1,2,3)")
	tdb.PrintAll()

}

func test_attach(){

	d := dot.CreateDot("/(a,b)")
	d_attach := dot.CreateDot("b(1,2,3,4.x)")
	d.Attach(d_attach)
	log.Println(d.ToString())

}