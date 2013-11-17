
function ImgList(){
	this.list = new Array()
	
	this.add = function(key,img){
		list[key] = img		
	}
		
	this.get = function(key){
		return list[key]		
	}
	
	
	
}


function ImgFullViewer(){
	this.obj = document.createElement("div")
	this.obj.id = "imgview"
	this.image = null
	
	this.open = function(img){
		debugger
		//this.obj.style.width = img.naturalWidth + "px"
		
				
		if(img != null){		
			this.obj.innerHTML = ""
			this.obj.appendChild(img)
		}
		
		img.style.position = "absolute"
		img.style.width = img.naturalWidth + "px"
		img.style.top =  window.innerHeight/2 - img.height/2 + "px"
		img.style.left = document.body.offsetWidth/2 - img.width/2 + "px"
		
		var first = document.body.childNodes[0]
		document.body.insertBefore(this.obj,first)
	}
	
	var me = this
	this.e_click = function(e){
		me.close.apply(me,[e])		
	}
	
	this.obj.onclick =  this.e_click
	
	this.close  = function(){
		document.body.removeChild(this.obj)		
	}
	
}