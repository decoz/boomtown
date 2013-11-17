
function addCanvas(obj){

	var cv = document.createElement('canvas')

	cv.width = obj.offsetWidth 
	cv.height  = obj.offsetHeight 

	//alert(obj+cv.width + "x" + cv.height)
	cv.style.position = "absolute"
	obj.appendChild(cv)
	return cv	
}

function addDiv(obj){

	var cv = document.createElement('div')

	cv.width = obj.offsetWidth 
	cv.height  = obj.offsetHeight 

	cv.style.position = "absolute"
	cv.className = "clist_div"
	obj.appendChild(cv)
	return cv	
}


function CircleList(parent,cx,cy,rad){
	this.pobj = parent
	this.baselayer = addCanvas(this.pobj) 
	this.gridlayer = addCanvas(this.pobj) 
	this.baselayer.style.zIndex = "1"
	this.gridlayer.style.zIndex = "2"

	this.div = addDiv(this.pobj)
	this.div.style.zIndex = "3"

	this.circle = new Circle(cx,cy,rad) 
	this.basecolor = "#cccccc"

	this.gridstart  = 20 
	this.angle = 20
	this.objangle = 20
	this.stepangle = 15


	this.nodelist = new Array() 

	this.setStartAngle = function(ang,step){
		this.gridstart = ang	
		this.angle = 0
		this.objangle = ang 
		if(step != null) this.stepangle  = step
	}

	this.addValue = function(key,value){
		var node = new Object()
		this.nodelist[key] =  node
		var box = this.insertBox(0,0,key)
		var pt = this.circle.r2x(this.objangle,this.circle.rad)

		node.inbox = box
		node.pt = pt
		node.ang = this.objangle

		box.style.left = ( pt.x - box.offsetWidth - 10 ) + "px"
		box.style.top = ( pt.y - box.offsetHeight ) + "px"

		this.objangle += this.stepangle
	}

	this.stroke = function(){
		this.drawBase() 
		this.drawGrid()
		this.drawText()
	}

	this.drawBase = function(){
		var ctx = this.baselayer.getContext('2d')
		 //ctx.globalCompositeOperation = ""
		this.circle.draw(ctx,10,this.basecolor)	
		ctx.stroke()
	}

	this.drawGrid = function(){
		var ctx = this.gridlayer.getContext('2d')
		ctx.clearRect(0,0,this.pobj.offsetWidth,this.pobj.offsetHeight)
		for(key in this.nodelist)	{
			//var pt = this.circle.r2x(ang,this.circle.rad)

			var obj_ang =   this.nodelist[key].ang + this.angle
			var pt = this.circle.r2x(obj_ang,this.circle.rad)
			//var pt = this.nodelist[key].pt
			new Circle(pt.x,pt.y,5).fill(ctx,"gray")

			//ang += this.stepangle

		}

		ctx.stroke()

	}


	this.drawText = function(){
		for(key in this.nodelist)	{
			//debugger
			var obj_ang =   this.nodelist[key].ang + this.angle
			var pt = this.circle.r2x(obj_ang,this.circle.rad)

			var box = this.nodelist[key].inbox
			box.style.left = ( pt.x - box.offsetWidth - 10 ) + "px"
			box.style.top = ( pt.y - box.offsetHeight ) + "px"
		}

	}


	this.insertBox = function(x,y,value) {
		var div	= this.div
		var obj =document.createElement("div")
		obj.innerHTML = value
		obj.style.position = "absolute"
		//obj.style.fontSize = "12px"
		obj.className = "clist_label"
		div.appendChild(obj)
		obj.style.top= ( y - obj.offsetHeight) + "px"
		obj.style.left = ( x - obj.offsetWidth - 10) +  "px"
		return obj
	}

	this.aniUp = function(){
		var ang = this.stepangle	
		this.animateAngle(ang,10,20)
	}

	this.aniDown = function(){
		var ang = this.stepangle * -1
		this.animateAngle(ang,10,20)
	}

	

	var me = this
	this.check_click = function(e){
		var func =  function(e){
			debugger
			var cx = this.circle.cx
			var cy = this.circle.cy
			var ox = e.offsetX
			var oy = e.offsetY

			var w = cx - ox 
			var h = cy - oy 

			var r = Math.sqrt(w*w + h*h) 
			var pt = this.circle.r2x(45,this.circle.rad)
			if( r < this.circle.rad){
				if( oy > pt.y)	this.aniUp() 
				else this.aniDown()
			} 
		}		
		func.apply(me,[e])
	}





	debugger;

	this.pobj.addEventListener("click",this.check_click)
	//this.pobj.onclick = alert



	this.animateAngle = function( ang, interval, step) {
		//debugger;
		var obj = this
		var cangle = 0 
		var anidraw = function(){
			this.angle  += ang / step  
			cangle  += ang / step  
			this.drawGrid()	
			this.drawText()
			debugger
			if( Math.abs(cangle) > Math.abs(ang)) window.clearInterval(this.interval)
		}

		var ontime = function(e){
			anidraw.apply(obj,[e])			
		}	

		this.interval = setInterval(ontime,interval) 

	}

	return this

}

function insert_box(x,y,value) {
	var obj =document.createElement("div")
	//obj.className = 'rbox'
	obj.className = "clist_label"
	obj.innerHTML = value
	obj.style.top= y + "px"
	obj.style.left = x + "px"
	obj.style.position = "absolute"
	obj.addEventListener("click",onclick_rbox,false)
	e_area.appendChild(obj)
	return obj
}	


function Circle(x,y,rad){
	this.cx = x
	this.cy = y
	this.rad = rad	

	this.r2x = function(rad,range){
	/* rad 는 0~360 도로 */
		var r  = rad * 2 * Math.PI  / 360 
		var  pt = new Object
		pt.y = Math.sin(r) * range + this.cy
		pt.x = Math.cos(r) * range + this.cx
		return pt
	}

	this.draw = function(context,linewidth,color){
		//alert(this.cx + "," + this.cy + ":" + context)
		if(linewidth !=  null)  {
			context .lineWidth = linewidth  
		}
		context.strokeStyle = color
		context.beginPath()
		context.arc(this.cx,this.cy,this.rad,0,2*Math.PI,false)
	}

	this.fill = function(context,color){
		context.beginPath()
		context.fillStyle = color 
		context.strokeStyle = color
		context.arc(this.cx,this.cy,this.rad,0,2*Math.PI,false)
		context.fill()
	}

	return this
}