package webservice

import (
	"log"
	"strings"
	"strconv"	
	"encoding/base64"
	
)

type Item struct {

	id	 	int
	owner 	string
	subject string
	content string
	wcount	int
	
}


type Board struct {	
	lastId int
	Contents map[int] *Item 
	WImages map[int] []*WebImage  
}


func CreateBoard() *Board{
	
	board := new(Board)
	board.initboard()
	return board
}

func (board *Board) initboard(){
	board.Contents = make(map[int] *Item)
	board.WImages = make(map[int] []*WebImage)
	board.lastId = 0
	
} 



func MakeFieldsByExcept(c,s,e rune) func(rune)bool {

	state := true

	return func(r rune)bool{		
		if state {
			switch r{
				case c : return true
				case s : state = false
					return true
				default: return false
			} 
		} else {
			if r == e {
				state = true 
				return true 
			} 		
		}
		return false		
	}
	
	
} 
	

func (board Board) listPage(page,pagesize int) string {
/* 	page 단위로 list 를 불러온다. 
  	page 시작은 0부터 시작하기로 한다.  
*/

	log.Println("board.listPage:",page,",",pagesize)
	if pagesize < 10 { pagesize = 10 }
	si := pagesize * page 
	
	if si >= len(board.Contents) { return "" }
			
	result := ""
	i := board.lastId
		
	for cnt:=0;cnt<si;i--{
		if _,ok := board.Contents[i];ok {
			cnt++ 
		}	
	}  
	
	max := pagesize	
	if ic := len(board.Contents)-si;pagesize > ic{
		max = ic
	}

	
	log.Println("now find ",max,"items")
	
	for cnt:=0;cnt<max;i--{
				
		if item,ok := board.Contents[i];ok {
						
			result +=   strconv.Itoa(item.id) + "/" + item.owner +
			 "/" + item.subject + "/" + strconv.Itoa(item.wcount) + "|"
			
			log.Println(result,"-",cnt)
			
			cnt++
		}			
	}
	
	log.Println("board.listPage:",result)
	return result
}

func (board Board) getRecordString(key int) (string){

	if item,ok := board.Contents[key];ok{	
		result :=   strconv.Itoa(item.id) + "/" + item.owner +
			 "/" + item.subject + "/" + item.content + "/" + 
			 strconv.Itoa(item.wcount)
		return  result 
	} else { 
		return ""
	}
	
			 
	
} 


func (board *Board) Request(key int,request []byte) [][]byte {
	
	req := string(request)
	//cmd := strings.TrimSpace(req)
	cmd := ""
	fields := strings.FieldsFunc(req,MakeFieldsByExcept('/','[',']'))
		
	
	fcount := len(fields)

	if fcount > 0 {
		cmd = strings.TrimSpace(fields[0])
	}
	
	result := ""
	log.Println("board_req:[",cmd,"]",len(req))
	switch cmd{	
	case "list":
		log.Println("board: list item subject -",len(board.Contents))
		result = "list:" 
		/*
		for _,item := range(board.Contents) {
			result +=   strconv.Itoa(item.id) + "/" + item.owner + "/" + item.subject + "|"					
		}
		log.Println("result:",result)
		*/
		
		page,pagesize := 0,10
		var err error   
		if fcount > 1 {
			page,err = strconv.Atoi(fields[1])			
		}
		if fcount > 2 {
			pagesize,err = strconv.Atoi(fields[2])
		}
		
		if err!=nil { 
			result = "error:bad format of list" 
		} else {
			result = "list:" + board.listPage(page,pagesize)
		}	
		
	case "kernel":
		log.Println("board: send kernel info")
		result =  board.KernelInfo()		
		
		
	case "read":
		if fcount < 2 {
			result =  "error:read need more argument"
		}
		key,err := strconv.Atoi(fields[1])
		if err != nil { 
			result =  "error:bad number" 
		} else {
			result = board.ReadItem(key)
			//rs.Ws.Write([]byte(result))
			
			result2 := "%cnt:" + strconv.Itoa(key) +
				"/" + strconv.Itoa(board.Contents[key].wcount)
			
			return [][]byte{ []byte(result), []byte(result2) }
		}
	case "openimg": 
		if fcount == 2 {			
			result = board.ReadImage(fields[1])			
		} else { result =  "error:read need more argument" } 
		
	case "write":
		if fcount < 4 {
			result = "error:not enough argument"		
		}
		if fcount == 5 {
			board.AddImage(fields[4])
		}
		result =  board.AddItem(fields[1],fields[2],fields[3])
		
	case "delete":
		if fcount < 2 {
			result =  "error:read need more argument"
		}
		key,err :=  strconv.Atoi(fields[1])
		if err != nil { 
			result =  "error:bad number" 
		} else {
			result = board.deleteItem(key)
		}
		 	
	default: 
		result = "error:bad command"
	}

	
	return [][]byte{[]byte(result)}
}

