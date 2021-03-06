package webservice

import (
	"dot"
	"log"
	"strconv"
	"io/ioutil"
	//"strings"

)


func (cmt Comment) ToDot() *dot.Dot {

	dot0 := dot.NewDot("")
	dot0.SetKChild("owner",cmt.owner)
	dot0.SetKChild("cmt",cmt.cmt)
			
	return dot0

}


func (l_cmt L_Cmt) ToDot() *dot.Dot{
	root := dot.NewDot("")   	
	for _,cmt := range(l_cmt) {
		root.AppendChild(cmt.ToDot())
	}
	return  root	
}


func (m_l_cmt M_L_Cmt) ToDot() *dot.Dot{
	
	root := dot.NewDot("comments")
	for key,l_cmt := range(m_l_cmt) {
		dot_l := l_cmt.ToDot()
		dot_l.SetValue(strconv.Itoa(key))
		root.AppendChild(dot_l)				
	}
	
	return  root
}


func (item Item) ToDot() *dot.Dot{
	root := dot.NewDot("")
	root.SetKChild("id",strconv.Itoa(item.id))
	root.SetKChild("content",item.content)
	root.SetKChild("owner",item.owner)
	root.SetKChild("subject",item.subject)
		
	return root
}  

func (m_item M_Item) ToDot() *dot.Dot{
	
	//log.Println("current count:",len(m_item))
	
	root := dot.NewDot("contents")	
	for key,item := range(m_item) {
		d :=  item.ToDot()
		d.SetValue(strconv.Itoa(key))
		root.AppendChild(d)
	}
	
	return root
}

func (wi *WebImage) ToDot() *dot.Dot{

	root := dot.NewDot("")
	root.SetKChild("img",string(wi.GetE64()))
	root.SetKChild("thumb",string(wi.GetThumbE64()))
		
	return root
}


func (l_wi L_WImage) ToDot() *dot.Dot {

	root := dot.NewDot("")
	for _,wi := range(l_wi){
		root.AppendChild(wi.ToDot())	
	}

	return root
}
 
 func (m_l_wi M_L_WImage) ToDot() *dot.Dot {
 	
 	root := dot.NewDot("wimages")  
 	for key,l_wi := range(m_l_wi){
 		d_l_wi := l_wi.ToDot()
 		d_l_wi.SetValue(strconv.Itoa(key))
 		root.AppendChild(d_l_wi) 	
 	}
 	
 	return root
 }
 
 
  
 

func (board Board) SaveBoard(fname string) {
	
	log.Println("now save board data to " + fname)
	root := dot.CreateDot("") 
	d_images := board.WImages.ToDot()
	d_contents := board.Contents.ToDot()
	d_comments := board.Comments.ToDot()
	root.AppendChild(d_contents,d_comments,d_images)
	
	
	ioutil.WriteFile(fname,[]byte(root.ToString()),0667)
	
	
}




