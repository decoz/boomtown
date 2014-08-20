package tag

import (
	

)



type Object struct {
	code int
	content string
	tags []int 
	children []int
}


type tagDB interface {
/*
 tag  , tag Index , item 의 저장을 담당하는 인터페이스 
 databse 이든 filesystem 이든 분산 객체이든 무관 
	 


	getTagByPath(string) Tag	
	getTag(int) 
*/
	addTag(string,Tag)	
} 




type objectDB interface {

/*
object 는 tag 값을 가진 객체 

*/


}






