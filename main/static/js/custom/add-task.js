/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..
document.getElementById("task").className += " active";
var pageType = vm.PageType;
console.log(vm);

$(function () {
    
    if (pageType == "edit") {
        document.getElementById("jobName").value = vm.JobName;
        document.getElementById("taskName").value = vm.TaskName;
        document.getElementById("taskLocation").value = vm.TaskLocation;
        document.getElementById("startDate").value = vm.StartDate;
        document.getElementById("endDate").value = vm.EndDate;
        document.getElementById("taskDescription").value = vm.TaskDescription;
        document.getElementById("users").value = vm.UserNumber;
        document.getElementById("log").value = vm.Log ;
        document.getElementById("userType").value = vm.UserType;
        document.getElementById("contacts").value = vm.Contact;
        document.getElementById("fitToWork").value = vm.FitToWork;
        document.getElementById("taskHead").innerHTML = "Edit Task";
    }
});

       
$().ready(function() {

   var loginTypeRadio = "";
   $(".radio-inline").change(function () {
       loginTypeRadio = $('.radio-inline:checked').val();
   });



   $("#taskDoneForm").validate({
       
       rules: {
           taskName: "required",
           jobName: "required"
       },
       
       submitHandler: function() {
           var taskId=vm.TaskId;
           var formData = $("#taskDoneForm").serialize() + "&loginType=" + loginTypeRadio;
           if(pageType == "edit"){
            
                $.ajax({
                    url: '/task/'+taskId+'/edit',
                    type: 'post',
                    datatype: 'json',
                    data: formData,
                    success : function(response) {
                        if (response =="true") {
                            window.location = '/task';
                        } else {
                        }
                    },
                    error: function (request,status, error) {
                        console.log(error);
                    }
                });
            
            } else {

                $.ajax({
                    url: '/task/add',
                    type: 'post',
                    datatype: 'json',
                    data: formData,
                    success : function(response) {
                        if (response =="true") {
                            window.location = '/task';
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
            window.location = '/task';
    });
});