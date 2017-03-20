/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..
document.getElementById("task").className += " active";
var pageType = vm.PageType;
var customerName = "";
var jobId = "";
var companyTeamName = vm.CompanyTeamName;
var selectedUserArray = []; // contains all selected users and groups
var selectedGroupArray = []; // contains all selected groups
var fitToWorkFromDynamicTextBox = []; // contains all fit to work
var mapLatitude = "";
var mapLongitude = "";
var startDateToCompare = "";
var endDateToCompare = "";
var i = 0;//function for editing
console.log("bfhbshb", vm.StartTime);
$(function () {

   
    if (pageType == "edit") {
        var selectArray =  vm.ContactNameToEdit;
       $("#contactId").val(selectArray);
        var selectArrayForGroup = vm.GroupMembersAndUserToEdit;
        $("#userOrGroup").val(selectArrayForGroup);
        document.getElementById("jobName").value = vm.JobName;
        document.getElementById("taskName").value = vm.TaskName;
        document.getElementById("taskLocation").value = vm.TaskLocation;
         //$('#startDate').val($.datepicker.formatDate('mm-dd-yy', vm.StartDate));
        document.getElementById("startDate").value = vm.StartDate;
        document.getElementById("endDate").value = vm.EndDate;
        document.getElementById("taskDescription").value = vm.TaskDescription;
        document.getElementById("users").value = vm.UserNumber;
        document.getElementById("log").value = vm.Log ;
        document.getElementById("fitToWork").value = vm.FitToWork;
        document.getElementById("startTime").value = vm.StartTime;
        document.getElementById("endTime").value = vm.EndTime;
        document.getElementById("taskHead").innerHTML = "Edit Task";
    }
    
    //google map
    
    
   
    //function for getting textbox dynamically
    
    
    $("#btnAdd").bind("click", function () {
        var div = $("<div class='plus'/>");
        div.html(GetDynamicTextBox(""));
        $("#TextBoxContainer").append(div);
    });
    $("#saveButton").bind("click", function () {
        
       
    });
    $("body").on("click", ".delete-decl", function () {
        $(this).closest("div").delete-decl();
    });
});
function GetDynamicTextBox(value) {
    return ' <input class="form-control"  name = "DynamicTextBox"  id=  "DynamicTextBox"  type="text" value = "" />&nbsp;' +
            '<button id="btnAdd"   class="delete-decl">+</button>'
    
}
 


//function to load add task
var addItem = $('<span>+</span>');
addItem.click(function() {
    window.location = "/"  +  companyTeamName +  "/task/add";
});

