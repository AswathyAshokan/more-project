/*Author: Sarath
Date:01/02/2017*/

$(function(){
    
    if(document.referrer != 'http://localhost:8080/'){
        history.pushState(null, null, '/login');
        window.addEventListener('popstate', function () {
            history.pushState(null, null, '/login');
        });
    }
    $("#signIn").click(function(){
        $.ajax({
                type    :   'POST',
                dataType: 'json',
                url     :   '/login',
                data    : {
                    'email'     :   $("#email").val(),
                    'password'  :   $("#password").val()
                },
                success :   function(data){
                   console.log(data);
                    if(data[0]=="true"){
                        if( localStorage.getItem('loginStatus') != null){
                            window.localStorage.clear();
                            window.location = '/plan'
                        } else{
                                window.location = '/'+ data[1] +'/invite';
                            }
                            
                        } else if(data[0] == "SuperAdmin"){
                            window.location ='/customer-management';
                    } else{
                        $("#login_err").css({"color": "red", "font-size": "15px"});
                        $("#login_err").html("Invalid Username or Password!").show().fadeOut( 4000 );
                    }
                }
            });
        
        return false;
    });

});

