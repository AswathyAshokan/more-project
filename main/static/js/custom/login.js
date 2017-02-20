/*Author: Sarath
Date:01/02/2017*/
$(function(){
    
    if(document.referrer != 'http://localhost:8080/'){
        history.pushState(null, null, '/');
        window.addEventListener('popstate', function () {
            history.pushState(null, null, '/');
        });
    }
    
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
                    window.location = '/invite';
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

