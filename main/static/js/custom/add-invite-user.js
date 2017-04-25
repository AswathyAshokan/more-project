/*Created By Farsana*/

//Below line is for adding active class to layout side menu..
console.log(vm.companyTeamName);
document.getElementById("user").className += " active";

var companyTeamName = vm.CompanyTeamName;

$().ready(function() {
    
    if(vm.PageType == "edit"){  
        
        document.getElementById("emailid").style.display='block'; 
        document.getElementById("firstname").value = vm.FirstName;
        document.getElementById("lastname").value = vm.LastName;
        document.getElementById("emailid").value = vm.EmailId;
        document.getElementById("usertype").value = vm.UserType;
        document.getElementById("pageTitle").innerHTML = "Edit Invited User"
        
    }
    if(vm.AllowInvitations == false){
        $("#InviteUserValidationError").css({"color": "red", "font-size": "15px"});
        $("#InviteUserValidationError").html("Your user invitation limit is exceeded.Please upgrade your plan").show();
        $("#saveButton").attr('disabled', true);
    }
    
    $("#adduserForm").validate({
	  rules: {
          firstname: "required",
          usertype :"required",
          emailid:{
              required:true,
              email:true,
               
          },
      },
        messages: {
            firstname:"please enter First name ",
            usertype :"Plese select UserType",
            emailid:"email Address is already inuse!"
        },
        submitHandler: function(){//to pass all data of a form serial
            $("#saveButton").attr('disabled', true);
            if (vm.PageType == "edit"){
                var formData = $("#adduserForm").serialize();
                var InviteId = vm.InviteId;
                $.ajax({
                    url:'/' + companyTeamName +'/invite/'+ InviteId +'/edit',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/' + companyTeamName +'/invite';
                        }else {
                            $("#emailValidationError").css({"color": "red", "font-size": "15px"});
                            $("#emailValidationError").html("please select location from map.").show();
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
            } else {
                var formData = $("#adduserForm").serialize();
                $.ajax({
                    url:'/' + companyTeamName +'/invite/add',
                    type:'post',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/' + companyTeamName +'/invite';
                        }else {
                              $("#saveButton").attr('disabled', false);
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
            }
            return false;
        }
    });
    
    
    /*$('#updateButton').on('click', function() {
        
         $("#adduserForm").validate({
             rules: {
                 firstname: "required",
                 usertype :"required",
                 emailid:{
                     required:true,
                     email:true,
                 },
             },
             messages: {
                 firstname:"please enter First name ",
                 usertype :"Plese select UserType",
                 emailid:"please enter a valid Email address!"
             },
             submitHandler: function(){
             $("#saveButton").attr('disabled', false);
             $("#updateButton").attr('disabled', true);
             });
         
         });
    
    });*/
    
    $("#cancel").click(function() {
            window.location = '/' + companyTeamName +'/invite';
    });
});






