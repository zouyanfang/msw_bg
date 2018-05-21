$("#msgpage .page").click(function(){
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
  			url: '/information/officiallist',
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
                      					"<td>"+result.Object[i].Title+"</td>"+
                      					"<td>"+result.Object[i].Date+"</td>"+
                      					"<td>"+result.Object[i].Content+"</td>"+
                      					"<td class=''><a class='btn btn-social-icon btn-vk del' value="+result.Object[i].Id+"><i class='fa fa-close'></i></a></td>"+
                    				"</tr>"
        				
  						$("#msgbody").append(text)
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

// 清空
$("#clear").click(function(event) {
	$("#title").val('')
	$('#content').val('')
});

$("#submit").click(function(){
	var title = $("#title").val()
	var content = $('#content').val()
	if (title == "" || content == ""){
		alert("消息不能为空！")
		return
	}
	$.ajax({
		url: './releaseofficialmsg',
		type: 'post',
		dataType: 'json',
		data: {"title":title,"content":content},
		success:function(result){
			if (result.Status){
				alert("发布成功")
				location.reload()
				return
			}
			alert(result.Msg)
		}
	})
});


$(document).on('click', '.del', function(event) {
	if (!confirm("确定要删除该公告吗？")){
		return
	}
	var id = $(this).attr("value")
	var $this = $(this)
	$.ajax({
		url: './delofficialmsg',
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
