/*Author: Sarath
Date:01/02/2017*/
//Below line is for adding active class to layout side menu..
document.getElementById("nfc").className += " active";

$(function(){
    console.log(array.Name);
    var pageType = array.PageType;
    
    if(pageType ==  "edit") {
        console.log(array);
            document.getElementById("customerName").value = array.CustomerName;
            document.getElementById("site").value = array.Site;
            document.getElementById("location").value = array.Location;
            document.getElementById("nfcNumber").value = array.NFCNumber;
            } 
    
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
                        var nfcId = array.NfcId;
                        if (pageType == "edit") {
                            $.ajax({
                                url: '/nfc/'+ nfcId +'/edit',
                                type: 'post',
                                datatype: 'html',
                                data: form_data,
                                success : function(response) {
                                    console.log(response);
                                    if (response == "true") {
                                        window.location = '/nfc';
                                    } else {

                                    }
                                },
                                error: function (request,status, error) {
                                    console.log(error);
                                }

                           });

                        } else {
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
                }
    });

});