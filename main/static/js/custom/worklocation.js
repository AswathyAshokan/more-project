
var companyTeamName = vm.CompanyTeamName;
$(document).ready(function() {
    var selectedUserArray = []; // contains all selected users and groups
    var selectedGroupArray = [];
    var groupKeyArray = [];
    $("#usersAndGroupId").on('change', function(evt, params) {
        console.log("inside group1");
        var tempArray = $(this).val();
        var clickedOption = "";
        console.log("array length",tempArray.length)
        if (selectedUserArray.length < tempArray.length) { // for selection
            for (var i = 0; i < tempArray.length; i++) {
                if (selectedUserArray.indexOf(tempArray[i]) == -1) {
                    console.log("clicked");
                    clickedOption = tempArray[i];
                }
            }
            for (var i = 0; i < vm.GroupMembers.length; i++) {
                if (vm.GroupMembers[i][0] == clickedOption) {
                    var memberLength = vm.GroupMembers[i].length;
                    groupKeyArray.push(clickedOption)
                    tempArray =[];
                    for (var j = 1; j < memberLength; j++) {
                        if (tempArray.indexOf(vm.GroupMembers[i][j]) == -1) {
                            tempArray.push(vm.GroupMembers[i][j])
                        }
                        console.log("values of temp array",tempArray);
                        $("#usersAndGroupId").val(tempArray);
                    }
                    selectedGroupArray.push(clickedOption);
                }
            }
            selectedUserArray = tempArray;
        } else if (selectedUserArray.length > tempArray.length) { // for deselection
            for (var i = 0; i < selectedUserArray.length; i++) {
                if (tempArray.indexOf(selectedUserArray[i]) == -1) {
                    clickedOption = selectedUserArray[i];
                    
                }
            }
        }
        console.log("group array",groupKeyArray);
        console.log("user array",selectedUserArray);
    });
    $("#usersAndGroupId").on('change', function(evt, params) {
        console.log("inside group1");
        var tempArray = $(this).val();
        var clickedOption = "";
        console.log("array length",tempArray.length)
        if (selectedUserArray.length < tempArray.length) { // for selection
            for (var i = 0; i < tempArray.length; i++) {
                if (selectedUserArray.indexOf(tempArray[i]) == -1) {
                    console.log("clicked");
                    clickedOption = tempArray[i];
                }
            }
            for (var i = 0; i < vm.GroupMembers.length; i++) {
                if (vm.GroupMembers[i][0] == clickedOption) {
                    var memberLength = vm.GroupMembers[i].length;
                    groupKeyArray.push(clickedOption)
                    tempArray =[];
                    for (var j = 1; j < memberLength; j++) {
                        if (tempArray.indexOf(vm.GroupMembers[i][j]) == -1) {
                            tempArray.push(vm.GroupMembers[i][j])
                        }
                        console.log("values of temp array",tempArray);
                        $("#userOrGroup").val(tempArray);
                    }
                    selectedGroupArray.push(clickedOption);
                }
            }
            selectedUserArray = tempArray;
        } else if (selectedUserArray.length > tempArray.length) { // for deselection
            for (var i = 0; i < selectedUserArray.length; i++) {
                if (tempArray.indexOf(selectedUserArray[i]) == -1) {
                    clickedOption = selectedUserArray[i];
                    
                }
            }
            selectedUserArray = tempArray;
        }
        console.log("group array",groupKeyArray);
        console.log("user array",selectedUserArray);
    });
    $("#workLocationForm").validate({
        rules: {
            usersAndGroupId:"required",
            taskLocation : "required",
        },
        messages: {
            usersAndGroupId: "Please select user and group",
            taskLocation:"please fill this column",
        },
        submitHandler: function(){//to pass all data of a form serial
             $("#saveButton").attr('disabled', true);
            var formData = $("#workLocationForm").serialize();
            //get the user's name corresponding to  keys selected from dropdownlist 
            var selectedUserAndGroupName = [];
              $("#usersAndGroupId option:selected").each(function () {
                  var $this = $(this);
                  if ($this.length) {
                      var selectedUserName = $this.text();
                      selectedUserAndGroupName.push( selectedUserName);
                  }
              });
              for(i = 0; i < selectedUserAndGroupName.length; i++) {
                  formData = formData+"&userAndGroupName="+selectedUserAndGroupName[i];
              }
            for(i = 0; i < groupKeyArray.length; i++) {
                formData = formData+"&groupArrayElement="+groupKeyArray[i];
            }
           // formData = formData+"&selectedUserNames="+selectedUserArray
            for(i = 0; i < selectedUserArray.length; i++) {
                formData = formData+"&selectedUserNames="+selectedUserArray[i];
            }
              console.log("formData",formData);
            $.ajax({
                
                    url:'/' + companyTeamName +'/worklocation',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                           console.log("llllll");
                        }else {
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
                return false;
        }
    });
    
    
    
    console.log("haiii",vm)
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
} );
