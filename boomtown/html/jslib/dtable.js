function dtable(obj){

	this.obj = obj	

	this.thead = null
	this.tbody = null
	this.selected = null
	this.onclick = null

	var me = this
	
	var first = this.obj.children[0]

	this.Clear = function(){
		this.obj.innerHTML = ""
	}


	this.SetHeader = function( arr, isshow){

		if(this.thead == null) this.addHeader() 
		this.thead.innerHTML = ""

		for(i in arr){
			var th = document.createElement("TH")
			th.innerHTML = arr[i]
			this.thead.appendChild(th)

		}
	}

	this.addHeader = function() {

		var  thd = document.createElement("THEAD")
		var first = this.obj.children[0]
		
		if(first) this.obj.insertBefore(thd,first)
		else this.obj.appendChild(thd)

		this.thead = thd

	}
	
	this.addTBody = function() {		
		var tbody = document.createElement("TBODY")
		var arr = new Array() 
		var children = this.obj.children
	
		for(i in children){
			if( children[i].tagName == "TR")  arr[i] = children[i]				
		}

		for(i in arr){
			tbod.appendChild(arr[i])
		}

		this.obj.appendChild(tbody)
		this.tbody = tbody

	}
	
	this.SetClass = function(row,col,cname){		

		var rarr = new Array() 
		var carr = new Array()

		if(row == -1)  rarr[0] = this.thead 
		else if( row == null)  {
			for( var i   = 0; i < this.tbody.children.length;  i++ ){
				var obj = this.tbody.children[i]
				rarr[i] = obj
			}			
		}
		else rarr[0] = this.tbody.children[row]

		if( col == null)  {
			for( var i = 0; i < this.thead.children.length; i++ ){
				var obj = this.thead.children[i]				
				carr[i] = i
			}			
		}
		else carr[0] = col

		for(r in rarr){
			var tr = rarr[r]
			if( tr  )  for(i in carr){
				var td = tr.children[i]
				if(td) td.className = cname
			}
		}		
	}


	this.getTRIndex = function(tr){
		var tbody = this.tbody
		for(i in tbody.children){
			if(tr == tbody.children[i]) return i
		}
		return null
	}

	this.createTR = function(arr){
		var tr = document.createElement("TR")
		tr.addEventListener("click",e_click)

		for(var i in arr) {
			var td = document.createElement("TD")			
			td.innerHTML = arr[i]
			tr.appendChild(td)
		}

		return tr		
	}



	this.AppendData = function(arr){
		debugger
		var tr = this.createTR(arr)
		this.tbody.appendChild(tr)
		return this.tbody.children.length - 1
	}

	this.InsertData = function(idx,arr){
		
		if( idx < this.tbody.children.length ){					
			var tr = this.createTR(arr)
			var next = this.tbody.children[i]
			this.tbody.insertBefore(tr,next)
		}
 		else this.AppendData(arr)
	}


	if(first && first.tagName == "THEAD")  this.thead = first 
	else this.addHeader()
		
	var second = this.obj.children[1]
	if(second && second.tagName == "TBODY") this.tbody = second
	else  this.addTBody()

	var e_click = function(e){
		debugger
		me.selected = me.getTRIndex(this)
		//alert(me.selected)
		if( me.onclick != null ) me.onclick.apply(me,[e])
	}

	for(i in this.tbody.children){
		var obj = this.tbody.children[i]
		if( obj.tagName == "TR") obj.addEventListener("click",e_click)

	}
	
}

