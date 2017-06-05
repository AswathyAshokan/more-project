/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..


$().ready(function() {
    var emailId = $("#emailAddress").val();
    $('#emailCheck').on('click', function(){
        $('#emailCheck').attr('type', 'submit');
         emailId = $("#emailAddress").val();
        var formData ="&emailId="+emailId;
        $("#forgotPassword").validate({
            rules: {
                name:"required",
                emailAddress:{
                    required:true,
                    email:true
                },
                phoneNumber: "required"
            },
            messages: {
                name:"Please enter your Name ",
                emailAddress: "please enter Email id",
                phoneNumber:"Please enter your Email id"
            },
            submitHandler: function(){//to pass all data of a form serial
                $.ajax({
                    url: "/forgot-password/email-checking",
                    type: 'post',
                    datatype: 'json',
                    data: formData,
                    success : function(data) {
                        var jsonData = JSON.parse(data)
                        if (jsonData[0] == "true") {
                            window.location = '/login';
                        } else {
//                            var responsearray = response.split(','); //to seperate verification key
                            window.localStorage.clear();
                            status = localStorage.setItem('verificationKey',jsonData[1]);//set value in the local host
                            
//                            console.log("error",jsonData[2]);
//                            setTimer( 300, {
//                                300: function()
//    });                             {
//                                    display( 'notifier', 'seconds');
//                                },
//                                1: function()
//                                {
//                                    display( 'notifier', 'second');
//                                }
////                        0: function()
////                        {
////                        display( 'notifier', 'seconds');
////                    }
//                            } );
//                            function display( notifier, str )
//                            {
//                                document.getElementById( notifier ).innerHTML = str;
//                            }
//                            function toMinuteAndSecond( x )
//                            {
//                                return Math.floor( x / 60 ) + ':' +  ( x % 60 );
//                  });           }
//                            function setTimer( remain, actions )
//                            {
//                                ( function countdown()
//                                 });{
//                                    display( 'countdown', toMinuteAndSecond( remain ) );
//                                    actions[remain] && actions[remain]();
//                                    (remain -= 1) >= 0 && setTimeout( arguments.callee, 1000 );
//                   return false;               })();
//                            }
                        }
                    },
                    error: function (request,status, error) {
                        console.log(error);
                    }
                });
                return false;
            }
             
        });
        //return false;
    });
    //return false;
    $('#verifyKey').click(function(){
        if( localStorage.getItem('verificationKey') != null){
            var localStorageValue = localStorage.getItem('verificationKey');
            console.log("local",localStorageValue);
            var verificationKey = $("#verificationKey").val();
            console.log("verification",verificationKey);
            if (localStorageValue == verificationKey  ){
//                $('#verifyKey').data("target") === "#change-pass";
                $('#verifyKey'<script src="https://ajax.googleapis.com/ajax/libs/angularjs/1.6.4/angular.min.js"></script>).attr('data-target','#change-pass');
            }
            else{
                $("#verifyValidationError").css({"color": "red", "font-size": "15px"});
                $("#verifyValidationError").html("verification code is incorrect.").show();
            }
        }
    });
   $('#updateAdminPassword').click(function(){
       $("#adminPasswordChangeModal").validate({
            rules: {
                newPassword:"required",
                confirmpassword:{
                    equalTo : "#newPassword"
                } ,
                
            },
            messages: {
                
                newPassword: "Please enter New Password",
                confirmpassword:"Retype password is incorrect"
            },
            submitHandler: function(){//to pass all data of a form serial
                var formData = $("#adminPasswordChangeModal").serialize()+"&emailId="+emailId;
                $.ajax({
                    url:'/forgot-password/passwordReset',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    success : function(response){
                        if(response == "true"){
                            window.location ='/login' ;
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
    $('#cancel').click(function(){
        window.location = '/login';
    });
});