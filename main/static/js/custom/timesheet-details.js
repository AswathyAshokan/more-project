
/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..
//add contact.js

console.log(vm);
$(function(){ 
    //checking plans
    
    if(vm.CompanyPlan == 'family' ){
        var parent = document.getElementById("menuItems");
        var contact = document.getElementById("contact");
        var job = document.getElementById("job");
        var crm = document.getElementById("crm");
        var leave = document.getElementById("leave");
        var timesheet  = document.getElementById("time-sheet");
        var consent = document.getElementById("consent")
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
   
    var Values =[[]];
    var workValues =[[]];
    var mainArray = [];
    var workMainArray =[];
    var taskArrayWithDate =[];
    var workArrayWithDate =[];
    var table = "";
    
    if (vm.TaskTimeSheetDetail != null){
        for(var i=0;i<vm.TaskTimeSheetDetail.length;i++){
             if (vm.TaskTimeSheetDetail[i] !=null){
                 var lateHours ="00:00";
                 var extraHours ="00:00";
                 var diffInStartTime ="00:00";
                 var diffInEndTime ="00:00";
                 var sumLateHours ="00:00";
                 var sumExtraHours ="00:00";
                 var timeSlice =[];
                 timeSlice.push(vm.TaskTimeSheetDetail[i][0]);
                 timeSlice.push(vm.TaskTimeSheetDetail[i][1]);
                 timeSlice.push(vm.TaskTimeSheetDetail[i][2]);
                 timeSlice.push(vm.TaskTimeSheetDetail[i][3]);
                 var utcTime = vm.TaskTimeSheetDetail[i][4];
                 console.log("user name",vm.TaskTimeSheetDetail[i][0]);
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
                 var taskStartTimeFromLog = (HH + ':' + min);
                 console.log("start task time from log",taskStartTimeFromLog);
                 var taskStartDateFromLog = (mm + '/' + dd + '/' + yyyy);
                 
                 var utcTime = vm.TaskTimeSheetDetail[i][5];
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
                 var taskStartTimeFromTask = (HH + ':' + min);
                 var taskStartDateFromTask = (mm + '/' + dd + '/' + yyyy);
                 console.log("task start time from task",taskStartTimeFromTask);
                 
                 var utcTime = vm.TaskTimeSheetDetail[i][6];
                 var dateFromDb = parseInt(utcTime)
                 var offset = new Date().getTimezoneOffset();

                 var d = new Date((dateFromDb * 1000)+offset*60000);
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
                 var DailyStartTime = (HH + ':' + min);
                 console.log("daily start time",DailyStartTime);
                 var utcTime = vm.TaskTimeSheetDetail[i][7];
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
                 var taskEndTimeFromLog = (HH + ':' + min);
                 console.log("task end time from log",taskEndTimeFromLog);
                 var taskEndDateFromLog = (mm + '/' + dd + '/' + yyyy);
                 
                 var utcTime = vm.TaskTimeSheetDetail[i][8];
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
                 var taskEndTimeFromTask = (HH + ':' + min);
                 console.log("task end time from task",taskEndTimeFromTask);
                 var taskEndDateFromTask = (mm + '/' + dd + '/' + yyyy);
                 var utcTime = vm.TaskTimeSheetDetail[i][9];
                 var dateFromDb = parseInt(utcTime)
                 var offset = new Date().getTimezoneOffset();
                
                 var d = new Date(dateFromDb * 1000+offset*60000);
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
                 var DailyEndTime = (HH + ':' + min);
                 console.log("daily end time",DailyEndTime);
                 

// Add two times in hh:mm format
                 function toSeconds(t) {
                     var bits = t.split(':');
                     return bits[0]*3600 + bits[1]*60 ;
                 }
                 var taskStartTimeFromLogDiff = toSeconds(taskStartTimeFromLog);
                 var taskStartTimeFromTaskDiff = toSeconds(taskStartTimeFromTask);
                 var DailyStartTimeDiff = toSeconds(DailyStartTime);
                 var taskEndTimeFromLogDiff = toSeconds(taskEndTimeFromLog);
                 var taskEndTimeFromTaskDiff = toSeconds(taskEndTimeFromTask);
                 var DailyEndTimeDiff = toSeconds(DailyEndTime);
                 if (taskStartDateFromLog ==taskStartDateFromTask ){
                     if (taskStartTimeFromLogDiff>taskStartTimeFromTaskDiff){
                         diffInStartTime =moment.utc(moment(taskStartTimeFromLog," HH:mm").diff(moment(taskStartTimeFromTask," HH:mm"))).format("HH:mm");
                         var t1 = diffInStartTime.split(':');
                         var t2 = sumLateHours.split(':');
                         var mins = Number(t1[1])+Number(t2[1]);
                         var hrs = Math.floor(parseInt(mins / 60));
                         hrs = Number(t1[0])+Number(t2[0])+hrs;
                         mins = mins % 60;
                         if ((hrs >= 0) && (hrs < 10) && (Math.floor(hrs) == hrs)) {
                             hrs ="0"+hrs;
                         }
                         if ((mins >= 0) && (mins < 10) && (Math.floor(mins) == mins)) {
                             mins ="0"+mins;
                         }
                         sumLateHours=hrs+':'+mins;
                         lateHours =sumLateHours;
//                         console.log("late hours of task",lateHours);
                     }
//                     else if (taskStartTimeFromLogDiff<taskStartTimeFromTaskDiff){
//                         diffInStartTime =moment.utc(moment(taskStartTimeFromTask," HH:mm").diff(moment(taskStartTimeFromLog," HH:mm"))).format("HH:mm");
//                         var t1 = diffInStartTime.split(':');
//                         var t2 = sumExtraHours.split(':');
//                         var mins = Number(t1[1])+Number(t2[1]);
//                         var hrs = Math.floor(parseInt(mins / 60));
//                         hrs = Number(t1[0])+Number(t2[0])+hrs;
//                         mins = mins % 60;
//                         if ((hrs >= 0) && (hrs < 10) && (Math.floor(hrs) == hrs)) {
//                             hrs ="0"+hrs;
//                         }
//                         if ((mins >= 0) && (mins < 10) && (Math.floor(mins) == mins)) {
//                             mins ="0"+mins;
//                         }
//                         sumExtraHours=hrs+':'+mins;
//                         extraHours =sumExtraHours;
////                         console.log("extra hours of task",extraHours);
//                     }
                 }else{
                     console.log("hhhh");
                     if (taskStartTimeFromLogDiff>DailyStartTimeDiff){
                         diffInStartTime =moment.utc(moment(taskStartTimeFromLog," HH:mm").diff(moment(DailyStartTime," HH:mm"))).format("HH:mm");
                         var t1 = diffInStartTime.split(':');
                         var t2 = sumLateHours.split(':');
                         var mins = Number(t1[1])+Number(t2[1]);
                         var hrs = Math.floor(parseInt(mins / 60));
                         hrs = Number(t1[0])+Number(t2[0])+hrs;
                         mins = mins % 60;
                         if ((hrs >= 0) && (hrs < 10) && (Math.floor(hrs) == hrs)) {
                             hrs ="0"+hrs;
                         }
                         if ((mins >= 0) && (mins < 10) && (Math.floor(mins) == mins)) {
                             mins ="0"+mins;
                         }
                         sumLateHours=hrs+':'+mins;
                         lateHours =sumLateHours;
//                         console.log("late hours of task 1",lateHours);
                     }else if (taskStartTimeFromLogDiff<DailyStartTimeDiff){
                         diffInStartTime =moment.utc(moment(DailyStartTime," HH:mm").diff(moment(taskStartTimeFromLog," HH:mm"))).format("HH:mm");
                         var t1 = diffInStartTime.split(':');
                         var t2 = sumExtraHours.split(':');
                         var mins = Number(t1[1])+Number(t2[1]);
                         var hrs = Math.floor(parseInt(mins / 60));
                         hrs = Number(t1[0])+Number(t2[0])+hrs;
                         mins = mins % 60;
                         if ((hrs >= 0) && (hrs < 10) && (Math.floor(hrs) == hrs)) {
                             hrs ="0"+hrs;
                         }
                         if ((mins >= 0) && (mins < 10) && (Math.floor(mins) == mins)) {
                             mins ="0"+mins;
                         }
                         sumExtraHours=hrs+':'+mins;
                         extraHours =sumExtraHours;
//                         console.log("extra hours of task 1",extraHours);
                     }
                 }
//                 if (taskStartDateFromLog !=taskStartDateFromTask ){
                     
//                 }
                
                 if (taskEndDateFromLog ==taskEndDateFromTask ){
                     if (taskEndTimeFromLogDiff>taskEndTimeFromTaskDiff){
                        
                         diffInStartTime =moment.utc(moment(taskEndTimeFromLog," HH:mm").diff(moment(taskEndTimeFromTask," HH:mm"))).format("HH:mm");
                         var t1 = diffInStartTime.split(':');
                         var t2 = sumExtraHours.split(':');
                         var mins = Number(t1[1])+Number(t2[1]);
                         var hrs = Math.floor(parseInt(mins / 60));
                         hrs = Number(t1[0])+Number(t2[0])+hrs;
                         mins = mins % 60;
                         if ((hrs >= 0) && (hrs < 10) && (Math.floor(hrs) == hrs)) {
                             hrs ="0"+hrs;
                         }
                         if ((mins >= 0) && (mins < 10) && (Math.floor(mins) == mins)) {
                             mins ="0"+mins;
                         }
                         sumExtraHours=hrs+':'+mins;
                         extraHours =sumExtraHours;
//                         console.log("extra hours3",extraHours);
                     }
//                     else if (taskEndTimeFromLogDiff<taskEndTimeFromTaskDiff){
//                        
//                         diffInStartTime =moment.utc(moment(taskEndTimeFromTask," HH:mm").diff(moment(taskEndTimeFromLog," HH:mm"))).format("HH:mm");
//                         var t1 = diffInStartTime.split(':');
//                         var t2 = sumLateHours.split(':');
//                         var mins = Number(t1[1])+Number(t2[1]);
//                         var hrs = Math.floor(parseInt(mins / 60));
//                         hrs = Number(t1[0])+Number(t2[0])+hrs;
//                         mins = mins % 60;
//                         if ((hrs >= 0) && (hrs < 10) && (Math.floor(hrs) == hrs)) {
//
//                             hrs ="0"+hrs;
//                         }
//                         if ((mins >= 0) && (mins < 10) && (Math.floor(mins) == mins)) {
//                             mins ="0"+mins;
//                         }
//                         sumLateHours=hrs+':'+mins;
//                         lateHours =sumLateHours;
//                         console.log("late hours3",lateHours);
//                     }
                 }else{
                      if (taskEndTimeFromLogDiff>DailyEndTimeDiff){
                         diffInStartTime =moment.utc(moment(taskEndTimeFromLog," HH:mm").diff(moment(DailyEndTime," HH:mm"))).format("HH:mm");
                         var t1 = diffInStartTime.split(':');
                         var t2 = sumExtraHours.split(':');
                         var mins = Number(t1[1])+Number(t2[1]);
                         var hrs = Math.floor(parseInt(mins / 60));
                         hrs = Number(t1[0])+Number(t2[0])+hrs;
                         mins = mins % 60;
                         if ((hrs >= 0) && (hrs < 10) && (Math.floor(hrs) == hrs)) {
                             hrs ="0"+hrs;
                         }
                         if ((mins >= 0) && (mins < 10) && (Math.floor(mins) == mins)) {
                             mins ="0"+mins;
                         }
                         sumExtraHours=hrs+':'+mins;
                         extraHours =sumExtraHours;
//                          console.log("extra hours 4",extraHours);

                     }
//                     else if (taskEndTimeFromLogDiff<DailyEndTimeDiff){
//                         diffInStartTime =moment.utc(moment(DailyEndTime," HH:mm").diff(moment(taskEndTimeFromLog," HH:mm"))).format("HH:mm");
//                         var t1 = diffInStartTime.split(':');
//                         var t2 = sumLateHours.split(':');
//                         var mins = Number(t1[1])+Number(t2[1]);
//                         var hrs = Math.floor(parseInt(mins / 60));
//                         hrs = Number(t1[0])+Number(t2[0])+hrs;
//                         mins = mins % 60;
//                         if ((hrs >= 0) && (hrs < 10) && (Math.floor(hrs) == hrs)) {
//                            hrs ="0"+hrs;
//                         }
//                         if ((mins >= 0) && (mins < 10) && (Math.floor(mins) == mins)) {
//                             mins ="0"+mins;
//                         }
//                         sumLateHours=hrs+':'+mins;
//                         lateHours =sumLateHours;
////                         console.log("late hours 4",lateHours);
//                         
//                     }
                 }
//                 if (taskEndDateFromLog !=taskEndDateFromTask ){
                
                    
//                 }
                 
                 timeSlice.push(lateHours);
                 timeSlice.push(extraHours);
                 console.log("total late hours",lateHours);
                 console.log("total extra hours",extraHours);
                 timeSlice.push(vm.TaskTimeSheetDetail[i][10]);
                 timeSlice.push(vm.TaskTimeSheetDetail[i][11]);
                 timeSlice.push(taskStartDateFromLog);
                 Values.push(timeSlice);
             }
        }
        
        
    }
    
    if (vm.WorkTimeSheeetDetails != null){
        for(var i=0;i<vm.WorkTimeSheeetDetails.length;i++){
            if (vm.WorkTimeSheeetDetails[i] !=null){
                var lateHours ="00:00";
                var extraHours ="00:00";
                var diffInStartTime ="00:00";
                var diffInEndTime ="00:00";
                var sumLateHours ="00:00";
                var sumExtraHours ="00:00";
                var workSlice =[];
                workSlice.push(vm.WorkTimeSheeetDetails[i][0]);
                workSlice.push(vm.WorkTimeSheeetDetails[i][1]);
                workSlice.push(vm.WorkTimeSheeetDetails[i][2]);
                workSlice.push(vm.WorkTimeSheeetDetails[i][3]);
                var utcTime = vm.WorkTimeSheeetDetails[i][4];
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
                var workStartTimeFromLog = (HH + ':' + min);
                var workStartDateFromLog = (mm + '/' + dd + '/' + yyyy);
                console.log("work start time from log",workStartTimeFromLog);
                var utcTime = vm.WorkTimeSheeetDetails[i][5];
                var dateFromDb = parseInt(utcTime)
                var offset = new Date().getTimezoneOffset();
                var d = new Date((dateFromDb * 1000)+offset*60000);
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
                var DailyStartTime = (HH + ':' + min);
                console.log("daily work time",DailyStartTime);
//                console.log("daily start time",DailyStartTime);
                var utcTime = vm.WorkTimeSheeetDetails[i][6];
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
                var workEndTimeFromLog = (HH + ':' + min);
                console.log("work end time from log",workEndTimeFromLog);
                var workEndDateFromLog = (mm + '/' + dd + '/' + yyyy);
                var utcTime = vm.WorkTimeSheeetDetails[i][7];
                var dateFromDb = parseInt(utcTime);
                var offset = new Date().getTimezoneOffset();
                var d = new Date(dateFromDb * 1000+offset*60000);
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
                 var DailyEndTime = (HH + ':' + min);
                console.log("daily  work end time",DailyEndTime);
                function toSeconds(t) {
                     var bits = t.split(':');
                     return bits[0]*3600 + bits[1]*60 ;
                 }
                 var workStartTimeFromLogDiff = toSeconds(workStartTimeFromLog);
                 var DailyStartTimeDiff = toSeconds(DailyStartTime);
                 var workEndTimeFromLogDiff = toSeconds(workEndTimeFromLog);
                 var DailyEndTimeDiff = toSeconds(DailyEndTime);
                 if (workStartTimeFromLogDiff>DailyStartTimeDiff){
                     
                     diffInStartTime =moment.utc(moment(workStartTimeFromLog," HH:mm").diff(moment(DailyStartTime," HH:mm"))).format("HH:mm");
                     var t1 = diffInStartTime.split(':');
                     var t2 = sumLateHours.split(':');
                     var mins = Number(t1[1])+Number(t2[1]);
                     var hrs = Math.floor(parseInt(mins / 60));
                     hrs = Number(t1[0])+Number(t2[0])+hrs;
                     mins = mins % 60;
                     if ((hrs >= 0) && (hrs < 10) && (Math.floor(hrs) == hrs)) {
                         hrs ="0"+hrs;
                     }
                     if ((mins >= 0) && (mins < 10) && (Math.floor(mins) == mins)) {
                         mins ="0"+mins;
                     }
                     sumLateHours=hrs+':'+mins;
                     lateHours =sumLateHours;
                 }
//                else if (workStartTimeFromLogDiff<DailyStartTimeDiff){
//                     diffInStartTime =moment.utc(moment(DailyStartTime," HH:mm").diff(moment(workStartTimeFromLog," HH:mm"))).format("HH:mm");
////                     console.log("c3",diffInStartTime);
//                     var t1 = diffInStartTime.split(':');
//                     var t2 = sumExtraHours.split(':');
//                     var mins = Number(t1[1])+Number(t2[1]);
//                     var hrs = Math.floor(parseInt(mins / 60));
//                     hrs = Number(t1[0])+Number(t2[0])+hrs;
//                     mins = mins % 60;
//                     if ((hrs >= 0) && (hrs < 10) && (Math.floor(hrs) == hrs)) {
//                         hrs ="0"+hrs;
//                     }
//                     if ((mins >= 0) && (mins < 10) && (Math.floor(mins) == mins)) {
//                         mins ="0"+mins;
//                     }
//                     sumExtraHours=hrs+':'+mins;
//                     extraHours =sumExtraHours;
////                     console.log("extra hours",extraHours);
//                 }
               
                if (workEndTimeFromLogDiff>DailyEndTimeDiff){
//                    console.log("c4");
                     diffInStartTime =moment.utc(moment(workEndTimeFromLog," HH:mm").diff(moment(DailyEndTime," HH:mm"))).format("HH:mm");
                     var t1 = diffInStartTime.split(':');
                     var t2 = sumExtraHours.split(':');
                     var mins = Number(t1[1])+Number(t2[1]);
                     var hrs = Math.floor(parseInt(mins / 60));
                     hrs = Number(t1[0])+Number(t2[0])+hrs;
                     mins = mins % 60;
                    if ((hrs >= 0) && (hrs < 10) && (Math.floor(hrs) == hrs)) {
                        hrs ="0"+hrs;
                    }
                    if ((mins >= 0) && (mins < 10) && (Math.floor(mins) == mins)) {
                        mins ="0"+mins;
                    }
                    sumExtraHours=hrs+':'+mins;
                    extraHours =sumExtraHours;
                }
//                else if (workEndTimeFromLogDiff<DailyEndTimeDiff){
//                    diffInStartTime =moment.utc(moment(DailyEndTime," HH:mm").diff(moment(taskEndTimeFromLog," HH:mm"))).format("HH:mm");
//                    var t1 = diffInStartTime.split(':');
//                    var t2 = sumLateHours.split(':');
//                    var mins = Number(t1[1])+Number(t2[1]);
//                    var hrs = Math.floor(parseInt(mins / 60));
//                    hrs = Number(t1[0])+Number(t2[0])+hrs;
//                    mins = mins % 60;
//                    if ((hrs >= 0) && (hrs < 10) && (Math.floor(hrs) == hrs)) {
//                        hrs ="0"+hrs;
//                    }
//                    if ((mins >= 0) && (mins < 10) && (Math.floor(mins) == mins)) {
//                        mins ="0"+mins;
//                    }
//                    sumLateHours=hrs+':'+mins;
//                    lateHours =sumLateHours;
//                }
//                console.log("jsdddddddd",extraHours);
                console.log("work extra hours",extraHours);
                console.log("work late hours",lateHours);
                workSlice.push(lateHours);
                workSlice.push(extraHours);
                workSlice.push(vm.WorkTimeSheeetDetails[i][8]);
                workSlice.push(workStartDateFromLog);
                workValues.push(workSlice);
                
            }
        }
        
        
    }
     
    function createDataArray(taskValues){
        var subArray = [];
        for(i = 1; i < taskValues.length; i++) {
            for(var propertyName in taskValues[i]) {
                subArray.push(taskValues[i][propertyName]);
            }
            mainArray.push(subArray);
            subArray = [];
        }
    }
    
    function createDataArrayForWork(workValues){
        var subArray = [];
        for(i = 1; i < workValues.length; i++) {
            for(var propertyName in workValues[i]) {
                subArray.push(workValues[i][propertyName]);
            }
            workMainArray.push(subArray);
            subArray = [];
        }
    }
    function createDataArrayTaskWithDate(workValues){
        var subArray = [];
        for(i = 1; i < workValues.length; i++) {
            for(var propertyName in workValues[i]) {
                subArray.push(workValues[i][propertyName]);
            }
            taskArrayWithDate.push(subArray);
            subArray = [];
        }
    }
    function createDataArrayWorkWithDate(workValues){
        var subArray = [];
        for(i = 1; i < workValues.length; i++) {
            for(var propertyName in workValues[i]) {
                subArray.push(workValues[i][propertyName]);
            }
            workArrayWithDate.push(subArray);
            subArray = [];
        }
    }
    function dataTableManipulate(mainArray){
        table =  $("#timeSheet_details").DataTable({
            data: mainArray,
            "columnDefs": [{ "title": "Task Name", "targets": 1 },{'visible': false, 'targets': [3],
            }]
        });
    }
    if(Values != null) {
        createDataArray(Values);
    }
    dataTableManipulate(mainArray);
    $("#taskDetail").on('click',function(){
        mainArray=[];
        $('#timeSheet_details').dataTable().fnDestroy();
        if(Values != null) {
            createDataArray(Values);
        }
        dataTableManipulate(mainArray);
    });
    if(workValues != null) {
            createDataArrayForWork(workValues);
        }
    function workdataTableManipulate(workMainArray){
            table =  $("#timeSheet_details").DataTable({
                data: workMainArray,
                "searching": true,
                "info": false,
                "lengthChange":false,
                "columnDefs": [ { "title": "Work Location", "targets": 1 },{
                    'visible': false, 'targets': [3],
                }]
            });
        $('#tbl_details_length').after($('.datepic-top'));
    }
    
    $("#workDetail").on( 'click', function () {
        
        $('#timeSheet_details').dataTable().fnDestroy();
        workdataTableManipulate(workMainArray);
    });
    
     
    
    
    //date filtering function
    $('#toDate').change(function () {
        toDateValue = $('#toDate').val();
        fromDateValue = $('#fromDate').val();
        if (toDateValue.length !=0 && fromDateValue.length !=0){
            var FinalArrayForDateFilter =[[]];
            var  ArrayForDateFilter =[];
            var FinalArrayForDateFilterOfTask =[[]];
            var  ArrayForDateFilterOfTask =[];
            if(document.getElementById('workDetail').clicked != true)
            {
                for (var k=0;k<mainArray.length;k++){
                    if (mainArray[k].length !=0){
                        var d1 = fromDateValue.split("/");
                        var d2 = toDateValue.split("/");
                        var c = mainArray[k][8].split("/");
                        var from = new Date(d1[2], parseInt(d1[1])-1, d1[0]);  // -1 because months are from 0 to 11
                        var to   = new Date(d2[2], parseInt(d2[1])-1, d2[0]);
                        var check = new Date(c[2], parseInt(c[1])-1, c[0]);
                        if (check >= from && check <= to){
                            ArrayForDateFilter.push(mainArray[k]);
                            for(var j=0;j<mainArray.length;j++){
                                if (j !=k){
                                    if (ArrayForDateFilter[0][6]==mainArray[j][6] && ArrayForDateFilter[0][7]==mainArray[j][7]){
                                        ArrayForDateFilter.push(mainArray[j]);
                                    }
                                }
                            }
                            FinalArrayForDateFilter.push(ArrayForDateFilter); 
                            var  ArrayForDateFilter =[];
                        }
                    }
                }
                var TaskTimeSheetRealArray =[[]];
                for (var i=1;i<FinalArrayForDateFilter.length;i++){
                    var TaskValues =[];
                    var sumLateHours="00:00";
                    var sumExtraHours="00:00";
                    TaskValues.push(FinalArrayForDateFilter[i][0][0]);
                    TaskValues.push(FinalArrayForDateFilter[i][0][1]);
                    TaskValues.push(FinalArrayForDateFilter[i].length);
                    TaskValues.push(FinalArrayForDateFilter[i][0][3]);
                    for  (var j=0;j<FinalArrayForDateFilter[i].length;j++){
                        var t1 = FinalArrayForDateFilter[i][j][4].split(':');
                        var t2 = sumLateHours.split(':');
                        var mins = Number(t1[1])+Number(t2[1]);
                        var hrs = Math.floor(parseInt(mins / 60));
                        hrs = Number(t1[0])+Number(t2[0])+hrs;
                        mins = mins % 60;
                        if ((hrs >= 0) && (hrs < 10) && (Math.floor(hrs) == hrs)) {
                            hrs ="0"+hrs;
                        }
                        if ((mins >= 0) && (mins < 10) && (Math.floor(mins) == mins)) {
                            mins ="0"+mins;
                        }
                        sumLateHours=hrs+':'+mins;
                        var extrat1 = FinalArrayForDateFilter[i][j][5].split(':');
                        var extrat2 = sumExtraHours.split(':');
                        var extramins = Number(extrat1[1])+Number(extrat2[1]);
                        var extrahrs = Math.floor(parseInt(extramins / 60));
                        extrahrs = Number(extrat1[0])+Number(extrat2[0])+extrahrs;
                        extramins = extramins % 60;
                        if ((extrahrs >= 0) && (extrahrs < 10) && (Math.floor(extrahrs) == extrahrs)) {
                            extrahrs ="0"+extrahrs;
                        }
                        if ((extramins >= 0) && (extramins < 10) && (Math.floor(extramins) == extramins)) {
                            extramins ="0"+extramins;
                        }
                        sumExtraHours=extrahrs+':'+extramins;
                    }
                    TaskValues.push(sumExtraHours);
                    TaskValues.push(sumLateHours);
                    TaskTimeSheetRealArray.push(TaskValues);
                }
                $('#timeSheet_details').dataTable().fnDestroy();
                if(TaskTimeSheetRealArray != null) {
                    createDataArrayTaskWithDate(TaskTimeSheetRealArray);
                }
                dataTableManipulate(taskArrayWithDate);
            }
            if(document.getElementById('workDetail').clicked == true){
                for (var k=0;k<workMainArray.length;k++){
                    if (workMainArray[k].length !=0){
                        var d1 = fromDateValue.split("/");
                        var d2 = toDateValue.split("/");
                        var c = workMainArray[k][7].split("/");
                        var from = new Date(d1[2], parseInt(d1[1])-1, d1[0]);  // -1 because months are from 0 to 11
                        var to   = new Date(d2[2], parseInt(d2[1])-1, d2[0]);
                        var check = new Date(c[2], parseInt(c[1])-1, c[0]);
                        if (check >= from && check <= to){
                            ArrayForDateFilterOfTask.push(workMainArray[k]);
                            for(var j=0;j<workMainArray.length;j++){
                              if (j !=k){
                                  if (ArrayForDateFilterOfTask[0][6]==workMainArray[j][6]){
                                      ArrayForDateFilterOfTask.push(workMainArray[j]);
                                  }
                              }
                            }
                            FinalArrayForDateFilterOfTask.push(ArrayForDateFilterOfTask); 
                            var  ArrayForDateFilterOfTask =[];
                        }
                    }
                }
                var WorkTimeSheetRealArray =[[]];
                for (var i=1;i<FinalArrayForDateFilterOfTask.length;i++){
                    var WorkValues =[];
                    var sumLateHours="00:00";
                    var sumExtraHours="00:00";
                    WorkValues.push(FinalArrayForDateFilterOfTask[i][0][0]);
                    WorkValues.push(FinalArrayForDateFilterOfTask[i][0][1]);
                    WorkValues.push(FinalArrayForDateFilterOfTask[i].length);
                    WorkValues.push(FinalArrayForDateFilterOfTask[i][0][3]);
                    for  (var j=0;j<FinalArrayForDateFilterOfTask[i].length;j++){
                        var t1 = FinalArrayForDateFilterOfTask[i][j][4].split(':');
                        var t2 = sumLateHours.split(':');
                        var mins = Number(t1[1])+Number(t2[1]);
                        var hrs = Math.floor(parseInt(mins / 60));
                        hrs = Number(t1[0])+Number(t2[0])+hrs;
                        mins = mins % 60;
                        if ((hrs >= 0) && (hrs < 10) && (Math.floor(hrs) == hrs)) {
                            hrs ="0"+hrs;
                        }
                        if ((mins >= 0) && (mins < 10) && (Math.floor(mins) == mins)) {
                            mins ="0"+mins;
                        }
                        sumLateHours=hrs+':'+mins;
                        var extrat1 = FinalArrayForDateFilter[i][j][5].split(':');
                        var extrat2 = sumExtraHours.split(':');
                        var extramins = Number(extrat1[1])+Number(extrat2[1]);
                        var extrahrs = Math.floor(parseInt(extramins / 60));
                        extrahrs = Number(extrat1[0])+Number(extrat2[0])+extrahrs;
                        extramins = extramins % 60;
                        if ((extrahrs >= 0) && (extrahrs < 10) && (Math.floor(extrahrs) == extrahrs)) {
                            extrahrs ="0"+extrahrs;
                        }
                        if ((extramins >= 0) && (extramins < 10) && (Math.floor(extramins) == extramins)) {
                            extramins ="0"+extramins;
                        }
                        sumExtraHours=extrahrs+':'+extramins;
                    }
                    WorkValues.push(sumExtraHours);
                    WorkValues.push(sumLateHours);
                    WorkTimeSheetRealArray.push(WorkValues);
                }
                $('#timeSheet_details').dataTable().fnDestroy();
                if(WorkTimeSheetRealArray != null) {
                    createDataArrayWorkWithDate(WorkTimeSheetRealArray);
                }
                workdataTableManipulate(workArrayWithDate);
            }
        }
    });
    
     $('#refreshButton').click(function(e) {
        $('#timeSheet_details').dataTable().fnDestroy();
        $('#fromDate').datepicker('setDate', null);
        $('#toDate').datepicker('setDate', null);
        dataTableManipulate(mainArray);
    });
});