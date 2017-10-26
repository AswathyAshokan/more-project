console.log("company name",vm);
  $(document).ready(function(){
    document.getElementById("dashBoard").className += " active";

    //for notification
      var PersentageOfStartedUser;
      var PersentageOfCompletedUsers;
      var persentageOfPendingUsers;
      var persentageOfStartedUserOnly;
      var tempStart;
      var returnString;
      var DynamicNotification ="";
      var TotalNoUsers;
      var today;
      var allData = [[]];
      var  dateIncrementDate =0;
      var subArray = [];
      
      //set default value on job 
      var temp="All"; 
      $("#jobName").val(temp);
      //select all in drop down
       if (vm.TaskDetailArray !=null){
         for(i = 0; i < vm.TaskDetailArray.length; i++) {
             subArray.push(vm.TaskDetailArray[i][1]);
            
        }  
       }
      var DynamicTaskListing ="";
      console.log("sub array",subArray);
      for (var i=0; i<subArray.length; i++){
          DynamicTaskListing+=' <p onclick="FunctionToChangeBarChart(event) " style="cursor:pointer;" class="active" >'+subArray[i]+'</p>'; 
          console.log("jjjjjjj",DynamicTaskListing)
         
      }
       $("#taskListing").append(DynamicTaskListing);
          //end
      
      
      
      function LoadBarChart(total,start,pending,complete,todayVal){
           document.getElementById('today').innerHTML = todayVal;
          $.jqplot.config.enablePlugins = true;
          s1 =[0,0,0,0]
          var  max_count = total ; 
          var plot1 = $.jqplot('chart1', [s1]);
          plot1.destroy();
          
          console.log('vvv', max_count);
          var s1 = [total, start, pending, complete];
            var ticks = ['total', 'started', 'pending','completed' ];
            plot1 = $.jqplot('chart1', [s1], {
                // Only animate if we're not using excanvas (not in IE 7 or IE 8)..
                
                animate: !$.jqplot.use_excanvas,
                seriesDefaults:{
                    renderer:$.jqplot.BarRenderer,
                    rendererOptions: {barMargin: 0 , varyBarColor : true},
                    pointLabels: { show: true }
                },
                title:{text:"Daily Task Status Report"},
                grid: {
                    background: 'transparent',      // CSS color spec for background color of grid.
                    drawBorder:false,
                    drawGridlines:false,
                    shadow:false
                },
                axes: {
                    xaxis: {
                        renderer: $.jqplot.CategoryAxisRenderer,
                        ticks: ticks,
                        
                        tickOptions : {
                              
                                          //  showGridline : false
                        }
                    },
                    yaxis: {
                        tickOptions : {
                            formatString:'%d',
                            formatter: $.jqplot.euroFormatter
                             
                        },
                        labelRenderer: $.jqplot.CanvasAxisLabelRenderer,
                        min: 0,
                        tickInterval: 1,
                       
                    }
                },
                seriesColors: [ "#000", "#ccc", "red","green"],
                highlighter: { show: false }
            });
            $('#chart1').bind('jqplotDataClick',
                        function (ev, seriesIndex, pointIndex, data) {$('#info1').html('series: '+seriesIndex+', point: '+pointIndex+', data: '+data);
            });
      }
      
      
    if (vm.NotificationNumber !=0){
        document.getElementById("number").textContent=vm.NotificationArray.length;
    }else{
        document.getElementById("number").textContent="";
    }
    
    
      var companyTeamName =vm.CompanyTeamName
   if(vm.CompanyPlan == 'family' ){
       $("#contact").remove();
       $("#crm").remove();
       $("#leave").remove();
       $("#fitToWork").remove();
       $("#time-sheet").remove();
       $("#consent").remove();
   } else if(vm.CompanyPlan == 'campus'){
       $("#contact").remove();
       $("#crm").remove();
       $("#leave").remove();
       $("#fitToWork").remove();
       $("#time-sheet").remove();
       $("#consent").remove();
    }
    document.getElementById("username").textContent=vm.AdminFirstName;
    document.getElementById("imageId").src=vm.ProfilePicture;
    if (vm.ProfilePicture ==""){
        document.getElementById("imageId").src="/static/images/default.png"
    }
    if(vm.CompanyPlan == "family"){
        $('#planChange').attr('data-target','#family');
    } else if (vm.CompanyPlan == "campus") {
        $('#planChange').attr('data-target','#campus');
    }else if (vm.CompanyPlan == "business") {
        $('#planChange').attr('data-target','#business');
    }else if (vm.CompanyPlan == "businessPlus") {
        $('#planChange').attr('data-target','#business-plus');
    }
  if(vm.CompletedTask ==0 &&vm.PendingTask ==0){
       jQuery("#pie1").radialPieChart("init", {
           'font-size': 13,
           'fill': 25,
           "size": 150,
           'text-color': "transparent",
           'data': [
               {'color': "#363433 ", 'perc': 100}
           ]
       });
   }else {
       jQuery("#pie1").radialPieChart("init", {
           'font-size': 13,
           'fill': 25,
           "size": 150,
           'text-color': "transparent",
           'data': [
               {'color': "#29a0ff ", 'perc': vm.CompletedTask},
               {'color': " #008000 ", 'perc': vm.PendingTask}
           ]
       });
   }
   if(vm.PendingUsers ==0 && vm.AcceptedUsers ==0 && vm.RejectedUsers ==0){
       jQuery("#pie2").radialPieChart("init", {
           'font-size': 13,
           'fill': 25,
           "size": 150,
           'text-color': "transparent",
           'data': [
               {'color': "#363433 ", 'perc': 100 }
               
           ]
       });
       
   }else {
       jQuery("#pie2").radialPieChart("init", {
           'font-size': 13,
           'fill': 25,
           "size": 150,
           'text-color': "transparent",
           'data': [
               {'color': "#5b93c2 ", 'perc': vm.PendingUsers },
               {'color': "#06599e ", 'perc': vm.AcceptedUsers},
               {'color': "#696969 ", 'perc':vm.RejectedUsers}
           ]
       });
   }
      
       //notification
   //notification
    var DynamicNotification ="";
    
    var expirycount =0;
    var documentNotifyArry =  vm.DocumentExpiryNotification;
    if( vm.DocumentExpiryNotification!=null){
        for(i = 0;i<vm.DocumentExpiryNotification.length;i++){
                console.log("haiiiii");
                console.log("document values",documentNotifyArry[i]);
                var today = new Date();
                var dd = today.getDate();
                var mm = today.getMonth()+1; //January is 0!
                var yyyy = today.getFullYear();
                if(dd<10) {
                    dd = '0'+dd;
                } 
                if(mm<10) {
                    mm = '0'+mm;
                }
                
                var tempArray = [];
                
                var CurrentMonth = mm;
                var currentDay = dd;
                var currentYear = yyyy;
                var localToday = (mm + '/' + dd + '/' + yyyy);
                var dateParts = documentNotifyArry[i][3].split("/");
                var dateFromDb = (dateParts[1]+'/'+ dateParts[0]+'/'+ dateParts[2]);
                if(CurrentMonth ==dateParts[1] && currentDay ==dateParts[0] &&currentYear ==  dateParts[2]){
                    if(documentNotifyArry[i][2] == "false"){
                       expirycount = expirycount+1;
                    }
            }
        }
        
    }
    console.log("count",vm.NotificationNumber+expirycount);
    vm.NotificationNumber = vm.NotificationNumber+expirycount;
    if (vm.NotificationNumber !=0){
        document.getElementById("number").textContent=vm.NotificationNumber;
    }else{
        document.getElementById("number").textContent="";
    }
    

var formDataClear = [[]];
var allIdArray = [[]];
var allExpiryId = [[]];
var allUserId = [[]];
    
    var notificationSorted =[[]];
    function sortByCol(arr, colIndex){
        notificationSorted=arr.sort(sortFunction);
        function sortFunction(a, b) {
            a = a[colIndex]
            b = b[colIndex]
            return (a === b) ? 0 : (a < b) ? -1 : 1
        }
    }
    
    myNotification= function () {
        if (vm.NotificationArray !=null){
            sortByCol(vm.NotificationArray, 6);
            var reverseSorted =[[]];
            reverseSorted=notificationSorted.reverse();
            document.getElementById("notificationDiv").innerHTML = "";
            var DynamicTaskListing="";
            if (reverseSorted !=null){
                DynamicTaskListing ="<h5>"+"Notifications"+ "<button class='no-button-style' method='post' onclick='clearNotification()'>"+"clear all"+"</button>"+"</h5>"+"<ul>";
                for(var i=0;i<reverseSorted.length;i++){
                    if(reverseSorted[i].length != 0){
                        if (reverseSorted[i][5]==""){
                             console.log("cpp1");
                            console.log("iam in first");
                            var utcTime =reverseSorted[i][3];
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
                            var startDate = (mm + '/' + dd + '/' + yyyy);
                            var utcTime =reverseSorted[i][4];
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
                                var endDate = (mm + '/' + dd + '/' + yyyy);
                                var timeDifferenceForLeave =moment(new Date(new Date(reverseSorted[i][6]*1000)), "YYYYMMDD").fromNow();
                                 DynamicTaskListing += "<li>"+"User"+" "+reverseSorted[i][2]+" "+"Applied Leave For "+"  "+reverseSorted[i][7]+" "+"Days"+" "+"From"+" "+startDate+" "+"to"+" "+endDate+" <span>"+timeDifferenceForLeave+"</span>"+"</li>";
                        } else if (reverseSorted[i][5] =="Expiry111@@&&EEE"){
                            console.log("cpp2");
                            var utcTime =reverseSorted[i][7];
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
                            var expiryDate = (mm + '/' + dd + '/' + yyyy);
                            var currentDate = reverseSorted[i][3].split("/");
                            console.log("currentDate[1]",currentDate[1]);
                            console.log("mm",mm);
                            console.log("currentDate[0]",currentDate[0]);
                            console.log("dd",dd);
                            console.log("currentDate[2]",currentDate[2]);
                            console.log("yyyy",reverseSorted[i][3]);
                           
                            if(currentDate[0] == mm && currentDate[1] ==dd && currentDate[2] == yyyy  ){
                                
                                var timeDifferenceForLeave =moment(new Date(new Date(reverseSorted[i][6]*1000)), "YYYYMMDD").fromNow();
                                DynamicTaskListing += "<li>"+"User"+" "+reverseSorted[i][0]+"'s"+" "+reverseSorted[i][4]+ " "+ "expired"+" "+"("+ expiryDate+")"+"<span>"+timeDifferenceForLeave+"</span>"+"</li>";
                                
                            } else{
                               
                                var timeDifferenceForLeave =moment(new Date(new Date(reverseSorted[i][6]*1000)), "YYYYMMDD").fromNow();
                                DynamicTaskListing += "<li>"+"User"+" "+reverseSorted[i][0]+"'s"+reverseSorted[i][4]+ " "+ "will be expired on" +" "+expiryDate+
                                    "<span>"+timeDifferenceForLeave+"</span>"+"</li>";
                            }
                        
                        } else if(reverseSorted[i][5] == "WorkLocationt!@#$%&*YTREFFDD"){
                            var timeDifference =moment(new Date(new Date(reverseSorted[i][6]*1000)), "YYYYMMDD").fromNow();
                            if(reverseSorted[i][8] == "After" ){
                               DynamicTaskListing += "<li>"+"User"+" "+reverseSorted[i][2]+" "+" will be reaching within"+" "+reverseSorted[i][3]+"  "+"at work location "+" "+reverseSorted[i][4]+" <span>"+timeDifference+"</span>"+"</li>"; 
                            } else{
                                DynamicTaskListing += "<li>"+"User"+" "+reverseSorted[i][2]+" "+"will be delayed "+reverseSorted[i][3]+"  "+"to reach at work location"+" "+reverseSorted[i][4]+" <span>"+timeDifference+"</span>"+"</li>";
                            }
                        }else{
                            var timeDifference =moment(new Date(new Date(reverseSorted[i][6]*1000)), "YYYYMMDD").fromNow();
                             if(reverseSorted[i][7] == "After" ){
                                 DynamicTaskListing += "<li>"+"User"+" "+reverseSorted[i][2]+" "+" will be reaching within"+" "+reverseSorted[i][3]+"  "+" at task location"+" "+reverseSorted[i][4]+" "+"for the task"+" "+reverseSorted[i][5]+" <span>"+timeDifference+"</span>"+"</li>";
                            } else{
                                DynamicTaskListing += "<li>"+"User"+" "+reverseSorted[i][2]+" "+"will be delayed"+" "+reverseSorted[i][3]+"  "+" to reach the task location "+" "+reverseSorted[i][4]+" "+"for the task"+" "+reverseSorted[i][5]+" <span>"+timeDifference+"</span>"+"</li>";
                            }
                        }
                    }
                    
                }
                $("#notificationDiv").prepend(DynamicTaskListing+"</ul>");
                document.getElementById("number").textContent="";
                $.ajax({
                    url:'/'+ companyTeamName + '/notification/update',
                    type: 'post',
                    data:formDataClear,
                    success : function(response) {
                        if (response == "true" ) {
                        } else {
                        }
                    },
                    error: function (request,status, error) {
                        console.log(error);
                    }
                }); 
            }else{
                document.getElementById("notificationDiv").innerHTML = "";
                DynamicTaskListing ="<h5>"+" No New Notifications"+"</h5>";
                $("#notificationDiv").prepend(DynamicTaskListing);
            }
        }else{
            document.getElementById("notificationDiv").innerHTML = "";
            DynamicTaskListing ="<h5>"+" No New Notifications"+"</h5>";
            $("#notificationDiv").prepend(DynamicTaskListing);
        }
    }
    clearNotification= function () {
        document.getElementById("notificationDiv").innerHTML = "";
        $.ajax({
            url:'/'+ companyTeamName + '/notification/delete',
            type: 'post',
            data :formDataClear,
            success : function(response) {
                if (response == "true" ) {
                    DynamicTaskListing ="<h5>"+" No New Notifications"+"</h5>";
                    $("#notificationDiv").prepend(DynamicTaskListing);
                } else {
                }
            },
            error: function (request,status, error) {
                console.log(error);
            }
        }); 
    }
    
    
    //this is for notification of expiredDetails
    var allexpiryNotification = [[]];
    var tempIdArry =[];
    var chechNotification;
    var documentNotifyArry =  vm.DocumentExpiryNotification;
        if( vm.DocumentExpiryNotification!=null){
            for(i = 0;i<vm.DocumentExpiryNotification.length;i++){
                console.log("haiiiii");
                console.log("document values",documentNotifyArry[i]);
                var today = new Date();
                var dd = today.getDate();
                var mm = today.getMonth()+1; //January is 0!
                var yyyy = today.getFullYear();
                if(dd<10) {
                    dd = '0'+dd;
                } 
                if(mm<10) {
                    mm = '0'+mm;
                }
                
                var tempArray = [];
                
                var CurrentMonth = mm;
                var currentDay = dd;
                var currentYear = yyyy;
                var localToday = (mm + '/' + dd + '/' + yyyy);
                var dateParts = documentNotifyArry[i][3].split("/");
                var dateFromDb = (dateParts[1]+'/'+ dateParts[0]+'/'+ dateParts[2]);
                if(CurrentMonth ==dateParts[1] && currentDay ==dateParts[0] &&currentYear ==  dateParts[2]){
                    
                    chechNotification = documentNotifyArry[i][0]+"111@@&&EEE";
                    tempArray.push(documentNotifyArry[i][5]);
                    tempArray.push(documentNotifyArry[i][1]);
                    tempArray.push(documentNotifyArry[i][2]);
                    tempArray.push(localToday);
                    tempArray.push(documentNotifyArry[i][6]);
                    tempArray.push(chechNotification);
                    tempArray.push(documentNotifyArry[i][4]);
                    tempArray.push(documentNotifyArry[i][7]);
                    tempArray.push(documentNotifyArry[i][8]);
                    allexpiryNotification.push(tempArray);
                    if(vm.NotificationArray !=null){
                        console.log("oh my goddddddddd");
                        vm.NotificationArray.push(tempArray);
                    }
                    tempArray = [];
                    allIdArray.push(documentNotifyArry[i][8]);
                    allExpiryId.push(documentNotifyArry[i][1]);
                    allUserId.push(documentNotifyArry[i][9]);
                }
                
                
            }
           if(vm.NotificationArray == null) {
               vm.NotificationArray = allexpiryNotification;
               allexpiryNotification = [[]];
           }
            
        }
    for(var i = 0; i<allIdArray.length;i++){
        formDataClear = formDataClear+"&DeletedId="+allIdArray[i];
    }
    for(var i = 0; i<allExpiryId.length;i++){
        formDataClear = formDataClear+"&DeletedExpiryId="+allExpiryId[i];
    }
    for(var i = 0; i<allUserId.length;i++){
        formDataClear = formDataClear+"&DeletedUserId="+allUserId[i];
    }
    
    
   getTaskDetails = function(){
        $("#taskListing").html("");
        var job = $("#jobName option:selected").val() ;
       if (vm.TaskDetailArray !=null){
          for(i = 0; i < vm.TaskDetailArray.length; i++) {
            if (vm.TaskDetailArray[i][0]==job) {
                subArray.push(vm.TaskDetailArray[i][1]);
            }
        } 
       }
        
        //select all in drop down
       if (vm.TaskDetailArray !=null){
         for(i = 0; i < vm.TaskDetailArray.length; i++) {
            if (job =="All") {
                subArray.push(vm.TaskDetailArray[i][1]);
            }
        }  
       }
        
        var DynamicTaskListing ="";
        for (var i=0; i<subArray.length; i++){
DynamicTaskListing+=' <p onclick="FunctionToChangeBarChart(event) " style="cursor:pointer;" class="active" >'+subArray[i]+'</p>';        }
        $("#taskListing").append(DynamicTaskListing);
        subArray = [];
    }
    var selectAJob = $("#jobName option:selected").val() ;
    console.log("default job",selectAJob);
      if (vm.TaskDetailArray !=null){
           for(i = 0; i < vm.TaskDetailArray.length; i++) {
        if (selectAJob =="SelectAJob") {
            subArray = [];
//            subArray.push(vm.TaskDetailArray[i][1]);
        }
      }
   
    }
//    var DynamicTaskListing ="";
//    for (var i=0; i<subArray.length; i++){
//DynamicTaskListing+=' <p onclick="FunctionToChangeBarChart(event) " style="cursor:pointer;" class="active" >'+subArray[i]+'</p>';    }
//    $("#taskListing").prepend(DynamicTaskListing);
    
    FunctionToChangeBarChart = function(event){
        var TaskName = $(event.target).text();
        console.log($(event.target).text());
        var formData = formData+"&TaskName="+TaskName;
        $.ajax({
            url:'/' + companyTeamName +'/dashboard/barchart',
            type:'post',
            //dataType: 'json',
            data: formData,
            //call back or get response here
            success : function(data){
                var jsonData = JSON.parse(data)
                console.log("data",jsonData);
                allData = jsonData
                console.log("allData",allData)
                if(jsonData[0] == "true"){
                    totalNoUsers = jsonData[5];
                    today = new Date();
                    console.log("today   $$$$$$$$$$$$$",today);
                    var dd = today.getDate();
                    var mm = today.getMonth()+1; //January is 0!
                    var yyyy = today.getFullYear();
                    if(dd<10) {
                        dd = '0'+dd
                    } 

                    if(mm<10) {
                        mm = '0'+mm
                    }
                    var CurrentMonth = mm;
                    var currentDay = dd;
                    var currentYear = yyyy;
                    var localToday = (mm + '/' + dd + '/' + yyyy);
                    
                    console.log("todayDate",localToday);
                    //for filtaring details of task started User
                    var startTaskArray = jsonData[1];
                    var startTaskCount = 0;
                    var tempArray = [];
                    if (startTaskArray !=null){
                        for (i = 0;i<startTaskArray.length;i++){
                            console.log("inner loop of ",startTaskArray[i][2]);
                             /*var returnValues = checkStartedUser(startTaskArray[i][3]);
                            if(returnValues =="true"){
                                startTaskCount =startTaskCount+1;
                            }*/
                            
                           // tempArray.push()
                            var startTaskDate = startTaskArray[i][1];
                            var startTaskDateFromDb = parseInt(startTaskDate)
                            var d = new Date(startTaskDateFromDb * 1000);
                            var dd = d.getDate();
                            var mm = d.getMonth() + 1; //January is 0!
                            var yyyy = d.getFullYear();
                            if (dd < 10) {
                                dd = '0' + dd;
                            }
                            if (mm < 10) {
                                mm = '0' + mm;
                            }
                            if (mm == CurrentMonth && currentDay == dd && currentYear == yyyy ){
                                tempArray.push(startTaskArray[i][2])
                               //startTaskCount = startTaskCount+1;
                            }
                            //console.log("startTaskCount 111",tempArray)
                        }
                    }
                    var uniqueArry = Array.from(new Set(tempArray));
                    console.log("uniqueArry",uniqueArry);
                    startTaskCount = uniqueArry.length;
                    
                    
                    //for filtering of Completed task
                    
                    var completedTask = jsonData[2];
                    var completedTaskCount = 0;
                    if (completedTask !=null){
                        for (i = 0;i<completedTask.length;i++){
                            var completedTaskDate = completedTask[i][1];
                            var completedTaskDateFromDb = parseInt(completedTaskDate)
                            var d = new Date(completedTaskDateFromDb * 1000);
                            var dd = d.getDate();
                            var mm = d.getMonth() + 1; //January is 0!
                            var yyyy = d.getFullYear();
                            if (dd < 10) {
                                dd = '0' + dd;
                            }
                            if (mm < 10) {
                                mm = '0' + mm;
                            }
                            if (mm == CurrentMonth && currentDay == dd && currentYear == yyyy ){
                               completedTaskCount = completedTaskCount+1;
                            }
                            console.log("completedTaskCount",completedTaskCount)
                        }
                    }
                    //for filtering of pending Task
                    var pendingTask = jsonData[3];
                    var pendingTaskCount = 0;
                    if (pendingTask !=null){
                        pendingTaskCount = pendingTask.length;
                    }
                    if(startTaskCount>completedTaskCount){
                        tempStart = startTaskCount - completedTaskCount;
                    } else{
                        tempStart = completedTaskCount -startTaskCount;
                    }
                   
                    LoadBarChart(totalNoUsers,tempStart,pendingTaskCount,completedTaskCount,localToday);
                    if(allData[4]!=null){
                        var startTaskDate = allData[4][0];
                        var startTaskDateUnix = parseInt(startTaskDate);
                        var d = new Date(startTaskDateUnix * 1000);
                        var dd = d.getDate();
                        var mm = d.getMonth() + 1; //January is 0!
                        var yyyy = d.getFullYear();
                        if (dd < 10) {
                            dd = '0' + dd;
                        }
                        if (mm < 10) {
                            mm = '0' + mm;
                        }
                        var LocalTaskStartDate = (mm+'/'+dd+'/'+yyyy);
                    }
                    /*if(localToday !=LocalTaskStartDate){
                        var previousDay = document.getElementById('previousDay');
                        previousDay.style.visibility = 'visible';
                    }*/
                }
                else{
                    console.log("Server Problem");
                }
            },
            error: function (request,status, error) {
            }
        });
       
    }
    
    getPreviousDayValues = function(Event){
        var d = new Date();
        dateIncrementDate = dateIncrementDate+1;
        d.setDate(d.getDate() - dateIncrementDate);
        console.log("yesterDay nnnnnn",d)
        /*var nextDay = document.getElementById('nextDay');
        nextDay.style.visibility = 'visible';*/
        var dd = d.getDate();
        var mm = d.getMonth()+1; //January is 0!
        var yyyy = d.getFullYear();
        if(dd<10) {
            dd = '0'+dd
        } 

        if(mm<10) {
            mm = '0'+mm
        }
        var CurrentMonth = mm;
        var currentDay = dd;
        var currentYear = yyyy;
        var localToday = (mm + '/' + dd + '/' + yyyy);
        
        if(allData[4]!=null){
            var startTaskDate = allData[4][0];
            var startTaskDateUnix = parseInt(startTaskDate);
            var d = new Date(startTaskDateUnix * 1000);
            var dd = d.getDate();
            var mm = d.getMonth() + 1; //January is 0!
            var yyyy = d.getFullYear();
            if (dd < 10) {
                dd = '0' + dd;
            }
            if (mm < 10) {
                mm = '0' + mm;
            }
            var LocalTaskStartDate = (mm+'/'+dd+'/'+yyyy);
            
        }
        console.log("LocalTaskStartDate",LocalTaskStartDate);
        console.log("localToday",localToday);
        
        /*if (LocalTaskStartDate ==localToday ){
            console.log("we are in if success condition");
           var previousDay = document.getElementById('previousDay');
            previousDay.style.visibility = 'hidden';
        }*/
        
        if(allData[4]!=null){
            var endTaskDate = allData[4][1];
            var endTaskDateUnix = parseInt(endTaskDate);
            var d = new Date(endTaskDateUnix * 1000);
            var dd = d.getDate();
            var mm = d.getMonth() + 1; //January is 0!
            var yyyy = d.getFullYear();
            if (dd < 10) {
                dd = '0' + dd;
            }
            if (mm < 10) {
                mm = '0' + mm;
            }
            var LocalTaskEndDate = (mm+'/'+dd+'/'+yyyy);
        }
        /*if(LocalTaskEndDate == localToday ){
            console.log("iam  here at lasrt");
            var previousDay = document.getElementById('nextDay');
            previousDay.style.visibility = 'hidden';
        }*/
        
        
       // for(var i=0;i<allData.length;i++){
        var totalUsers = allData[5]
        /*for(var k=0;k<allData[i].length;k++){*/
        var startTaskArray = allData[1];
        var startTaskCount = 0;
        
        if (startTaskArray !=null){
           for (i = 0;i<startTaskArray.length;i++){
                console.log("inner loop of ",startTaskArray[i]);
                 /*var returnValues = checkStartedUser(startTaskArray[i][3]);
                if(returnValues =="true"){
                    startTaskCount =startTaskCount+1;
                }*/

               // tempArray.push()
                var startTaskDate = startTaskArray[i][1];
                var startTaskDateFromDb = parseInt(startTaskDate)
                var d = new Date(startTaskDateFromDb * 1000);
                var dd = d.getDate();
                var mm = d.getMonth() + 1; //January is 0!
                var yyyy = d.getFullYear();
                if (dd < 10) {
                    dd = '0' + dd;
                }
                if (mm < 10) {
                    mm = '0' + mm;
                }
                if (mm == CurrentMonth && currentDay == dd && currentYear == yyyy ){

                   startTaskCount = startTaskCount+1;
                }
                console.log("startTaskCount 111",startTaskCount)
           }
        }
            
            var completedTask = allData[2];
            var completedTaskCount = 0;
            if (completedTask !=null){
                for (i = 0;i<completedTask.length;i++){
                    var completedTaskDate = completedTask[i][1];
                    var completedTaskDateFromDb = parseInt(completedTaskDate)
                    var d = new Date(completedTaskDateFromDb * 1000);
                    var dd = d.getDate();
                    var mm = d.getMonth() + 1; //January is 0!
                    var yyyy = d.getFullYear();
                    if (dd < 10) {
                        dd = '0' + dd;
                    }
                    if (mm < 10) {
                        mm = '0' + mm;
                    }
                    if (mm == CurrentMonth && currentDay == dd && currentYear == yyyy ){
                       completedTaskCount = completedTaskCount+1;
                    }
                    console.log("completedTaskCount",completedTaskCount)
                }
            }
            var pendingTask = allData[3];
            var pendingTaskCount = 0;
            if (pendingTask !=null){
                pendingTaskCount = pendingTask.length;
            }
            if(startTaskCount>completedTaskCount){LoadBarChart
                tempStart = startTaskCount - completedTaskCount;
            } else{
                tempStart = completedTaskCount -startTaskCount;
            }
        
        LoadBarChart(totalUsers,tempStart,pendingTaskCount,completedTaskCount,localToday);
    }
    
    
    LoadNextDayValues = function(Event){
        var today= new Date();
        today = new Date();
        console.log("today   $$$$$$$$$$$$$",today);
        var dd = today.getDate();
        var mm = today.getMonth()+1; //January is 0!
        var yyyy = today.getFullYear();
        if(dd<10) {
            dd = '0'+dd
        } 

        if(mm<10) {
            mm = '0'+mm
        }
        var CurrentMonth = mm;
        var currentDay = dd;
        var currentYear = yyyy;
        var localToday = (mm + '/' + dd + '/' + yyyy);
        var totalUsers = allData[5]
            /*for(var k=0;k<allData[i].length;k++){*/
        var startTaskArray = allData[1];
        var startTaskCount = 0;

        if (startTaskArray !=null){
           for (i = 0;i<startTaskArray.length;i++){
                console.log("inner loop of ",startTaskArray[i]);
                 /*var returnValues = checkStartedUser(startTaskArray[i][3]);
                if(returnValues =="true"){
                    startTaskCount =startTaskCount+1;
                }*/

               // tempArray.push()
                var startTaskDate = startTaskArray[i][1];
                var startTaskDateFromDb = parseInt(startTaskDate)
                var d = new Date(startTaskDateFromDb * 1000);
                var dd = d.getDate();
                var mm = d.getMonth() + 1; //January is 0!
                var yyyy = d.getFullYear();
                if (dd < 10) {
                    dd = '0' + dd;
                }
                if (mm < 10) {
                    mm = '0' + mm;
                }
                if (mm == CurrentMonth && currentDay == dd && currentYear == yyyy ){

                   startTaskCount = startTaskCount+1;
                }
                console.log("startTaskCount 111",startTaskCount)
           }
        }

        var completedTask = allData[2];
        var completedTaskCount = 0;
        if (completedTask !=null){
            for (i = 0;i<completedTask.length;i++){
                var completedTaskDate = completedTask[i][1];
                var completedTaskDateFromDb = parseInt(completedTaskDate)
                var d = new Date(completedTaskDateFromDb * 1000);
                var dd = d.getDate();
                var mm = d.getMonth() + 1; //January is 0!
                var yyyy = d.getFullYear();
                if (dd < 10) {
                    dd = '0' + dd;
                }
                if (mm < 10) {
                    mm = '0' + mm;
                }
                if (mm == CurrentMonth && currentDay == dd && currentYear == yyyy ){
                   completedTaskCount = completedTaskCount+1;
                }
                console.log("completedTaskCount",completedTaskCount)
            }
        }
        var pendingTask = allData[3];
        var pendingTaskCount = 0;
        if (pendingTask !=null){
            pendingTaskCount = pendingTask.length;
        }
        if(startTaskCount>completedTaskCount){
            tempStart = startTaskCount - completedTaskCount;
        } else{
            tempStart = completedTaskCount -startTaskCount;
        }
        LoadBarChart(totalUsers,tempStart,pendingTaskCount,completedTaskCount,localToday);
         /*var previousDay = document.getElementById('previousDay');
        previousDay.style.visibility = 'visible';*/
        
        
        if(allData[4]!=null){
            var endTaskDate = allData[4][1];
            var endTaskDateUnix = parseInt(endTaskDate);
            var d = new Date(endTaskDateUnix * 1000);
            var dd = d.getDate();
            var mm = d.getMonth() + 1; //January is 0!
            var yyyy = d.getFullYear();
            if (dd < 10) {
                dd = '0' + dd;
            }
            if (mm < 10) {
                mm = '0' + mm;
            }
            var LocalTaskEndDate = (mm+'/'+dd+'/'+yyyy);
            
        }
        /*if(LocalTaskEndDate == localToday ){
            var previousDay = document.getElementById('nextDay');
            previousDay.style.visibility = 'hidden';
        }*/
        
        
        
       
    }
    
    
    
  });