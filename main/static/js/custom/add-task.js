/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..
document.getElementById("task").className += " active";

console.log(array.Key)
console.log(array.PageType)
console.log(array.CustomerName)
console.log(array.JobName)
var pageType = array.PageType;

function test(id) {
    alert("Hi: " + id);
}


$(function () {
    
    if(pageType == "edit") {
        document.getElementById("jobName").value = array.JobName;
        document.getElementById("taskName").value = array.TaskName;
        document.getElementById("taskLocation").value = array.TaskLocation;
        document.getElementById("startDate").value = array.StartDate;
        document.getElementById("endDate").value = array.EndDate;
        document.getElementById("taskDescription").value = array.TaskDescription;
        document.getElementById("users").value = array.UserNumber;
        document.getElementById("log").value = array.Log ;
        document.getElementById("userType").value = array.UserType;
        document.getElementById("contacts").value = array.Contact;
        document.getElementById("fitToWork").value = array.FitToWork;
        document.getElementById("taskHead").innerHTML = "Edit Task";
    }
});

       
$().ready(function() {

   var val;
   $(".radio-inline").change(function () {
       val = $('.radio-inline:checked').val();
   });



   $("#taskDoneForm").validate({
       
       rules: {
           taskName: "required",
           jobName: "required"
       },
       
       submitHandler: function() {
           var taskId=array.TaskId;
           var formData = $("#taskDoneForm").serialize();
           if(pageType == "edit"){
            
                $.ajax({
                    url: '/task/'+taskId+'/edit',
                    type: 'post',
                    datatype: 'json',
                    data: formData + "&loginType=" + val,
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
                    data: formData + "&loginType=" + val,
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
});