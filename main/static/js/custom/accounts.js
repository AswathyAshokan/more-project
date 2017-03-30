document.getElementById("accounts").className += " active";

$().ready(function() {
    
    document.getElementById("superAdminName").value = vm.FirstName;
    document.getElementById("superadminEmail").value = vm.Email;
    document.getElementById("superAdminPhone").value = vm.PhoneNo;
    
    $('#edit-txt').on('click', function() {
        var btntxt = $("#edit-txt").text();
        if (btntxt == 'Edit') {
            
            $(".edit-account input").prop( "disabled", false );
            $(".edit-account input").toggleClass("dis-txt");	
            $('#edit-txt').text("Save");
            $('#edit-txt').attr('type', 'submit');
            return false;
        }
        $("#superAdminDetailsForm").validate({
            rules: {
                superAdminName:"required",
                superadminEmail:{
                    required:true,
                    email:true
                },
                superAdminPhone: "required"
            },
            messages: {
                superAdminName:"Please enter your Name ",
                superAdminPhone: "Please enter Phone Number",
                superadminEmail:"Please enter your Email id"
            },
            submitHandler: function(){//to pass all data of a form serial
                var formData = $("#superAdminDetailsForm").serialize();
                $.ajax({
                    url:'/accounts',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            $('#edit-txt').text("Edit");
                            $(".edit-account input").prop( "disabled", true );
                            $('#edit-txt').attr('type', 'button');
                        } else {
                            $('#edit-txt').text("Edit");
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
                return false;
            }
        });
    });

    $('#updatePassword').on('click', function() {
        $("#passwordChangeModal").validate({
            rules: {
                newPassword:"required",
                confirmpassword:{
                    equalTo : "#newPassword"
                } ,
                oldPassword: {
                required: true,
                remote:{
                    url: "/isOldPasswordCorrect/" + oldPassword,
                    type: "post"
                }
            },
            },
            messages: {
                 oldPassword:{
                     required: "Please enter Old Password ",
                     remote: "The password entered is not correct !!!"
                 },
                newPassword: "Please enter New Password",
                confirmpassword:"Retype password is incorrect"
            },
            submitHandler: function(){//to pass all data of a form serial
                alert("ttttttttt");
                //$('#passwordChangeModal').hide(); 
                var formData = $("#passwordChangeModal").serialize();
                $.ajax({
                    url:'/changePassword',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    success : function(response){
                        if(response == "true"){
                            window.location = '/accounts';
                        } else {
                            alert("password incorrect");
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
                return false;
            }
        });
    });
});