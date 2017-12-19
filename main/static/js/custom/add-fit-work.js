/*created by Aswathy Ashok*/
var companyTeamName = vm.CompanyTeamName;

$().ready(function() {
    if(vm.PageType == "edit"){ 
        console.log("instructions",vm.InstructionArrayToEdit[0]);
        console.log("fit name",vm.FitToWorkName);
        document.getElementById("fitWorkName").value = vm.FitToWorkName;
        var dynamicTextBox= "";
        for (var i = 1; i < vm.InstructionArrayToEdit.length; i++) {
            dynamicTextBox+= '<div class="plus"><input class="form-control"  name = "DynamicTextBox"  id=  "DynamicTextBox"  type="text" value = "' + vm.InstructionArrayToEdit[i] + '" />&nbsp;' + "<span class='add-decl'>+</span>" + "</div>" ;
        }
        document.getElementById("addFitToWorkValue").value = vm.InstructionArrayToEdit[0];

        $( ".wrp-plus" ).prepend( dynamicTextBox);
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
   
    
    
    function  addleveldata(){
        var repeat =  "<div class='plus'>" + "<input class='form-control' name='DynamicTextBox' id='DynamicTextBox' type='text'>" + "<span class='add-decl'>+</span>" + "</div>" ;
        $( ".wrp-plus" ).prepend( repeat );
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

 