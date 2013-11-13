package dot

import (
	//"log"
)


type Dots []*Dot 

type Dot struct{	
	value []byte
	child Dots	
}

func NewDot(str string) *Dot{
	return newDot([]byte(str))
}

func newDot(value []byte) *Dot{
	n := new( Dot )
	n.value = value	  
	return n
}

func CreateDot(str string)*Dot {
	p := CreateParser([]byte(str))
	
	dot := p.Parse([]byte(""),0)
	if len(dot.child) == 1 { return dot.child[0] 
	} else { return dot }
	
}


func (n *Dot) add(c *Dot) {
	
	if n.child == nil { n.child =  make([]*Dot,5)[0:0] }
	if len(c.child) > 0 || len(c.value) > 0 {
		//log.Println("add ",string(c.value),"with", len(c.child),"children")
		n.child = append(n.child,c)
	}
} 

func (dot Dot) ToString() string{
	
	return dot.str()
	
}

func (dot *Dot) str() string{

	str := ""
	ccnt := len(dot.child) 
		
	if dot.value != nil {
		str += string(dot.value) 	
	}
	
	switch  {
	
	case ccnt == 0:
		
	
	case ccnt == 1:		
		str += "." + dot.child[0].str()
					
	case ccnt > 1: 
		
		str +=  "("
		
		isfirst := true	
		
		for _,child := range(dot.child) {
			if isfirst { isfirst = false
			} else { str += "," }
			
			str += child.str()	
		}		
		
		str += ")"	
	}


	return str

}

func (dot *Dot) GetChild(i int) *Dot{
	
	if i < len(dot.child) { return dot.child[i]
	} else { return nil }  
}


func (dot *Dot) GetChildList() []*Dot{
	return dot.child

}

func (dot *Dot) AppendChild(children... *Dot){
	for _,child := range(children) {  
		dot.add(child)
	}
}

func (dot *Dot) SetValue(str string){
	dot.value = []byte(str)
	
}

func (dot *Dot) GetValue() string {
	return string(dot.value)
}

func (dot *Dot) GetCValue() string {
/*
	node 의 child 벨류 값을 리턴한다. 	현재는 퍼스트 노드의 값을 리턴하되 
	차후에 하위 노드중 터미널 노드 의 값을 종합해서 리턴하는 것으로 업그레이드예정  
*/

	c := dot.GetChild(0)
	
	if c != nil { return c.GetValue() 
	} else { return  "" }
	
	
}




func (dot *Dot) GetKChild(k string) *Dot{	
	for _,d := range dot.child {
		if string(d.value) == k { return d }		
	}	
	return nil
}


