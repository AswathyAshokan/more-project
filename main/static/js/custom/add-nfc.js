/*Author: Sarath
Date:01/02/2017*/
$(function(){
    $("#save").click(function(){
        $.ajax({
            type : 'POST',
            url  : '/nfc/add',
            data : {
                'customerName' : $("#customerName").val(),
                'site'         : $("#site").val(),
                'location'     : $("#location").val(),
                'nfcNumber'    : $("#nfcNumber").val()
            },
            success : function(data){
                            if(data=="true"){
                                window.location ='/nfc';

                            }
                            else{

                            }
                        }
        });
        return false;
    });

});