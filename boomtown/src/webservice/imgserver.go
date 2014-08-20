package webservice

import (
	"log"
	//"io/ioutil"
	"encoding/base64"	
	"image"
	"strconv"
	"dot"
	"rshell"
	
)


type ImgServe struct{
	 
	msgout *chan rshell.RsMsg
	imglist map[string] *image.Image
	last int
	
	 	
}

func CreateImgserve() *ImgServe {

	imgsrv := new(ImgServe) 
	imgsrv.imglist = make(map[string] *image.Image)
	imgsrv.last = 0
	 
	return imgsrv

}


func (imgsrv ImgServe) KernelInfo() string{

	return  "imgServe v0.1" 

} 

func (imgsrv ImgServe) Unlink(key int){
	 

} 

func (imgsrv *ImgServe) SetChannel(rsmch *chan rshell.RsMsg){
	imgsrv.msgout = rsmch
}

/*

func (imgsrv Imgserve) LinkRs( rs *RShell ) {



}


func (imgsrv Imgserve) UnLinkRs(key int) {



}
*/
func (imgsrv *ImgServe) send(cnum int, content string){
	log.Println("send to",cnum,":"+content)
	smsg := rshell.RsMsg{Key:cnum,Msg:[]byte(content)} 
	*imgsrv.msgout <- smsg
}

func (imgsrv *ImgServe) Request(key int ,data []byte){
	log.Println("request for imgsrv")
	 
	//cmd,_  := getString(data)
	
	/*		
	buff := Decode(nil,data,base64.StdEncoding)	
	img,_ := ByteToImage(buff)	
	wi := CreateWebImage(img,"img"+strconv.Itoa(imgsrv.lastindex))
	imgsrv.imglist[imgsrv.lastindex] = wi 
	imgsrv.lastindex++
	log.Println("img count:",imgsrv.lastindex)			
	//SaveImage(wi.Thumb,"last_thumb.jpg","jpg")	
				
	//ioutil.WriteFile("last.png",buff,0644)
	result := wi.GetThumbE64()  
	log.Println("return byte:",len(result))
	*/
	
	
	d_req := dot.CreateDot(string(data))
	//result = d_req.ToString()
	
	cmd := d_req.GetKChild("cmd")
	param := d_req.GetKChild("param")
	attr := d_req.GetKChild("attr")
	option := d_req.GetKChild("option")
		
	imgsrv.ExeCmd(key,cmd,param,attr,option)	
	//imgsrv.send(key,result)	
	
	
}

func (imgsrv *ImgServe) ExeCmd(id int,cmd,param,attr,option *dot.Dot){

	//log.Println("req from",id,":","["+cmd.GetCValue()+"]")
	switch(cmd.GetCValue()){
		
		case "put": 
			//log.Println("processing save transfered image") 
			d_src := attr.GetChild("img.src")
			if d_src == nil { return }
 			
			src := d_src.GetCValue() 
			buff := Decode(nil,[]byte(src),base64.StdEncoding)
			img,_ := ByteToImage(buff)
						
			d_as := attr.GetChild("as")
			name := ""
			
			if d_as != nil {
				name = d_as.GetCValue()							
			} else {
				name = "img" + strconv.Itoa(imgsrv.last) 			
				imgsrv.last++
			} 
			
			imgsrv.imglist[name] = &img
			
			imgsrv.send(id,"image saved  as " + name)			
		
		case "get":	
		
	
	}
	
}


func Decode(decBuf, enc []byte, e64 *base64.Encoding) []byte { 
        maxDecLen := e64.DecodedLen(len(enc)) 
        if decBuf == nil || len(decBuf) < maxDecLen { 
                decBuf = make([]byte, maxDecLen) 
        } 
        n, err := e64.Decode(decBuf, enc) 
        _ = err 
        return decBuf[0:n] 
} 

func Encode(encBuf, bin []byte, e64 *base64.Encoding) []byte { 
        maxEncLen := e64.EncodedLen(len(bin)) 
        if encBuf == nil || len(encBuf) < maxEncLen { 
                encBuf = make([]byte, maxEncLen) 
        } 
        e64.Encode(encBuf, bin) 
        return encBuf[0:] 
} 

/*

func getString(buff []byte) (string,int) {
	
	pos := 0
	for ptr,c := range(buff){
		if c == 0 {
			pos = ptr
			break;		
		}
	}
	
	
	return string(buff[:pos]),pos
}  

*/

