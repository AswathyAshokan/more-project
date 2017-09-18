console.log("company name",vm.CompanyTeamName);
$(function () {
    var companyTeamName =vm.CompanyTeamName
    if(vm.CompanyPlan == 'family' ){
        var parent = document.getElementById("menuItems");
        var contact = document.getElementById("contact");
        var job = document.getElementById("job");
        var crm = document.getElementById("crm");
        var leave = document.getElementById("leave");
        var time  = document.getElementById("time-sheet");
        var consent = document.getElementById("consent");
         var workLocation = document.getElementById("workLocation")
         parent.removeChild(workLocation);
        parent.removeChild(time);
        parent.removeChild(consent);
        parent.removeChild(leave);
        parent.removeChild(contact);
        parent.removeChild(job);
        parent.removeChild(crm);
    } else if(vm.CompanyPlan == 'campus'){
            var campusParent = document.getElementById("menuItems");
            var contact = document.getElementById("contact");
            var job = document.getElementById("job");
            var crm = document.getElementById("crm");
            var leave = document.getElementById("leave");
            var time  = document.getElementById("time-sheet");
            var consent = document.getElementById("consent")
            var workLocation = document.getElementById("workLocation")
            parent.removeChild(workLocation);
            campusParent.removeChild(time);
            campusParent.removeChild(consent);
            campusParent.removeChild(leave);
            campusParent.removeChild(contact);
            campusParent.removeChild(job);
            campusParent.removeChild(crm);
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
    if(vm.CompletedTask ==0 &&vm.PendingTask ==0){
        jQuery("#pie1").radialPieChart("init", {
            'font-size': 13,
            'fill': 25,
            "size": 150,
            'text-color': "transparent",
            'data': [
                {'color': "#363433", 'perc': 100}
            ]
        });
    }else {
        jQuery("#pie1").radialPieChart("init", {
            'font-size': 13,
            'fill': 25,
            "size": 150,
            'text-color': "transparent",
            'data': [
                {'color': "#29a0ff", 'perc': vm.CompletedTask},
                {'color': "#6abdff", 'perc': vm.PendingTask}
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
                {'color': "#363433", 'perc': 100 }
                
            ]
        });
        
    }else {
        jQuery("#pie2").radialPieChart("init", {
            'font-size': 13,
            'fill': 25,
            "size": 150,
            'text-color': "transparent",
            'data': [
                {'color': "#5b93c2", 'perc': vm.PendingUsers },
                {'color': "#06599e", 'perc': vm.AcceptedUsers},
                {'color': "#8fb4d3", 'perc':vm.RejectedUsers}
            ]
        });
    }
    window.onload = function () {
        CanvasJS.addColorSet("colors",
                             [
            "#857198"
        ]);
        var chart = new CanvasJS.Chart("chartContainer", {
            height: 435,
            backgroundColor: "transparent",
            colorSet: "colors",
            axisY:{
                title: "Status",
                titleFontSize: 14,
                lineThickness: 1,
                gridThickness: 0,
                labelFontSize: 14,
                },
                axisX:{
                    title: "Users",
                    titleFontSize: 14,
                    lineThickness: 1,
                    labelFontSize: 14,
                    },
            data: [{
                type: "column",
                dataPoints: [
                    { y: 22, label: "User 1" },
                    { y: 31, label: "User 2" },
                    { y: 52, label: "User 3" },
                    { y: 60, label: "User 4" },
                ]
            }]
        });
        chart.render();
        $(".canvasjs-chart-credit").hide();
    }
    
    var subArray = [];
    getTaskDetails = function(){
        $("#taskListing").html("");
        var job = $("#jobName option:selected").val() ;
        for(i = 0; i < vm.TaskDetailArray.length; i++) {
            if (vm.TaskDetailArray[i][0]==job) {
                subArray.push(vm.TaskDetailArray[i][1]);
            }
        }
        //select all in drop down
        for(i = 0; i < vm.TaskDetailArray.length; i++) {
            if (job =="All") {
                subArray.push(vm.TaskDetailArray[i][1]);
            }
        }
        var DynamicTaskListing ="";
        for (var i=0; i<subArray.length; i++){
            DynamicTaskListing+=' <p onclick="FunctionToChangeBarChart(event)">'+subArray[i]+'</p>';
        }
        $("#taskListing").prepend(DynamicTaskListing);
        subArray = [];
    }
    var selectAJob = $("#jobName option:selected").val() ;
    console.log("default job",selectAJob);
    for(i = 0; i < vm.TaskDetailArray.length; i++) {
        if (selectAJob =="SelectAJob") {
            subArray = [];
//            subArray.push(vm.TaskDetailArray[i][1]);
        }
    }
    var DynamicTaskListing ="";
    for (var i=0; i<subArray.length; i++){
        DynamicTaskListing+=' <p onclick="FunctionToChangeBarChart(event)">'+subArray[i]+'</p>';
    }
    $("#taskListing").prepend(DynamicTaskListing);
    
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
                
                
                if(jsonData[0] == "true"){
                    var persentageOfAcceptedUser;
                    var persentageOfRejectedUsers;
                    var PersentageOfStartedUser;
                    var PersentageOfCompletedUsers;
                    var persentageOfPendingUsers;
                    var TotalNoUsers = jsonData[7];
                    var today = new Date();
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
                    
                    console.log("todayDate",today);
                   
                    //for filtering details of task started users
                     
                     var AcceptedWOrk = jsonData[3];
                    var acceptedCount = 0;
                    if (AcceptedWOrk.length !=null){
                        for (i = 0;i<AcceptedWOrk.length;i++){
                            console.log("inner loop of ",AcceptedWOrk[i][3]);
                            var acceptedDate = AcceptedWOrk[i][1];
                            var acceptedDateFromDb = parseInt(acceptedDate)
                            var d = new Date(acceptedDateFromDb * 1000);
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
                               acceptedCount = acceptedCount+1;
                            }
                            
                            
                        }
                    }
                    
                    //for filtaring details of task accepted User
                    var startTaskArray = jsonData[1];
                    var startTaskCount = 0;
                    if (startTaskArray.length !=null){
                        for (i = 0;i<startTaskArray.length;i++){
                            console.log("inner loop of ",startTaskArray[i][3]);
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
                            
                            
                        }
                    }
                    
                    //for filtering of Completed task
                    
                    var completedTask = jsonData[2];
                    var completedTaskCount = 0;
                    if (completedTask.length !=null){
                        for (i = 0;i<completedTask.length;i++){
                            console.log("inner loop of ",completedTask[i][3]);
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
                        }
                    }
                    //for filtering of pending Task
                    var pendingTask = jsonData[4];
                    var pendingTaskCount = 0;
                    if (pendingTask.length !=null){
                        pendingTaskCount = pendingTask.length;
                    }
                    
                    //for fitering of rejected Users
                    var rejectedUsers = jsonData[5];
                    var rejectedTaskCount = 0;
                    if(rejectedUsers.length !=null){
                        rejectedTaskCount = rejectedUsers.length;
                    }
                    
                    persentageOfAcceptedUser = (acceptedCount/TotalNoUsers)*100;
                    persentageOfRejectedUsers = (rejectedTaskCount/TotalNoUsers)*100;
                    PersentageOfStartedUser = (startTaskCount/TotalNoUsers)*100;
                    PersentageOfCompletedUsers = (completedTaskCount/TotalNoUsers)*100;
                    persentageOfPendingUsers = (pendingTaskCount/TotalNoUsers)*100;
                    console.log("persentageOfAcceptedUser",persentageOfAcceptedUser);
                    console.log("persentageOfRejectedUsers",persentageOfRejectedUsers);
                    console.log("PersentageOfStartedUser",PersentageOfStartedUser);
                    console.log("persentageOfPendingUsers",persentageOfPendingUsers);
                    console.log("PersentageOfCompletedUsers",PersentageOfCompletedUsers);
                }
                else{
                    console.log("Server Problem");
                }
            },
            error: function (request,status, error) {
            }
        });
        
    }
    
    
});