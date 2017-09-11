
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
                     }else if (taskStartTimeFromLogDiff<taskStartTimeFromTaskDiff){
                         diffInStartTime =moment.utc(moment(taskStartTimeFromTask," HH:mm").diff(moment(taskStartTimeFromLog," HH:mm"))).format("HH:mm");
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
                 }
                 if (taskStartDateFromLog !=taskStartDateFromTask ){
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
                     }
                 }
                
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
                     }else if (taskEndTimeFromLogDiff<taskEndTimeFromTaskDiff){
                         diffInStartTime =moment.utc(moment(taskEndTimeFromTask," HH:mm").diff(moment(taskEndTimeFromLog," HH:mm"))).format("HH:mm");
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
                 }
                 if (taskEndDateFromLog !=taskEndDateFromTask ){
                
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

                     }else if (taskEndTimeFromLogDiff<DailyEndTimeDiff){
                         diffInStartTime =moment.utc(moment(DailyEndTime," HH:mm").diff(moment(taskEndTimeFromLog," HH:mm"))).format("HH:mm");
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
                 }
                 timeSlice.push(lateHours);
                 timeSlice.push(sumExtraHours);
                 timeSlice.push(vm.TaskTimeSheetDetail[i][10]);
                 timeSlice.push(vm.TaskTimeSheetDetail[i][11]);
                 timeSlice.push(vm.TaskTimeSheetDetail[i][12]);
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
                function toSeconds(t) {
                     var bits = t.split(':');
                     return bits[0]*3600 + bits[1]*60 ;
                 }
                 var workStartTimeFromLogDiff = toSeconds(workStartTimeFromLog);
                 var DailyStartTimeDiff = toSeconds(DailyStartTime);
                 var workEndTimeFromLogDiff = toSeconds(workEndTimeFromLog);
                 var DailyEndTimeDiff = toSeconds(DailyEndTime);
                 if (workStartTimeFromLogDiff>DailyStartTimeDiff){
                     console.log("c1");
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
                 }else if (workStartTimeFromLogDiff<DailyStartTimeDiff){
                     diffInStartTime =moment.utc(moment(DailyStartTime," HH:mm").diff(moment(workStartTimeFromLog," HH:mm"))).format("HH:mm");
                     console.log("c3",diffInStartTime);
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
                console.log("error2");
                if (workEndTimeFromLogDiff>DailyEndTimeDiff){
                    console.log("c4");
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
                }else if (workEndTimeFromLogDiff<DailyEndTimeDiff){
                    diffInStartTime =moment.utc(moment(DailyEndTime," HH:mm").diff(moment(taskEndTimeFromLog," HH:mm"))).format("HH:mm");
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
                workSlice.push(lateHours);
                workSlice.push(sumExtraHours);
                workSlice.push(vm.WorkTimeSheeetDetails[i][8]);
                workSlice.push(vm.WorkTimeSheeetDetails[i][9]);
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
    function dataTableManipulate(mainArray){
        table =  $("#timeSheet_details").DataTable({
            data: mainArray,
            "columnDefs": [{
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
    
    $("#workDetail").on( 'click', function () {
        if(workValues != null) {
            createDataArrayForWork(workValues);
        }
        $('#timeSheet_details').dataTable().fnDestroy();
        workdataTableManipulate(workMainArray);
        function workdataTableManipulate(workMainArray){
            table =  $("#timeSheet_details").DataTable({
            data: workMainArray,
            "searching": true,
            "info": false,
            "lengthChange":false,
            "columnDefs": [{
                'visible': false, 'targets': [1] ,
            }]
        });
        $('#tbl_details_length').after($('.datepic-top'));
    }
                workMainArray =[];

    });
    
    
    //date filtering function
    $('#toDate').change(function () {
        toDateValue = $('#toDate').val();
        fromDateValue = $('#fromDate').val();
        if (toDateValue.length !=0 && fromDateValue.length !=0){
            console.log("hhhhhhh");
            if(document.getElementById('workDetail').clicked != true)
            {
                console.log("our main array",mainArray);
                for (var i:=0;i<mainArray.length;i++){
                    for (var j:=0;j<mainArray[i].length;j++){
                        
//                        if (mainArray[i][j])
//                        
                        
                        
                    }
                }
            }
        }
        
        

      
  });
    
    
    
    
    
    
    
    
    
   
   
});