/*Created By Farsana*/

//Below line is for adding active class to layout side menu..
console.log(vm.FirstName);
document.getElementById("user").className += " active";

var companyTeamName = vm.CompanyTeamName;
var selectedCompanyPlan = vm.CompanyPlan;

$().ready(function() {
    
    $( "#noOfUsers" ).keyup(function() {
        var noOfUserFromUser = document.getElementById("noOfUsers").value;
        var total = 5*noOfUserFromUser;
        document.getElementById("result").innerHTML="$"+total;
    });
    
    if(vm.PageType == "edit"){
        document.getElementById("firstname").value = vm.FirstName;
        document.getElementById("lastname").value = vm.LastName;
        document.getElementById("usertype").value = vm.UserType;
        document.getElementById("emailid").value = vm.EmailId;
        document.getElementById("pageTitle").innerHTML = "Edit Invited User"
        if(vm.UserResponse == "Rejected"){
            $("#saveButton").attr('disabled', false);
            $("#updateButton").attr('disabled', true);
        } 
        else{
            $("#updateButton").attr('disabled', false);
            $("#saveButton").attr('disabled', true);
            $("#emailid").attr("disabled", "disabled");
        }
        
    }
    else{
         $("#saveButton").attr('disabled', false);
         $("#updateButton").attr('disabled', true);
        
    }
    if(vm.AllowInvitations == false){
        
        $("#limitModel").modal();
        $("#InviteUserValidationError").css({"color": "red", "font-size": "15px"});
        $("#InviteUserValidationError").html("Your user invitation limit is exceeded.Please upgrade your plan").show();
        $("#saveButton").attr('disabled', true);
        $("#limitModalCancel").click(function(){
            window.location = '/' + companyTeamName +'/invite';
        })
         $("#payNowBtn").click(function(){
              //$('#closemodal').modal('hide');
             $('#limitModel').modal('hide');
             var numberOfUsers  = document.getElementById("noOfUsers").value
             window.location = '/'+selectedCompanyPlan+'/payment'
             //window.location ='/' + companyTeamName +'/invite/'+numberOfUsers+'/AddExtraUserByUpgradePlan';
         });
    }
    
    
    /*var date = new Date();
    var datum = (Date.parse(date))/1000;
    */
    $('#saveButton').on('click', function() {
        $("#adduserForm").validate({
          rules: {
              firstname: "required",
              lastname:"required",
              usertype :"required",
              emailid:{
                  required:true,
                  email:true,
              },
          },
            messages: {
                firstname:"Enter first name ",
                lastname:"Enter last name",
                usertype :"Select UserType",
                emailid:"Enter a valid Email address"
            },
            submitHandler: function(){//to pass all data of a form serial
                $("#saveButton").attr('disabled', true);
                $("#updateButton").attr('disabled', true);
                var formData = $("#adduserForm").serialize();
                    $.ajax({
                        url:'/' + companyTeamName +'/invite/add',
                        type:'post',
                        data: formData,
                        //call back or get response here
                        success : function(data){
                            console.log("error",data);
//                            var jsonData = JSON.parse(data)
                            if(data == "true"){
                                
                                window.location='/' + companyTeamName +'/invite';
                            }else if(data = "false"){
                                $("#emailValidationError").css({"color": "red", "font-size": "15px"});
                                  $("#emailValidationError").html("email already in use.").show();
                                  $("#saveButton").attr('disabled', false);
                            } else{
                                console.log("Server Problem");
                            }
                        },
                        error: function (request,status, error) {
                        }
                    });
                return false;
            }
        });
     });
    
    $('#updateButton').on('click', function() {
        $("#adduserForm").validate({
             rules: {
                 firstname: "required",
                 lastname:"required",
                 usertype :"required",
                 emailid:{
                     required:true,
                     email:true,
                 },
             },
             messages: {
                 firstname:"Enter first name ",
                 lastname:"Enter last name",
                 usertype :"Select UserType",
                 emailid:"Enter a valid Email address"
                 
             },
             submitHandler: function(){
                 $("#updateButton").attr('disabled', true);
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
                            $("#updateButton").attr('disabled', false);
                        }
                    },
                    error: function (request,status, error) {
                    }
                 });
             }
         
         });
    
    });
    
    $("#cancel").click(function() {
            window.location = '/' + companyTeamName +'/invite';
    });
});






