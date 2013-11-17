
function initEventFunction(){
	
	record_write.ondragover = function(){
		this.className = 'hover'
		return false
	}
	
    record_write.ondragend = function(){
    	this.className = ''    		
    	return false 
	}
    
    record_write.ondrop = function(e){
        debugger;
        this.className = ''
        e.preventDefault()
        readfiles(e.dataTransfer.files) 
    }

}

function onClickRecord(){
    if(lastSelected != null) lastSelected.className = ''
    this.className = 'selected'
    lastSelected = this

    var key = this.getAttribute("key") 
    wsocket.doSend("read/"+key)
}


function onMessage(evt, msg_handler){	

   	arr = evt.data.split(":")
    writeToScreen('log: cmd - [' + arr[0] + ']')

    var handler = msg_handler[arr[0]]
    if(handler){
    	handler(arr[1])        	
    }
    else errlog( arr[1])


}


function onOverRecord(){

 	if(this.className != 'selected') this.className = 'mouseover'

}

function onOutRecord(){

 	if(this.className != 'selected') this.className = 'mouseout'

}


function onImgClick(e){
	var key = e.srcElement.getAttribute("key")
	debugger
	if(key in rvimage_list){
		var obj = document.getElementById("cur_image")
		if(obj && obj.getAttribute("key") == key)		
			rv_image.innerHTML = ""
	    else show_rv_image(key,rvimage_list[key])       
	}
	else wsocket.doSend("openimg/"+key)

}


function onButtonWrite(){
	record_write.className = 'show' 
	rv.className = 'hidden'        
}




function onComment(){
    if(cmtinput.value != ""){
     var key = rv.getAttribute("key")
     var id =  userid.innerHTML
     wsocket.doSend("comment/" + key + "/"+ id + "/" +  cmtinput.value ) 

     cmtinput.value = ""
    }
}


function onCommand(){

        var msgarr = msginput.value.split("/")
    var attrs = new Array()

    for(var i=1;i<msgarr.length;i++) attrs[i-1]  = msgarr[i];
    /*
    var msg = msginput.value 
    msginput.value = "" 
    */
    var cmd = msgarr[0];
    var pre = ''
    switch(cmd[0]){
            case '!' :
            case ':' :
            case '/' : 
            pre = cmd[0]
            cmd = cmd.substring(1,cmd.length)
            break;
    }


    switch(pre) {
            case '!':  execCmd(cmd,attrs)
                        break;
            case ':': 
                        break;
            case '/':
                        break;
            case '':
                        execNCmd(cmd,attrs)
                        break; 
            default: 
                        break;
    }


    msginput.value = ''

}



function check_enter(e){
	if(e && e.keyCode== 13) {
        if(event.srcElement.id == "msginput") onCommand()
    
        if(event.srcElement.id == "cmtinput") onComment()
	}	
}

function check_tab(e){    
    if(e && e.keyCode == 27 ) {   
    	if(console_area.style.display == "block" ) console_area.style.display = "none"
    	else console_area.style.display = "block"
    }
}
	
