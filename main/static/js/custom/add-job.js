/* Author :Aswathy Ashok */
$(function () {
    
    if(array.PageType == "2") {
            
            document.getElementById("customerName").value = array.CustomerName;
            document.getElementById("jobName").value = array.JobName;
            document.getElementById("jobNumber").value = array.JobNumber;
            document.getElementById("numberOfTask").value = array.NumberOfTask;
            
           
    }
});
     
$().ready(function() {
   

    $("#jobForm").validate({
                rules: {
                    numberOfTask:"required",
                    jobName: {
                        required: true,
                       
                    },
		            jobNumber: {
                        required: true
                    }
                    
                },
           
                messages: {
                    jobName: "Please enter your name",

                    emailAddress: "Please enter a valid email address",
                    jobName: {
                        required: "Please provide a job name",
                        
                    },
		            jobNumber:{
			            required:"please provide a job number"
			        },

                },
	            submitHandler: function() {

				    var form_data = $("#jobForm").serialize();
                    if(array.PageType == 2){
                         $.ajax({
		                       url: '/job/:jobId/edit',
		                       type: 'post',
		                       datatype: 'json',

				               data: form_data,


		                      success : function(response) {

                                        if (response =="true") {


                                            window.location = '/job';
                               			} else {

                               			    }

		                       },
				              error: function (request,status, error) {
       					            console.log(error);
    				          }
		               });
                    
                        
                    } else {
                
                        $.ajax({
		                       url: '/job/add',
		                       type: 'post',
		                       datatype: 'json',

				               data: form_data,


		                      success : function(response) {

                                        if (response =="true") {


                                            window.location = '/job';
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