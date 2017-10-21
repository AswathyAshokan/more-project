console.log("viwe model value",vm);
console.log("end time",vm.DailyEndTime);
var loginTypeForEdit = vm.LoginType;
var companyTeamName = vm.CompanyTeamName;



//aswathy
$(function(){
    document.getElementById("imageId").src=vm.ProfilePicture;
    if (vm.ProfilePicture ==""){
        document.getElementById("imageId").src="/static/images/default.png"
    }
});

$().ready(function() {
     if(vm.CompanyPlan == 'family' ){
         $("#contact").remove();
         $("#crm").remove();
         $("#leave").remove();
         $("#fitToWork").remove();
         $("#time-sheet").remove();
         $("#consent").remove();
         $("#job").remove();
         var rbtn = document.getElementById('loginType1');
        rbtn.style.display = 'none';
         var rbtnfc = document.getElementById('nfcbutton');
        rbtnfc.style.display = 'none';
   } else if(vm.CompanyPlan == 'campus'){
       $("#contact").remove();
       $("#crm").remove();
       $("#leave").remove();
       $("#fitToWork").remove();
       $("#time-sheet").remove();
       $("#consent").remove();
       $("#job").remove();
    }
    var loginTypeRadio = "";
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
   var uniqueIdArray = [];
   var condition ="";
  
   var editConditionArry = [];
   var countInEdit = 0;
   var condintionInEdit ="";
   var localDate="";
   var localEndDate="";
   var exposureSlice =[];
   var exposureTimeArray =[];
   var exposureWorkSlice =[];
   var exposureWorkTimeArray =[];
   var exposureHour ="";
   var exposureMinute ="";
   var TotalBreakTime ="";
   var exposureWorkHour ="";
   var exposureWorkMinute ="";
   var TotalWorkTime ="";
   var fitWork = "";
   if (vm.UsersKey !=null){
       selectedUserArray=vm.UsersKey;
   }
    
    
    
    
    var repeat ="";
   function  addleveldata(){
        repeat = "<div class='exposureId plus' style='padding-right: 17px;' >"+"<label for='workExplosureText'>Break Time</label>"+"<input type='text' class='form-control break-duration'  data-timepicker placeholder='12:00'   name='breakTime'size='5' id='breakTime' >"+"<label>After</label>"+ "<input type='text'  class='form-control break-duration-after'  placeholder='12:00'   name='workTime' size='5' id='workTime' data-timepicker>"+
       "<span class='add-decl'  >+</span>"+"</div>"
         $( "#exposureTextBoxAppend" ).append( repeat );
     
  }
   
   $(document).on('click', '.add-decl', function () {
      if ($(this).closest('.plus').is(':last-child')) {
          addleveldata();
      }
      else {
          $(this).closest('.plus').remove();
      }
   });

   
   
   $(document).on("keypress",".break-duration", function(){
       console.log("damn my ........faith");
       var value=$(this).val();
       if(value.length==2){
       value=value+":";
       }
       if(value.length>=5){
           
           return false;
       }
       $(this).val(value)  
   });
    $(document).on("keypress",".break-duration-after", function(){
        console.log("damn my ........faith11");
        var value=$(this).val();
       if(value.length==2){
       value=value+":";
       }
        if(value.length>=5){
           
           return false;
       }
       $(this).val(value)  
   });
    
    
    function checkUserId(userId) {
        if(vm.DateValues !=null){
           for(var j=0 ;j<vm.DateValues.length;j++ ){
           if(vm.DateValues[j][0] !=userId ){
               console.log("in if func")
               returnString ="true"
               
           }else{
                 console.log("in else func")
               returnString ="false"
               break;
           }
       }
           return returnString
       }
   }
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
   // get loginType details
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
    
    if(vm.PageType == "edit"){ 
       
        document.title = 'Edit WorkLocation'
        if(vm.FitToWorkCheck =="EachTime") {
            document.getElementById("fitToWorkCheck").checked = true;
        }
        if(vm.WorkTime !=null){
            if(vm.WorkTime.length !=0){
            document.getElementById("workExplosure").checked = true;
            var div = document.getElementById('work');
            div.style.visibility = 'visible';
            div.style.display ='inline';
            document.getElementById("workTime").value =vm.WorkTime[0];
            document.getElementById("breakTime").value =vm.BreakTime[0];
            var repeat ="";
            for (var i=1; i<vm.WorkTime.length; i++){
                repeat = "<div class='exposureId plus wrp-plus' style='padding-right: 17px;' >"+"<label for='workExplosureText'>Break Time</label>"+"<input type='text' class='form-control break-duration'  data-timepicker placeholder='12:00'   name='breakTime'size='5' id='breakTime'value="+ vm.BreakTime[i] +" >"+"<label>After</label>"+ "<input type='text'  class='form-control break-duration-after'  placeholder='12:00'   name='workTime' size='5' id='workTime' data-timepicker  value="+ vm.WorkTime[i] +">"+"<span class='add-decl'  >+</span>"+"</div>"
             $( "#exposureTextBoxAppend" ).append( repeat );
            }
            }
        }
        
        fitWork  =  vm.FitToWorkName;
        var fitToWorkName = vm.FitToWorkName;
        if(fitToWorkName !=null){
            if (fitToWorkName.length !=0){
                $('#WorkLocationFitToWork option:contains(' + fitToWorkName + ')').prop({selected: true});
            }
        }
        var element = document.getElementById('logInMinutes');
        element.value = vm.LogInMin;
        
        if(vm.LoginType =="NFC" ){
            document.getElementById("loginType1").checked = true;
            var div = document.getElementById('nfcTagId');
            div.style.visibility = 'visible';
            div.style.display ='inline';
            document.getElementById("nfcTagForTask").value =vm.NFCTagId;
        }else{
            document.getElementById("loginType2").checked = true;
        }
        
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
        localDate = (mm + '/' + dd + '/' + yyyy);
       
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
        localEndDate = (endmm + '/' + enda + '/' + endyyyy);
       
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
   });
    $("#WorkLocationFitToWork").on('change', function(evt, params) {
       //var doc =$('#WorkLocationFitToWork').find('option:selected').val()
       var doc = document.getElementById("WorkLocationFitToWork");
       if(doc.length !=0){
           fitWork =doc.options[doc.selectedIndex].text;;
       }
   });
    $("#workLocationForm").validate({
       rules: {
          // usersAndGroupId:"required",
           taskLocation : "required",
           startDate:"required",
           endDate:"required",
           dailyStartTime:"required",
           dailyEndTime:"required",
           loginType : "required"
       },
       messages: {
           usersAndGroupId: "Please select user or group",
           taskLocation:"please fill this column",
           loginType:"select login type"
       },
        submitHandler: function(){//to pass all data of a form serial
            
            if(loginTypeRadio.length ==0)
           {
               loginTypeRadio = loginTypeForEdit;
           } else {
               loginTypeRadio = loginTypeRadio;
           }
           
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
             var existingWorkLocationUserId =[];
            if(vm.PageType == "edit"){
                console.log("selected",selectedUserArray);
                 var currentWorkLocationId =vm.WorkLogId 
                 var WorklocationCondition=false;
                var isDateExist=false;
                editConditionArry=[];
                for(var i=0;i<selectedUserArray.length;i++){
                    
//                    WorklocationCondition=0;
                    
                    for(var j=0;j<vm.DateValues.length;j++){
                        if(selectedUserArray[i]==vm.DateValues[j][0]){
                            
                            if(currentWorkLocationId!=vm.DateValues[j][3]){
                                var utcTime = vm.DateValues[j][1];
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
                                    console.log("work start from db",workStartDateFromDb);
                                    var utcTime =vm.DateValues[j][2];
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
                                    console.log("work end from db",workEndDateFromDb);
                                    var StartDateOfTask = document.getElementById("startDate").value ;
                                    var EndDateOfTask = document.getElementById("endDate").value;
                                    console.log("start date from task",StartDateOfTask);
                                    console.log("end date from task",EndDateOfTask);
                                    var from = Date.parse(workStartDateFromDb);
                                    var to   = Date.parse(workEndDateFromDb);
                                    var StartDateOfTaskCheck = Date.parse(StartDateOfTask );
                                    var EndDateOfTaskCheck = Date.parse(EndDateOfTask );
                                
                                    console.log("StartDateOfTaskCheck",StartDateOfTaskCheck)
                                    console.log("EndDateOfTaskCheck",EndDateOfTaskCheck)
                                    console.log("from",from)
                                    console.log("to",to)
                                    console.log("---------------------")
                                    
//                                        if (StartDateOfTaskCheck >= from && StartDateOfTaskCheck <= to && EndDateOfTaskCheck >= from && EndDateOfTaskCheck <= to ||from <StartDateOfTaskCheck  && to<EndDateOfTaskCheck){
                                            
                                    if((StartDateOfTaskCheck>=from && StartDateOfTaskCheck<=to) || (EndDateOfTaskCheck>=from && EndDateOfTaskCheck<=to)){
                                            
                                        isDateExist=true;
                                        console.log("dateExist",true)
                                    }
                                else{
                                    console.log("dateExist",false)
//                                     WorklocationCondition++;
//                                   if (WorklocationCondition==1){
//                                        editConditionArry.push("true");                                       
//                                      
//                                    }
                                   //if WorklocationCondition push
                                   //then WorklocationCondition=false;
                                }
                            }
                            else{
                                    console.log("dateExist",false)
//                                    WorklocationCondition++;
//                                   if (WorklocationCondition==1){
//                                        editConditionArry.push("true");                                       
//                                      
//                                    }
                                   //if WorklocationCondition push
                                   //then WorklocationCondition=false;
                                }
                        }
                    }
                    
                    //
                     if (!isDateExist){                                       
                         editConditionArry.push("true");   
                     }
                }
                console.log("kkkkkk",editConditionArry);
                if (editConditionArry.length ==selectedUserArray.length&&editConditionArry.length >0){
                    taskLocationCondition="true";
                }else{
                    taskLocationCondition="false";
                } 
            } else{
                var count =0;
                var conditionArray =[];
                console.log("kkksssssssssssssssssss",selectedUserArray);
                if(vm.DateValues != null){
                    if (selectedUserArray.length !=0){
                        taskWorkLocation=[];
                        for ( var x=0;x<vm.DateValues.length;x++){
                            for( var y=0;y<selectedUserArray.length;y++){
                                if (vm.DateValues[x][0] == selectedUserArray[y]){
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
                                    console.log("work start from db",workStartDateFromDb);
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
                                    console.log("work end from db",workEndDateFromDb);
                                    var StartDateOfTask = document.getElementById("startDate").value ;
                                    var EndDateOfTask = document.getElementById("endDate").value;
                                    console.log("start date from task",StartDateOfTask);
                                    console.log("end date from task",EndDateOfTask);
                                    var from = Date.parse(workStartDateFromDb);
                                    var to   = Date.parse(workEndDateFromDb);
                                    var StartDateOfTaskCheck = Date.parse(StartDateOfTask );
                                    var EndDateOfTaskCheck = Date.parse(EndDateOfTask );
                                    if((StartDateOfTaskCheck>=from && StartDateOfTaskCheck<=to) || (EndDateOfTaskCheck>=from && EndDateOfTaskCheck<=to)){
                                        condition="true";
                                        console.log("i am in success of ifff");
                                        break;
                                    } else{
                                        condition="false";
                                        conditionArray.push("false");
                                        existingWorkLocationUserId.push(vm.DateValues[x][3]);
                                        console.log("iam in else part");
                                    }
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
                console.log("lllllllllioioioioioii",existingWorkLocationUserId);
                console.log("condition array",conditionArray);
                if(vm.DateValues != null){
                    if (selectedUserArray.length !=0){
                         for ( var x=0;x<vm.DateValues.length;x++){
                            for( var y=0;y<selectedUserArray.length;y++){
                                if (vm.DateValues[x][0] == selectedUserArray[y]){
                                count=count+1;

                            }
                            }
                        }
                    }
                }
                 var conditionInLoop ="";
                console.log("count",count);
                if(count !=0){
                    if (conditionArray !=null){
                        if (conditionArray.length ==count){
                            console.log("in here1");
                            for(var i=0;i<conditionArray.length;i++){
                                 conditionInLoop="true";
                                
                            }
                        } 
                    }
                }
                if (conditionInLoop=="true"){
                    taskWorkLocation.push("true");
                }
                console.log("task 562",taskWorkLocation);
                var selecetUserArrayLength = selectedUserArray.length;
                for(var i=0;i<selecetUserArrayLength;i++){
                    var returnValues = checkUserId(selectedUserArray[i]);
                    if(returnValues =="true"){
                        idArray.push(selectedUserArray[i]);
                    } else{
                         existingWorkLocationUserId.push(selectedUserArray[i]);
                        
                    }
                }
                for(var i=0;i<idArray.length;i++){
                    taskWorkLocation.push("true");
                }
//           console.log("hai existingWorkLocationUserId @@@@@@@@@@@@@@@@@@@@@@@@@@@@@ ",existingWorkLocationUserId);
            console.log("final taskLocation",taskWorkLocation);
            if (selectedUserArray.length !=0){
                if (taskWorkLocation.length ==selectedUserArray.length&&taskWorkLocation.length >0){
                    taskLocationCondition="true"
                }else{
                    taskLocationCondition="false"
                } 
            }
            }
            
            
            console.log("enter restricted or not",taskLocationCondition)
           var nfcTagId =  document.getElementById("nfcTagForTask").value;
           var starDateString = document.getElementById('startDate').value;
           var endDateString = document.getElementById('endDate').value;
           $("#saveButton").attr('disabled', true);
           
           var startdatum = Date.parse(starDateString)/1000;
           var endDatum = Date.parse(endDateString)/1000;
         
           var startDateInDate = new Date(starDateString);
           var dailyStartTime = document.getElementById('dailyStartTime').value;
           
           var endDateInDate = new Date(endDateString);
           var dailyEndTime = document.getElementById('dailyEndTime').value;
           
           var startTimeArray = dailyStartTime.split(':');
           var startHour = parseInt(startTimeArray[0]);
           var startMin = parseInt(startTimeArray[1]);
           startDateInDate.setHours(startHour);
           startDateInDate.setMinutes(startMin);
           var endTimeArray = dailyEndTime.split(':');
           var endHour = parseInt(endTimeArray[0]);
           var endMin = parseInt(endTimeArray[1]);
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
          // check check box of fit to work
           var chkPassport = document.getElementById("fitToWorkCheck");
           if (chkPassport.checked) {
               fitToWorkCheck ="EachTime";
           }else {
               fitToWorkCheck ="OnceADay";
           }
           
           
           startDateString = startDateInDate;
           var date = new Date(Date.parse(startDateString));
           var startDateOfWork = formatDate(date);
           var endDateStringInUtc = endDateInDate;
           var endDateData = new Date(Date.parse(endDateStringInUtc));
           var endDateOfWork = formatDate(endDateData);
           var formData = $("#workLocationForm").serialize();
           //get the user's name corresponding to  keys selected from dropdownlist
           formData = formData+"&startDateTimeStamp="+startdatum+"&endDateTimeStamp="+endDatum +"&dailyStartTimeString="+startDateOfWork+"&dailyEndTimeString="+endDateOfWork+"&fitToWorkCheck="+fitToWorkCheck +"&fitToWorkName="+fitWork+ "&loginType=" + loginTypeRadio+"&exposureBreakTime="+ exposureSlice+"&exposureWorkTime="+ exposureWorkSlice;
           
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
            
            if(taskLocationCondition=="true"){
                var ConcatinatedUser ;
                if (vm.PageType == "edit"){
//                    for(i=0;i<vm.UsersKey.length;i++){
//                        formData = formData+"&oldUsers="+vm.UsersKey[i];
//                    }
//                   
//                    for(i = 0; i < vm.UsersKey.length; i++) {
//                        formData = formData+"&selectedUserNames="+vm.UsersKey[i];
//                    }
                    var workLocationId =vm.WorkLogId 
                    $("#saveButton").attr('disabled', true);
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
                    $("#saveButton").attr('disabled', true);
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
                console.log("hai existingWorkLocationUserId ########## ",existingWorkLocationUserId);
                var uniqueItems = Array.from(new Set(existingWorkLocationUserId));
                console.log("after filtering of existing unique user id ",uniqueItems);
                $("#myModalForUniqueTest").modal();
                $("#saveButton").attr('disabled', false);
//                $("#cancelForCheckUnique").click(function(){
//                    window.location = '/'+companyTeamName+'/worklocation/add';
//                });
            }
            
        }
    });
     $("#cancel").click(function() {
            window.location = '/'+companyTeamName+'/worklocation';
    });
    document.getElementById("WorkLocation").className += " active";
    
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