/*Author: Sarath
Date:01/02/2017*/
$(function(){
    $("#signIn").click(function(){
         //alert("hi");
        $.ajax({
            type    :   'POST',
            url     :   '/',
            data    : {
                'email'     :   $("#email").val(),
                'password'  :   $("#password").val()
            },
            success :   function(data){
                if(data=="true"){
                    window.location = '/job';
                }
                else{
                    $("#login_err").css({"color": "red", "font-size": "15px"});
					$("#login_err").html("Invalid Username or Password!").show().fadeOut( 4000 );
                }
            }
        });
        return false;
    });
    
    
});

