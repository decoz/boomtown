package tag

import (
	"dot"
	"fmt"
	"strconv"
)


type Tag struct{
	id int
	value string	
	parent *Tag	
	children []*Tag
}


type DotTagDB struct {

	root *Tag;
	lastid int;
		
	/*
	AddTag(string,Tag)
	GetTagByPath(string) Tag	
	GetTag(int)
	*/	

} 

func CreateDotTagDB() *DotTagDB {

	tdb := new( DotTagDB )
	tdb.root = tdb.NewTag("")
	tdb.lastid = 0
	
	return tdb
		
}

func (tdb *DotTagDB) getNewId() int {
	rv := tdb.lastid	
	tdb.lastid += 1
	return rv
}

 
func (tdb *DotTagDB) NewTag(str string) *Tag {

	ntag := new( Tag )
	ntag.value = str	
	ntag.id = tdb.getNewId()	
	return ntag

}

func (tdb *DotTagDB) AddChildK(tag *Tag, value string) *Tag {

	if tag.children == nil { tag.children = make([]*Tag,5)[0:0] }
	child  := tdb.NewTag(value)	
	tag.children = append(tag.children, child)		
		
	return child
}

func (tdb *DotTagDB) AddTag( tagstr string) {

	tdb.attach(tdb.root, dot.CreateDot(tagstr) )

}


func (tag *Tag) GetChild(value string) *Tag {
	
	for _,child := range(tag.children) {
		if(child.value == value) { return child }   	
	}
	return nil
}

func (tag *Tag) ToStr() string {
	str := tag.value + "[" + strconv.Itoa(tag.id) + "/" +  strconv.Itoa(len(tag.children)) + "]"		
	return str
}

func (tag *Tag) Prints(depth int) {

	for i:=0;i<depth;i++ {fmt.Print("\t")}
	fmt.Println(tag.ToStr())
	for _,child := range(tag.children) {
		child.Prints(depth + 1) 	
	}		

}


func (tdb *DotTagDB) attach(tag *Tag,path *dot.Dot) {
	
	r := tag.GetChild(path.GetValue())	
	if r == nil {
		r = tdb.AddChildK(tag,path.GetValue())		
	} 
	
	for _,child := range( path.GetChildList()) {
		tdb.attach(r,child)	
	}
		
}

func (tdb *DotTagDB) PrintAll(){

	tdb.root.Prints(0)


}

func (tdb *DotTagDB) GetTagID(path string){
	
				

}



