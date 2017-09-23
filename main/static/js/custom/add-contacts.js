
/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..
//add contact.js

var DynamicNotification ="";
    if (vm.NotificationNumber !=0){
        document.getElementById("number").textContent=vm.NotificationNumber;
    }else{
        document.getElementById("number").textContent="";
    }
document.getElementById("contact").className += " active";
var pageType = vm.PageType;
 var selectedCustomerNames = [];
//console.log( "page type",document.getElementById("phoneNumber").value);
var companyTeamName = vm.CompanyTeamName
$(function () {
    
    
    
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
    
     
    
    if( pageType  ==  "edit") {
        var selectArray =  vm.EditCustomerKey;
        console.log("contact",selectArray);
        $("#customerId").val(selectArray);
        document.getElementById("name").value =vm.Name;
        document.getElementById("address").value =vm.Address;
        document.getElementById("state").value =vm.State;
        document.getElementById("zipcode").value =vm.ZipCode;
        document.getElementById("emailAddress").value =vm.Email;
        document.getElementById("phoneNumber").value =vm.PhoneNumber;
        document.getElementById("country").value =vm.Country;
        document.getElementById("contactHead").innerHTML = "Edit Contact";
    }
});
  $().ready(function() {
     if( pageType  ==  "edit") {
         
            $("#contactForm").validate({
          
          rules: {
              phoneNumber: {
                  required: true,
                  remote:{
                      
//                      url: "/isPhoneNumberUsed/" + phoneNumber,
                      url: "/isPhoneNumberUsed/" + phoneNumber+ "/" +vm.PageType+ "/" + vm.PhoneNumber,
                      type: "post"
                  }
              },
              emailAddress: {
                  required: true,
                  email: true,
                  remote:{
//                      url: "/isemailAddressUsed/" + emailAddress,
               url: "/isemailAddressUsed/" + emailAddress+ "/" + vm.PageType + "/" + vm.Email,
                      type: "post"
                  }
              },
              name: "required",
              address:"required",
              zipcode:"required",
              country:"required",
              state:"required"
              
          },
          messages: {
               phoneNumber: {
                  required: "Enter phone number",
                  remote: "Phone number already exists!"
              },
              emailAddress: {
                  required: "Enter email Address",
                  remote: "Email Address already exists!"
              },
              name: "Enter name",
              address:"Enter address",
              zipcode:"Enter zipcode"
          },
          submitHandler: function() {
               $("#saveButton").attr('disabled', true);
              var form_data = $("#contactForm").serialize();
              var contactId =vm.ContactId
              $("#customerId option:selected").each(function () {
                  var $this = $(this);
                  if ($this.length) {
                      var selectedCustomerName = $this.text();
                      selectedCustomerNames.push( selectedCustomerName);
                  }
              });
              console.log("customer name",selectedCustomerNames);
              for(i = 0; i < selectedCustomerNames.length; i++) {
                  form_data = form_data+"&customerName="+selectedCustomerNames[i];
              }
                  $.ajax({
                      url:'/'+ companyTeamName + '/contact/'+contactId+'/edit',
                      type: 'post',
                      datatype: 'json',
                      data: form_data,
                      success : function(response) {
                          if (response =="true") {
                              
                              window.location = '/' + companyTeamName +'/contact';
                              
                              
                          } else {
                              $("#saveButton").attr('disabled', false);
                          }
                          
                      },
                      error: function (request,status, error) {
                          console.log(error);
                      }
                  });
              
          }
      });
     }
      if( pageType  ==  "add") {
         
            $("#contactForm").validate({
          
          rules: {
              phoneNumber: {
                  required: true,
                  remote:{
                      
                      url: "/isPhoneNumberUsed/" + phoneNumber,
//                      url: "/isPhoneNumberUsed/" + phoneNumber+ "/" +vm.PageType+ "/" + vm.PhoneNumber,
                      type: "post"
                  }
              },
              emailAddress: {
                  required: true,
                  email: true,
                  remote:{
                      url: "/isemailAddressUsed/" + emailAddress,
//               url: "/isemailAddressUsed/" + emailAddress+ "/" + vm.PageType + "/" + vm.Email,
                      type: "post"
                  }
              },
              name: "required",
              address:"required",
              zipcode:"required",
              country:"required",
              state:"required"
              
          },
          messages: {
               phoneNumber: {
                  required: "Enter phone number",
                  remote: "Phone number already exists!"
              },
              emailAddress: {
                  required: "Enter email Address",
                  remote: "Email Address already exists!"
              },
              name: "Enter name",
              address:"Enter address",
              zipcode:"Enter zipcode"
          },
          submitHandler: function() {
               $("#saveButton").attr('disabled', true);
              var form_data = $("#contactForm").serialize();
              var contactId =vm.ContactId
              $("#customerId option:selected").each(function () {
                  var $this = $(this);
                  if ($this.length) {
                      var selectedCustomerName = $this.text();
                      selectedCustomerNames.push( selectedCustomerName);
                  }
              });
              console.log("customer name",selectedCustomerNames);
              for(i = 0; i < selectedCustomerNames.length; i++) {
                  form_data = form_data+"&customerName="+selectedCustomerNames[i];
              }
              $.ajax({
                      url: '/'+ companyTeamName +'/contact/add',
                      type: 'post',
                      datatype: 'json',
                      data: form_data,
                      success : function(response) {
                          console.log("dgggd",response);
                          if (response =="true") {
                              window.location = '/' + companyTeamName +'/contact';
                          } else {
                              $("#saveButton").attr('disabled', false);
                          }
                      },
                      error: function (request,status, error) {
                          console.log(error);
                      }
                  });
              
          }
      });
     }
   
      $("#cancel").click(function() {
          window.location = '/'+ companyTeamName +'/contact';
      });
  });