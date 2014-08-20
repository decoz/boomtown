package webservice

import (
	"encoding/base64"	
	"image"
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