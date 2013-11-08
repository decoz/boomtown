

    var myInterval

  function stopbot(){
   	myInterval = setInterval(write_random,5000)	
  } 

    function startbot(){

        setInterval(write_random,5000)

    } 

    function write_random() {

        var dmcnt = dummy_list.length 
        var idx = (Math.floor(Math.random() * 100)  )
        idx = idx % dmcnt
        log(idx)
        log(dummy_list[idx])
        wsocket.doSend( "write/bot/" + dummy_list[idx]);


    }

       var dummy_list  = [
        "분홍신/길을 잃었다, 어딜 가야 할까 \
            gireul irheotda, eodil gaya halkka",
        "분홍신2/열두 개로 갈린 조각난 골목길 \
            yeoldu gaero gallin jogangnan golmokgil",
        "분홍신3/어딜 가면 너를 다시 만날까 \
         eodil gamyeon neoreul dasi mannalkka",
        "분홍신4/운명으로 친다면, 내 운명을 고르자면 \
            unmyeongeuro chindamyeon, nae unmyeongeul goreujamyeon",
        "모던타임즈1/난 그댈 알아 그댄 날 몰라 늘 같은 시간 우린 같은 길을 걷는데\
        시간 좀 내요 얘기 좀 해요 ",
        "모던타임즈2/바람에 날린 동그란 저 모자 달리기를 하나 왜 서둘러 시계톱니처럼 똑같은 모퉁일 돌아돌아 뛰네요 ",
        "모던타임즈3/어디서 와서 어디로 갈까 말고도 궁금한 게 이렇게 난 많은데 잠시만 서서 얘기 좀 해요 ",
        "기다려/이 느낌이 아냐 깊숙이 숨겨놓은 그 아일 불러줘 조금 더 내게 불친절 해도 돼 다문 \
        입술이 열리는 순간을 난 기다려 착한 얼굴이 일그러지는 순간을 기다려 기다려 ",
        "입술사이1/Oh darling 넘지 말아요 두 입술 사이 거린 아직까진 50cm 달콤한 말로 뻔한 말로 착한 나를 유혹하려 하진 말아주세요",
        "입술사이2/사랑은, 이 사랑은 완벽할 거에요 Largo, adagio 조급해 말아요 Slowly, baby slowly, very slow 느린 템포로",
        "입술사이3/우리 사랑은, 이 사랑은 짜릿할 거에요 Love ya, baby love ya 그대 윗입술에 빨간 나의 아랫 입술이 닿을 때 쯤엔",
       ]


