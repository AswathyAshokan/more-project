/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..

console.log("jjjj",vm.FitToWorkArray);
console.log("gsgsgsgs",vm.GroupMembers);
console.log("job name",vm.WorkLocationArray);
console.log("job name from url",vm.JobNameFormUrl);
document.getElementById("task").className += " active";
var date = new Date();
var datum = (Date.parse(date))/1000;
console.log("greenn",datum);
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
 var repeat= "";
var fitWork= "";
var jobNameWithUrl ="";
var customerNameWithUrl ="";
var contactName =[];
var contactId =[];
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
        document.getElementById("ExplosureDiv").style.display = "none";
        $("#contactDiv").hide();
    } else if(vm.CompanyPlan == "campus"){
        
        document.getElementById("jobNamelabel").style.display = "none";
        document.getElementById("workExplosureLabel").style.display = "none";
        document.getElementById("minUsersLabel").style.display = "none";
        document.getElementById("minUsers").style.display = "none";
        document.getElementById("jobName").style.display = "none";
        document.getElementById("ExplosureDiv").style.display = "none";
        document.getElementById("contactIdLabel").style.display = "none";
         $("#contactDiv").hide();
    }else{
         document.getElementById("minUsers").style.display = "block";
         document.getElementById("jobName").style.display = "block";
         document.getElementById("ExplosureDiv").style.display = "block";
          $("#contactDiv").show();
    }
    
    /*var date = new Date();
    var datum = (Date.parse(date))/1000;*/
    
    
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
    
    //functionfor automatically fill task location textbox
    
    $('#taskLocation').keyup(function(){
        var valThis = $(this).val().toLowerCase();
        if(valThis == ""){
            console.log("nulllll");
            //$('.navList > li').show();  
        } else {
            /*for (i=0;i<vm.WorkLocationArray.length;i++){
                var text = vm.WorkLocationArray[i];
                (text.indexOf(valThis) >= 0) ? $(this).show() : $(this).hide();
                
            }*/
            console.log("not nullll");
            /*$('.navList > li').each(function(){
                var text = $(this).text().toLowerCase();
                (text.indexOf(valThis) >= 0) ? $(this).show() : $(this).hide();
            });*/
        };
    });
    
    
    
    //function to setting jobna me when loaded add and continue button
    if (vm.JobNameFormUrl.length !=0){
        document.getElementById("jobName").value = vm.JobNameFormUrl;
        for (var i = 0; i < vm.ContactUser.length; i++) {
            for (var j=0; j<vm.ContactUser[i].length ;j++){
                for ( var k=0;k<vm.ContactUser[i][j].CustomerName.length;k++){
                    if (vm.ContactUser[i][j].CustomerName[k] ==vm.CustomerNameFormUrl){
                        contactName.push(vm.ContactUser[i][j].ContactName);
                        contactId.push(vm.ContactUser[i][j].ContactId);
                    }
                }
            }
        }
        function removeOptions(selectbox)
        {
            var i;
            for(i = selectbox.options.length - 1 ; i >= 0 ; i--)
            {
                selectbox.remove(i);
            }
        }
        console.log("contact name",contactName)
        removeOptions(document.getElementById("contactId"));
        var sel = document.getElementById('contactId');
        for(var i = 0; i < contactName.length; i++) {
            var opt = document.createElement('option');
            opt.innerHTML = contactName[i];
            opt.value = contactId[i];
            sel.appendChild(opt);
        }
        
    }
    if (pageType == "edit") {
        document.getElementById("saveAndContinue").disabled = true;
        document.title = 'Edit Task'
        console.log("log",vm.Log);
        var element = document.getElementById('logInMinutes');
        element.value = vm.Log;
        document.getElementById("jobName").value = vm.JobName;
        document.getElementById("taskName").value = vm.TaskName;
        document.getElementById("taskLocation").value = vm.TaskLocation;
        document.getElementById("startDate").value = vm.StartDate;
        document.getElementById("endDate").value = vm.EndDate;
        document.getElementById("taskDescription").value = vm.TaskDescription;
        document.getElementById("taskLocation").value =vm.TaskLocation;
        var fitToWorkName = vm.FitToWorkName;
        fitWork =vm.FitToWorkName;
        if (fitToWorkName.length !=0){
            $('#TaskFitToWork option:contains(' + fitToWorkName + ')').prop({selected: true});
        }
        var selectArrayForGroup = vm.GroupMembersAndUserToEdit;
        $("#userOrGroup").val(selectArrayForGroup);
        document.getElementById("startTime").value = vm.StartTime;
        document.getElementById("endTime").value = vm.EndTime;
        var element = document.getElementById('minUsers');
        element.value = vm.UserNumber;
        if(vm.LoginType =="NFC" ){
            document.getElementById("loginType1").checked = true;
            var div = document.getElementById('nfcTagId');
            div.style.visibility = 'visible';
            div.style.display ='inline';
            console.log("nfc",vm.NFCTagId);
            document.getElementById("nfcTagForTask").value =vm.NFCTagId;
            
        }else{
            document.getElementById("loginType2").checked = true;
        }
        loginTypeForEdit = vm.LoginType;
        for (var i = 0; i < vm.ContactUser.length; i++) {
            for (var j=0; j<vm.ContactUser[i].length ;j++){
                for ( var k=0;k<vm.ContactUser[i][j].CustomerName.length;k++){
                    if (vm.ContactUser[i][j].CustomerName[k] ==vm.CustomerNameToEdit){
                        contactName.push(vm.ContactUser[i][j].ContactName);
                        contactId.push(vm.ContactUser[i][j].ContactId);
                    }
                }
            }
        }
        var sel = document.getElementById('contactId');
        for(var i = 0; i < contactName.length; i++) {
            var opt = document.createElement('option');
            opt.innerHTML = contactName[i];
            opt.value = contactId[i];
            sel.appendChild(opt);
        }
        var optionValues =[];
        $('#contactId option').each(function(){
            if($.inArray(this.value, optionValues) >-1){
                $(this).remove()
            }else{
                optionValues.push(this.value);
            }
        });
        var eid = document.getElementById('contactId');
        for (var i = 0; i < eid.options.length; ++i) {
            for (var j=0;j<vm.ContactNameToEdit.length;j++){
                if (eid.options[i].text === vm.ContactNameKeyToEdit[j]){
                    console.log("log1");
                    eid.options[i].selected = true;
                 }
            }
        }
        document.getElementById("taskHead").innerHTML = "Edit Task";
        $("body").on("click", ".delete-decl", function () {
            $(this).closest("div").remove();
        });
        minUserForTaskEdit = vm.UsersToEdit.length;
        if(vm.FitToWorkCheck =="EachTime") {
            document.getElementById("fitToWorkCheck").checked = true;
        }
        console.log("ffff",vm.BreakTime[0]);
        console.log("ggg",vm.WorkTime[0]);
        if(vm.WorkTime.length !=0){
            document.getElementById("workExplosure").checked = true;
            var div = document.getElementById('work');
            div.style.visibility = 'visible';
            div.style.display ='inline';
            document.getElementById("workTime").value =vm.WorkTime[0];
            document.getElementById("breakTime").value =vm.BreakTime[0];
            var DynamicExposureTextBox ="";
            for (var i=1; i<vm.WorkTime.length; i++){
                DynamicExposureTextBox+=        '<div class="exposureId"> <label for="workExplosureText" class="">Break Time</label>'+
                    '<input type="text"    placeholder="12:00"  id="breakTime" name="breakTime" size="5" value="'+ vm.BreakTime[i] +'">'+ '<label>'+'After'+'</label>'+'<input type="text"    placeholder="12:00"  id="workTime" name="workTime" size="5" value="'+ vm.WorkTime[i] +'" >'+'<img  id="exposureDelete" src="/static/images/exposureCancel.jpg" width="20" height="20" style= "float:right; margin-top:0em; margin-right:0em;"  class="delete-exposure" /></div>';
            }
            $("#exposureTextBoxAppend").append(DynamicExposureTextBox);
        }
    }
});

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
        '<input type="text"    placeholder="12:00" id="breakTime" name="breakTime" size="5">'+ '<label>'+'After'+'</label>'+'<input type="text"    placeholder="12:00"  id="workTime" name="workTime" size="5" >'+'<img  id="exposureDelete" src="/static/images/exposureCancel.jpg"  class="delete-exposure" />'
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
        contactName = [];
        contactId = [];
        var job = $("#jobName option:selected").val() + " (";
        jobNameWithUrl =$("#jobName option:selected").val();
       
        var jobAndCustomer = $("#jobName option:selected").text();
        var tempName = jobAndCustomer.replace(job, '');
        customerName = tempName.replace(')', '');
        customerNameWithUrl =tempName.replace(')', '');
       for (var i = 0; i < vm.ContactUser.length; i++) {
           for (var j=0; j<vm.ContactUser[i].length ;j++){
               for ( var k=0;k<vm.ContactUser[i][j].CustomerName.length;k++){
                   if (vm.ContactUser[i][j].CustomerName[k] ==customerName){
                       console.log("jjjj",vm.ContactUser[i][j].ContactName);
                       contactName.push(vm.ContactUser[i][j].ContactName);
                       contactId.push(vm.ContactUser[i][j].ContactId);
                   }
               }
           }
       }
        removeOptions(document.getElementById("contactId"));
        var sel = document.getElementById('contactId');
        for(var i = 0; i < contactName.length; i++) {
            var opt = document.createElement('option');
            opt.innerHTML = contactName[i];
            opt.value = contactId[i];
            sel.appendChild(opt);
        }
        if ($("#jobName option:selected").val()== ""){
            var sel = document.getElementById('contactId');
            for(var i = 0; i < vm.ContactNameArray.length; i++) {
                var opt = document.createElement('option');
                opt.innerHTML = vm.ContactNameArray[i];
                opt.value = vm.ContactKey[i];
                sel.appendChild(opt);
            }
        }
        var jobDropdownId = document.getElementById("jobName");
        jobId = jobDropdownId.options[jobDropdownId.selectedIndex].id;
    }
    function removeOptions(selectbox)
    {
        var i;
        for(i = selectbox.options.length - 1 ; i >= 0 ; i--)
        {
            selectbox.remove(i);
        }
    }
