/* Author :Aswathy Ashok */

       $().ready(function() {

            $("#projectForm").validate({


                rules: {

                    projectName : "required",
                    projectNumber: "required",
                    numberOfTask: "required",



                    //lastName: "required",

                },
                messages: {
                    projectName: "Please enter your name",

                    emailAddress: "Please enter a valid email address"

                },
	submitHandler: function() {

				var form_data = $("#projectForm").serialize();
				//console.log(form_data);
				$.ajax({
		                       url: '/project',
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
