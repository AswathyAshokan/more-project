console.log("viwe model value",vm);
console.log("end time",vm.DailyEndTime);
var companyTeamName = vm.CompanyTeamName;
 var selectedUserArray = [];
var startDateInUnix;
var endDateInUnix;
var dailyStartTimeUnix;
var dailyEndTimeUnix;
var taskWorkLocation = [];
var taskLocationCondition="";
var  startDateString ;

$(document).ready(function() {
    // contains all selected users and groups
    if(vm.PageType == "edit"){ 
        var selectArray =[];
        selectArray = vm.UsersKey;
        $("#usersAndGroupId").val(selectArray);
        startDateInUnix = vm.StartDate
        var dateFromDb = parseInt(startDateInUnix)
        var d = new Date(dateFromDb * 1000);
        var dd = d.getDate();
        var mm = d.getMonth() + 1; //January is 0!
        var yyyy = d.getFullYear();
        if (dd < 10) {
                dd = '0' + dd;
            }
        if (mm < 10) {
            mm = '0' + mm;
        }
        var localDate = (mm + '/' + dd + '/' + yyyy);
        
        endDateInUnix = vm.EndDate
        var endDateFromDb = parseInt(endDateInUnix)
        var end = new Date(endDateFromDb * 1000);
        var enda = end.getDate();
        var endmm = end.getMonth() + 1; //January is 0!
        var endyyyy = end.getFullYear();if (enda < 10) {
            enda = '0' + enda;
        }
        if (endmm < 10) {
            endmm = '0' + endmm;
        }
        var localEndDate = (endmm + '/' + enda + '/' + endyyyy);
        
        document.getElementById("taskLocation").value = vm.WorkLocation;
        document.getElementById("startDate").value = localDate;
        document.getElementById("endDate").value = localEndDate;
        document.getElementById("dailyStartTime").value = vm.DailyStartTime;
        document.getElementById("dailyEndTime").value = vm.DailyEndTime;
        document.getElementById("latitudeId").value = vm.LatitudeForEditing;
        document.getElementById("longitudeId").value = vm.LongitudeForEditing;
        document.getElementById("workLocationId").innerHTML = "Edit WorkLocation";//for display heading of each webpage
        var selectedGroupArray = [];
        var groupKeyArray = [];
        for(var i=0;i<vm.UsersKey.length;i++){
            selectedUserArray.push(vm.UsersKey[i]);
        }
        console.log("selectedUserArray",selectedUserArray);
    }
    
    
    var selectedGroupArray = [];
    var groupKeyArray = [];
    $("#usersAndGroupId").on('change', function(evt, params) {
        var tempArray = $(this).val();
        var clickedOption = "";
        console.log("array length",tempArray);
        if (selectedUserArray.length < tempArray.length) { // for selection
            for (var i = 0; i < tempArray.length; i++) {
                if (selectedUserArray.indexOf(tempArray[i]) == -1) {
                    console.log("clicked");
                    clickedOption = tempArray[i];
                }
            }
            if (vm.GroupMembers !=null){
                for (var i = 0; i < vm.GroupMembers.length; i++) {
                    if (vm.GroupMembers[i][0] == clickedOption) {
                    var memberLength = vm.GroupMembers[i].length;
                    groupKeyArray.push(clickedOption)
                    tempArray =[];
                    for (var j = 1; j < memberLength; j++) {
                        if (tempArray.indexOf(vm.GroupMembers[i][j]) == -1) {
                            tempArray.push(vm.GroupMembers[i][j])
                        }
                        $("#usersAndGroupId").val(tempArray);
                    }
                    selectedGroupArray.push(clickedOption);
                }
                }
            }
            selectedUserArray = tempArray;
        } else if (selectedUserArray.length > tempArray.length) { // for deselection
            for (var i = 0; i < selectedUserArray.length; i++) {
                if (tempArray.indexOf(selectedUserArray[i]) == -1) {
                    clickedOption = selectedUserArray[i];
                    
                }
            }
            selectedUserArray = tempArray;
        }
        console.log("group array",groupKeyArray);
        console.log("user array",selectedUserArray);
    });
    
    $("#workLocationForm").validate({
        rules: {
           // usersAndGroupId:"required",
            taskLocation : "required",
            startDate:"required",
            endDate:"required",
            dailyStartTime:"required",
            dailyEndTime:"required"
        },
        messages: {
            usersAndGroupId: "Please select user or group",
            taskLocation:"please fill this column",
        },
        submitHandler: function(){//to pass all data of a form serial
            
            
            if(vm.DateValues != null){
                console.log("cp1");
                for ( var x=0;x<vm.DateValues.length;x++){
                    for( var y=0;y<selectedUserArray.length;y++){
                         console.log("cp2",selectedUserArray[y],"cp3",vm.DateValues[x][0]);
                        if (vm.DateValues[x][0] == selectedUserArray[y]){
                            console.log(" both id of users ARE equal");
                            var utcTime = vm.DateValues[x][1];
                            var dateFromDb = parseInt(utcTime);
                            var d = new Date(dateFromDb * 1000);
                            var dd = d.getDate();
                            var mm = d.getMonth() + 1; //January is 0!
                            var yyyy = d.getFullYear();
                            var HH = d.getHours();
                            var min = d.getMinutes();
                            var sec = d.getSeconds();
                            if (dd < 10) {
                                dd = '0' + dd;
                            }
                            if (mm < 10) {
                                mm = '0' + mm;
                            }
                            if (HH < 10) {
                                HH = '0' + HH;
                            }
                            if (min < 10) {
                                min = '0' + min;
                            }
                            if (sec < 10) {
                                sec = '0' + sec;
                            }
                            var workStartDateFromDb = (mm + '/' + dd + '/' + yyyy);
                            var utcTime =vm.DateValues[x][2];
                            var dateFromDb = parseInt(utcTime)
                            var d = new Date(dateFromDb * 1000);
                            var dd = d.getDate();
                            var mm = d.getMonth() + 1; //January is 0!
                            var yyyy = d.getFullYear();
                            var HH = d.getHours();
                            var min = d.getMinutes();
                            var sec = d.getSeconds();
                            if (dd < 10) {
                                dd = '0' + dd;
                            }
                            if (mm < 10) {
                                mm = '0' + mm;
                            }
                            if (HH < 10) {
                                HH = '0' + HH;
                            }
                            if (min < 10) {
                                min = '0' + min;
                            }
                            if (sec < 10) {
                                sec = '0' + sec;
                            }
                            var workEndDateFromDb = (mm + '/' + dd + '/' + yyyy);
                            /*var newStartDate =document.getElementById("startDate").value 
                            var newEndDate =document.getElementById("endDate").value*/
                            
                            
                            
                             var StartDateOfTask = document.getElementById("startDate").value 
                                    var EndDateOfTask = document.getElementById("endDate").value
                                    var workStartDate1 = workStartDateFromDb.split("/");
                                    var workEndDate1 = workEndDateFromDb.split("/");
                                    var StartDateOfTask1 = StartDateOfTask.split("/");
                                    var EndDateOfTask1 = EndDateOfTask.split("/");
                                    var from = new Date(workStartDate1[2], parseInt(workStartDate1[1])-1, workStartDate1[0]);  // -1 because months are from 0 to 11
                                    var to   = new Date(workEndDate1[2], parseInt(workEndDate1[1])-1, workEndDate1[0]);
                                    var StartDateOfTaskCheck = new Date(StartDateOfTask1[2], parseInt(StartDateOfTask1[1])-1, StartDateOfTask1[0]);
                                    var EndDateOfTaskCheck = new Date(EndDateOfTask1[2], parseInt(EndDateOfTask1[1])-1, EndDateOfTask1[0]);
                                    if (StartDateOfTaskCheck >= from && StartDateOfTaskCheck <= to && EndDateOfTaskCheck >= from && EndDateOfTaskCheck <= to){
                                        console.log("condition is true")
                                        taskWorkLocation.push("true")
                                    }
                        } else{
                             taskWorkLocation.push("true");
                        }
                    }
                }
            }else{
                for( var z=0;z<selectedUserArray.length;z++){
                    taskWorkLocation.push("true");
                }
            }
            
            if (taskWorkLocation.length ==selectedUserArray.length&&taskWorkLocation.length >0){
                taskLocationCondition="true"
            }else{
                taskLocationCondition="false"
            } 
            
            
            
            var starDateString = document.getElementById('startDate').value;
            var endDateString = document.getElementById('endDate').value;
            $("#saveButton").attr('disabled', true);
            
            var startdatum = Date.parse(starDateString)/1000;
            var endDatum = Date.parse(endDateString)/1000;
           
            var startDateInDate = new Date(starDateString);
            var dailyStartTime = document.getElementById('dailyStartTime').value;
            
            var endDateInDate = new Date(endDateString);
            var dailyEndTime = document.getElementById('dailyEndTime').value;
            
            startTimeArray = dailyStartTime.split(':');
            startHour = parseInt(startTimeArray[0]);
            startMin = parseInt(startTimeArray[1]);
            startDateInDate.setHours(startHour);
            startDateInDate.setMinutes(startMin);
            endTimeArray = dailyEndTime.split(':');
            endHour = parseInt(endTimeArray[0]);
            endMin = parseInt(endTimeArray[1]);
            endDateInDate.setHours(endHour);
            endDateInDate.setMinutes(endMin);
            //function to convert  date to mm/dd/yyyy format
            function formatDate(d){
                function addZero(n){
                    return n < 10 ? '0' + n : '' + n;
                }
                return addZero(d.getMonth()+1)+"/"+ addZero(d.getDate()) + "/" + d.getFullYear() + " " + 
            addZero(d.getHours()) + ":" + addZero(d.getMinutes());
            }
            startDateString = startDateInDate;
            var date = new Date(Date.parse(startDateString));
            var startDateOfWork = formatDate(date);
            var endDateStringInUtc = endDateInDate;
            var endDateData = new Date(Date.parse(endDateStringInUtc));
            var endDateOfWork = formatDate(endDateData);
            
            console.log("localStartDate",starDateString);
            console.log("localEndDate",endDateString);
            
            var formData = $("#workLocationForm").serialize();
            //get the user's name corresponding to  keys selected from dropdownlist 
            formData = formData+"&startDateTimeStamp="+startdatum+"&endDateTimeStamp="+endDatum +"&dailyStartTimeString="+startDateOfWork+"&dailyEndTimeString="+endDateOfWork;
            
            var selectedUserAndGroupName = [];
              $("#usersAndGroupId option:selected").each(function () {
                  var $this = $(this);
                  if ($this.length) {
                      var selectedUserName = $this.text();
                      selectedUserAndGroupName.push( selectedUserName);
                  }
              });
              for(i = 0; i < selectedUserAndGroupName.length; i++) {
                  formData = formData+"&userAndGroupName="+selectedUserAndGroupName[i];
              }
            for(i = 0; i < groupKeyArray.length; i++) {
                formData = formData+"&groupArrayElement="+groupKeyArray[i];
            }
           // formData = formData+"&selectedUserNames="+selectedUserArray
            for(i = 0; i < selectedUserArray.length; i++) {
                formData = formData+"&selectedUserNames="+selectedUserArray[i];
            }
            //for checking worklocation for a user is unique
            console.log("vm.DateValues",vm.DateValues);
            console.log("selectedUserArray",selectedUserArray);
            
            if(taskLocationCondition=="true"){
                var ConcatinatedUser ;
                if (vm.PageType == "edit"){
                    for(i=0;i<vm.UsersKey.length;i++){
                        formData = formData+"&oldUsers="+vm.UsersKey[i];
                    }
                    var workLocationId =vm.WorkLogId  
                    $.ajax({
                        url:'/' + companyTeamName +'/worklocation/'+ workLocationId + '/edit',
                        type:'post',
                        datatype: 'json',
                        data: formData,
                        //call back or get response here
                        success : function(response){
                            if(response == "true"){
                                window.location='/' + companyTeamName +'/worklocation';
                            }else {
                                $("#saveButton").attr('disabled', false);
                            }
                        },
                        error: function (request,status, error) {
                        }
                    });
                } else {
                    $.ajax({
                        url:'/' + companyTeamName +'/worklocation/add',
                        type:'post',
                        datatype: 'json',
                        data: formData,
                        //call back or get response here
                        success : function(response){
                            console.log("response",response);
                            if(response == "true"){
                               window.location = '/'+companyTeamName+'/worklocation';
                            }
                            else{
                                 $("#saveButton").attr('disabled', false);
                            }
                        },
                        error: function (request,status, error) {
                        }
                    });
                    return false;
                }
            }else{
                $("#myModalForUniqueTest").modal();
                $("#cancelForCheckUnique").click(function(){
                    window.location = '/'+companyTeamName+'/worklocation';
                });
            }
            
        }
    });
     $("#cancel").click(function() {
            window.location = '/'+companyTeamName+'/worklocation';
    });
    
    
    
    if(vm.CompanyPlan == 'family' ){
        var parent = document.getElementById("menuItems");
        var contact = document.getElementById("contact");
        var job = document.getElementById("job");
        var crm = document.getElementById("crm");
        var leave = document.getElementById("leave");
        var timesheet  = document.getElementById("time-sheet");
        var consent = document.getElementById("consent")
        var workLocation = document.getElementById("WorkLocation")
        parent.removeChild(workLocation)
        parent.removeChild(timesheet);
        parent.removeChild(consent);
        parent.removeChild(leave);
        parent.removeChild(contact);
        parent.removeChild(job);
        parent.removeChild(crm);
    } else if(vm.CompanyPlan == 'campus'){
        var parent = document.getElementById("menuItems");
        var contact = document.getElementById("contact");
        var job = document.getElementById("job");
        var crm = document.getElementById("crm");
        var leave = document.getElementById("leave");
        var timesheet  = document.getElementById("time-sheet");
        var consent = document.getElementById("consent")
        var workLocation = document.getElementById("WorkLocation")
        parent.removeChild(workLocation)
        parent.removeChild(timesheet);
        parent.removeChild(consent);
        parent.removeChild(leave);
        parent.removeChild(contact);
        parent.removeChild(job);
        parent.removeChild(crm);
    }
    document.getElementById("username").textContent=vm.AdminFirstName;
    document.getElementById("imageId").src=vm.ProfilePicture;
    if (vm.ProfilePicture ==""){
        document.getElementById("imageId").src="/static/images/default.png"
    }
    if(vm.CompanyPlan == "family")
    {
        $('#planChange').attr('data-target','#family');
    } else if (vm.CompanyPlan == "campus") {
        $('#planChange').attr('data-target','#campus');
    }else if (vm.CompanyPlan == "business") {
        $('#planChange').attr('data-target','#business');
    }else if (vm.CompanyPlan == "businessPlus") {
        $('#planChange').attr('data-target','#business-plus');
    }
} );
