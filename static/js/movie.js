
$("#moviepage .page").click(function(){
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
  			url: './movielist',
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
                      					"<td>"+result.Object[i].MovieName+"</td>"+
                      					"<td>"+result.Object[i].ReleaseDate+"</td>"+
                                "<td><a class='btn btn-social-icon btn-vk update' value="+result.Object[i].Id+"><i class='fa  fa-wrench'></i></a></td>"
              if (result.Object[i].Source == "exist"){
                  text += "<td class=''><a class='btn btn-social-icon btn-github  disabled' value="+result.Object[i].Id+" style='background-color: #4CAF50'><i class='fa fa-send-o'></i></a></td>"
              }else {
                  text += "<td class=''><a class='btn btn-social-icon btn-github upmovie2' value="+result.Object[i].Id+"><i class='fa fa-send-o'></i></a></td>"
              }
              text += "</tr>"
        				
  						$("#moviebody").append(text)
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

// 查看电影详情
$(document).on('click', '.update', function(event) {
  console.log("function")
  var $this = $(this)
  var id = $(this).attr("value")
  var moviename = $("#movieName")
  var director = $("#director")
  var area = $("#area")
  var type = $("#type")
  var language = $("#language")
  var releasedate = $("#releasedate")
  var content = $("#content")
  var short = $("#short")
  var sumbit = $("#submit")
  var s = $("#area option")
  var t = $("#type option")
  var l = $("#language option")
  $.ajax({
    url: './findmovie',
    type: 'get',
    dataType: 'json',
    data: {"id": id},
    success:function(result){
      console.log(result)
      if (result.Status){
        moviename.val(result.Object.MovieName)
        director.val(result.Object.Director)
        var i = 0
        for(var j = 0;j<=s.length;j++){
            if($(s[j]).text() == result.Object.Area){
                i = j+1
                break;
              }
        }
        area.val(i)
        for(var j = 0;j<=t.length;j++){
          if($(t[j]).text() == result.Object.Type){
              i = j+1
              break;
          }
        }
        type.val(i)
        for(var j = 0;j<=l.length;j++){
          if($(l[j]).text() == result.Object.Language){
              i = j+1
              break;
          }
          console.log(l[j])
        }
        language.val(i)
        releasedate.val(result.Object.ReleaseDate)
        content.val(result.Object.Plot)
        short.val(result.Object.Introuduce)
        sumbit.attr("value",id)
        $("#modifybox").removeClass('hidden')
      }
    }
  })
});

//上传电影
$("#submit").click(function(){
  console.log(123)
  var formData = new FormData()
  var id = $(this).attr("value")
  var moviename = $("#movieName").val()
  var director = $("#director").val()
  var area = $("#area").find("option:selected").text()
  var type = $("#type").find("option:selected").text()
  var language = $("#language").find("option:selected").text()
  var releasedate = $("#releasedate").val()
  var content = $("#content").val()
  var short = $("#short").val()
  console.log(456)
  if (id == "" || moviename == "" || director == "" || area == "" || type == "" || language == "" || releasedate == "" || content == "" || short== ""){
    alert("有选项为空请填写完毕")
    return
  }
  console.log(789)
  if ($("#img").val() != ""){
    formData.append('img',$('#img')[0].files[0])  
    formData.append('status',1)
  }
  formData.append('moviename',moviename)
  formData.append('director',director)
  formData.append('area',area)
  formData.append('type',type)
  formData.append('language',language)
  formData.append('releasedate',releasedate)
  formData.append('content',content)
  formData.append('short',short)
  formData.append('id',id)
  $.ajax({
    url: './modifymovie',
    type: 'post',
    dataType: 'json',
    data:formData,
    processData:false,
    contentType:false,
    success:function(result){
      if (result.Status){
        alert("修改成功")
        location.reload()
        return
      }
      alert(result.Msg)
    }
  })
})

// 上传电影资源
$(document).on('click', '.upmovie2', function(event) {
  var id = $(this).attr("value")
  $.ajax({
    url: './findmovie',
    type: 'post',
    dataType: 'json',
    data: {"id":id},
    success:function(result){
      if (result.Status){
        console.log(result)
        $("#name").text(result.Object.MovieName)
        $("#movieId").val(result.Object.Id)
        console.log($("#movieId").val()+result.Object.Id)
        $("#movieresource").removeClass("hidden")
      }
    }
  })  
});


//添加电影
$("#submits").click(function(){
  console.log(123)
  $("#submits").text("正在上传...")
  var formData = new FormData()
  var id = $(this).attr("value")
  var moviename = $("#movieName").val()
  var director = $("#director").val()
  var area = $("#area").find("option:selected").text()
  var type = $("#type").find("option:selected").text()
  var language = $("#language").find("option:selected").text()
  var releasedate = $("#releasedate").val()
  var content = $("#content").val()
  var short = $("#short").val()
  console.log(456)
  if ( moviename == "" || director == "" || area == "" || type == "" || language == "" || releasedate == "" || content == "" || short== ""){
    alert("有选项为空请填写完毕")
    return
  }
  console.log(789)
  if ($("#img").val() != ""){
    formData.append('img',$('#img')[0].files[0])  
    formData.append('status',1)
  }
  if ($("#video").val() != ""){
    formData.append('video',$('#video')[0].files[0])  
    formData.append('videostatus',1)
  }
  formData.append('moviename',moviename)
  formData.append('director',director)
  formData.append('area',area)
  formData.append('type',type)
  formData.append('language',language)
  formData.append('releasedate',releasedate)
  formData.append('content',content)
  formData.append('short',short)
  formData.append('id',id)
  $.ajax({
    url: './releasemovie',
    type: 'post',
    dataType: 'json',
    data:formData,
    processData:false,
    contentType:false,
    success:function(result){
      if (result.Status){
        alert("成功添加电影！")
        $("#submits").text('上传成功')
        $("#close").click()
        location.reload()
        return
      }
      alert(result.Msg)
    }
  })
})