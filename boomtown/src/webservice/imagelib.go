package webservice

import (
	"image"
	"image/png"
	"image/jpeg"	
	"bytes"	
	"github.com/nfnt/resize"
	"log"
	"os"
)

func ByteToImage(buff []byte) (image.Image,error){

	reader := bytes.NewReader(buff)
	img,imgtype,err := image.Decode(reader)
	log.Println("decoded size:",len(buff), " type:",imgtype)
		
	if err != nil {
		 log.Println("auto file decode failure")
		 img,err = png.Decode(reader)
		 
	 }
	
	if err != nil { 
		log.Println("png decode fail")
		img,err = jpeg.Decode(reader)
	}
	
	if err != nil {
		log.Println("jpeg decode fail")
	}
	
	//if err!=nil { panic("bad file format") }
	
	if err == nil {
		log.Println(imgtype, " decoding succeded")		 
	
	}
	
	
	return img,err
} 

func ImageToByte(img image.Image,imgtype string) []byte{
	
	
	//buff := make([]byte ,1024*1024)
	//buffer := bytes.NewBuffer(buff)
	out := new(bytes.Buffer)
	
	
	//log.Println("img to byte:",buffer.Len())
	
		
	switch imgtype {	
	case "jpg","jpeg":	
		jpeg.Encode(out,img,nil)
	case "png":
		png.Encode(out,img)
	default:
		jpeg.Encode(out,img,nil)
	}	

	log.Println("img to byte:",out.Len())
	
	return out.Bytes()
	
}


func CreateThumbnail(img image.Image,width,height uint) image.Image{
		
	//thumb := resize.Resize(width,height,img,resize.Lanczos3)
	thumb := resize.Resize(width,height,img,resize.Bilinear)
	return thumb
}


func SaveImage(img image.Image, fname,imgtype string){

	out, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}	
	defer out.Close()
	
	switch imgtype {
	
	case "jpg","jpeg":	
		jpeg.Encode(out,img,nil)
	case "png":
		png.Encode(out,img)
	default:
		jpeg.Encode(out,img,nil)
	}	
	 
}