func (board Board) KernelInfo() string{
	return "web board kernel  v0.2"
}



func (board *Board) ReadItem( key int) string{
	
	item,ok := board.Contents[key]
	
	if !ok { return "error:item not exist" }
	
	cmd := "open"
	owner := item.owner
	subj := item.subject 
	content := item.content
	cnt :=  strconv.Itoa(item.wcount)
	
	item.wcount += 1
	log.Println(key, "th item count now:" , item.wcount)
	
	result := cmd + ":" +  strconv.Itoa(key) + "/" + owner + "/" +
		subj + "/" + content + "/" + cnt
		
	wilist,ok := board.WImages[key]
	if ok {
		var thumb_str = "["	
		for i,wi := range(wilist) {
			if i != 0 { thumb_str += "," } 
			thumb_str += string(wi.GetThumbE64())		
		}
		thumb_str += "]"
		
		result += "/" + thumb_str		
	}
	

		
	return result
		 
}


func (board *Board) deleteItem( key int) string{
		
	delete(board.Contents,key)
	result := "%delete:" + strconv.Itoa(key) 	
	return result
		
}

func (board *Board) AddItem(isend,isubj,icont string )string{
	
	id := board.lastId
	item := new( Item)
	item.id = id
	item.owner = isend
	item.subject = isubj 
	item.content  = icont
	
	//board.Contents[id] = new( Item{id,isend,isubj,icont,0} )
	board.Contents[id] = item
	board.lastId = board.lastId + 1		
	
	log.Println("last_id:",board.lastId)
	
	result := "%add:" + strconv.Itoa(id) + "/" + isend + "/"  + isubj + "/0"
	return result

}

func (board *Board) AddImage(imgInput string){

	ifields := strings.Split(imgInput,",")
	
	wilist := make([]*WebImage , len(ifields))
	for i,imgstr := range(ifields){
		buff := Decode(nil,[]byte(imgstr),base64.StdEncoding)
		img,_ := ByteToImage(buff)
		wilist[i] = CreateWebImage(img,"")
	}
	
	board.WImages[board.lastId] = wilist	
	log.Println("saved ",len(wilist)," images at " ,board.lastId) 

}

func (board *Board) ReadImage(imgkey string) string {
	arr := strings.Split(imgkey,"-")
	
	rkeys := arr[0]
	ikeys := arr[1]
	rkey,_ := strconv.Atoi(rkeys)
	ikey,_ := strconv.Atoi(ikeys)
	
	result := ""
	
	wlist,ok := board.WImages[rkey]
	if ok {
		wi := wlist[ikey]
		if ok {
			result = "image:"+imgkey+"/[" + string(wi.GetE64()) + "]"		
		}
	}
	
	log.Println("openimg:" + result)

	return result
}

func (board *Board) DummyInit(){

	board.AddItem("kwak","good day",
		"오늘 날씨 꽤 괜찬흔 것 같네요 좋은 날이 계속 되기를" +
		"\n 바람도 좋은 것 같고 " )
		
	board.AddItem("jung","일단 테스트 게시물입니다.",
		"짧게 쓸게요.. 귀찮으니까.. 한 3-4개만 적어두면 될듯?" +
		"\n 그나저나 댓글 기능도 넣어야 하는데.. 후... " +
		"\n\n 언젠간 되겠지.. ")
		
	board.AddItem("yoo","이것도 이니셜 게시물",
		"아직 db에 연결 안된 관계로 기능 테스트용 이니셜 게시물" +
		"\n 라이브웹? 다이나믹웹? 여튼 그런 웹페이지를 구상중 " +
		"\n 우선 간단한 게시판으로 데모중입니다. ")
		
	log.Println("dummy contents ",len(board.Contents), " added")
}
