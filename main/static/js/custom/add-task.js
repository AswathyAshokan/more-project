/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..
document.getElementById("task").className += " active";
var pageType = vm.PageType;
var customerName = "";
var jobId = "";
var companyTeamName = vm.CompanyTeamName

$(function () {

   
    if (pageType == "edit") {
        var selectArray = vm.ContactNameToEdit;
         console.log(selectArray);
        $("#contactId").val(selectArray);
        document.getElementById("jobName").value = vm.JobName;
        document.getElementById("taskName").value = vm.TaskName;
        document.getElementById("taskLocation").value = vm.TaskLocation;
        document.getElementById("startDate").value = vm.StartDate;
        document.getElementById("endDate").value = vm.EndDate;
        document.getElementById("taskDescription").value = vm.TaskDescription;
        document.getElementById("users").value = vm.UserNumber;
        document.getElementById("log").value = vm.Log ;
        document.getElementById("fitToWork").value = vm.FitToWork;
        document.getElementById("taskHead").innerHTML = "Edit Task";
    }
});

var addItem = $('<span>+</span>');
addItem.click(function() {
    window.location = "/"  +  companyTeamName +  "/task/add";
});

$().ready(function() {
    var loginTypeRadio = "";
   $(".radio-inline").change(function () {
       loginTypeRadio = $('.radio-inline:checked').val();
   });
    
    getJobAndCustomer = function(){
        var job = $("#jobName option:selected").val() + " (";
        var jobAndCustomer = $("#jobName option:selected").text();
        var tempName = jobAndCustomer.replace(job, '');
        customerName = tempName.replace(')', '');
        var jobDropdownId = document.getElementById("jobName");
        jobId = jobDropdownId.options[jobDropdownId.selectedIndex].id;
        
    }
     
       
    $("#taskDoneForm").validate({
        rules: {
            taskName: "required",
            loginType: "required",
        },
        submitHandler: function() {
            
            var taskId=vm.TaskId;
           var jobnew = $("#jobName option:selected").val()
           console.log("job id",jobnew);
           var formData = $("#taskDoneForm").serialize() + "&loginType=" + loginTypeRadio + "&customerName=" + customerName + "&jobId=" + jobId;
        //if job name is not selected
//            if (!$("#jobName option:selected").length) {
//                var job
//                var formData = $("#taskDoneForm").serialize() + "&jobName=" + loginTypeRadio 
//            }
           var selectedContactNames = [];

           //get the user's name corresponding to  keys selected from dropdownlist
            $("#contactId option:selected").each(function () {
                var $this = $(this);
                if ($this.length) {
                    var selectedContactName = $this.text();
                    selectedContactNames.push( selectedContactName);
                }
            });
              for(i = 0; i < selectedContactNames.length; i++) {
                formData = formData+"&contactName="+selectedContactNames[i];
            }
           
           //function to get all users and group
           
           var selectedUserAndGroupName = [];
           $("#userOrGroup option:selected").each(function () {
               var $this = $(this);
                if ($this.length) {
                    var selectedUserName = $this.text();
                    console.log(selectedUserName);
                    selectedUserAndGroupName.push( selectedUserName);
                }
            });
    
       for(i = 0; i < selectedUserAndGroupName.length; i++) {
               formData = formData+"&userAndGroupName="+selectedUserAndGroupName[i];
           }
           if(pageType == "edit"){
               $.ajax({
                    url: '/' +  companyTeamName  + '/task/'+taskId+'/edit',
                    type: 'post',
                    datatype: 'json',
                    data: formData,
                    success : function(response) {
                        if (response =="true") {
                            window.location ='/'  +  companyTeamName  + '/task';
                        } else {
                            
                        }
                    },
                    error: function (request,status, error) {
                        console.log(error);
                    }
                });
            } else {
                $.ajax({
                    url:'/'+ companyTeamName + '/task/add',
                    type: 'post',
                    datatype: 'json',
                    data: formData,
                    success : function(response) {
                        if (response =="true") {
                            window.location = '/' + companyTeamName + '/task';
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
    
    $("#cancel").click(function() {
            window.location = '/' + companyTeamName + '/task';
    });
});