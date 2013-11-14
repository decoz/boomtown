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
	return newDot(Enc(str))
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
	dot.value = Enc(str)
	
}

func (dot *Dot) GetValue() string {
	return Dec(dot.value)
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

func (dot *Dot) SetKChild(k string, val string) *Dot{
/*
	k key 값의 노드를 찾아 그 아래 벨류값을 변경한다. 
	만일 k 노드나 value 노드가 없을 경우 생성한다.
	
	수행후 k 노드를 리턴한다.
	
*/

	d := dot.GetKChild(k)
	
	if d == nil { 
		d = NewDot(k)
		dot.AppendChild(d) 
	}
	
	if len(d.child) > 0 {d.child[0].SetValue(val)
	} else  {
		d.AppendChild( NewDot(val) )
	}
			
	return d
}


func Enc(str string) []byte{

	src := []byte(str) 
	dst := make([]byte, len(str) + 10 )[0:0]
	
	for _,c := range src {
		switch c {
		case '.' : dst = append(dst,'&','d') 
		case ',' : dst = append(dst,'&','c')
		case '(' : dst = append(dst,'&','s')
		case ')' : dst = append(dst,'&','e')
		case '&' : dst = append(dst,'&','&')
		default : dst = append(dst,c)
		}
	}
	
	return dst	 			
}

func Dec(src []byte) string{
	 
	dst := make([]byte, len(src))[0:0]
	
	flag := false
	for _,c := range src {
		if flag {
			switch c {
			case 'd' : dst = append(dst,'.')
			case 'c' : dst = append(dst,',')
			case 's' : dst = append(dst,'(')
			case 'e' : dst = append(dst,')')
			case '&' : dst = append(dst,'&')
			}
			flag = false 
		} else {
			if c == '&' { flag = true 
			} else {
				dst = append(dst,c)
			}
		}	
	
	}
	
	return string(dst)	 			
}





