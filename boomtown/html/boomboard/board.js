
/*********************** record _wrtie code **********************/

    function send_write(){
       var id =  userid.innerHTML
       var subj = rw_subj.value
       var cont = rw_cont.value
       var cmd = 'write/'+id+'/'+subj+'/'+cont

       var first = true;
       var imgdata  = ""

        for( var i=0;i<lastIndex;i++){
            img = wimage_list[i];
            if( i > 0 ) {imgdata += ","}
            debugger;
            imgdata += imageToMessage(img)
        }

        if( imgdata != "") cmd += "/[" + imgdata + "]"

       if(subj != '' && cont != '') doSend(cmd)

        rw_subj.value = ""
        rw_cont.value = ""
        clearWImage()
    }

function imageToMessage(img){

    var e64 = img.src.split(",")[1]
    return e64

}

function clearWImage (){
        for( var i=0;i<lastIndex;i++){
            rw_imgview.removeChild(wimage_list[i])
            //lete(wimage_list[i])
        }
        lastIndex = 0
}

/********************** record _write hoover code **********************/


var wimage_list = new Array()
var rvimage_list = new Array()

var lastIndex = 0


function readfiles(files) {
    for (var i = 0; i < files.length; i++) {      
      previewfile(files[i]);
    }
}

function previewfile(file) {

    var reader = new FileReader();

    reader.onload = function (event) {
        var image = new Image();

        image.src = event.target.result;
        if(image.width > image.height * 2) {
        image.width = 100; // a fake resize
        }
        else image.height = 50

        wimage_list[lastIndex] = image
        lastIndex++
        rw_imgview.appendChild(image);

    }

    reader.readAsDataURL(file);
}

/********************** record _write hoover code **********************/


function setAnimateClass(element, cname, end_cname ,endtime){
    if(endtime == null) endtime = 1000
    if ( element != null) {
        var backup = element.className
        element.className = cname

       if(end_cname != null) setTimeout( resetClass,endtime,element,end_cname)
    } 

}

function cntRecord(key_cnt){
    var arr = key_cnt.split("/")
    var key = arr[0]
    var cnt = arr[1]
    log(key + " : " + cnt)
    var obj = document.getElementById("cnt_" + key )
    if(obj != null) {
        obj.innerHTML = cnt
        setAnimateClass(obj,'watching','',4000)
    }

}

function resetClass(element, cname){

    element.className = cname    

}




function addComment(str){

//rv_comment.innerHTML += str + "<br>"
var flds =  str.split("|")
if(flds.length == 2){
    var p= document.createElement("p")
    var owner= document.createElement("span")
    var content = document.createElement("span")

    p.className = "comment"
    owner.className = "cmt_owner"
    owner.innerHTML =  flds[0]
    content.className = "cmt_cont"
    content.innerHTML =  flds[1]
    p.appendChild(owner)
    p.appendChild(content)

    rv_comment.appendChild(p)

    return p 
}
else return null


}



function show_rv_image(data){
    var image = new Image()
    image.src = "data:image/jpeg;base64," + data
    image.id = "cur_image"

    image.onclick = function(event){
       rv_image.removeChild(event.srcElement ) 
    }
    rv_image.innerHTML = ""
    rv_image.appendChild (image)
    setAnimateClass(image,"new_img")

}


function clearTable(){
    for(var i=list.childElementCount -1;i>=0;i--){
        list.removeChild(list.childNodes[0])
    }
}



lastSelected = null 

function onClickRecord(){
    if(lastSelected != null) lastSelected.className = ''
    this.className = 'selected'
    lastSelected = this

    var key = this.getAttribute("key") 
    doSend("read/"+key)
}


function insertRecord(record,cname,pos){

    fields = record.split("/")
    log(record)
    //e_tr = document.createElement("tr")
    if(fields.length > 5) {
        var e_tr = list.insertRow(pos)
        var key = fields[0]
        e_tr.setAttribute("key",fields[0])
     

        var e_td = document.createElement("td")
        e_td.id = "no_" + key 
        e_td.innerHTML = fields[0]
        e_tr.appendChild(e_td)

        e_td = document.createElement("td")
        e_td.id = "sub_" + key 
        e_td.innerHTML = fields[2]
        e_tr.appendChild(e_td)

        var ccnt = document.createElement("span")
        ccnt.id = "ccnt_" + key
        ccnt.className = "ccnt"
        ccnt.innerHTML = fields[3]
        e_td.appendChild(ccnt)


        e_td = document.createElement("td")
        e_td.id = "owner_" + key 
        e_td.innerHTML = fields[1]
        e_tr.appendChild(e_td)

        e_td = document.createElement("td")
        e_td.id = "cnt_" + key 
        e_td.innerHTML = fields[4]
        e_tr.appendChild(e_td)

        e_tr.onclick = onClickRecord
        e_tr.onmouseover = onOverRecord
        e_tr.onmouseout = onOutRecord

        e_tr.className = cname 
    


    }

}




function addRecord(record){

    //log("{"+record+"}")
   // list.className = 'scroll' 
   //var psize = 10 

   if(page == 0 ) {
       var rcount  = list.childNodes.length 
       

       setAnimateClass(list,'scroll','')
        insertRecord(record,'record_new')
        if( rcount > psize ) list.deleteRow(psize)
    }
}

function deleteRecord(key){
    for(var i =0;i<list.childNodes.length;i++){
        var item = list.childNodes[i]
        if(item.getAttribute("key") == key)  {
            list.deleteRow(i)
            break;
        }
    }

    doSend("list/"+page+"/"+psize)
}


function deleteRows(start, delcount) {
    for(var i=0;i<delcount;i++) list.deleteRow(start)
}

function log(msg){
    writeToScreen('log:'  + msg)
}

function prevpage(){
      page -- 
        if(page < 0){ page = 0; }
        else {
            doSend("list/"  + page + "/" + psize )
            pagenum.innerHTML = "[" + (page+1) + "]"
        }
}


function nextpage(){

     page ++ 
    doSend("list/"  + page  + "/" + psize)
    pagenum.innerHTML = "[" + (page+1) + "]"

}









function writeToScreen(message) { 

    var pre = document.createElement("p"); 
    pre.style.wordWrap = "break-word"; 
    pre.innerHTML = message; output.appendChild(pre); 

    output.scrollTop = output.scrollHeight
}  

