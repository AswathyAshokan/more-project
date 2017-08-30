
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
    var timeSlice =[];
//    var mainArray = [];   
//    var table = "";
//    var unixFromDate = 0;
//    var unixToDate = 0;
//    var mainArray = [];   
//    var table = "";
//    var selectedToDate;
//    var actualToDate;
//    var selectFromDate;
//    var actualFromDate;
//    var completeTable =[];
//    var lattitude;
//    var longitude;
//    var Values =[][];
//    function createDataArray(values, keys){
//        console.log("inside create");
//        var subArray = [];
//        for(i = 0; i < values.length; i++) {
//            for(var propertyName in values[i]) {
//                subArray.push(values[i][propertyName]);
//            }
//            subArray.push(keys[i])
//            mainArray.push(subArray);
//            subArray = [];
//            
//        }
//    }
//    completeTable = mainArray;
//    $('#refreshButton').click(function(e) {
//        $('#log-details').dataTable().fnDestroy();
//        $('#fromDate').datepicker('setDate', null);
//        $('#toDate').datepicker('setDate', null);
//        dataTableManipulate(completeTable);
//     });
//    function listLogDetails(unixFromDate,unixToDate){
//        var tempArray = [];
//        var startDate =0;
//        var unixStartDate = 0;
//        for (i =0;i<vm.Values.length;i++){
//            startDate = new Date(vm.Values[i][6]);
//            unixStartDate = Date.parse(startDate)/1000;
//           if( (unixFromDate <= unixStartDate && unixStartDate <= unixToDate) || (unixFromDate <= unixStartDate && unixStartDate <= unixToDate) || (unixFromDate >= startDate && unixStartDate >= unixToDate)) {
//               
//                tempArray.push(mainArray[i]);
//           }
//            $('#timeSheet_details').dataTable().fnDestroy();
//            dataTableManipulate(tempArray);
//        }
//    } 
//    function dataTableManipulate(mainArray){
//        table =  $("#timeSheet_details").DataTable({
//            data: mainArray,
//            "searching": true,
//            "info": false,
//            "lengthChange":false,
//           "columnDefs": [{
//               
//           }]
//        });
//        $('#tbl_details_length').after($('.datepic-top'));
//        
//    }
//     if(vm.Values != null) {
//        for( i=0;i<vm.Values.length;i++){
//            var utcTime = vm.Values[i][3];
//            var utcInDateForm = new Date(utcTime);
//            console.log("utcInDateForm",utcInDateForm);
//            var localTime = (utcInDateForm.toLocaleTimeString());
//            var localDate = (utcInDateForm.toLocaleDateString());
//            var timeSplitArray = localTime.split(":");
//            var hours = timeSplitArray[0];
//            var minutes = timeSplitArray[1];
//            var duration = vm.Values[i][2];
//            var durationSplitArray = duration.split(":");
//            var duartionHours = durationSplitArray[0];
//            var durationMinutes = durationSplitArray[1];
//            var localTimeInMinutes = parseFloat(hours)*60 + parseFloat(minutes);
//            if (localTimeInMinutes>durationMinutes){
//                var loggedTime = localTimeInMinutes - durationMinutes
//                var loggedHours = window.parseInt(loggedTime/60);
//                var loggedMins = loggedTime%60;
//            }
//            var actualloggedTime =loggedHours +   ":" +loggedMins
//            var between = actualloggedTime + " &nbspto&nbsp" +hours +    ":"    +minutes;
//            vm.Values[i][2]= localDate;
//            vm.Values[i][3] = hours +    ":"    +minutes;
//            lattitude = vm.Values[i][4];
//            longitude= vm.Values[i][5];
//            
//        }
//         createDataArray(vm.Values, vm.Keys);
//    }
    if (vm.TaskDetails != null){
        for(var i=0;i<vm.TaskDetails.length;i++){
            var userName =vm.TaskDetails[i][5];
            var taskName =vm.TaskDetails[i][2];
            var userId =vm.TaskDetails[i][4];
            console.log("username",taskName);
            timeSlice.push(userName);
            timeSlice.push(taskName);
            var taskStartDate =vm.TaskDetails[i][0];
            var taskEndDate =vm.TaskDetails[i][1];
            var daysWorked =0;
            var extraHours =00:00;
            var sumExtraHours =00:00;
            var sumLateHours =00:00;
            var lateHours =00:00;
            var diffInStartTime =00:00;
            var leave =0;
            var sumOfLeave =0;
            var sumOfWorkingDays =0
            var startDateInFormat = moment.unix(taskStartDate).format("MM/DD/YYYY");
            var endDateInFormat =   moment.unix(taskEndDate).format("MM/DD/YYYY");
            var utcTimeOfStartDate = taskStartDate;
            var dbStartTime = parseInt(utcTimeOfStartDate)
            var startDate = new Date(dbStartTime * 1000);
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
            var utcTimeOfEndDate = taskEndDate;
            var dbEndTime = parseInt(utcTimeOfEndDate)
            var endDate = new Date(dbEndTime * 1000);
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
            var startTimeFromWorkLocation ="";
            var endTimeFromWorkLocation = "";
            for (var)
            
            
            
            
            
            for (var j=0;j<vm.LogArray.length;j++){
                for (var k=0; k<vm.LogArray[j].length ;k++){
                    if userId ==vm.LogArray[j][k].UserID{
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
                        var localDate = (mm + '/' + dd + '/' + yyyy);
                       
                        var difference = taskStartDate - taskEndDate;
                        var noOfDaysWorked = Math.floor(difference/1000/60/60/24);
                        console.log("days difference",noOfDaysWorked);
                        var startDate = new Date(taskStartDate*1000);
                        var startHours = startDate.getHours();
                        var startMinutes = "0" + startDate.getMinutes();
                        var formattedStartTime = startHours + ':' + startMinutes.substr(-2) ;
                        console.log("formated start time",formattedStartTime);
                        var endDate = new Date(taskEndDate*1000);
                        var endHours = endDate.getHours();
                        var endMinutes = "0" + endDate.getMinutes();
                        var formattedEndTime = startHours + ':' + startMinutes.substr(-2) ;
                        var localTimeDiff = new Date('00','00',localTime.split(':')[0],localTime.split(':')[1]);
                        var startTimeDiff = new Date('00','00',startTime.split(':')[0],startTime.split(':')[1]);
                        var endTimeDiff = new Date('00','00',endTime.split(':')[0],endTime.split(':')[1]);
                        console.log("formated end time",formattedEndTime);
                        if (noOfDaysWorked ==0){
                            if (localDate ==startDateInFormat){
                                daysWorked =1;
                                leave=0;
                                sumOfLeave =leave+sumOfLeave;
                               
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
                startDateInFormat =startDateInFormat.setDate(startDateInFormat.getDate()+1);
            }
            timeSlice.push(daysWorked);
        }
    }
    
    
    
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