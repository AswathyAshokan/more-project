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
var count =0;
var returnString;
var dbString;
var idArray = [];
var successIdArray = [];
var uniqueIdArray = [];
var uniqueSucessArray =[];


$(document).ready(function() {
    
    function checkUserId(userId) {
       if(vm.DateValues !=null){
           for(var j=0 ;j<vm.DateValues.length;j++ ){
           if(vm.DateValues[j][0] !=userId ){
               console.log("in if func")
               returnString ="true"
           }else{
                 console.log("in else func")
               returnString ="false"
           }
       }
           return returnString
       }
   }
    // contains all selected users and groups
    /*function CheckUserIdFromDb(userId) {
        if(vm.DateValues !=null){
            for(var j=0 ;j<vm.DateValues.length;j++ ){
                console.log("hi all first me");
            if(vm.DateValues[j][0] != userId ){
                returnString ="true"
            }else{
                returnString ="false"
            }
        }
            return returnString
        }
    }
    function checkUserId(userId){
        if(uniqueSucessArray.length !=0){
            for(var i=0;i< uniqueSucessArray.length;i++){
                console.log("hi all second me",uniqueSucessArray,userId);
                if( uniqueSucessArray[i] != userId ){
                    dbString ="true"
                }
                else{
                    dbString = "false"
                }
            }
        
        }else{
            console.log("opopopopp");
            dbString ="true"
        }
        console.log("dbString",dbString)
        return dbString
        
    }*/
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
        
        // for checking the uniqueness Of work loccation
        
         if(vm.DateValues != null){
                if (vm.UsersKey.length !=0){
                    taskWorkLocation=[];
                    for ( var x=0;x<vm.DateValues.length;x++){
                        for( var y=0;y<vm.UsersKey.length;y++){
                            if (vm.DateValues[x][0] == vm.UsersKey[y]){
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
                                var StartDateOfTask = document.getElementById("startDate").value ;
                                var EndDateOfTask = document.getElementById("endDate").value;
                                console.log("StartDateOfTask",StartDateOfTask);
                                console.log("EndDateOfTask",EndDateOfTask);
                                console.log("workStartDateFromDb",workStartDateFromDb);
                                console.log("workEndDateFromDb",workEndDateFromDb);
                                var workStartDate1 = workStartDateFromDb.split("/");
                                var workEndDate1 = workEndDateFromDb.split("/");
                                var StartDateOfTask1 = StartDateOfTask.split("/");
                                var EndDateOfTask1 = EndDateOfTask.split("/");
                                console.log("workStartDate1",workStartDate1);
                                console.log("workEndDate1",workEndDate1);
                                console.log("StartDateOfTask1",StartDateOfTask1);
                                console.log("EndDateOfTask1",EndDateOfTask1)
                                var from = new Date(workStartDate1[2], parseInt(workStartDate1[1])-1, workStartDate1[0]);  // -1 because months are from 0 to 11
                                    var to   = new Date(workEndDate1[2], parseInt(workEndDate1[1])-1, workEndDate1[0]);
                                    var StartDateOfTaskCheck = new Date(StartDateOfTask1[2], parseInt(StartDateOfTask1[1])-1, StartDateOfTask1[0]);
                                    var EndDateOfTaskCheck = new Date(EndDateOfTask1[2], parseInt(EndDateOfTask1[1])-1, EndDateOfTask1[0]);
                                    if (StartDateOfTaskCheck >= from && StartDateOfTaskCheck <= to && EndDateOfTaskCheck >= from && EndDateOfTaskCheck <= to){
                                        console.log("condition is true")
                                        taskWorkLocation.push("true")
                                    }
                                console.log("inside id equal testlocation",taskWorkLocation);
                            } 
                        }
                    }
            }
            }else{
                /*for( var z=0;z<vm.UsersKey.length;z++){
                    taskWorkLocation.push("true");
                }*/
            }
            var selecetUserArrayLength = vm.UsersKey.length;
            for(var i=0;i<selecetUserArrayLength;i++){
                console.log("selectedUserArray[i]",vm.UsersKey[i]);
               var returnValues = checkUserId(vm.UsersKey[i]);
                if(returnValues =="true"){
                    count = count+1;
                }
            }
            for(var i=0;i<count;i++){
                taskWorkLocation.push("true")
            }
            console.log("count",count);
            console.log("taskWorkLocation in final",taskWorkLocation);
            
            if (taskWorkLocation.length ==vm.UsersKey.length&&taskWorkLocation.length >0){
                taskLocationCondition="true"
            }else{
                taskLocationCondition="false"
            }
        console.log("taskLocationCondition in editing.....",taskLocationCondition)
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
        /*for(var i=0;i<vm.UsersKey.length;i++){
            selectedUserArray.push(vm.UsersKey[i]);
        }*/
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
                if (selectedUserArray.length !=0){
                    taskWorkLocation=[];
                    for ( var x=0;x<vm.DateValues.length;x++){
                        for( var y=0;y<selectedUserArray.length;y++){
                            if (vm.DateValues[x][0] == selectedUserArray[y]){
                                //successIdArray.push(selectedUserArray[y]);
                                console.log(" both id of users ARE equal from db",vm.DateValues[x][0]);
                                console.log("selested from dropdown",selectedUserArray[y]);
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
                                var StartDateOfTask = document.getElementById("startDate").value ;
                                var EndDateOfTask = document.getElementById("endDate").value;
                                console.log("StartDateOfTask",StartDateOfTask);
                                console.log("EndDateOfTask",EndDateOfTask);
                                console.log("workStartDateFromDb",workStartDateFromDb);
                                console.log("workEndDateFromDb",workEndDateFromDb);
                                var workStartDate1 = workStartDateFromDb.split("/");
                                var workEndDate1 = workEndDateFromDb.split("/");
                                var StartDateOfTask1 = StartDateOfTask.split("/");
                                var EndDateOfTask1 = EndDateOfTask.split("/");
                                console.log("workStartDate1",workStartDate1);
                                console.log("workEndDate1",workEndDate1);
                                console.log("StartDateOfTask1",StartDateOfTask1);
                                console.log("EndDateOfTask1",EndDateOfTask1)
                                var from = new Date(workStartDate1[2], parseInt(workStartDate1[1])-1, workStartDate1[0]);  // -1 because months are from 0 to 11
                                    var to   = new Date(workEndDate1[2], parseInt(workEndDate1[1])-1, workEndDate1[0]);
                                    var StartDateOfTaskCheck = new Date(StartDateOfTask1[2], parseInt(StartDateOfTask1[1])-1, StartDateOfTask1[0]);
                                    var EndDateOfTaskCheck = new Date(EndDateOfTask1[2], parseInt(EndDateOfTask1[1])-1, EndDateOfTask1[0]);
                                    if (StartDateOfTaskCheck >= from && StartDateOfTaskCheck <= to && EndDateOfTaskCheck >= from && EndDateOfTaskCheck <= to){
                                        console.log("condition is true");
                                    } else{
                                        taskWorkLocation.push("true")
                                        console.log("iam in else part");
                                    }
                                
                                console.log("inside id equal testlocation",taskWorkLocation);
                            }/*else{
                                 //idArray.push(selectedUserArray[y]);
                            }*/
                           
                        }
                    }
            }
            }else{
                for( var z=0;z<selectedUserArray.length;z++){
                    taskWorkLocation.push("true");
                }
            }
            var selecetUserArrayLength = selectedUserArray.length;
           for(var i=0;i<selecetUserArrayLength;i++){
               console.log("selectedUserArray[i]",selectedUserArray[i]);
              var returnValues = checkUserId(selectedUserArray[i]);
               console.log("returnValues",returnValues);
               if(returnValues =="true"){
                   idArray.push(selectedUserArray[i]);
                   //count = count+1;
               }
           }
            console.log("idArray",idArray);
           for(var i=0;i<idArray.length;i++){
               taskWorkLocation.push("true");
           }
            
//            uniqueIdArray = Array.from(new Set(idArray));
//            console.log("unique ",uniqueIdArray);
//            uniqueSucessArray = Array.from(new Set(successIdArray)); 
            /*if(uniqueSucessArray.length !=0){
                for(var i=0;i<uniqueSucessArray.length;i++){
                    console.log("successIdArray[i]",uniqueSucessArray[i]);
                    var returnValues = checkUserId(uniqueSucessArray[i]);
                    if(returnValues =="true"){
                        var dbReturnValue = CheckUserIdFromDb(uniqueSucessArray[i]);
                        console.log("returnValues",returnValues);
                        if( dbReturnValue =="true"){
                            console.log("dbReturnValue",dbReturnValue);
                            count = count+1;
                        }
                    }
                }
            } else{*/
            /*if(uniqueIdArray.length !=1 && uniqueSucessArray.length !=1){
                for(var i =0;i<uniqueIdArray.length;i++){
                    var dbReturnValue = CheckUserIdFromDb(uniqueIdArray[i]);
                    console.log("dbReturnValue",dbReturnValue);
                    if( dbReturnValue =="true"){
                        var returnValues = checkUserId(uniqueIdArray[i]);
                        console.log("returnValues",returnValues);
                        if(returnValues == "true"){
                            console.log("iam in second loop");
                            console.log("dbReturnValue",dbReturnValue);
                            count = count+1;
                        }
                        
                    }
                }
            } else{
                console.log("iam fars second loop")
                for(var i =0;i<uniqueIdArray.length;i++){
                    var dbReturnValue = CheckUserIdFromDb(uniqueIdArray[i]);
                     if( dbReturnValue =="true"){
                          count = count+1;
                     }
                }
                
            }*/
                
          /*  }*/
            
            console.log("count",count)
            for(var i=0;i<count;i++){
                taskWorkLocation.push("true")
            }
            console.log("selectedUserArray",selectedUserArray);
            console.log("taskWorkLocation in final",taskWorkLocation);
            if (selectedUserArray.length !=0){
                if (taskWorkLocation.length ==selectedUserArray.length&&taskWorkLocation.length >0){
                    taskLocationCondition="true"
                }else{
                    taskLocationCondition="false"
                } 
            }
            console.log("taskLocationCondition",taskLocationCondition);
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
