/*Author: Sarath
Date:01/02/2017*/
$(function(){
    console.log(array.Name);
    if(array.PageType == "2") {
                document.getElementById("customerName").value = array.CustomerName;
                document.getElementById("site").value = array.Site;
                document.getElementById("location").value = array.Location;
                document.getElementById("nfcNumber").value = array.NFCNumber;
    }
    /*$("#save").click(function(){
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
    });*/
    $("#addNfcForm").validate({
                    
                    rules: {
                        customerName : "required",
                        nfcNumber: "required",
                        location: "required"
                    },
                    messages: {
                        customerName: "Please enter a Customer Name",
                        nfcNumber: "Please enter a valid NFC Number",
                        location: "Please enter a Location"
                    },
    	            submitHandler: function() {
                        var form_data = $("#addNfcForm").serialize();
                        //alert(form_data);
    				    $.ajax({
                                type : 'POST',
                                url  : '/nfc/add',
                                data : form_data,
                                success : function(data){
                                                if(data=="true"){
                                                    window.location ='/nfc';
                                                }
                                                else{
                                                }
                                },
                                error: function (request,status, error) {
           					            console.log(error);
        				        }
                        });
                }
    });

});