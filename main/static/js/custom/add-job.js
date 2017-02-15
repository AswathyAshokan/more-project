/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..
document.getElementById("job").className += " active";
$().ready(function() {
    
    var pageType = vm.PageType;
    
    if(pageType == "edit") {
        console.log(vm);
            console.log("Hai");
            $("#customerName").val(vm.CustomerName);
            document.getElementById("jobName").value = vm.JobName;
            document.getElementById("jobNumber").value = vm.JobNumber;
            document.getElementById("numberOfTask").value = vm.NumberOfTask;
            document.getElementById("jobHead").innerHTML = "Edit Job";
            
           
    } 
   

    $("#jobForm").validate({
                rules: {
                    numberOfTask:"required",
                    jobName: {
                        required: true,
                        remote:{
                            url: "/isJobNameUsed/" + jobName,
                            type: "post"
                        }

                    },
		            jobNumber: {
                        required: true,
                        remote:{
                             url: "/isJobNumberUsed/" + jobNumber,
                             type: "post"
                        },
                        number: true
                    }
                    
                },
           
                messages: {

                    jobName: {
                        required: "Please provide a job name",
                         remote: "Job name already exists!"

                    },
		            jobNumber:{
			            required:"please provide a job number",
			            remote: "Job number already exists!"
			        },

                },
                    
	            submitHandler: function() {
                    var formData = $("#jobForm").serialize();
                    var customerName = $('#customerId option:selected').text();
                    formData = formData +"&customerName="+customerName;
                    console.log(formData);
                    var jobId = vm.JobId;
                    
                    if (pageType == "edit") {
                        
                        $.ajax({
                            url: '/job/'+ jobId +'/edit',
                            type: 'post',
                            datatype: 'html',
                            data: formData,
                            success : function(response) {
                                console.log(response);
                                if (response == "true") {
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
                            data: formData,
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

          $("#cancel").click(function() {
                     window.location = '/job';
             });
 });