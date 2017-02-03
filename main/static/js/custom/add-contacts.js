
/* Author :Aswathy Ashok */
       $().ready(function() { 
	
            $("#contactForm").validate({

                rules: {
                    name: "required",

                    //lastName: "required",
                   emailAddress: {
                        required: true,
                        email: true
                    	},
			phoneNumber: {
				required: true,
				minlength : 10
			},
                    password: {
                        required: true,
                        minlength: 8
                    },
		 confirmpassword: {
                        required: true,
			equalTo :"#password"
                    }
                },
                messages: {
                    firstName: "Please enter your firstname",
                    lastName: "Please enter your lastname",
                    password: {
                        required: "Please provide a password",
                        minlength: "Password at least have 8 characters"
                    },
		confirmpassword:{
			required:"please provide a password",
			equalTo:"please enter the password as above"
			},


               phoneNumber:{
			required:"please provide a phone number",
			minlength:"your phone number at least 10 digit long"
		},
                    emailAddress: "Please enter a valid email address"

                },
	submitHandler: function() {

				var form_data = $("#contactForm").serialize();
				//console.log(form_data);
				$.ajax({
		                       url: '/contact',
		                       type: 'post',
		                       datatype: 'json',


						//name: $('#name').val(),
					 	//phoneNumber: $('#phoneNumber').val(),
						//emailAddress: $('#emailAddress').val(),
						//address: $('#address').val(),
						//state: $('#state').val(),
						//zipcode: $('#zipcode').val(),
						//request: serverName
				         data: form_data,


		                      success : function(response) {




		                       },
				        error: function (request,status, error) {
       					 console.log(error);
    				}
		                     });


                          }
                    });
 });
