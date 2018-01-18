/*_____________________________________________________________________________________________*/


console.log("vmmm",vm);
var DynamicNotification ="";
    
document.getElementById("consent").className += "active";
var companyTeamName = vm.CompanyTeamName;
var selectedUserArray = [];
console.log("vm.UsersKey",vm.UsersKey);
if (vm.UsersKey !=null){
       selectedUserArray = vm.UsersKey;
}
console.log("selected array id",selectedUserArray);


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
    if(vm.PageType == "edit"){ 
        /*console.log("instructions",vm.InstructionArrayToEdit[0]);
        var selectArray = vm.UsersKey;
        document.getElementById("recieptName").value = vm.ReceiptName;
        document.getElementById("addConsentValue").value = vm.InstructionArrayToEdit[0];
        document.getElementById("consentHead").innerHTML = "Edit Consent Receipt";//for display heading of each webpage
        $("#selectedUserIds").val(selectArray);
        for (var i = 1; i < vm.InstructionArrayToEdit.length; i++) {
                console.log("cp1");
                var dynamicTextBox = "<div class='plus'>"+"<input class='form-control'  name = 'DynamicTextBox'  id=  'DynamicTextBox'  type='text' value = " + vm.InstructionArrayToEdit[i] + ">" + "<span class='add-decl'>+</span>" + "</div>";
             $( ".wrp-plus" ).append( dynamicTextBox );
            }*/
        
        
        console.log("vm.SelectedUsersKey",vm.SelectedUsersKey);
        var selectArray = vm.SelectedUsersKey;
        document.getElementById("recieptName").value = vm.ReceiptName;
       
        document.getElementById("consentHead").innerHTML = "Edit Consent Receipt";//for display heading of each webpage
        $("#selectedUserIds").val(selectArray);
       var dynamicTextBox= "";
        for (var i = 1; i <vm.InstructionArrayToEdit.length; i++) {
            dynamicTextBox+= '<div class="plus"><input class="form-control"  name = "DynamicTextBox"  id=  "DynamicTextBox"  type="text" value = "' + vm.InstructionArrayToEdit[i] + '"/>' + "<span class='add-decl'>+</span>" + "</div>" ;
        }
        $( ".wrp-plus" ).prepend(dynamicTextBox);
         document.getElementById("addConsentValue").value = vm.InstructionArrayToEdit[0];
        
        
        }
    
    
    var selectedGroupArray = [];
    var groupKeyArray = [];
    $("#selectedUserIds").on('change', function(evt, params) {
        var tempArray = $(this).val();
        var clickedOption = "";
        if (selectedUserArray.length < tempArray.length) { // for selection
            for (var i = 0; i < tempArray.length; i++) {
                if (selectedUserArray.indexOf(tempArray[i]) == -1) {
                    console.log("clicked");
                    clickedOption = tempArray[i];
                }
            }
             console.log("tempArray",tempArray);
            console.log("clickedOption",clickedOption);
             console.log("vm.GroupMembers",vm.GroupMembers);
            if (vm.GroupMembers !=null){
                for (var i = 0; i < vm.GroupMembers.length; i++) {
                    if (vm.GroupMembers[i][0] == clickedOption) {
                        var memberLength = vm.GroupMembers[i].length;
                        groupKeyArray.push(clickedOption);
                        tempArray =[];
                        for (var j = 1; j < memberLength; j++) {
                            if (tempArray.indexOf(vm.GroupMembers[i][j]) == -1) {
                                tempArray.push(vm.GroupMembers[i][j]);
                            }
                            $("#selectedUserIds").val(tempArray);
                        }
                        selectedGroupArray.push(clickedOption);
                    }
                }
                console.log("hai iam waiting for ur coming",groupKeyArray);
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
    });
        
        
   
       /* $( "#TextBoxContainer" ).append( dynamicTextBox );
    }*/
    $("#addConsentForm").validate({
        rules: {
            recieptName:"required",
            selectedUserIds:"required"
        },
        messages: {
            recieptName:"Please enter consent Reciept Name",
            selectedUserIds: "Please select Users"
        },
        submitHandler: function(){//to pass all data of a form serial
            console.log("hai iam in submit handler");
             $("#saveButton").attr('disabled', true);
            var formData = $("#addConsentForm").serialize();
            var selectedUsersNames = [];
            var instructionFromDynamicTextBox = [];
          //  var replaceCharacter = []
            //get the user's name corresponding to  keys selected from dropdownlist 
            var selectedUserAndGroupName = [];
             $("#selectedUserIds option:selected").each(function () {
                 var $this = $(this);
                 if ($this.length) {
                     var selectedUserName = $this.text();
                     selectedUserAndGroupName.push( selectedUserName);
                 }
             });
            for(i = 0; i < groupKeyArray.length; i++) {
                formData = formData+"&groupArrayElement="+groupKeyArray[i];
            }
           
           for(i = 0; i < selectedUserAndGroupName.length; i++) {
               formData = formData+"&userAndGroupName="+selectedUserAndGroupName[i];
           }
            
            for(i = 0; i < selectedUserArray.length; i++) {
               formData = formData+"&selectedUserNames="+selectedUserArray[i];
           }
            
            
             var ConsentValue = document.getElementById("addConsentValue").value;
            if(ConsentValue.length !=0){
                var replaceString = ConsentValue.split(';').join(':');
                instructionFromDynamicTextBox.push(replaceString+"/@@");
                //instructionFromDynamicTextBox.push("&&");
            }
            $("input[name=DynamicTextBox]").each(function () {
                 if($(this).val().length !=0){
                     var valueOfTextBox = $(this).val();
                     var replaceString =  valueOfTextBox.split(';').join(':');
                     instructionFromDynamicTextBox.push(replaceString+"/@@");
                     // instructionFromDynamicTextBox.push("&&");
                     
                 }
            });
            console.log("instructionFromDynamicTextBox in js",instructionFromDynamicTextBox)
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
        console.log("hai iam in cancel section");
        window.location = '/' + companyTeamName +'/consent';
    });
});





























































