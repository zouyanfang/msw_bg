
 //     $(".next").click(function(){
 //      alert("这是外部引入样式")
	// })
   
  $(function(){

  	// 网站建设的分页
  	$("#buildpage .page").click(function(){
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
  			url: 'user/userlist',
  			type: 'get',
  			dataType: 'json',
  			data: {"page":page},
  			success:function(result){
  				console.log(result.Object)
  				if (result.Status){
  					$(".th").remove()
  					for (var i = 0 ;i <result.Object.length ; i++ ){
  						var text =                     
  									"<tr class='th'>"+
                      					"<td>"+result.Object[i].Id+"</td>"+
                      					"<td>"+result.Object[i].Uid+"</td>"+
                      					"<td>"+result.Object[i].AdviseDate+"</td>"+
                      					"<td>"+result.Object[i].Title+"</td>"+
                      					"<td class='' ><span class='label label-success pointer checks' value="+result.Object[i].Id+">查看</span></td>"+
                    				"</tr>"
        				
  						$("#buildbody").append(text)
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

  	// 网站建设的分页
  	$("#resourcepage .page").click(function(){
  		if ($(this).hasClass('disabled')){
  			return
  		}
  		var page = $("#pages1").text()
  		if ($(this).attr("value") == "0" ){

  			page = parseInt(page)-1
  		}else{
  			page = parseInt(page)+1
  		}
  		$.ajax({
  			url: 'user/resourcelist',
  			type: 'get',
  			dataType: 'json',
  			data: {"page":page},
  			success:function(result){
  				console.log(result.Object)
  				if (result.Status){
  					$(".th1").remove()
  					for (var i = 0 ;i <result.Object.length ; i++ ){
  						var text =                     
  									"<tr class='th1'>"+
                      					"<td>"+result.Object[i].Id+"</td>"+
                      					"<td>"+result.Object[i].Uid+"</td>"+
                      					"<td>"+result.Object[i].AdviseDate+"</td>"+
                      					"<td>"+result.Object[i].Title+"</td>"+
                      					"<td class='' ><span class='label label-success pointer checks' value="+result.Object[i].Id+">查看</span></td>"+
                    				"</tr>"
        				
  						$("#resourcebody").append(text)
  					}
  					// 页面的判断
  					$("#pages1").text(result.Page)
  					if (result.Next){
  							$("#nexts").removeClass('disabled')
  						
  					}else {
  						$("#nexts").addClass('disabled')
  					}
  					if(result.Pref){
						$("#prevs").removeClass('disabled')
  						
  					}else {
  						$("#prevs").addClass('disabled')
  					}
  				}
  			}
  		})  		
  	});

  	// 查看详情
  	$(document).on('click', '.checks', function(event) {
  		var id = $(this).attr("value")
	    $.ajax({
	    	url: 'user/detail',
	    	type: 'get',
	    	dataType: 'json',
	    	data: {"id":id},
	    	success:function(result){
	    		if (result.Status){
			    	var text = `
      			<div class="col-sm-12 me" style="margin: 0px;padding: 0px;">
                    <div class="box" >
                      <div class="box-header with-border">`
                          	text += "<h3 class='box-title'>"+result.Object.Title+"</h3>"
                        	text +=`<div class="box-tools pull-right">
	                          	<button type="button" class="btn btn-box-tool remove" data-widget="remove" data-toggle="tooltip" title="" data-original-title="Remove">
	                            	<i class="fa fa-times" ></i>
	                          	</button>
                  			</div>
                    	</div>`
	                    text += "<div class='box-body'>"+
	                        	result.Object.Content+
	                    "</div>"
	                    text +="<div class='box-footer'>"+
	                        result.Object.AdviseDate+
	                    "</div>"+
               		"</div>"+
            	 "</div>"
            	 console.log(text)
            	 $("#detatil").append(text)
            	 return
	    		}
	    		alert(result.Msg)
	    	}
	    })	    
  	});


  	$(document).on('click', '.remove', function(event) {
		$(this).parent().parent().parent().parent().remove()
  	});
  })

