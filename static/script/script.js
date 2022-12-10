
function stampAnime(visitCount){
	return new Promise(resolve => {
	setTimeout(function(){
		$('#visit-stamp td:eq('+(visitCount)+') span')
			.css('transition','all 0.5s ease-in')
			.addClass('visited');
			resolve('resolved');
	},300);
	
	});
}



async function stamp() {
	
	//Cookieの読み込み
	var visitCount = $.cookie('visitCount');
	//alert(visitCount)
	//スタンプの処理
	if($('#visit-stamp td:eq('+visitCount+') span').length){ //指定のtd要素があるか判定
		//今回訪問したぶんのスタンプをアニメーションで表示
			await stampAnime(visitCount);
			visitCount++;
			//Cookieに訪問数を保存
			$.cookie('visitCount', visitCount, {expires: 7});
	}else{
		//訪問回数がtd要素の数を超えたらすべて表示
		visitCount++;
		$('#visit-stamp td:lt('+visitCount+') span').addClass('visited');
	}

	//訪問数の表示
	$('#visitnum').text((visitCount)+'回訪問しました。');
	
};


const latitude = 35.5827712;
const longitude = 139.7030912;

function getPosition() {
	// 現在地を取得
	navigator.geolocation.getCurrentPosition(
	  // 取得成功した場合
	  function(position) {
		
		  //alert("緯度:"+position.coords.latitude+",経度"+position.coords.longitude+" "+latitude+" "+longitude);
		  if (position.coords.latitude == latitude && position.coords.longitude == longitude){
		  stamp();
		  }
	  },
	  // 取得失敗した場合
	  function(error) {
		switch(error.code) {
		  case 1: //PERMISSION_DENIED
			alert("位置情報の利用が許可されていません");
			break;
		  case 2: //POSITION_UNAVAILABLE
			alert("現在位置が取得できませんでした");
			break;
		  case 3: //TIMEOUT
			alert("タイムアウトになりました");
			break;
		  default:
			alert("その他のエラー(エラーコード:"+error.code+")");
			break;
		}
	  }
	);
  }



$(function(){
	
	//Cookieの読み込み
	var visitCount = $.cookie('visitCount');
	//訪問数のカウント
	if(visitCount == null){ //最初の訪問
		visitCount = 0;
	}
	//Cookieに訪問数を保存
	$.cookie('visitCount', visitCount, {expires: 7});
	
	//スタンプの処理
	if($('#visit-stamp td:eq('+visitCount+') span').length){ //指定のtd要素があるか判定
		//過去に訪問したぶんのスタンプを表示
		if($('#visit-stamp td:lt('+visitCount+') span').length){
			$('#visit-stamp td:lt('+visitCount+') span').addClass('visited');
		}
	}else{
		//訪問回数がtd要素の数を超えたらすべて表示
		$('#visit-stamp td:lt('+visitCount+') span').addClass('visited');
	}
	
	//訪問数の表示
	$('#visitnum').text((visitCount)+'回訪問しました。');
	
	//Cookieのリセットクリック時の処理
	$('#reset').click(function(){
		$.removeCookie('visitCount');
		alert("Cookieをリセットしました。")
	});
	
});
