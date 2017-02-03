/*Author: Sarath
Date:01/02/2017*/
$(function(){
    $("#register").click(function(){
        $.ajax({
            type    : 'POST',
            url     : '/register',
            data    : {
                'firstName'     : $("#firstName").val(),
                'lastName'      : $("#lastName").val(),
                'phoneNo'       : $("#phoneNo").val(),
                'emailId'       : $("#emailId").val(),
                'password'      : $("#password").val(),
                'companyName'   : $("#companyName").val(),
                'address'       : $("#address").val(),
                'state'         : $("#state").val(),
                'zipCode'       : $("#zipCode").val()
            },
            success : function(data){
                if(data=="true"){
                    window.location='';
                }
                else{

                }
            }

        });
        return false;
    });
});