/*Author: Sarath
Date:01/02/2017*/
//Below line is for adding active class to layout side menu..
document.getElementById("nfc").className += " active";
var companyTeamName = vm.CompanyTeamName;
var DynamicNotification ="";
    if (vm.NotificationNumber !=0){
        document.getElementById("number").textContent=vm.NotificationNumber;
    }else{
        document.getElementById("number").textContent="";
    }
$(function(){
    //notification
   //notification
    var notificationSorted =[[]];
    function sortByCol(arr, colIndex){
    notificationSorted=arr.sort(sortFunction);
    function sortFunction(a, b) {
        a = a[colIndex]
        b = b[colIndex]
        return (a === b) ? 0 : (a < b) ? -1 : 1
    }
}

    
     myNotification= function () {
         if (vm.NotificationArray !=null){
        console.log("hiiii");
         sortByCol(vm.NotificationArray, 6);
         console.log("jjjjj",notificationSorted);
         var reverseSorted =[[]];
         reverseSorted=notificationSorted.reverse();

        document.getElementById("notificationDiv").innerHTML = "";
        var DynamicTaskListing="";
        if (reverseSorted !=null){
            DynamicTaskListing ="<h5>"+"Notifications"+ "<button class='no-button-style' method='post' onclick='clearNotification()'>"+"clear all"+"</button>"+"</h5>"+"<ul>";
        for(var i=0;i<reverseSorted.length;i++){
            console.log("sp1");
            var timeDifference =moment(new Date(new Date(reverseSorted[i][6]*1000)), "YYYYMMDD").fromNow();
            DynamicTaskListing += "<li>"+"User"+" "+reverseSorted[i][2]+" "+reverseSorted[i][3]+"  "+"delay to reach location"+" "+reverseSorted[i][4]+" "+"for task"+" "+reverseSorted[i][5]+" <span>"+timeDifference+"</span>"+"</li>";
            
            
        }
            $("#notificationDiv").prepend(DynamicTaskListing+"</ul>");
            document.getElementById("number").textContent="";
            $.ajax({
                url:'/'+ companyTeamName + '/notification/update',
                type: 'post',
                success : function(response) {
                    if (response == "true" ) {
                    } else {
                    }
                },
                error: function (request,status, error) {
                    console.log(error);
                }
            }); 
        }else{
             document.getElementById("notificationDiv").innerHTML = "";
            DynamicTaskListing ="<h5>"+" No New Notifications"+"</h5>";
                        $("#notificationDiv").prepend(DynamicTaskListing);
            
        }
        
        }else{
             document.getElementById("notificationDiv").innerHTML = "";
            DynamicTaskListing ="<h5>"+" No New Notifications"+"</h5>";
            $("#notificationDiv").prepend(DynamicTaskListing);
        }
}
     
     
     
     clearNotification= function () {
          document.getElementById("notificationDiv").innerHTML = "";
          $.ajax({
                url:'/'+ companyTeamName + '/notification/delete',
                type: 'post',
                success : function(response) {
                    if (response == "true" ) {
                        DynamicTaskListing ="<h5>"+" No New Notifications"+"</h5>";
                        $("#notificationDiv").prepend(DynamicTaskListing);
                    } else {
                    }
                },
                error: function (request,status, error) {
                    console.log(error);
                }
            }); 
         
         
         
     }
    
     
    
    var pageType = vm.PageType;
    //Chech whether Pagtype is Add or Edit NFC Tag 
    if(pageType ==  "edit") {
        console.log(vm);
            document.getElementById("customerName").value = vm.CustomerName;
            document.getElementById("site").value = vm.Site;
            document.getElementById("location").value = vm.Location;
            document.getElementById("nfcNumber").value = vm.NFCNumber;
            document.getElementById("pageTitle").innerHTML = "Edit NFC Tag"
            } 
    //Add new NFC Tag and perform Validation
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
                         $("#save").attr('disabled', true);
                        var form_data = $("#addNfcForm").serialize();
                        //alert(form_data);
                        var nfcId = vm.NfcId;
                        if (pageType == "edit") {
                            $.ajax({
                                url: '/'+ companyTeamName +'/nfc/'+ nfcId +'/edit',
                                type: 'post',
                                datatype: 'json',
                                data: form_data,
                                success : function(response) {
                                    console.log(response);
                                    if (response == "true") {
                                        window.location = '/'+companyTeamName+'/nfc';
                                    } else {
                                        $("#save").attr('disabled', false);
                                    }
                                },
                                error: function (request,status, error) {
                                    console.log(error);
                                }

                           });

                        } else {
                            $.ajax({
                                    type : 'POST',
                                    url  : '/'+companyTeamName+'/nfc/add',
                                    data : form_data,
                                    success : function(data){
                                                    if(data=="true"){
                                                        window.location ='/'+companyTeamName+'/nfc';
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

    $("#cancel").click(function() {
            window.location = '/'+companyTeamName+'/nfc';
    });

});