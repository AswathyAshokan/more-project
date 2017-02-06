/*Created By Farsana*/

$().ready(function() {

	$("#adduserForm").validate({
	  rules: {

		            firstname: "required",
                  	emailid:{
            			required:true,
            			email:true
            		},

   	  },
	messages: {

		            firstname:"please enter first name ",
                	emailid:"please enter currect email id"



    },
	    submitHandler: function(){//to pass all data of a form serial
		    var formData = $("#adduserForm").serialize();
	         $.ajax({
		    	url:'/invitate/add',
			    type:'post',
			    datatype: 'json',
			    data: formData,
			    //call back or get response here
			    success : function(response){
			    	 if(response == "true"){

                    	 window.location='/invitate';
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






