console.log(vm);
//if (vm.NotificationNumber !=0){
//        document.getElementById("number").textContent=vm.NotificationArray.length;
//
//    }else{
//        document.getElementById("number").textContent="";
//    }
$(document).ready(function() {
    
    //checking plans
    
    if(vm.CompanyPlan == 'family' ){
        var parent = document.getElementById("menuItems");
        var contact = document.getElementById("contact");
        var job = document.getElementById("job");
        var crm = document.getElementById("crm");
        var leave = document.getElementById("leave");
        var timesheet  = document.getElementById("time-sheet");
        var consent = document.getElementById("consent");
        var workLocation = document.getElementById("workLocation");
        parent.removeChild(workLocation);
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
            var consent = document.getElementById("consent");
            var workLocation = document.getElementById("workLocation");
            parent.removeChild(workLocation);
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
    
//    
//      myNotification= function () {
//        console.log("hiiii");
//        var DynamicTaskListing="";
//        DynamicTaskListing ="<h5>"+"Notifications"+"</h5>"+"<ul>";
//        for(var i=0;i<vm.NotificationArray.length;i++){
//            console.log("sp1");
//            var timeDifference =moment(new Date(new Date(vm.NotificationArray[i][6]*1000)), "YYYYMMDD").fromNow();
//            DynamicTaskListing += "<li>"+"User"+" "+vm.NotificationArray[i][2]+" "+vm.NotificationArray[i][3]+"  "+"delay to reach location"+" "+vm.NotificationArray[i][4]+" "+"for task"+" "+vm.NotificationArray[i][5]+" <span>"+timeDifference+"</span>"+"</li>";
//            
//            
//        }
//            $("#notificationDiv").prepend(DynamicTaskListing+"</ul>");
//            document.getElementById("number").textContent="";
//            $.ajax({
//                url:'/'+ companyTeamName + '/notification/update',
//                type: 'post',
//                success : function(response) {
//                    if (response == "true" ) {
//                    } else {
//                    }
//                },
//                error: function (request,status, error) {
//                    console.log(error);
//                }
//            });
//        }
} );
