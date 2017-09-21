/*Created By Farsana*/

//Below line is for adding active class to layout side menu..
console.log("nnnn",vm.GroupMembersToEdit)
document.getElementById("group").className += " active";
var companyTeamName = vm.CompanyTeamName;
var DynamicNotification ="";
    if (vm.NotificationNumber !=0){
        document.getElementById("number").textContent=vm.NotificationArray.length;
    }else{
        document.getElementById("number").textContent="";
    }

$().ready(function() {
    
    
    
    myNotification= function () {
        console.log("hiiii");
        document.getElementById("notificationDiv").innerHTML = "";
        var DynamicTaskListing="";
        if (vm.NotificationArray !=null){
            DynamicTaskListing ="<h5>"+"Notifications"+"</h5>"+"<ul>";
        for(var i=0;i<vm.NotificationArray.length;i++){
            console.log("sp1");
            var timeDifference =moment(new Date(new Date(vm.NotificationArray[i][6]*1000)), "YYYYMMDD").fromNow();
            DynamicTaskListing += "<li>"+"User"+" "+vm.NotificationArray[i][2]+" "+vm.NotificationArray[i][3]+"  "+"delay to reach location"+" "+vm.NotificationArray[i][4]+" "+"for task"+" "+vm.NotificationArray[i][5]+" <span>"+timeDifference+"</span>"+"</li>";
            
            
        }
            $("#notificationDiv").prepend(DynamicTaskListing+"</ul>");
            document.getElementById("number").textContent="";
            $.ajax({
                url:'/'+ companyTeamName + '/notification/update',
                type: 'post',
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
            DynamicTaskListing ="<h5>"+" No New Notifications"+"</h5>";
            $("#notificationDiv").prepend(DynamicTaskListing);
            
        }
        
        }
     
     
     
     clearNotification= function () {
          document.getElementById("notificationDiv").innerHTML = "";
          $.ajax({
                url:'/'+ companyTeamName + '/notification/delete',
                type: 'post',
                success : function(response) {
                    if (response == "true" ) {
                    } else {
                    }
                },
                error: function (request,status, error) {
                    console.log(error);
                }
            }); 
         
         
         
     }
    
    if(vm.PageType == "edit"){ 
        var selectArray =[];
        for(var u = 0 ; u <= vm.GroupMembersToEdit.length;u++){
            selectArray.push(vm.GroupMembersToEdit[u]);
        }
        
        $("#selectedUserIds").val(selectArray);
        console.log("array",selectArray);
        document.getElementById("groupName").value = vm.GroupNameToEdit;
        document.getElementById("groupHead").innerHTML = "Edit Group";//for display heading of each webpage
       
    }

    $("#addgroupForm").validate({
        rules: {
            groupName:{
                required:true,
                remote:{
                    url: "/isgroupnameused/" + groupName + "/" + vm.PageType + "/" + vm.GroupNameToEdit ,
                    type: "post"
                },
            },
            selectedUserIds : "required",
        },
        messages: {
            groupName:{
                required: "Please enter Group Name",
                remote: "The Group Name is already in use!"
                },
            selectedUserIds:"please fill this column",
        },
        submitHandler: function(){//to pass all data of a form serial
             $("#saveButton").attr('disabled', true);
            var formData = $("#addgroupForm").serialize();
            var selectedUsersNames = [];
            
//get the user's name corresponding to  keys selected from dropdownlist 
            $("#selectedUserIds option:selected").each(function () {
                var $this = $(this);
                if ($this.length) {
                    var selectedUsersName = $this.text();
                    selectedUsersNames.push(selectedUsersName);
                }
            });
            
// Serialialize all the selected invite user name from dropdown list with form data
            for(i = 0; i < selectedUsersNames.length; i++) {
                formData = formData+"&selectedUserNames="+selectedUsersNames[i];
            }
            var groupId = vm.GroupId;
            if (vm.PageType == "edit"){
                
                $.ajax({
                    url:'/' + companyTeamName +'/group/'+ groupId  +'/edit',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/' + companyTeamName +'/group';
                        }else {
                             $("#saveButton").attr('disabled', false);
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
            
            } else {
                $.ajax({
                
                    url:'/' + companyTeamName +'/group/add',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/' + companyTeamName +'/group';
                        }else {
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
                return false;
            }
        }
    });
    
    $("#cancel").click(function() {
            window.location = '/' + companyTeamName +'/group';
    });
});




























