/*created by Aswathy Ashok*/
var companyTeamName = vm.CompanyTeamName;
$().ready(function() {
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
            fitWorkName:"required",
            addFitToWorkValue:"required"
        },
        messages: {
            fitWorkName:"Please enter fit to work Name",
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

 