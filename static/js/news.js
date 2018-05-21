$("#save").click(function(){
	var formData = new FormData()
	var title = $("#title").val()
	var point = $("#point").val()
	var img = $("#upload").val()
	var text = $(".textarea").val()
	if (title == "" || point == "" || img == "" || text == "" ) {
		console.log(title)
		console.log(point)
		console.log(img)
		console.log(text)
		return
	}
	formData.append('img',$('#upload')[0].files[0])
	formData.append('title',title)
	formData.append('point',point)
	formData.append('content',text)
	console.log(name)
	$.ajax({
		url: './createnews',
		type: 'post',
		dataType: 'json',
		data:formData,
		processData:false,
	    contentType:false,
		beforeSend:function(){
			console.log("正在进行")	
		},
		success:function(result){
			if (result.Status){
				alert("新闻已发布！")
				location.reload()
			}else {
				alert(result.Msg)
			}
		}
	})	
});

$(document).on('click', '.del', function(event) {
	if (!confirm("确定要删除该新闻吗？")){
		return
	}
	var id = $(this).attr("value")
	var $this = $(this)
	$.ajax({
		url: './detelenews',
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

$("#newspage .page").click(function(){
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
  			url: '/information/newslist',
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
                      					"<td>"+result.Object[i].Point+"</td>"+
                      					"<td class=''><a class='btn btn-social-icon btn-vk del' value="+result.Object[i].Id+"><i class='fa fa-close'></i></a></td>"+
                    				"</tr>"
        				
  						$("#newsbody").append(text)
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

