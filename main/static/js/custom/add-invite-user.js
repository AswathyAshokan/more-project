/*Created By Farsana*/
console.log(vm);
//document.getElementById("firstname").value = vm.firstname;


$().ready(function() {
    
    if(vm.PageType == "edit"){        
                document.getElementById("firstname").value = vm.FirstName;
                document.getElementById("lastname").value = vm.LastName;
                document.getElementById("emailid").value = vm.EmailId;
                document.getElementById("usertype").value = vm.UserType;
    }

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
	        if (vm.PageType == "2"){

	                var formData = $("#adduserForm").serialize();
                	         $.ajax({

                		    	url:'/invitate/'+ InviteId+ '/edit',
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
	        } else {
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
	        }

	        return false;
     	}
	});

});