//using the function:

    //getting instructions of fit to work
    getInstructions =function(){
        var doc = document.getElementById("TaskFitToWork");
        
        if(doc.length !=0){
            fitWork =doc.options[doc.selectedIndex].value;
        }
        
        
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
    var groupKeyArray = [];
    $("#userOrGroup").on('change', function(evt, params) {
        console.log("inside group1");
        var tempArray = $(this).val();
        var clickedOption = "";
        console.log("array length",tempArray.length)
       
        
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
                    groupKeyArray.push(clickedOption)
                    tempArray =[];
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
        console.log("group array",groupKeyArray);
        console.log("user array",selectedUserArray);
    });
    console.log("lattitude kkkkk",document.getElementById('latitudeId').value);
    $("#saveButton").click(function() {
            $("#taskDoneForm").validate({
                rules: {
                    taskName : "required",
                    loginType : "required",
                    startDate : "required",
                    endDate : "required",
                    taskDescription : "required",
                    taskLocation : "required",
                    startTime : "required",
                    endTime : "required"
                },
                messages: {
                    startTime:{
                        required: "time required"
                    },
                    endTime:{
                        required: "time required"
                    },
                    taskName:{
                        required: "task name required"
                    },
                    loginType:{
                        required: "login type name required"
                    },
                    startDate:{
                        required: "date required"
                    },
                    endDate:{
                        required: "date required"
                    },
                    taskLocation:{
                        required: "task location required"
                    },
                    taskDescription:{
                        required: "task description required"
                    },
                    userOrGroup:{
                        required: "select user/group"
                    },
                },
                submitHandler: function() {
                    var nfcTagId =  document.getElementById("nfcTagForTask").value;
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
//                                      $("input[name=DynamicTextBox]").each(function () {
//                                          
//                                          if($(this).val().length !=0){
//                                              fitToWorkFromDynamicTextBox.push($(this).val())
//                                          }
//                                      });
                                      if (document.getElementById('jobName').length !=0)
                                          {
                                              var job = $("#jobName option:selected").val() + " (";
                                              
                                              var jobAndCustomer = $("#jobName option:selected").text();
                                              var tempName = jobAndCustomer.replace(job, '');
                                              customerName = tempName.replace(')', ''); 
                                              var jobDropdownId = document.getElementById("jobName");
                                              jobId = jobDropdownId.options[jobDropdownId.selectedIndex].id;
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
                                      var formData = $("#taskDoneForm").serialize() + "&loginType=" + loginTypeRadio + "&customerName=" + customerName + "&jobId=" + jobId +"&latitude=" +  mapLatitude +"&longitude=" +  mapLongitude +"&startDateFomJs="+ startDateOfTask +"&endDateFromJs="+ endDateOfTask+"&fitToWorkCheck="+ fitToWorkCheck+"&exposureBreakTime="+ exposureSlice+"&exposureWorkTime="+ exposureWorkSlice+"&fitToWorkName="+ fitWork+"&dateOfCreation="+datum;
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
                            for(i = 0; i < groupKeyArray.length; i++) {
                                          formData = formData+"&groupArrayElement="+groupKeyArray[i];
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
                                    console.log("seleceeeeeee",selectedUserAndGroupName)
                            
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
        });
    $("#cancel").click(function() {
        window.location = '/' + companyTeamName + '/task';
    });
    $("#saveAndContinue").click(function() {
         $('#saveAndContinue').attr('type', 'submit');
        $('#saveButton').attr('type', 'button');
        console.log("inside save and continue");

         $("#taskDoneForm").validate({
        rules: {
            taskName : "required",
            loginType : "required",
            startDate : "required",
            endDate : "required",
            taskDescription : "required",
            taskLocation : "required",
            startTime : "required",
            endTime : "required"
        },
         messages: {
             startTime:{
                 required: "time required"
             },
             endTime:{
                 required: "time required"
             },
             taskName:{
                 required: "task name required"
             },
             loginType:{
                 required: "login type name required"
             },
             startDate:{
                 required: "date required"
             },
             endDate:{
                 required: "date required"
             },
             taskLocation:{
                 required: "task location required"
             },
              taskDescription:{
                 required: "task description required"
             },
              userOrGroup:{
                 required: "select user/group"
             },
         },
        submitHandler: function() {
            
             var nfcTagId =  document.getElementById("nfcTagForTask").value;
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
                                       $("#saveAndContinue").attr('disabled', true);
                                      var taskId=vm.TaskId;
                                      var jobnew = $("#jobName option:selected").val()
                                      if ($("#jobName ")[0].selectedIndex <= 0) {
                                          document.getElementById('jobName').innerHTML = "";
                                      }
                                      //get all values of fit to work
                                      
                                      var values = "";
//                                      $("input[name=DynamicTextBox]").each(function () {
//                                          
//                                          if($(this).val().length !=0){
//                                              fitToWorkFromDynamicTextBox.push($(this).val())
//                                          }
//                                      });
                                      if (document.getElementById('jobName').length !=0)
                                          {
                                              var job = $("#jobName option:selected").val() + " (";
                                              var jobAndCustomer = $("#jobName option:selected").text();
                                              var tempName = jobAndCustomer.replace(job, '');
                                              customerName = tempName.replace(')', ''); 
                                              var jobDropdownId = document.getElementById("jobName");
                                              jobId = jobDropdownId.options[jobDropdownId.selectedIndex].id;
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
                                      var formData = $("#taskDoneForm").serialize() + "&loginType=" + loginTypeRadio + "&customerName=" + customerName + "&jobId=" + jobId +"&latitude=" +  mapLatitude +"&longitude=" +  mapLongitude +"&startDateFomJs="+ startDateOfTask +"&endDateFromJs="+ endDateOfTask+"&fitToWorkCheck="+ fitToWorkCheck+"&exposureBreakTime="+ exposureSlice+"&exposureWorkTime="+ exposureWorkSlice+"&fitToWorkName="+ fitWork+"&dateOfCreation="+datum;
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
                            for(i = 0; i < groupKeyArray.length; i++) {
                                          formData = formData+"&groupArrayElement="+groupKeyArray[i];
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
                                        if (jobNameWithUrl.length ==0){
                                            jobNameWithUrl="unDefined"
                                        }
                                        if (customerNameWithUrl.length ==0){
                                            customerNameWithUrl ="unDefined"
                                        }
                                    console.log("seleceeeeeee",selectedUserAndGroupName)
                            
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
                                                      $("#saveAndContinue").attr('disabled', false);
                                                  }
                                              },
                                              error: function (request,status, error) {
                                                  console.log(error);
                                              }
                                          });
                                      } else {
                                          $.ajax({
                                              url:'/'+ companyTeamName + '/task/add/'+jobNameWithUrl+'/'+customerNameWithUrl,
                                              type: 'post',
                                              datatype: 'json',
                                              data: formData,
                                              success : function(response) {
                                                  if (response == "true" ) {
                                                      window.location = '/' + companyTeamName + '/task/add/'+jobNameWithUrl+'/'+customerNameWithUrl;
                                                  } else {
                                                      $("#saveAndContinue").attr('disabled', false);
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

         
    });
});