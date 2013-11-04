/*
need (
	rv
		.rv_imgs
		.rv_image
		.rv_comment
		 rv_subj

	list

	record_write


	rvimage_list[]

	show_rv_iamge()
	setAnimateClass()
	addComment()

	wsocket.doSend()
	
)
*/
	



function h_add(record){

       if(page == 0 ) {
           var rcount  = list.childNodes.length 
           

           setAnimateClass(list,'scroll','')
            insertRecord(record,'record_new')
            if( rcount > psize ) list.deleteRow(psize)
        }
}


function h_error(msg){
        writeToScreen('<span style="color: red;">error:' + msg + '</span> '); 
}


function h_addcmt(str){
        var trr = str.split("/")
        if( rv.getAttribute("key") == trr[0]) {
            var obj = addComment(trr[1])
            if(obj) setAnimateClass(obj,"comment_new","comment")
        }

}

function  h_image(imgstr){

    var arr = imgstr.split("/[")
    var key = arr[0]
    var  data = arr[1].substr(0,arr[1].length -1)
    rvimage_list [key] = data
    show_rv_image(data)

}

function h_delete(key){
      for(var i =0;i<list.childNodes.length;i++){
          var item = list.childNodes[i]
          if(item.getAttribute("key") == key)  {
              list.deleteRow(i)
              break;
          }
	}
	  wsocket.doSend("list/"+page+"/"+psize)
}

 function h_open(record){

    //record_view.innerHTML = record
    rv_imgs.innerHTML = ""
    rv_image.innerHTML = ""
    rv_comment.innerHTML = ""
    var arr  = record.split("/[") 
  

    var fields = arr[0].split("/")
    var key = fields[0]
    if(fields.length > 2) {            
        rv.setAttribute("key",key)
        rv_subj.innerHTML =  fields[2] 
        rv_owner.innerHTML =  fields[1] 
        rv_cont.innerHTML =  fields[3] 

        if(fields[4] != "" ) {
            var cmts =  fields[5].split(",")
            for(i in cmts ){
                addComment(cmts[i])
            }
        }

    }

    if( arr.length > 1) {

        var str = arr[1].substr(0,arr[1].length-1)
        var imgarr = str.split(",")
        for(var i  =0;i < imgarr.length;i++ ) {
            var imgstr = imgarr[i]
           var image = new Image();
           image.id = "img" + key + "-"+i
           image.setAttribute("key",key + "-"+i)
           image.onclick  = onImgClick
           image.src = "data:image/jpeg;base64," + imgstr
            rv_imgs.appendChild(image)
        }
     
    }
    record_write.className = 'hidden'
    setAnimateClass(rv,'show','')

}




function h_list(data){
   
    clearTable() 
    var cnt = list.childNodes.length 
    records = data.split("|")

    for(var i=0; i<records.length;i++ ) {
         record = records[records.length - i - 1]
         insertRecord(record,'',0)

    }

    /* 
    setAnimateClass(list,'list_down','')
    deleteRows(0,cnt-1)
    list.style.top = '0'
    */

}



function h_update(info){

    var arr = info.split("/")
    var key = arr[0]
    var cmt = arr[1]
    var cnt = arr[2]

    log("updateCnt:" + key + "," + cmt + "," + cnt)
    var obj = document.getElementById("cnt_" + key )
    if(obj != null) {
        var last  =  obj.innerHTML
        obj.innerHTML = cnt
        if( last != cnt)  setAnimateClass(obj,'watching','',4000)
    }                
    obj = document.getElementById("ccnt_" + key )
    if(obj != null) {
        var last  =  obj.innerHTML
        obj.innerHTML =   cmt 
        if(last != cmt) setAnimateClass(obj,'watching','ccnt',4000)
    }                

}