document.getElementById("consent").className += "active";

console.log(vm);


var companyTeamName = vm.CompanyTeamName;
$().ready(function() {
    if(vm.PageType == "edit"){ 
        console.log("instructions",vm.InstructionArrayToEdit[0]);
        var selectArray = vm.UserNameToEdit;
        $("#selectedUserIds").val(selectArray);
        document.getElementById("recieptName").value = vm.ReceiptName;
        document.getElementById("addConsentValue").value = vm.InstructionArrayToEdit[0];
        document.getElementById("consentHead").innerHTML = "Edit Consent Receipt";//for display heading of each webpage
        var dynamicTextBox= "";
        for (var i = 1; i < vm.InstructionArrayToEdit.length; i++) {
            console.log("cp1");
            dynamicTextBox+= '<div class="plus"><input class="form-control"  name = "DynamicTextBox"  id=  "DynamicTextBox"  type="text" value = "' + vm.InstructionArrayToEdit[i] + '" />&nbsp;'+'<img  id="exposureDelete" src="/static/images/exposureCancel.jpg" width="25" height="25" style= "float:right; margin-top:-1em; margin-right:-1em;"  class="delete-exposure" />';
        }
        $("#TextBoxContainer").prepend(dynamicTextBox);
    }
    $("#btnAdd").on("click", function () {
        var div = $("<div class='plus'/>");
        div.html(GetDynamicTextBox(""));
        $("#TextBoxContainer").prepend(div);
    });
    $("body").on("click", ".delete-decl", function () {
        $(this).closest("div").remove();
    });
    
    function GetDynamicTextBox(value) {
    return ' <input class="form-control"  name = "DynamicTextBox"  id=  "DynamicTextBox"  type="text" value = "" />&nbsp;' +'<button id="btnAdd"class="delete-decl">+</button>'
    }
    
    $("#addConsentForm").validate({
        rules: {
            recieptName:"required",
            selectedUserIds:"required"
        },
        messages: {
            recieptName:"Please enter Contact Person",
            selectedUserIds: "Please enter Phone Number"
        },
        submitHandler: function(){//to pass all data of a form serial
             $("#saveButton").attr('disabled', true);
            var formData = $("#addConsentForm").serialize();
            var selectedUsersNames = [];
            var instructionFromDynamicTextBox = [];
            
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
            
            var ConsentValue = document.getElementById("addConsentValue").value;
            if(ConsentValue.length !=0){
                instructionFromDynamicTextBox.push(ConsentValue);
            }
            $("input[name=DynamicTextBox]").each(function () {
                 if($(this).val().length !=0){
                     instructionFromDynamicTextBox.push($(this).val())
                 }
            });
            formData = formData+"&instructionsForUser="+instructionFromDynamicTextBox;
            var ConsentId = vm.ConsentId;
            if (vm.PageType == "edit"){
                $.ajax({
                    url:'/' + companyTeamName +'/consent/'+ ConsentId  +'/edit',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/' + companyTeamName +'/consent';
                        }else {
                             $("#saveButton").attr('disabled', false);
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
            
            } else {
                $.ajax({
                
                    url:'/' + companyTeamName +'/consent/add',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/' + companyTeamName +'/consent';
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
            window.location = '/' + companyTeamName +'/consent';
    });
});
































