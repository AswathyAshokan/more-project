/*Author: Sarath
Date:01/02/2017*/
$(function(){
    $("#signIn").click(function(){
        $.ajax({
            type    :   'POST',
            url     :   '/',
            data    : {
                'email'     :   $("#email").val(),
                'password'  :   $("#password").val()
            },
            success :   function(data){
                if(data=="true"){
                    window.location = '';
                }
                else{

                }
            }
        });
        return false;
    });
});