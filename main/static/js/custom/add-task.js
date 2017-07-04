/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..

console.log(vm.CompanyPlan);
document.getElementById("task").className += " active";
var pageType = vm.PageType;
var customerName = "";
var jobId = "";
var companyTeamName = vm.CompanyTeamName;
var selectedUserArray = []; // contains all selected users and groups
var selectedGroupArray = []; // contains all selected groups
var fitToWorkFromDynamicTextBox = []; // contains all fit to work
var fitToWorkFromDynamicTextBoxValue =[];
var mapLatitude = "";
var mapLongitude = "";
var startDateToCompare = "";
var endDateToCompare = "";
var minUserForTaskEdit ="";
var loginTypeForEdit ="";
var i = 0;//function for editing
var fitToWorkCheck ="";
var exposureSlice =[];
var exposureTimeArray =[];
var exposureWorkSlice =[];
var exposureWorkTimeArray =[];
console.log("log",vm.Log);
//if group members is null ,group member array is initialised
if(vm.GroupMembers == null) {
    vm.GroupMembers = [];
}
$(function () {
    if(vm.CompanyPlan == "family"){
        document.getElementById("jobNamelabel").style.display = "none";
        document.getElementById("workExplosureLabel").style.display = "none";
        document.getElementById("minUsersLabel").style.display = "none";
        document.getElementById("nfcbutton").style.display = "none";
        //document.getElementById("contactIdLabel").style.display = "none";
        document.getElementById("minUsers").style.display = "none";
        document.getElementById("loginType1").style.display = "none";
        document.getElementById("jobName").style.display = "none";
        document.getElementById("workExplosure").style.display = "none";
        $("#contactDiv").hide();
    } else if(vm.CompanyPlan == "campus"){
        
        document.getElementById("jobNamelabel").style.display = "none";
        document.getElementById("workExplosureLabel").style.display = "none";
        document.getElementById("minUsersLabel").style.display = "none";
        document.getElementById("minUsers").style.display = "none";
        document.getElementById("jobName").style.display = "none";
        document.getElementById("workExplosure").style.display = "none";
        document.getElementById("contactIdLabel").style.display = "none";
         $("#contactDiv").hide();
    }else{
         document.getElementById("minUsers").style.display = "block";
         document.getElementById("jobName").style.display = "block";
         document.getElementById("workExplosure").style.display = "block";
          $("#contactDiv").show();
    }
    
    // date picker 
    $( "#startDate" ).datepicker({ minDate: 0});
//    $( "#endDate" ).datepicker({ minDate: 0});
  $('#startDate').change(function () {
      selectedToDate = $('#startDate').val();
      var year = selectedToDate.substring(6, 10);
      var day = selectedToDate.substring(3, 5);
      var month = selectedToDate.substring(0, 2);
      $('#endDate').datepicker("option", "minDate", new Date(year, month-1, day));
       actualToDate = new Date(selectFromDate);
        actualToDate.setHours(23);
        actualToDate.setMinutes(59);
        actualToDate.setSeconds(59);
  });
    
    
    if (pageType == "edit") {
        $.getScript( '/static/js/timepicker.js', function( data, textStatus, jqxhr ) {
        console.log("jobname",vm.NFCTagId);
        var element = document.getElementById('minUsers');
        element.value = vm.UserNumber;
        var logUser =document.getElementById("log");
        logUser.value=vm.Log;
        if(vm.LoginType =="NFC" ){
            document.getElementById("loginType1").checked = true;
            var div = document.getElementById('nfcTagId');
            div.style.visibility = 'visible';
            div.style.display ='inline';
            document.getElementById("nfcTagId").value =vm.NFCTagId;
        }else{
            document.getElementById("loginType2").checked = true;
        }
        loginTypeForEdit = vm.LoginType;
        var selectArray =  vm.ContactNameToEdit;
        $("#contactId").val(selectArray);
        var selectArrayForGroup = vm.GroupMembersAndUserToEdit;
        $("#userOrGroup").val(selectArrayForGroup);
        document.getElementById("jobName").value = vm.JobName;
        document.getElementById("taskName").value = vm.TaskName;
        document.getElementById("taskLocation").value = vm.TaskLocation;
        document.getElementById("startDate").value = vm.StartDate;
        document.getElementById("endDate").value = vm.EndDate;
        document.getElementById("taskDescription").value = vm.TaskDescription;
        document.getElementById("taskLocation").value =vm.TaskLocation
        document.getElementById("addFitToWorkValue").value = vm.FitToWork[0];
        var dynamicTextBox= "";
        for (var i = 1; i < vm.FitToWork.length; i++) {
            dynamicTextBox+= '<div class="plus"><input class="form-control"  name = "DynamicTextBox"  id=  "DynamicTextBox"  type="text" value = "' + vm.FitToWork[i] + '" />&nbsp;' +
            '<button    class="delete-decl" >+</button></div>';
        }
        $("#TextBoxContainer").append(dynamicTextBox);
        document.getElementById("startTime").value = vm.StartTime;
        document.getElementById("endTime").value = vm.EndTime;
        document.getElementById("taskHead").innerHTML = "Edit Task";
        $("body").on("click", ".delete-decl", function () {
            $(this).closest("div").remove();
        });
        minUserForTaskEdit = vm.UsersToEdit.length;
        if(vm.FitToWorkCheck =="EachTime") {
            document.getElementById("fitToWorkCheck").checked = true;
        }
        if(vm.WorkTime.length !=0){
            console.log("inside work");
            document.getElementById("workExplosure").checked = true;
            var div = document.getElementById('work');
            div.style.visibility = 'visible';
            div.style.display ='inline';
            document.getElementById("workTime").value =vm.WorkTime[0];
            document.getElementById("breakTime").value =vm.BreakTime[0];
            var DynamicExposureTextBox ="";
            for (var i=1; i<vm.WorkTime.length; i++){
                DynamicExposureTextBox+=        '<div class="exposureId"> <label for="workExplosureText" class="">Break Time</label>'+
                    '<input type="text"    placeholder="12:00" data-timepicker id="breakTime" name="breakTime" size="5" value="'+ vm.BreakTime[i] +'">'+ 'After'+'<input type="text"    placeholder="12:00" data-timepicker id="workTime" name="workTime" size="5" value="'+ vm.WorkTime[i] +'" >'+'<img  id="exposureDelete" src="/static/images/exposureCancel.jpg" width="20" height="20" style= "float:right; margin-top:0em; margin-right:0em;"  class="delete-exposure" /></div>';
            }
            $("#exposureTextBoxAppend").append(DynamicExposureTextBox);
             
        }
            });
    }
    
    
   
   
    //function for getting textbox dynamically
    $("#btnAdd").bind("click", function () {
        var div = $("<div class='plus'/>");
        div.html(GetDynamicTextBox(""));
        $("#TextBoxContainer").append(div);
    });
    $("body").on("click", ".delete-decl", function () {
        $(this).closest("div").remove();
    });
});
function GetDynamicTextBox(value) {
    return ' <input class="form-control"  name = "DynamicTextBox"  id=  "DynamicTextBox"  type="text" value = "" />&nbsp;' +
            '<button    class="delete-decl">+</button>'
}
 

