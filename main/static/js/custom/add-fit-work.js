/*created by Aswathy Ashok*/
var companyTeamName = vm.CompanyTeamName;
var DynamicNotification ="";
    if (vm.NotificationNumber !=0){
        document.getElementById("number").textContent=vm.NotificationNumber;
    }else{
        document.getElementById("number").textContent="";
    }
$().ready(function() {
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
             $("#addFitToWorkForm").validate({
        rules: {
             fitWorkName: {
                required: true,
                remote:{
                    url: '/' + companyTeamName+"/isFitToWorkNameUsedEdit/" + fitWorkName+'/'+vm.FitToWorkName,
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
    }else{
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
        }          
   //notification
    var notificationSorted =[[]];
    function sortByCol(arr, colIndex){
    notificationSorted=arr.sort(sortFunction);
    function sortFunction(a, b) {
        a = a[colIndex]
        b = b[colIndex]
        return (a === b) ? 0 : (a < b) ? -1 : 1
    }
}

    
     myNotification= function () {
         if (vm.NotificationArray !=null){
        console.log("hiiii");
         sortByCol(vm.NotificationArray, 6);
         console.log("jjjjj",notificationSorted);
         var reverseSorted =[[]];
         reverseSorted=notificationSorted.reverse();

        document.getElementById("notificationDiv").innerHTML = "";
        var DynamicTaskListing="";
        if (reverseSorted !=null){
            DynamicTaskListing ="<h5>"+"Notifications"+ "<button class='no-button-style' method='post' onclick='clearNotification()'>"+"clear all"+"</button>"+"</h5>"+"<ul>";
        for(var i=0;i<reverseSorted.length;i++){
            console.log("sp1");
            var timeDifference =moment(new Date(new Date(reverseSorted[i][6]*1000)), "YYYYMMDD").fromNow();
            DynamicTaskListing += "<li>"+"User"+" "+reverseSorted[i][2]+" "+reverseSorted[i][3]+"  "+"delay to reach location"+" "+reverseSorted[i][4]+" "+"for task"+" "+reverseSorted[i][5]+" <span>"+timeDifference+"</span>"+"</li>";
            
            
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
             document.getElementById("notificationDiv").innerHTML = "";
            DynamicTaskListing ="<h5>"+" No New Notifications"+"</h5>";
                        $("#notificationDiv").prepend(DynamicTaskListing);
            
        }
        
        }else{
             document.getElementById("notificationDiv").innerHTML = "";
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
                        DynamicTaskListing ="<h5>"+" No New Notifications"+"</h5>";
                        $("#notificationDiv").prepend(DynamicTaskListing);
                    } else {
                    }
                },
                error: function (request,status, error) {
                    console.log(error);
                }
            }); 
         
         
         
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
    
    
    $("#cancel").click(function() {
            window.location = '/' + companyTeamName +'/fitToWork';
    });

});

 