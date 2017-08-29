
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
   
    var values =[[]];
     var mainArray = []; 
    var originalArray =[[]];
    var table = "";
    function createDataArray(values){
        var subArray = [];
        for(i = 1; i < values.length; i++) {
            console.log("array",values[i]);
            console.log("lengthhhh",values[i].length);
           
            for(var propertyName in values[i]) {
                
                     subArray.push(values[i][propertyName]);
                
            }
            mainArray.push(subArray);
            subArray = [];
            
        }
    }
    
    function dataTableManipulate(){
        table =  $("#timeSheet_details").DataTable({
            data: mainArray,
            "searching": true,
            "info": false,
            "lengthChange":false,
            "columnDefs": [{
//                       "targets": -1,
//                       "width": "5%",
//                       "data": null,
//                       "defaultContent": '<div class="edit-wrapper"><span class="icn"><i class="fa fa-pencil-square-o" aria-hidden="true" id="edit"></i><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
            }]
        });
    }
        
    
    
    
    
    
    
    
    

    if (vm.TaskDetails != null){
        for(var i=0;i<vm.TaskDetails.length;i++){
             var timeSlice =[];
            var userName =vm.TaskDetails[i][5];
            var taskName =vm.TaskDetails[i][2];
            var userId =vm.TaskDetails[i][4];
//            console.log("username",taskName);
            timeSlice.push(userName);
            timeSlice.push(taskName);
            var taskStartDate =vm.TaskDetails[i][0];
            var taskEndDate =vm.TaskDetails[i][1];
            var daysWorked =0;
            var extraHours ="";
            var sumExtraHours ="";
            var sumLateHours ="";
            var lateHours ="";
            var diffInStartTime ="";
            var leave =0;
            var sumOfLeave =0;
            var sumOfWorkingDays =0;
           
            var utcTimeOfStartDate = taskStartDate;
            var dbStartTime = parseInt(utcTimeOfStartDate)
            var d = new Date();
//            date = new Date(timestamp*1000 + d.getTimezoneOffset() * 60000)
            var startDate = new Date(dbStartTime * 1000+d.getTimezoneOffset() * 60000);
            var dd = startDate.getDate();
            var mm = startDate.getMonth() + 1; //January is 0!
            var yyyy = startDate.getFullYear();
            var HH = startDate.getHours();
            var min = startDate.getMinutes();
            var sec = startDate.getSeconds();
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
            var startTime = (HH + ':' + min);
//            console.log("start time",startTime);
            var startDateInFormat = (mm + '/' + dd + '/' + yyyy);
//            console.log("start date",startDateInFormat);
            var utcTimeOfEndDate = taskEndDate;
            var dbEndTime = parseInt(utcTimeOfEndDate)
            var d = new Date();
            var endDate = new Date(dbEndTime * 1000+d.getTimezoneOffset() * 60000);
            var dd = endDate.getDate();
            var mm = endDate.getMonth() + 1; //January is 0!
            var yyyy = endDate.getFullYear();
            var HH = endDate.getHours();
            var min = endDate.getMinutes();
            var sec = endDate.getSeconds();
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
            var endTime = (HH + ':' + min);
//            console.log("end time",endTime);
            var endDateInFormat =  (mm + '/' + dd + '/' + yyyy);
//            console.log("end date",endDateInFormat);
            for (var j=0;j<vm.LogArray.length;j++){
                for (var k=0; k<vm.LogArray[j].length ;k++){
                    if (userId ==vm.LogArray[j][k].UserID){
                        var utcTime = vm.LogArray[j][k].LogTime;
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
                        var localTime = (HH + ':' + min);
//                        console.log("log time",localTime);
                        var localDate = (mm + '/' + dd + '/' + yyyy);
                        var difference = taskEndDate - taskStartDate;
                        var noOfDaysWorked = Math.floor(difference/1000/60/60/24);
                        var localTimeDiff = new Date('00','00',localTime.split(':')[0],localTime.split(':')[1]);
                        var startTimeDiff = new Date('00','00',startTime.split(':')[0],startTime.split(':')[1]);
                        var endTimeDiff = new Date('00','00',endTime.split(':')[0],endTime.split(':')[1]);
                        var Splittedstart = startTime.split(":");
                        var Splittedend = endTime.split(":");
                        var SplittedlogTime =localTime.split(":");
                        var SplittedstartDate = new Date(0, 0, 0, Splittedstart[0], Splittedstart[1], 0);
                        var SplittedendDate = new Date(0, 0, 0, Splittedend[0], Splittedend[1], 0);
                        var SplittedlocalTime = new Date(0, 0, 0, Splittedend[0], Splittedend[1], 0);
                       
                        if (noOfDaysWorked ==0){
                            if (localDate ==startDateInFormat&& vm.LogArray[j][k].LogDescription =="Work Started"){
                                daysWorked =1;
                                leave=0;
                                sumOfLeave =leave+sumOfLeave;
                                if (localTime>startTime){
                                    var diffInStartTime = moment.utc(moment(localTime, "HH:mm").diff(moment(startTime, "HH:mm"))).format("HH:mm");
                                    sumLateHours =sumLateHours+diffInStartTime;
                                    lateHours =sumLateHours;
                                    console.log("late hours",lateHours); 
                                }
                            }
                            if (localDate ==startDateInFormat&& vm.LogArray[j][k].LogDescription =="Work Started"){
                                if (localTime<startTime){
                                    var  diffInStartTime =moment.utc(moment(startTime," HH:mm").diff(moment(localTime," HH:mm"))).format("HH:mm");
                                    sumExtraHours =sumExtraHours+diffInStartTime;
                                    extraHours =sumExtraHours;
                                    console.log("sum of extra hours",extraHours);
                                }
                                   
                                }
                               if (vm.LogArray[j][k].LogDescription =="End of work day"){
                                  if (localTime>endTime){
                                      var  diffInStartTime =moment.utc(moment(localTime," HH:mm").diff(moment(endTime," HH:mm"))).format("HH:mm");
                                      sumExtraHours =sumExtraHours+diffInStartTime;
                                      extraHours =sumExtraHours;
                                      console.log("break1",extraHours);
                                  } 
                               }
                            if (vm.LogArray[j][k].LogDescription =="End of work day"){
                                 if (localTime<endTime){
                                    diffInStartTime =moment.utc(moment(endTime," HH:mm").diff(moment(localTime," HH:mm"))).format("HH:mm");
                                    sumLateHours =sumLateHours+diffInStartTime;
                                    lateHours =sumLateHours;
                                     console.log("break2",lateHours);
                                }
                            }
                            if (vm.LogArray[j][k].LogDescription =="Completed"){
//                                console.log("enddddd");
                                  if (localTime>endTime){
                                      console.log("inside3");
                                      var  diffInStartTime =moment.utc(moment(localTime," HH:mm").diff(moment(endTime," HH:mm"))).format("HH:mm");
                                      sumExtraHours =sumExtraHours+diffInStartTime;
                                      extraHours =sumExtraHours;
                                      console.log("break3",extraHours);
                                  } 
                               }
                            
                            
                        }else{
                            
                                 if (localDate ==startDateInFormat){
                                     leave=0;
                                     sumOfLeave =leave+sumOfLeave;
                                     var dateTime1 = new Date(startDateInFormat).getTime();
                                     var  dateTime2 = new Date(endDateInFormat).getTime();
                                     var diff = dateTime2 - dateTime1;
                                     if (diff >= 0) {
                                         sumOfWorkingDays=sumOfWorkingDays+1;
                                         daysWorked =sumOfWorkingDays;
                                         if (localTimeDiff.getTime()>startTimeDiff.getTime()){
                                             diffInStartTime =moment.utc(moment(localTime," HH:mm").diff(moment(startTime," HH:mm"))).format("HH:mm");
                                             sumLateHours =sumLateHours+diffInStartTime;
                                             lateHours =sumLateHours;
                                         }else{
                                             diffInStartTime =moment.utc(moment(startTime," HH:mm").diff(moment(localTime," HH:mm"))).format("HH:mm");
                                             sumExtraHours =sumExtraHours+diffInStartTime;
                                             extraHours =sumExtraHours;
                                         }
                                         if (localTimeDiff.getTime()>endTimeDiff.getTime()){
                                             diffInStartTime =moment.utc(moment(localTime," HH:mm").diff(moment(endTime," HH:mm"))).format("HH:mm");
                                             sumExtraHours =sumExtraHours+diffInStartTime;
                                             extraHours =sumExtraHours;
                                         }else{
                                             diffInStartTime =moment.utc(moment(endTime," HH:mm").diff(moment(localTime," HH:mm"))).format("HH:mm");
                                             sumLateHours =sumLateHours+diffInStartTime;
                                             lateHours =sumLateHours;
                                         }
                                     }
                                 }else{
                                      leave=1;
                                     sumOfLeave =leave+sumOfLeave;
                                 }
                        }
                        var logExtraDay = new Date(localDate).getTime();
                        var extraEndDay =  new Date(endDateInFormat).getTime();
                        var extraDifference =logExtraDay-extraEndDay;
                        if (extraDifference>0){
                            extraWorkingDay =sumOfWorkingDays+1;
                            daysWorked =extraWorkingDay;
                            if (localTimeDiff.getTime()>startTimeDiff.getTime()){
                                    diffInStartTime =moment.utc(moment(localTime," HH:mm").diff(moment(startTime," HH:mm"))).format("HH:mm");
                                    sumLateHours =sumLateHours+diffInStartTime;
                                    lateHours =sumLateHours;
                                    
                                    
                                }else{
                                    diffInStartTime =moment.utc(moment(startTime," HH:mm").diff(moment(localTime," HH:mm"))).format("HH:mm");
                                    sumExtraHours =sumExtraHours+diffInStartTime;
                                    extraHours =sumExtraHours;
                                }
                                if (localTimeDiff.getTime()>endTimeDiff.getTime()){
                                    diffInStartTime =moment.utc(moment(localTime," HH:mm").diff(moment(endTime," HH:mm"))).format("HH:mm");
                                   sumExtraHours =sumExtraHours+diffInStartTime;
                                    extraHours =sumExtraHours;
                                    
                                    
                                }else{
                                    diffInStartTime =moment.utc(moment(endTime," HH:mm").diff(moment(localTime," HH:mm"))).format("HH:mm");
                                    sumLateHours =sumLateHours+diffInStartTime;
                                    lateHours =sumLateHours;
                                }
                        }
                    }
                }
                newDate      = new Date(startDateInFormat);
                startDateInFormat = new Date(newDate.setDate(newDate.getDate() + 1));
//                startDateInFormat =startDateInFormat.setDate(startDateInFormat.getDate()+1);
            }
            timeSlice.push(daysWorked);
            timeSlice.push(sumOfLeave);
            timeSlice.push(lateHours);
            timeSlice.push(extraHours);
            console.log("length of",timeSlice.length);
            if (timeSlice[0].length >0){
                 values.push(timeSlice);
            }
            console.log("final",values);
           
        }
        }
        
