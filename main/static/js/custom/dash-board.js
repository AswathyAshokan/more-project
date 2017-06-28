console.log(vm);
$(function () {
    
     if(vm.CompanyPlan == 'family' ){
        var parent = document.getElementById("menuItems");
        var contact = document.getElementById("contact");
        var job = document.getElementById("job");
        var crm = document.getElementById("crm");
        var leave = document.getElementById("leave");
        var time  = document.getElementById("time-sheet");
        var consent = document.getElementById("consent")
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
            DynamicTaskListing+=' <p>'+subArray[i]+'</p>';
        }
        $("#taskListing").prepend(DynamicTaskListing);
        subArray = [];
    }
    var selectAJob = $("#jobName option:selected").val() ;
    console.log("default job",selectAJob);
    for(i = 0; i < vm.TaskDetailArray.length; i++) {
        if (selectAJob =="SelectAJob") {
            subArray.push(vm.TaskDetailArray[i][1]);
        }
    }
    var DynamicTaskListing ="";
    for (var i=0; i<subArray.length; i++){
        DynamicTaskListing+=' <p>'+subArray[i]+'</p>';
    }
    $("#taskListing").prepend(DynamicTaskListing);
    
    
    
    
});