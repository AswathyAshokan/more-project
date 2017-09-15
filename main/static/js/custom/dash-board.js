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
                    console.log("jsonData[1]",jsonData[1])
                    console.log("jsonData[2]",jsonData[2])
                    console.log("jsonData[3]",jsonData[3])
                    var taskStartDate = jsonData[3][0];
                    var taskEndDate =  jsonData[3][1];
                    var startdateFromDb = parseInt(taskStartDate)
                    var d = new Date(startdateFromDb * 1000);
                    var dd = d.getDate();
                    var mm = d.getMonth() + 1; //January is 0!
                    var yyyy = d.getFullYear();
                    if (dd < 10) {
                        dd = '0' + dd;
                    }
                    if (mm < 10) {
                        mm = '0' + mm;
                    }
                    var starDate = ( dd +'/'+mm + '/' + yyyy);
                    console.log("starDate",starDate);
                    
                    var enddateFromDb = parseInt(taskEndDate);
                    var endDay = new Date(enddateFromDb * 1000);
                    var enddd = endDay.getDate();
                    var endmm = endDay.getMonth() + 1; //January is 0!
                    var endyyyy =endDay.getFullYear();
                    if (enddd < 10) {
                        enddd = '0' + enddd;
                    }
                    if (endmm < 10) {
                        endmm = '0' + endmm;
                    }
                    var endDate = (enddd+'/'+endmm +'/'+endyyyy);
                    console.log("endDate",endDate);
                    
                    
                    var mdy = starDate.split('/');
                    return new Date(mdy[2], mdy[0]-1, mdy[1]);
                    
                    
                    
                    
                    
                    
                    
                    
                    var usersArray = [[]];
                    var LogArray = [[]];
                    usersArray =  jsonData[1];
                    LogArray = jsonData[2];
                    var logTimeArray = []; 
                    var logTimeArrayForAllUser = [[]];
                    console.log("usersArray",usersArray);
                    console.log("LogArray",LogArray);
                    for(i=0;i<usersArray.length;i++){
                        console.log("user array for dash ", usersArray[i][2]);
                            for ( k=0;k<LogArray .length;k++){
                                console.log("LogArray[k][m]",LogArray[k][0]);
                                    if (usersArray[i][2] == LogArray[k][0] ){
                                        logTimeArray.push(LogArray[k][1]);
                                        logTimeArray.push(LogArray[k][0]);
                                        logTimeArrayForAllUser.push(logTimeArray);
                                        logTimeArray = [];
                                    }
                            }
                    }
                    console.log("kkkkkkk",logTimeArrayForAllUser);
                    
                    //window.location='/' + companyTeamName +'/invite';
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