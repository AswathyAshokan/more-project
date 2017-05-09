/*Created By Farsana*/

//Below line is for adding active class to layout side menu..
document.getElementById("group").className += " active";

console.log(vm);

var companyTeamName = vm.CompanyTeamName;
$().ready(function() {
    if(vm.PageType == "edit"){ 
        var selectArray = vm.GroupMembersToEdit;
        console.log("array",selectArray);
        document.getElementById("groupName").value = vm.GroupNameToEdit;
        document.getElementById("groupHead").innerHTML = "Edit Group";//for display heading of each webpage
        $("#selectedUserIds").val(selectArray);
    }

    $("#addgroupForm").validate({
        rules: {
            groupName:{
                required:"required",
                remote:{
                    url: "/isgroupnameused/" + groupName + "/" + vm.PageType + "/" + vm.GroupNameToEdit ,
                    type: "post"
                }
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




























