/* Author :Aswathy Ashok */

console.log(array.Key)
console.log(array.PageType)
console.log(array.CustomerName)
console.log(array.JobName)
var pageType = array.PageType;
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
var contactsValue;
 function getContact()
{
  var x=document.getElementById("contacts");
  for (var i = 0; i < x.options.length; i++) {
     if(x.options[i].selected){
       contactsValue=x.options[i].value;
  }
  }
}

       
$().ready(function() {

       var val;
       $(".radio-inline").change(function () {

            val = $('.radio-inline:checked').val();


        });
       
      
       $("#taskDoneForm").validate({
         rules: {
            taskName: "required",
            jobName: "required",
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

         submitHandler: function() {
             var taskId=array.TaskId;
             if(pageType == "edit"){
                         $.ajax({
                             
                             url: '/task/'+taskId+'/edit',
                             type: 'post',
                             datatype: 'json',
                             data: $("#taskDoneForm").serialize() + "&loginType=" + val,
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
                        data: $("#taskDoneForm").serialize() + "&loginType=" + val,
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