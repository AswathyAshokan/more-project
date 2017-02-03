/*Created By Farsana*/

$().ready(function() {

	$("#addgroupForm").validate({

	  rules: {
		        	groupname: "required",
                    addgroup:"required"
	},
	messages: {
		            groupname:"please enter group name ",
                    addgroup:"please fill this column"

	  },
	submitHandler: function(){//to pass all data of a form serial
		var formData = $("#addgroupForm").serialize();
	         $.ajax({
			url:'/add-customer',
			type:'post',
			datatype: 'json',
			data: formData,
			//call back or get response here
			success : function(response){
				console.log(response);

			},
		error: function (request,status, error) {
				}


		});
	return false;
     	}
	});

});



























