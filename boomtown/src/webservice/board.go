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

type Comment struct {		
	owner	string
	cmt  	string	
}

type Client struct {
	cur_view 	int
	cur_page  	int
	cur_psize 	int 
	id 			string		
}


type L_Cmt []Comment
type M_L_Cmt map[int] L_Cmt
type M_Item map[int] *Item
type L_WImage []*WebImage
type M_L_WImage map[int] L_WImage  

type Board struct {	
	lastId int
	Contents M_Item 
	WImages M_L_WImage
	Comments M_L_Cmt
	Clients	map[int] *Client	
	Watching map[int] int
}


func CreateBoard() *Board{
	
	board := new(Board)
	board.initboard()
	return board
}

func (board *Board) initboard(){
	board.Contents = make(map[int] *Item)
	board.WImages = make(map[int] L_WImage)
	board.Comments = make(M_L_Cmt)
	board.Clients = make(map[int] *Client)
	board.Watching = make(map[int] int)
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
			if cnt != 0 { result += "|" } 			
			result +=   strconv.Itoa(item.id) + "/" + item.owner +
			 "/" + item.subject + "/" + board.ItemInfo(item.id)
			
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


func (board *Board) ItemInfo(key int) string{
/* cmt/view[watching]/imgcnt */
	
	view := strconv.Itoa(board.Contents[key].wcount)
	
	watching := ""
	if w,ok := board.Watching[key];ok && w > 0{	
		watching = "[" + strconv.Itoa(w) + "]"
	}
	
	cmt := ""
	if arr,ok := board.Comments[key];ok{
		cmt = "[" + strconv.Itoa(len(arr)) + "]"
	}
	
	imgcnt := ""
	if arr,ok := board.WImages[key];ok{
		imgcnt = strconv.Itoa(len(arr))
	}
 	
	return  cmt + "/" + view +	watching +	"/" + imgcnt
	
	
}


func (board *Board) Request(cnum int,request []byte) [][]byte {
	
	
	if _,ok := board.Clients[cnum];!ok{
		board.Clients[cnum] = &Client{-1,-1,-1,""}			
	} 
	
	client := board.Clients[cnum]
		
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
			result = "error:bad argument for list" 
		} else {
			client.cur_page = page
			client.cur_psize = pagesize 
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
		
			last := client.cur_view			
			client.cur_view = key	
				
			
			if _,ok := board.Watching[key];!ok {
				board.Watching[key] = 0
			}
			board.Watching[key] += 1			
			if last != -1 { board.Watching[last] -= 1 }
						
									
			result = board.ReadItem(key)
			//rs.Ws.Write([]byte(result))
			/*
			result2 := "%cnt:" + strconv.Itoa(key) +
				"/" + strconv.Itoa(board.Contents[key].wcount)
			*/
			
			
			result2 := "%update:" + strconv.Itoa(key) + "/" + board.ItemInfo(key)
						
			if last != -1 {
				result3 := "%update:" + strconv.Itoa(last) + "/" + board.ItemInfo(last)
				return [][]byte{ []byte(result), []byte(result2) , []byte(result3)}
			}
				 
			
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
	case "comment":
		if fcount == 4 {
			log.Println("comment message",fields)
			key,_ := strconv.Atoi(fields[1])
			
			result = board.addComment(key,fields[2],fields[3])
			result2 := "%update:" + strconv.Itoa(key) + "/" + board.ItemInfo(key)
			
			return [][]byte{ []byte(result), []byte(result2) }
			
			
		}		
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
	
	case "saveboard":
		fname := "default.dot"
		if fcount >  1 {
			fname = fields[1]+".dot"
		}
		
		board.SaveBoard(fname)  
		result = "msg:board save complete"	
		
	case "loadboard":
		fname := "default.dot"
		if fcount >  1 {
			fname = fields[1]+".dot"
		}
						
		board.LoadBoard(fname)  
		result1 := "msg:board load complete"		
		result2 := "%list:" + board.listPage(client.cur_page,client.cur_psize)
		
		log.Println("----loaded ",fname,"----")
		
		return [][]byte{[]byte(result1),[]byte(result2)}
		
	case "clearboard":
		board.clearboard()  
		result1 := "msg:board is cleared"		
		result2 := "%list:" + board.listPage(client.cur_page,client.cur_psize)
		return [][]byte{[]byte(result1),[]byte(result2)}
		 	
	default: 
		result = "error:bad command"
	}

	
	return [][]byte{[]byte(result)}
}

func (board Board) addComment(key int, owner, cont string ) string {
	
	comments,ok := board.Comments[key]
	
	if !ok {
		comments = make([]Comment,0,50)				
	}				
	board.Comments[key] = append(comments, Comment{owner,cont})  
	
	list := ""
	for i,client := range(board.Clients) {
		if client.cur_view == key {
			if len(list)  > 0 { list += "," } 
			list += strconv.Itoa(i)				
		}		 
	}
	
	return "[" + list + "]" + "addcmt:" + strconv.Itoa(key) + "/" + owner + "|" + cont		
}




func (board Board) KernelInfo() string{
	return "web board kernel  v0.51"
}


func (board *Board) ReadComments( key int) string{

	result := ""	
	if comments,ok := board.Comments[key];ok{		
		for _,cmt := range(comments){			
			if len(result) > 0 { result += ","	} 			
			result += cmt.owner  + "|" + cmt.cmt		
		}						
	}
	
	log.Println("read comments for ", key , ":" , result)
	return result
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
		subj + "/" + content + "/" + cnt + "/" + board.ReadComments(key)
		
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

func (board *Board) clearboard() {
	log.Println("board cleared by request")

	board.Comments = make(map[int]L_Cmt) 
	board.Contents = make(map[int] *Item)
	board.WImages = make(map[int] L_WImage)
	
	board.lastId = 0
	
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
	
	result := "%add:" + strconv.Itoa(id) + "/" + isend + "/"  + isubj + "//0/"
	return result

}

func (board *Board) AddImage(imgInput string){

	ifields := strings.Split(imgInput,",")
	
	wilist := make([]*WebImage , len(ifields))
	for i,imgstr := range(ifields){
		buff := Decode(nil,[]byte(imgstr),base64.StdEncoding)
		img,_ := ByteToImage(buff)
		wilist[i] = CreateWebImage(img,"")
		log.Println("create wi complete")
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
	
	//log.Println("openimg:" + result)

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


func (board *Board) Unlink(key int) {
	cl := board.Clients[key]
	board.Watching[cl.cur_view]-- 
	delete(board.Clients,key)

}

