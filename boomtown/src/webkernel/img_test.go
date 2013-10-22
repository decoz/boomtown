package webkernel

import (

)

type ImgFile struct {
	


}


type ImgServe struct{
	 
	imglist map[string] *ImgFile 	
}


func CreateImgserve() *ImgServe {

	imgsrv := new(ImgServe) 
	imglist = make(map[string] *ImgFile)
	

}


func (imgsrv ImgServe) KernelInfo() string{

	return  "imgServe v0.1" 

} 


func (imgsrv Imgserve) LinkRs( rs *RShell ) {



}


func (imgsrv Imgserve) UnLinkRs(key int) {



}


func (imgsrv Imgserve) Require(string, *RShell) string {




}
