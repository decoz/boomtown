
     
function BoomSocket(wsUrl) { 

	this.ws = new WebSocket(wsUrl); 
	this.msg_handler = new Array()

	
	this.printmsg = null
	this.e_open = null
	
	this.handle = function(evt,handler){
		//debugger
		var obj = this
		var s_handler = function(e){
			handler.apply(obj,[e])			
		} 	
		this.ws.addEventListener(evt,s_handler,false)
		
	}

		
	this.onMessage = function(evt){
       		 this.selectHandler(evt.data)
	}; 
	

	this.writeToScreen = function(msg){
		if( this.printmsg ) this.printmsg(msg)
	}


	this.onOpen = function(evt){ 
		this.writeToScreen("CONNECTED"); 
		if (this.e_open) this.e_open(evt)
		
	}  

	this.onClose = function(evt) { 
		this.writeToScreen("DISCONNECTED"); 
	}  

 
	this.onError = function(evt){ 
		writeToScreen('<span style="color: red;">ERROR:</span> ' + evt.data); 
	}  

	this.handle("open",this.onOpen)
	this.handle("close",this.onClose)
	this.handle("message",this.onMessage)
	this.handle("error",this.onError)
	
	
	return this

}  

BoomSocket.prototype.selectHandler = function(msg){	
	var arr = msg.split(":")
    	this.writeToScreen('log: cmd - [' + arr[0] + ']')
    	if(this.msg_handler[arr[0]] != undefined){

       		var handler = this.msg_handler[arr[0]]
    		handler(arr[1])        	

    	}
    	else errlog(  "no handler for " + arr[1])	
	
}

BoomSocket.prototype.doSend = function(message) {
	if(message.length > 100)  this.writeToScreen("SENT: " + message.length + " byte" );
	else writeToScreen("SENT-: " + message);  
	this.ws.send(message); 
}



