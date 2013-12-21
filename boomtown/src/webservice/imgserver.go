package webservice

import (
	"log"
	//"io/ioutil"
	"encoding/base64"	
	"image"
	//"strconv"
	"dot"
	"rshell"
	
)

type WebImage struct {
	
	Img image.Image
	Thumb image.Image
	SrcName string
		
}



func (wi WebImage) GetE64() []byte {
	
	buff := ImageToByte(wi.Img,"jpg")
	rbuff := Encode(nil,buff,base64.StdEncoding)
	return rbuff
			
}

func (wi WebImage) GetThumbE64() []byte {

	buff := ImageToByte(wi.Thumb,"jpg")
	rbuff := Encode(nil,buff,base64.StdEncoding)
	return rbuff

}

func CreateWebImage(img image.Image,name string) *WebImage{

	wi := new(WebImage)
	wi.Img = img
	wi.Thumb = CreateThumbnail(img,50,50)	
	wi.SrcName = name
		
	return wi
}

type ImgServe struct{
	 
	msgout *chan rshell.RsMsg
	imglist map[int] *WebImage
	lastindex int
	 	
}

func CreateImgserve() *ImgServe {

	imgsrv := new(ImgServe) 
	imgsrv.imglist = make(map[int] *WebImage)
	imgsrv.lastindex = 0
	 
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
	smsg := rshell.RsMsg{Key:cnum,Msg:[]byte(content)} 
	*imgsrv.msgout <- smsg
}

func (imgsrv *ImgServe) Request(key int ,data []byte){
	log.Println("request for imgsrv")
	result := "" 
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
	result = d_req.ToString()
	
	imgsrv.send(key,result)
	
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



