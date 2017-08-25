console.log("gsgsgsgs",vm.UserNameToEdit);
var companyTeamName = vm.CompanyTeamName;
 var selectedUserArray = [];
$(document).ready(function() {
    // contains all selected users and groups
    if(vm.PageType == "edit"){ 
        var selectArray =[];
        selectArray = vm.UsersKey;
        $("#usersAndGroupId").val(selectArray);
        console.log("array",selectArray);
        document.getElementById("taskLocation").value = vm.WorkLocation;
        document.getElementById("workLocationId").innerHTML = "Edit WorkLocation";//for display heading of each webpage
        var selectedGroupArray = [];
    var groupKeyArray = [];
    $("#usersAndGroupId").on('change', function(evt, params) {
        console.log("inside group1");
        var tempArray = $(this).val();
        
        var clickedOption = "";
        console.log("array length",tempArray);
        if (selectedUserArray.length < tempArray.length) { // for selection
            for (var i = 0; i < tempArray.length; i++) {
                if (selectedUserArray.indexOf(tempArray[i]) == -1) {
                    console.log("clicked");
                    clickedOption = tempArray[i];
                }
            }
            if (vm.GroupMembers !=null){
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
       
    }
    
    
    var selectedGroupArray = [];
    var groupKeyArray = [];
    $("#usersAndGroupId").on('change', function(evt, params) {
        console.log("inside group1");
        var tempArray = $(this).val();
        
        var clickedOption = "";
        console.log("array length",tempArray);
        if (selectedUserArray.length < tempArray.length) { // for selection
            for (var i = 0; i < tempArray.length; i++) {
                if (selectedUserArray.indexOf(tempArray[i]) == -1) {
                    console.log("clicked");
                    clickedOption = tempArray[i];
                }
            }
            if (vm.GroupMembers !=null){
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
            if (vm.PageType == "edit"){
                var workLocationId =vm.WorkLogId  
                $.ajax({
                    url:'/' + companyTeamName +'/worklocation/'+ workLocationId + '/edit',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/' + companyTeamName +'/worklocation';
                        }else {
                             $("#saveButton").attr('disabled', false);
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
            
            } else {
                $.ajax({

                        url:'/' + companyTeamName +'/worklocation/add',
                        type:'post',
                        datatype: 'json',
                        data: formData,
                        //call back or get response here
                        success : function(response){
                            if(response == "true"){
                               window.location = '/'+companyTeamName+'/worklocation';
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
            window.location = '/'+companyTeamName+'/worklocation';
    });
    
    
    
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