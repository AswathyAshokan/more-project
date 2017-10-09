/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..

document.getElementById("job").className += " active";
var companyTeamName = vm.CompanyTeamName;
console.log("company team namwe",companyTeamName);

$().ready(function() {
    var pageType = vm.PageType;
    if(pageType == "edit") {
       var selectArray =vm.CustomerId;
        console.log("customer",selectArray);
        document.getElementById("jobName").value = vm.JobName;
        document.getElementById("orderNumber").value = vm.OrderNumber;
         document.getElementById("orderDate").value = vm.OrderDate;
        document.getElementById("jobNumber").value = vm.JobNumber;
        $("#customerId").val(selectArray);
        //$("#customerId option[text=selectArray]").attr("selected","selected");
        document.getElementById("jobHead").innerHTML = "Edit Job";
        
          $("#jobForm").validate({
        rules: {
            customerId:"required",
            jobName: {
                required: true,
                remote:{
                    url: "/"+companyTeamName+"/isJobNameUsed/" + jobName+ "/" +vm.PageType+ "/" + vm.JobName,
                    type: "post"
                }
            },
            
            jobNumber: {
                required: true,
                remote:{
                    url: "/"+companyTeamName+"/isJobNumberUsed/" + jobNumber+ "/" +vm.PageType+ "/" + vm.JobNumber,
                    type: "post"
                }
            }
        },
        messages: {
            jobName: {
                required: "Enter job name",
                remote: "Job name already exists!"
            },
            jobNumber:{
                required:"Enter job number",
                remote: "Job number already exists!",
            },
        },
        submitHandler: function() { 
            
             $("#saveButton").attr('disabled', true);
            var formData = $("#jobForm").serialize();
            var customerName = $('#customerId option:selected').text();
            formData = formData +"&customerName="+customerName;
            console.log(formData);
            var jobId = vm.JobId;
            if (pageType == "edit") {
                $.ajax({
                    url: '/' + companyTeamName  + '/job/'+ jobId +'/edit',
                    type: 'post',
                    datatype: 'json',
                    data: formData,
                    success : function(response) {
                        console.log(response);
                        if (response == "true") {
                            window.location = '/' + companyTeamName + '/job';
                        } else {
                            $("#saveButton").attr('disabled', false);
                        }
                    },
                    error: function (request,status, error) {
                        console.log(error);
                    }
                });
            } 
        }
    });
    }
    
     if(pageType == "add") {
         $("#jobForm").validate({
        rules: {
            customerId:"required",
            jobName: {
                required: true,
                remote:{
                    url: "/"+companyTeamName+"/isJobNameUsed/"+  jobName,
                    type: "post"
                }
            },
            
            jobNumber: {
                required: true,
                remote:{
                    url: "/"+companyTeamName+"/isJobNumberUsed/" + jobNumber,
                    type: "post"
                }
            }
        },
        messages: {
            jobName: {
                required: "Enter job name",
                remote: "Job name already exists!"
            },
            jobNumber:{
                required:"Enter job number",
                remote: "Job number already exists!",
            },
        },
        submitHandler: function() { 
            
             $("#saveButton").attr('disabled', true);
            var formData = $("#jobForm").serialize();
            var customerName = $('#customerId option:selected').text();
            formData = formData +"&customerName="+customerName;
            console.log(formData);
            var jobId = vm.JobId;
             
                $.ajax({
                    url: '/' + companyTeamName + '/job/add',
                    type: 'post',
                    datatype: 'json',
                    data: formData,
                    success : function(response) {
                        if (response =="true") {
                            window.location ='/' + companyTeamName + '/job'
                        } else {
                            $("#saveButton").attr('disabled', false);
                        }
                    },
                    error: function (request,status, error) {
                        console.log(error);
                    }
                });
        }
    });
    }
    
    
    //adding alphanumeric validation 
    

  
    $("#cancel").click(function() {
        window.location = '/' +  companyTeamName  + '/job';
    });
});