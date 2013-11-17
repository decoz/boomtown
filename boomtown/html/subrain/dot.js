function dot(str){
	this.value = str 
	this.child = new Array()

	this.AppendChild = function(dot){
		var cnt = this.child.length
		this.child[cnt] = dot
	}

	this.GetChildList = function(){
		return this.child
	}

	this.GetChild  = function(i){
		return this.child[i]
	}

	this.SetValue = function(val){
		this.value  = dot.Enc(val)
	}


	this.GetValue = function(){
		return dot.Dec(this.value )
	}

}



dot.Enc = function(str){
	result = ""
	for(var i = 0 ;i  < str.length; i++){
		var c = str[i]
		switch c {
			case ',' : result += "&c"
				break;
			case '.' : result += "&d"
				break;
			case '(': result += "&s"
				break;
			case ')': result += "&e"
				break;
			case '&' result += "&&"
				break;
			default:  result += c
		}

	}
}

dot.Dec = function(str){
	result = ""
	var flag = false
	for(var i = 0 ;i  < str.length; i++){
		var c = str[i]
		if(flag){
			switch c {
				case 'c' : result += ","
					break;
				case 'd' : result += "."
					break;
				case 's': result += "("
					break;
				case 'e': result += ")"
					break;
				case '&' result += "&"
					break;
				default:  result += c
			}

			flag = false
		}
		else {
			if( c == '&') flag  = true
			else result += c
		}

	}
}