/*Created By Farsana*/

//Below line is for adding active class to layout side menu..
document.getElementById("group").className += " active";

$().ready(function() {
    if(vm.PageType == "edit"){ 
        var selectArray = vm.GroupMembersToEdit;
        document.getElementById("groupName").value = vm.GroupNameToEdit;
        document.getElementById("groupHead").innerHTML = "Edit Group";//for display heading of each webpage
        $("#selectedUserIds").val(selectArray);
    }
    
    $("#addgroupForm").validate({
        rules: {
            groupName: "required",
            selectedUserIds : "required"
        },
        messages: {
            groupName:"please enter group name ",
            selectedUserIds:"please fill this column"
        },
        submitHandler: function(){//to pass all data of a form serial
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
            if (vm.PageType == "edit"){
                var groupId = vm.GroupId;
                $.ajax({
                    url:'/group/'+ groupId  +'/edit',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/group';
                        }else {
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
            
            } else {
                $.ajax({
                    url:'/group/add',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/group';
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
});




























