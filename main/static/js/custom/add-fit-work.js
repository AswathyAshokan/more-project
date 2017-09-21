/*created by Aswathy Ashok*/
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
        console.log("instructions",vm.InstructionArrayToEdit[0]);
        console.log("fit name",vm.FitToWorkName);
        document.getElementById("fitWorkName").value = vm.FitToWorkName;
        document.getElementById("addFitToWorkValue").value = vm.InstructionArrayToEdit[0];
        var dynamicTextBox= "";
        for (var i = 1; i < vm.InstructionArrayToEdit.length; i++) {
            dynamicTextBox+= '<div class="plus"><input class="form-control"  name = "DynamicTextBox"  id=  "DynamicTextBox"  type="text" value = "' + vm.InstructionArrayToEdit[i] + '" />&nbsp;' + "<span class='add-decl'>+</span>" + "</div>" ;
        }
        $( ".wrp-plus" ).append( dynamicTextBox);
    }
    function  addleveldata(){
        var repeat =  "<div class='plus'>" + "<input class='form-control' name='DynamicTextBox' id='DynamicTextBox' type='text'>" + "<span class='add-decl'>+</span>" + "</div>" ;
        $( ".wrp-plus" ).append( repeat );
    }
    $(document).on('click', '.add-decl', function () {
        if ($(this).closest('.plus').is(':last-child')) {
            addleveldata();
        }
        else {
            $(this).closest('.plus').remove();
        }
    });
     $("#addFitToWorkForm").validate({
        rules: {
             fitWorkName: {
                required: true,
                remote:{
                    url: '/' + companyTeamName+"/isFitToWorkNameUsed/" + fitWorkName,
                    type: "post"
                }
            },
            addFitToWorkValue:"required"
        },
        messages: {
            fitWorkName: {
                required: "Please enter fit to work Name",
                remote: "fit to work already exists!"
            },
            addFitToWorkValue: "Please add instruction"
        },
        submitHandler: function(){//to pass all data of a form serial
             $("#saveButton").attr('disabled', true);
            var formData = $("#addFitToWorkForm").serialize();
            var instructionFromDynamicTextBox = [];
            
            
            var FitToWorkValue = document.getElementById("addFitToWorkValue").value;
            if(FitToWorkValue.length !=0){
                instructionFromDynamicTextBox.push(FitToWorkValue+"/@@");
                //instructionFromDynamicTextBox.push("&&");
            }
            $("input[name=DynamicTextBox]").each(function () {
                 if($(this).val().length !=0){
                     instructionFromDynamicTextBox.push($(this).val()+"/@@")
                     // instructionFromDynamicTextBox.push("&&");
                     
                 }
            });
            console.log("instructionFromDynamicTextBox",instructionFromDynamicTextBox)
            formData = formData+"&instructionsForUser="+instructionFromDynamicTextBox;
            var fitToWorkId = vm.FitToWorkId;
            if (vm.PageType == "edit"){
                $.ajax({
                    url:'/' + companyTeamName +'/fitToWork/'+ fitToWorkId  +'/edit',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/' + companyTeamName +'/fitToWork';
                        }else {
                             $("#saveButton").attr('disabled', false);
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
            
            } else {
                $.ajax({
                
                    url:'/' + companyTeamName +'/fitToWork/add',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/' + companyTeamName +'/fitToWork';
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
            window.location = '/' + companyTeamName +'/fitToWork';
    });

});

 