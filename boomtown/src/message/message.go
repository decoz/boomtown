package message

import (

)

/*
	data parse 컨셉
	
	cmd string
	{
		"insert": 
			name string 
			contents string
		"file":
			name string
			data byte
		"list" : 
			target[,] string			
	
	}
			 	  
	
	배열 예제			  		 
			
	target[int] string
	target[string] 10
	target[byte] byte	
		
	
	


*/

type Message struct {
	data []byte
	contents map[string] struct{ start int;size int }
}

type Parser  struct {
		
}






