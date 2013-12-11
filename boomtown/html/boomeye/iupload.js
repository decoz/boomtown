function iupload(obj){	
	//var me = this

	this.obj = obj	
	this.files = new Array()
	
	this.e_dragstart = function(e){
		//  e.dataTransfer.setData('text/html', null); 
	}

	this.e_dragover = function(e){			
		e.preventDefault()		
		this.obj.className = 'hover'
		
		return false				
	}	
	
	this.e_dragend = function(e){		
		this.obj.className = '';
		this.obj.style.background = 'white'
		return false
	}

	this.e_drop = function(e){					
		e.preventDefault()	
		this.readfiles(e.dataTransfer.files)	
		this.obj.className = ''
	}

	this.handle = function(evt,handler){		
		var obj = this
		var t_handler = function(e){									
			handler.apply(obj,[e])
		}
		this.obj.addEventListener(evt,t_handler,false)	

	} 


	this.handle("dragstart",this.e_dragstart)
	this.handle("dragover",this.e_dragover)
	this.handle("dragout",this.e_drop)
	this.handle("drop",this.e_drop)	


	this.onfile = null

}



iupload.prototype.readfiles = function(files){
 	for (var i = 0; i < files.length; i++) {
 		this.loadfiles(files[i]);
	}
}


iupload.prototype.loadfiles = function(file){
	var reader = new FileReader()
	var me = this
	
	//this.obj.innerHTML += "<br>" +   file.name


	reader.onload =function(e) {
		var cnt = me.files.length
		me.files[cnt] = e.target.result
		/*
		var img = new Image()
		img.src = me.files[cnt]
		img.width = 100
		img.height = 100

		
		me.obj.appendChild(img)
		*/

		if(me.onfile != null)  me.onfile(file, e.target.result)


		cnt++
	}

	reader.readAsDataURL(file)


}




