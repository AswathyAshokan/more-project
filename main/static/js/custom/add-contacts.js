
/* Author :Aswathy Ashok */
console.log(array.Name);
$(function () {
    if(array.PageType == "2") {
            
                document.getElementById("name").value = array.Name;
                document.getElementById("address").value = array.Address;
                document.getElementById("state").value = array.State;
                document.getElementById("zipcode").value = array.ZipCode;
                document.getElementById("emailAddress").value = array.Email;
                document.getElementById("phoneNumber").value = array.PhoneNumber;
                
                }
});
  $().ready(function() {

            $("#contactForm").validate({
                

                rules: {
                   name: "required",


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
		            confirmPassword: {
                        required: true,
			            equalTo :"#password"
                    }
                },

                messages: {
                    firstName: "Please enter your firstName",
                    lastName: "Please enter your lastName",
                    password: {
                        required: "Please provide a password",
                        minlength: "Password at least have 8 characters"
                    },
		            confirmPassword:{
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
                 if(array.PageType == 2){
                         $.ajax({
		                       url: '/contact/:contactId/edit',
		                       type: 'post',
		                       datatype: 'json',

				               data: form_data,


		                      success : function(response) {

                                        if (response =="true") {


                                            window.location = '/contact';
                               			} else {

                               			    }

		                       },
				              error: function (request,status, error) {
       					            console.log(error);
    				          }
		               });
                    
                        
                    } else {

				    var form_data = $("#contactForm").serialize();
                            

				    $.ajax({

		                url: '/contact/add',
		                type: 'post',
		                datatype: 'json',
                        data: form_data,
                        success : function(response) {
                            if (response =="true") {

                                    window.location = '/contact';
                            } else {

                            }
                         },
				        error: function (request,status, error) {
       					    console.log(error);
    				     }
		             });
            }
         }
        });
  });