//    console.log("final",timeSlice);
//    console.log("values",values)
//  for (var p=0;p<values.length;p++){
//      for (var g=0;g<values[p].length;g++){
//          if (values[p][g].length !=0){
//              
//              
//          }
//      }
//  }
    
    if(values != null) {
        createDataArray(values);
    }
    dataTableManipulate(); 

//    dataTableManipulate(mainArray);   
//     $('#fromDate').change(function () {
//        selectFromDate = $('#fromDate').val();
//        var fromYear = selectFromDate.substring(6, 10);
//        var fromDay = selectFromDate.substring(3, 5);
//        var fromMonth = selectFromDate.substring(0, 2);
//        $('#toDate').datepicker("option", "minDate", new Date(fromYear, fromMonth-1, fromDay));
//        actualFromDate = new Date(selectFromDate);
//        actualFromDate.setHours(0);
//        actualFromDate.setMinutes(0);
//        actualFromDate.setSeconds(0);
//        unixFromDate = Date.parse(actualFromDate)/1000;
//        listLogDetails(unixFromDate,unixToDate);
//    });
//    
//    $('#toDate').change(function () {
//        selectedToDate = $('#toDate').val();
//        var year = selectedToDate.substring(6, 10);
//        var day = selectedToDate.substring(3, 5);
//        var month = selectedToDate.substring(0, 2);
//        $('#fromDate').datepicker("option", "maxDate", new Date(year, month-1, day));
//        actualToDate = new Date(selectedToDate);
//        actualToDate.setHours(23);
//        actualToDate.setMinutes(59);
//        actualToDate.setSeconds(59);
//        unixToDate = Date.parse(actualToDate)/1000;
//        listLogDetails(unixFromDate,unixToDate);
//    });
//    
});