$().ready(function() {
    var loginTypeRadio = "";
   $(".radio-inline").change(function () {
       loginTypeRadio = $('.radio-inline:checked').val();
   });
    
    //Functiion for getting job and customer separate
    getJobAndCustomer = function(){
        var job = $("#jobName option:selected").val() + " (";
        var jobAndCustomer = $("#jobName option:selected").text();
        var tempName = jobAndCustomer.replace(job, '');
        customerName = tempName.replace(')', '');
        var jobDropdownId = document.getElementById("jobName");
        jobId = jobDropdownId.options[jobDropdownId.selectedIndex].id;
    }
    
    
    
    
    /*Function will ceck if the selected value is a group name, and if so 
    function will auto select all users in that group*/
    $("#userOrGroup").on('change', function(evt, params) {
        var tempArray = $(this).val();
        var clickedOption = "";
        if (selectedUserArray.length < tempArray.length) { // for selection
            for (var i = 0; i < tempArray.length; i++) {
                if (selectedUserArray.indexOf(tempArray[i]) == -1) {
                    clickedOption = tempArray[i];
                    
                }
            }
            
            for (var i = 0; i < vm.GroupMembers.length; i++) {
                if (vm.GroupMembers[i][0] == clickedOption) {
                    var memberLength = vm.GroupMembers[i].length;
                    for (var j = 1; j < memberLength; j++) {
                        if (tempArray.indexOf(vm.GroupMembers[i][j]) == -1) {
                            tempArray.push(vm.GroupMembers[i][j])
                        }
                        $("#userOrGroup").val(tempArray);
                    }
                    
                    // Inserting group into group array for validating min. no. of users
                    selectedGroupArray.push(clickedOption);
                }
            }           
            
            selectedUserArray = tempArray;
            
            
        } else if (selectedUserArray.length > tempArray.length) { // for deselection
            for (var i = 0; i < selectedUserArray.length; i++) {
                if (tempArray.indexOf(selectedUserArray[i]) == -1) {
                    clickedOption = selectedUserArray[i];
                    
                }
            }
            
            for (var i = 0; i < vm.GroupMembers[i][0].length; i++) {
                if (vm.GroupMembers[i][0] == clickedOption) {
                    var memberLength = vm.GroupMembers[i].length;
                    for (var j = 1; j < memberLength; j++) {
                        var userIndex = tempArray.indexOf(vm.GroupMembers[i][j]);
                        if (userIndex != -1) {
                            tempArray.splice(userIndex, 1);
                        }
                        $("#userOrGroup").val(tempArray);
                    }           
                    
                    
                    // Removing group from group array for validating min. no. of users
                    var deleteGroupKeyIndex = selectedGroupArray.indexOf(clickedOption);
                    selectedGroupArray.splice(deleteGroupKeyIndex, 1);
                }
            }            
            
            selectedUserArray = tempArray;
            
        }
        
        
        
    });
     
       
    $("#taskDoneForm").validate({
        rules: {
            taskName: "required",
            loginType: "required",
        },
        submitHandler: function() {
            
            //code for date and time conversion
            var startDate = new Date($("#startDate").val());
            var startTime =  document.getElementById("startTime").value;
            var endDate = new Date($("#endDate").val());
            var endTime =  document.getElementById("endTime").value;
            
            //setting the time in start date and end date
            
            startTimeArray = startTime.split(':');
            startHour = parseInt(startTimeArray[0]);
            startMin = parseInt(startTimeArray[1]);
            startDate.setHours(startHour);
            startDate.setMinutes(startMin);
            endTimeArray = endTime.split(':');
            endHour = parseInt(endTimeArray[0]);
            endMin = parseInt(endTimeArray[1]);
            endDate.setHours(endHour);
            endDate.setMinutes(endMin);
            //function to convert  date to mm/dd/yyyy format
            
            function formatDate(d){
                function addZero(n){
                    return n < 10 ? '0' + n : '' + n;
                }
                return addZero(d.getMonth()+1)+"/"+ addZero(d.getDate()) + "/" + d.getFullYear() + " " + 
                    addZero(d.getHours()) + ":" + addZero(d.getMinutes());
            }
            var startDateString = startDate;
            var date = new Date(Date.parse(startDateString));
            var startDateOfTask = formatDate(date);
            var endDateString = endDate;
            var endDateData = new Date(Date.parse(endDateString));
            var endDateOfTask = formatDate(endDateData);
            
            
           
            var minUsers = $("#minUsers option:selected").val();
            
            //getting map longitude and latitude
            mapLatitude = document.getElementById("latitudeId").value;// variable to store map latitude
            mapLongitude = document.getElementById("longitudeId").value;// variable to store map longitude
            startDateToCompare = document.getElementById("startDate").value;
            endDateToCompare = document.getElementById("endDate").value;
            if (selectedUserArray.length - selectedGroupArray.length >= minUsers) {
                if(loginTypeRadio.length != 0)
                    {
                      if( endDateToCompare > startDateToCompare) 
                          {
                              if( mapLatitude.length  !=0)
                                  {
                                       $("#saveButton").attr('disabled', true);
                                      var taskId=vm.TaskId;
                                      var jobnew = $("#jobName option:selected").val()
                                      if ($("#jobName ")[0].selectedIndex <= 0) {
                                          document.getElementById('jobName').innerHTML = "";
                                      }
                                      //get all values of fit to work
                                      
                                      var values = "";
                                      $("input[name=DynamicTextBox]").each(function () {
                                          fitToWorkFromDynamicTextBox.push($(this).val())
                                      });
                                      var fitToWorkValue = document.getElementById("addFitToWorkValue").value;
                                      fitToWorkFromDynamicTextBox.push(fitToWorkValue);
                                      alert(fitToWorkFromDynamicTextBox);
                                      var formData = $("#taskDoneForm").serialize() + "&loginType=" + loginTypeRadio + "&customerName=" + customerName + "&jobId=" + jobId +"&addFitToWork=" + fitToWorkFromDynamicTextBox +"&latitude=" +  mapLatitude +"&longitude=" +  mapLongitude +"&startDateFomJs="+ startDateOfTask +"&endDateFromJs="+ endDateOfTask;
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
                                              url: '/' +  companyTeamName  + '/task/' + taskId + '/edit',
                                              type: 'post',
                                              datatype: 'json',
                                              data: formData,
                                              success : function(response) {
                                                  if (response == "true" ) {
                                                      window.location ='/'  +  companyTeamName  + '/task';
                                                  } else {
                                                      $("#saveButton").attr('disabled', false);
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
                                                  if (response == "true" ) {
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
                              else{
                                  $("#mapValidationError").css({"color": "red", "font-size": "15px"});
                                  $("#mapValidationError").html("please select location from map.").show();
                              }
                          }
                        else {
                            $("#dateValidationError").css({"color": "red", "font-size": "15px"});
                           $("#dateValidationError").html("please enter a valid date.").show();
                        }
                    }
                else {
                    $("#loginTypeValidationError").css({"color": "red", "font-size": "15px"});
                    $("#loginTypeValidationError").html("please select a login type.").show();
                }
                
            }
            else {
                $("#minUserValidationError").css({"color": "red", "font-size": "15px"});
                $("#minUserValidationError").html("More users need to start this Task.").show();
                }
        }
    });
    $("#cancel").click(function() {
        window.location = '/' + companyTeamName + '/task';
    });
});