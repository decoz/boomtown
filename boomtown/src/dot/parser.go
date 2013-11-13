package dot

import (
	//"log"
)




type Parser struct{
	pos int	
	data []byte
	symbol []byte
	err chan error	
}


const  (
	nt_value = 1 << iota
	nt_map
	nt_list
)



func CreateParser(idata []byte) *Parser{
	
	parser := new(Parser)
	parser.data = idata	
	parser.err = make(chan error)
	parser.symbol = []byte{'.', '(', ')', ','}
	
	return parser 

}

func (parser *Parser) SetData(input []byte){
	parser.data = input 
	parser.pos = 0
}

func (parser *Parser) Parse(name []byte, depth int) *Dot{
	
	n := new(Dot)
	n.value = name
	n.child = make([]*Dot,5)[0:0] 
	
	var (
		token []byte
		symbol byte
	)
	
	token,symbol = parser.next()
	
				
	for parser.pos < len(parser.data) {
	
		//log.Print("[",parser.pos,":",string(parser.data[parser.pos]),"]")
		
		switch symbol {
		
		case ',' : 
					//if len(token) > 0 { n.add(newDot(token)) }
					n.add(newDot(token))		
					if depth > 0 {						
						//log.Println("back with",string(token))						
						return n						 
					} else {						
						parser.pos++
					} 
					
					
		case '(' :  parser.pos++
					n.add(parser.Parse(token,0))   
			
		case ')' :  n.add(newDot(token)) 
					if depth > 0 {											
						return n 		
					} else {		
						parser.pos++				
						return n
					}					
		case '.' : parser.pos++
				   n.add(parser.Parse(token,depth+1))
					
		}		
		
		token,symbol = parser.next()
	}
	
	if len(token) > 0 { n.add( newDot(token) ) }
	return n
}

func (parser *Parser) next() ([]byte,byte){
/*
	특별 심볼까지 파싱을 진행한다. 
*/
	buff := make([]byte,1000)[0:0]
	for parser.pos < len(parser.data) {
		c := parser.data[parser.pos]		
		switch c {
			case '.', '(', ')', ',':
				//log.Println(string(buff),":",len(buff),"<-",string(c))
				return buff,c			
			default:
				buff = append(buff,c)		
		}	
		parser.pos++
	}
		
	return buff,0
}
