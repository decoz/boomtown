function execNCmd(cmd,attrs){
    switch(cmd){

        case "next" : 
            page ++ 
            //alert("list/"+page)
            wsocket.doSend ("list/"  + page + "/" + psize)
            pagenum.innerHTML = "[" + (page+1) + "]"
            break;
        case "prev" : 
            page -- 
            if(page < 0){ page = 0; }
            else {
                wsocket.doSend ("list/"  + page + "/" + psize )
                pagenum.innerHTML = "[" + (page+1) + "]"
            }
            break;
        case "delete":
                if(attrs.length != 1) error("parameter error for delete")
                else wsocket.doSend ("delete/"+attrs[0])
                break;
        case "saveboard":
            	if(attrs.length != 1) wsocket.doSend("saveboard")
            	else wsocket.doSend ("saveboard/"+attrs[0])
            	break;
        case "loadboard":
        	if(attrs.length != 1) wsocket.doSend("loadboard")
        	else wsocket.doSend ("loadboard/"+attrs[0])
            	
        default:
           // wsocket.doSend (cmd)
    }
}


    function execCmd(cmd,attrs){
            log('innerCmd:' + cmd)

            switch(cmd){
                case "startbot":
                    myInterval = setInterval(write_random,5000)
                    break;
                case "stopbot":
                    clearInterval(myInterval)
                    break;
                case "iam":
                    log("attr_cnt "+attrs[0])
                    if(attrs.length == 1 )  {
                        uid = attrs[0]
                        userid.innerHTML = uid
                    }
                    break; 
            }

    }


