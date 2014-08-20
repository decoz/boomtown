/*
 * dtable : data 선택, 출력을 쉽게 하기 위한 dot 기반 테이블
 * 테이블 생성후에 다음의 형식으로 사용한다.
 * 
 * var tbl = new dtable(object)
 * 
 * attr: 
 * 
 * 	selected : 셀렉트 된 위치  
 * 
 * method:
 * 	 
 * 	SetHeader(arr,isshow) header 의 id 값을 넣는다. 
 * 	SetClass(row,col,cname) 테이블의 특정 영역을 classname 으로 설정한다.
 * 	AppendData(arr) 테이블에 자료를 추가한다. 
 * 	InsertData(idx,arr) 테이블에 자료를 삽입한다. 
 *	GetRowArr(idx) idx 번째 row 의 값들을 배열로 받는다. 
 * 
 * event: 
 * 
 * 	onclick : 행 선택시  
 *  
 */


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

	this.GetRowArr = function(idx){

		var arr = new Array() 
		var tr = this.tbody.children[idx]
		for(var i =0;i <  tr.children.length;i++){
			arr[i] = tr.children[i].innerHTML
		}
		return arr
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