//function for getting exposure dynamically
$("#btnAddForExposure").bind("click", function () {
    $.getScript( '/static/js/timepicker.js', function( data, textStatus, jqxhr ) {
        var div = $("<div class='exposureId'/>");
        div.html(GetDynamicTextBoxForExposure(""));
        $("#exposureTextBoxAppend").append(div);
    } );
});
//$("#exposureDelete").bind("click", function () {
    $("body").on("click", ".delete-exposure", function () {
        $(this).closest("div").remove();
    });
function GetDynamicTextBoxForExposure(value) {
    return ' <label for="workExplosureText" class="">Break Time</label>'+
        '<input type="text"    placeholder="12:00" data-timepicker id="breakTime" name="breakTime" size="5">'+ 'After'+'<input type="text"    placeholder="12:00" data-timepicker id="workTime" name="workTime" size="5" >'+'<img  id="exposureDelete" src="/static/images/exposureCancel.jpg"  class="delete-exposure" />'
}

//function to load add task
var addItem = $('<span>+</span>');
addItem.click(function() {
    window.location = "/"  +  companyTeamName +  "/task/add";
});

$().ready(function() {
    var loginTypeRadio = "";
 $("input[type='radio']").change(function(){
     loginTypeRadio = $('.radio-inline:checked').val();
       if (loginTypeRadio =="NFC"){
           var div = document.getElementById('nfcTagId');
           div.style.visibility = 'visible';
           div.style.display ='inline';
           
       }else{
           var div = document.getElementById('nfcTagId');
           div.style.visibility = 'hidden';
           div.style.display ='none'
       }
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
    
    //function to show break time when checkbox is clicked
  $('#workExplosure').click(function () {
      if ($(this).is(":checked")) {
          var div = document.getElementById('work');
          div.style.visibility = 'visible';
          div.style.display ='inline';
      }else {
          var div = document.getElementById('work');
          div.style.visibility = 'hidden';
          div.style.display ='none';
      }
  });
    
    
    /*Function will ceck if the selected value is a group name, and if so 
    function will auto select all users in that group*/
    $("#userOrGroup").on('change', function(evt, params) {
        console.log("inside group1");
        var tempArray = $(this).val();
        var clickedOption = "";
        console.log("selected array",selectedUserArray);
        console.log("temp array",tempArray);
        if (selectedUserArray.length < tempArray.length) { // for selection
            for (var i = 0; i < tempArray.length; i++) {
                if (selectedUserArray.indexOf(tempArray[i]) == -1) {
                    console.log("clicked");
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
                        console.log("values of temp array",tempArray);
                        $("#userOrGroup").val(tempArray);
                    }
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
//            for (var i = 0; i < vm.GroupMembers[i].length; i++) {
//                if (vm.GroupMembers[i][0] == clickedOption) {
//                    var memberLength = vm.GroupMembers[i].length;
//                    for (var j = 1; j < memberLength; j++) {
//                        var userIndex = tempArray.indexOf(vm.GroupMembers[i][j]);
//                        if (userIndex != -1) {
//                            tempArray.splice(userIndex, 1);
//                        }
//                        $("#userOrGroup").val(tempArray);
//                    }
//                    // Removing group from group array for validating min. no. of users
//                    var deleteGroupKeyIndex = selectedGroupArray.indexOf(clickedOption);
//                    selectedGroupArray.splice(deleteGroupKeyIndex, 1);
//                }
//            }
            selectedUserArray = tempArray;
        }
    });
     
       
    $("#taskDoneForm").validate({
        rules: {
            taskName: "required",
            loginType: "required",
            startDate :"required",
            endDate :"required",
            taskDescription:"required"
            
        },
        submitHandler: function() {
            
            //code for date and time conversion
            var startDate = new Date($("#startDate").val());
            
            var startTime =  document.getElementById("startTime").value;
            var endDate = new Date($("#endDate").val());
            var endTime =  document.getElementById("endTime").value;
            
            var exposureHour ="";
            var exposureMinute ="";
            var TotalBreakTime ="";
            var exposureWorkHour ="";
            var exposureWorkMinute ="";
            var TotalWorkTime ="";
            

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
            endDateToCompare = document.getElementById("endDate").value
            //check minimum number of users during editing
//             minUserForTask =selectedUserArray.length - selectedGroupArray.length;
//            if(minUserForTask == 0)
//            {
//                minUserForTask = minUserForTaskEdit;
//            }
//            else {
//                minUserForTask = minUserForTask;
//            }
            //check the login type during editing
            if(loginTypeRadio.length ==0)
            {
                loginTypeRadio = loginTypeForEdit;
            } else {
                loginTypeRadio = loginTypeRadio;
            }
//            if (minUserForTask >= minUsers) {
//                if(loginTypeRadio.length != 0)
//                    {
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
                                      var fitToWorkValue = document.getElementById("addFitToWorkValue").value;
                                    
                                      if(fitToWorkValue.length !=0)
                                          {
                                              
                                              fitToWorkFromDynamicTextBox.push(fitToWorkValue);
                                          }
                                      $("input[name=DynamicTextBox]").each(function () {
                                          
                                          if($(this).val().length !=0){
                                              fitToWorkFromDynamicTextBox.push($(this).val())
                                          }
                                      });
                                      if (document.getElementById('jobName').length !=0)
                                          {
                                              getJobAndCustomer(); 
                                          }
                                       
                                      // function to get values of exposure dynamic text box
                                      $("input[name=breakTime]").each(function () {
                                          
                                          if($(this).val().length !=0){
                                              exposureTimeArray = $(this).val().split(':');
                                              exposureHour = parseInt(exposureTimeArray[0]);
                                              exposureMinute = parseInt(exposureTimeArray[1]);
                                              TotalBreakTime =exposureMinute+(exposureHour*60);
                                              exposureSlice.push(TotalBreakTime);
                                          }
                                      });
                                      
                                      $("input[name=workTime]").each(function () {
                                          
                                          if($(this).val().length !=0){
                                              exposureWorkTimeArray = $(this).val().split(':');
                                              exposureWorkHour = parseInt(exposureWorkTimeArray[0]);
                                              exposureWorkMinute = parseInt(exposureWorkTimeArray[1]);
                                              TotalWorkTime =exposureWorkMinute+(exposureWorkHour*60);
                                              exposureWorkSlice.push(TotalWorkTime);
                                          }
                                      });
                                      
                                      
                                      //function to get fit to work 
                                      var chkPassport = document.getElementById("fitToWorkCheck");
                                      if (chkPassport.checked) {
                                          fitToWorkCheck ="EachTime";
                                      }else {
                                          fitToWorkCheck ="OnceADay";
                                      }
                                      var formData = $("#taskDoneForm").serialize() + "&loginType=" + loginTypeRadio + "&customerName=" + customerName + "&jobId=" + jobId +"&addFitToWork=" + fitToWorkFromDynamicTextBox +"&latitude=" +  mapLatitude +"&longitude=" +  mapLongitude +"&startDateFomJs="+ startDateOfTask +"&endDateFromJs="+ endDateOfTask+"&fitToWorkCheck="+ fitToWorkCheck+"&exposureBreakTime="+ exposureSlice+"&exposureWorkTime="+ exposureWorkSlice;
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
                          
//                    }
//            else {
//                $("#loginTypeValidationError").css({"color": "red", "font-size": "15px"});
//                $("#loginTypeValidationError").html("please select a login type.").show();
//            }
//                
//            }
//            else {
//                $("#minUserValidationError").css({"color": "red", "font-size": "15px"});
//                $("#minUserValidationError").html("More users need to start this Task.").show();
//                }
        }
    });
    $("#cancel").click(function() {
        window.location = '/' + companyTeamName + '/task';
    });
});