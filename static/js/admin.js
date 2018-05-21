$("#register").click(function(){
	var name = $("#username").val()
	var pwd = $("#password").val()
	if (name == "" || pwd == ""){
		alert("用户名或密码为空")
		return
	}
	$.ajax({
		url: './registeradmin',
		type: 'get',
		dataType: 'json',
		data: {"uid": name,"password":pwd},
		success:function(result){
			if (result.Status){
				alert("添加成功")
				location.reload()
				return
			}
			alert(result.Msg)
		}
	})
})

$(document).on('click', '.forzen', function(event) {
	var id = $(this).attr("value")
	var type = $(this).attr("type")
	if (type == 1){
		type = 2
	}else {
		type = 1
	}
	var $this = $(this)
	$.ajax({
		url: './forzenadmin',
		type: 'get',
		dataType: 'json',
		data: {"id": id,"status":type},
		success:function(result){
			if (result.Status){
				if (result.Object == 1){
					$this.removeClass("btn btn-social-icon btn-twitter")
					$this.addClass('btn btn-social-icon btn-google')
				}else{
					$this.removeClass('btn btn-social-icon btn-google')
					$this.addClass("btn btn-social-icon btn-twitter")
				}
				$this.attr("type",type)
			}
		}
	})
});

$("#adminpage .page").click(function(){
		console.log($(this))
  		if ($(this).hasClass('disabled')){
  			return
  		}
  		var page = $("#pages").text()
  		if ($(this).attr("value") == "0" ){

  			page = parseInt(page)-1
  		}else{
  			page = parseInt(page)+1
  		}
  		$.ajax({
  			url: './systemlist',
  			type: 'get',
  			dataType: 'json',
  			data: {"page":page},
  			success:function(result){
  				if (result.Status){
  					$(".th").remove()
  					for (var i = 0 ;i <result.Object.length ; i++ ){
  						var text =                     
  									"<tr class='th'>"+
                      					"<td>"+result.Object[i].Id+"</td>"+
                      					"<td>"+result.Object[i].Uid+"</td>"+
                      					"<td>普通管理员</td>"
		          					if (result.Object[i].Status == 1){
		          						text+="<td><a class='btn btn-social-icon btn-google forzen' value="+result.Object[i].Id+" type="+result.Object[i].Status+"><i class='fa  fa-expeditedssl'></i></a></td>"
		          					}else {
		          						text+="<td><a class='btn btn-social-icon btn-twitter forzen' value="+result.Object[i].Id+" type="+result.Object[i].Status+"><i class='fa  fa-expeditedssl'></i></a></td>"
		          					}
                  					text+="<td class=''><a class='btn btn-social-icon btn-vk del' value="+result.Object[i].Id+"><i class='fa fa-close'></i></a></td>"+
                    					  "</tr>"
        				
  						$("#adminbody").append(text)
  					}
  					// 页面的判断
  					$("#pages").text(result.Page)
  					if (result.Next){
  							$("#next").removeClass('disabled')
  						
  					}else {
  						$("#next").addClass('disabled')
  					}
  					if(result.Pref){
						$("#prev").removeClass('disabled')
  						
  					}else {
  						$("#prev").addClass('disabled')
  					}
  				}
  			}
  		})  		
  	});

$(document).on('click', '.del', function(event) {
	if (!confirm("确定要删除该系统用户吗？")){
		return
	}
	var id = $(this).attr("value")
	var $this = $(this)
	$.ajax({
		url: './deladmin',
		type: 'post',
		dataType: 'json',
		data: {"id": id},
		success:function(result){
			if(result.Status){
				$this.parent().parent().remove()
				return
			}
			alert(result.Msg)
		}
	})
});