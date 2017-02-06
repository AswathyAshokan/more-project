/*Created By Farsana*/

$().ready(function() {

	$("#addgroupForm").validate({

	  rules: {
		        	groupname: "required",
                    groupMember:"required"
	},
	messages: {
		            groupname:"please enter group name ",
                    groupMember:"please fill this column"

	},
	submitHandler: function(){//to pass all data of a form serial
		var formData = $("#addgroupForm").serialize();
	         $.ajax({
			    url:'/group/add',
			    type:'post',
			    datatype: 'json',
			    data: formData,
			    //call back or get response here
			    success : function(response){
			        if(response == "true"){

            	         window.location='/group';
                    }else {
                    }
			    },
		        error: function (request,status, error) {
				}


		     });
	    return false;
    }
	});

});



























