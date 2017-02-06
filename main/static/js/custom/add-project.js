/* Author :Aswathy Ashok */

       $().ready(function() {

            $("#projectForm").validate({
                rules: {

                    projectName : "required",
                    projectNumber: "required",
                    numberOfTask: "required",


                },
                messages: {
                    projectName: "Please enter your name",

                    emailAddress: "Please enter a valid email address"

                },
	            submitHandler: function() {

				    var form_data = $("#projectForm").serialize();

				    $.ajax({
		                       url: '/project/add',
		                       type: 'post',
		                       datatype: 'json',

				               data: form_data,


		                      success : function(response) {

                                        if (response =="true") {


                                            window.location = '/project';
                               			} else {

                               			    }

		                       },
				              error: function (request,status, error) {
       					            console.log(error);
    				          }
		            });


            }
         });
 });