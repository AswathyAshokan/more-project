console.log("start time",vm.UsersKey);
console.log("end time",vm.DailyEndTime);
var companyTeamName = vm.CompanyTeamName;
 var selectedUserArray = [];
var startDateInUnix;
var endDateInUnix;
var dailyStartTimeUnix;
var dailyEndTimeUnix;

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
        /*alert("selectedUserArray",selectedUserArray);
        /*$("#usersAndGroupId").on('change', function(evt, params) {
            console.log("inside group1");
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
                            console.log("values of temp array",tempArray);
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
            console.log("user array in editing",selectedUserArray);
            alert("selectedUserArray in change",selectedUserArray)
        });
        console.log("user array in editing out side",selectedUserArray);*/
       
    }
    
    
    var selectedGroupArray = [];
    var groupKeyArray = [];
    $("#usersAndGroupId").on('change', function(evt, params) {
        console.log("inside group1");
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
                        console.log("values of temp array",tempArray);
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
            console.log("lattitude",document.getElementById('latitudeId').value);
             console.log("longitude",document.getElementById('longitudeId').value);
           
            
            
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
            var startDateString = startDateInDate;
            var date = new Date(Date.parse(startDateString));
            var startDateOfWork = formatDate(date);
            var endDateString = endDateInDate;
            var endDateData = new Date(Date.parse(endDateString));
            var endDateOfWork = formatDate(endDateData);
            console.log("startDateOfWork",startDateOfWork)
            console.log("endDateOfWork",endDateOfWork)
            
            
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
             //formData = formData+"&lattitude="+;
              console.log("formData",formData);
            var ConcatinatedUser ;
            if (vm.PageType == "edit"){
                for(i=0;i<vm.UsersKey.length;i++){
                    formData = formData+"&oldUsers="+vm.UsersKey[i];
                    /*console.log("second values",vm.UsersKey[i+1]);
                    var UserIdSring = vm.UsersKey[i].concat(",");
                    ConcatinatedUser = UserIdSring.concat(vm.UsersKey[i+1]);*/
                    
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
                        }else  if(response == "falseAlreadyExist"){
                            $("#myModalForUniqueTest").modal();
                            $("#cancelForCheckUnique").click(function(){
                                window.location = '/'+companyTeamName+'/worklocation';
                            });
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
                            }else  if(response == "falseAlreadyExist"){
                             $("#myModalForUniqueTest").modal();
                                $("#cancelForCheckUnique").click(function(){
                                     window.location = '/'+companyTeamName+'/worklocation';
                                });
                                // window.location = '/'+companyTeamName+'/worklocation';
                            }else{
                                 $("#saveButton").attr('disabled', false);
                            }
                        },
                        error: function (request,status, error) {
                        }
                    });
                    return false;
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
