
     
function createWebsocket(wsUrl) { 

	this.ws = new WebSocket(wsUrl); 
	this.msg_handler = new Array()

	this.ws.onopen = onOpen
	this.ws.onclose = onClose
	this.ws.onmessage =  onMessage
	this.ws.onerror =  onError

	this.printmsg = null


	this.doSend = function(message) {
		if(message.length > 100)  writeToScreen("SENT: " + message.length + " byte" );
		else writeToScreen("SENT: " + message);  
		this.ws.send(message); 
	}  

	function onMessage(evt){

	        selectHandler(evt.data)
	

	}; 


	function selectHandler(msg){

	        debugger 
		var arr = msg.split(":")
	        writeToScreen('log: cmd - [' + arr[0] + ']')
	        if(this.msg_handler[arr[0]] != undefined){

		        var handler = this.msg_handler[arr[0]]
	        	handler(arr[1])        	

	        }
	        else errlog(  "no handler for " + arr[1])				
	}

	function writeToScreen(msg){
		if( this.printmsg ) this.printmsg(msg)
	}


	function onOpen(evt) { 
		writeToScreen("CONNECTED"); 
		doSend("list/0/"+psize); 
	}  

	function onClose(evt) { 
		writeToScreen("DISCONNECTED"); 
	}  


	function onError(evt) { 
		writeToScreen('<span style="color: red;">ERROR:</span> ' + evt.data); 
	}  


	return this

}  


