package webservice

import (
	"dot"
	"log"
	"io/ioutil"
	"strconv"
	"encoding/base64"
)

func (item *Item) FromDot(d *dot.Dot){
	
	if d_id := d.GetKChild("id");d_id!=nil {
		id,err := strconv.Atoi(d_id.GetCValue())
		if err == nil { item.id = id }			
	}
	
	if d_cont := d.GetKChild("content");d_cont!=nil {
		item.content = d_cont.GetCValue()			
	}
	
	if d_owner := d.GetKChild("owner");d_owner!=nil {
		item.owner = d_owner.GetCValue()			
	}
	
	if d_subj := d.GetKChild("subject");d_subj!=nil {		 	
		item.subject = d_subj.GetCValue()
	
	}
	
	log.Println("{",item.id,"}")

}


func (m_item *M_Item) FromDot(d *dot.Dot){
	
	for k := range *m_item{ delete(*m_item,k) }
		
	l_item := d.GetChildList()
	log.Println("loading",len(l_item),"item")
	
	for _,d_item := range l_item {
		item := new(Item)
		item.FromDot(d_item)
		
		kstr := d_item.GetValue()
		key,err := strconv.Atoi(kstr)
		log.Println("key:",key)
		if err == nil {
			(*m_item)[key] = item
		} else { log.Println("[",[]byte(kstr),"]") }
		
				
	}
	
	log.Println("loaded ",len(*m_item),"item")
		
} 

func (cmt *Comment) FromDot(d *dot.Dot) {
			
	if d_owner :=  d.GetKChild("owner");d_owner != nil {
		cmt.owner = d_owner.GetCValue()
	}
	
	if d_cmt :=  d.GetKChild("cmt");d_cmt != nil {
		cmt.cmt = d_cmt.GetCValue()
	}
}


func (l_cmt L_Cmt) FromDot(d *dot.Dot) L_Cmt{

	l_cmt = make([]Comment,0,10)
	//l_cmt = l_cmt[0:0]

	l_child := d.GetChildList()
	for _,d_cmt := range l_child {
		cmt := Comment{"",""}
		cmt.FromDot(d_cmt)
		l_cmt = append(l_cmt,cmt)			
	}
	
	log.Println("loaded",len(l_cmt),"comments ") 
	return l_cmt
}


func (m_l_cmt M_L_Cmt) FromDot(d *dot.Dot) {

	for k := range m_l_cmt{ delete(m_l_cmt,k) }
		
	l_child := d.GetChildList()
	//log.Println("loading",len(l_child),"item")
	
	for _,d_lcmt := range l_child {
		var l_cmt L_Cmt						
		l_cmt = l_cmt.FromDot(d_lcmt)
				
		kstr := d_lcmt.GetValue()
		key,err := strconv.Atoi(kstr)
		//log.Println("key:",key)
		if err == nil {
			m_l_cmt[key] = l_cmt
		} else { log.Println("[",[]byte(kstr),"]") }		
				
	}

}

func (wi *WebImage) FromDot(d *dot.Dot) {

	d_img := d.GetKChild("img")
	d_thumb := d.GetKChild("thumb")
	
	if d_img != nil {
		imgstr := d_img.GetCValue()
		buff := Decode(nil,[]byte(imgstr),base64.StdEncoding)
		img,err := ByteToImage(buff)	
		if err == nil {
			wi.Img = img		
		}
	} 
	
	if d_thumb != nil {
		imgstr := d_thumb.GetCValue()
		buff := Decode(nil,[]byte(imgstr),base64.StdEncoding)
		img,err := ByteToImage(buff)	
		if err == nil {
			wi.Thumb = img		
		}
	} 
}


func (l_wi L_WImage) FromDot(d *dot.Dot) L_WImage {

	l_wi = make([]*WebImage,0,10)
	
	l_child := d.GetChildList()
	for _,d_wi := range l_child {
		wi := new(WebImage)
		wi.FromDot(d_wi)
		l_wi = append(l_wi,wi)			
	}
	
	log.Println("loaded",len(l_wi),"images ") 
	return l_wi

} 


func (m_l_wi M_L_WImage ) FromDot(d *dot.Dot) {

	for k := range m_l_wi{ delete(m_l_wi,k) }
		
	l_child := d.GetChildList()
		
	for _,d_lwi := range l_child {
		var l_wi L_WImage						
		l_wi = l_wi.FromDot(d_lwi)
				
		kstr := d_lwi.GetValue()
		key,err := strconv.Atoi(kstr)
		//log.Println("key:",key)
		if err == nil {
			m_l_wi[key] = l_wi
		} else { log.Println("[",[]byte(kstr),"]") }		
				
	}

}




func (board *Board) LoadBoard(fname string)string{

	log.Println("now restored board from",fname )
	buff,err := ioutil.ReadFile(fname)
	
	if err != nil {
		return "error:" + err.Error()	
	}
	 
	root := dot.CreateDot(string(buff))
	
	d_contents := root.GetKChild("contents")
	d_comments := root.GetKChild("comments")
	d_wimages := root.GetKChild("wimages")
	
	
	if d_contents != nil  { 
		board.Contents.FromDot(d_contents)	
	}
		
		
	if d_comments != nil  { 
		board.Comments.FromDot(d_comments)	
	}
	
	if d_wimages != nil  { 
		board.WImages.FromDot(d_wimages)	
	}
	
	board.lastId = 0
	for k,_ := range board.Contents {
		log.Print(k,"-")
		if k > board.lastId { board.lastId = k }	
	}
	
	board.lastId++
	
	
	log.Println("current item cnt:",len(board.Contents)," lastid:",board.lastId)
		
	return ""

}






