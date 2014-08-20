
     
function BoomSocket(wsUrl) { 

	this.ws = new WebSocket(wsUrl); 
	this.msg_handler = new Array()
	
	this.printmsg = null
	this.ondisconnect = null
	this.onmessage = null
	this.onconnect = null
	
	this.handle = function(evt,handler){		
		var obj = this
		var s_handler = function(e){
			handler.apply(obj,[e])			
		} 	
		this.ws.addEventListener(evt,s_handler,false)		
	}

		
	this.e_message = function(evt){
		if(this.onmessage) this.onmessage(evt)
	}; 

	this.e_connect = function(evt){
		this.writeToScreen("CONNECTED"); 
		if(this.onconnect )  this.onconnect(evt)
	}
	
	this.e_disconnect = function(evt) { 
		this.writeToScreen("DISCONNECTED"); 
		if(this.ondisconnect) this.ondisconnect(evt)
	}  

	this.writeToScreen = function(msg){
		if( this.printmsg ) this.printmsg(msg)
	}


	this.onOpen = function(evt){ 
		this.writeToScreen("CONNECTED"); 
		if (this.e_open) this.e_open(evt)
		
	}  

	
 
	this.onError = function(evt){ 
		writeToScreen('<span style="color: red;">ERROR:</span> ' + evt.data); 
	}  

	this.handle("open",this.e_connect)
	this.handle("close",this.e_disconnect)
	this.handle("message",this.e_message)
		
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



