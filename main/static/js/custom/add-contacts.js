
/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..
//add contact.js

document.getElementById("contact").className += " active";
var pageType = vm.PageType;
console.log( "page type",pageType);
var companyTeamName = vm.CompanyTeamName
$(function () {
    if( pageType  ==  "edit") {
            
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
      $("#contactForm").validate({
          rules: {
              name: "required",
              country:"required",
              state:"required ",
              address:"required",
              zipcode:"required",
              emailAddress: {
                  required: true,
                  email: true
              },
              phoneNumber: {
                  required: true
                 
              }
          },
          messages: {
              name: "Enter name",
              country:"Enter country",
              state:"Enter State ",
              address:"Enter address",
              zipcode:"Enter zipcode",
              phoneNumber:{
                  required:"Enter phone number"
                  
              },
              emailAddress: "Enter a valid email address"
          },
          submitHandler: function() {
               $("#saveButton").attr('disabled', true);
              var form_data = $("#contactForm").serialize();
              var contactId =vm.ContactId
              if(pageType ==  "edit"){
                  $.ajax({
                      url:'/'+ companyTeamName + '/contact/'+contactId+'/edit',
                      type: 'post',
                      datatype: 'json',
                      data: form_data,
                      success : function(response) {
//                          var jsonData = JSON.parse(data)
                          if (response =="true") {
                              
                              window.location = '/' + companyTeamName +'/contact';
                              
                              console.log("listRes",jsonData[1]);
                              console.log("res",jsonData[2]);
                              console.log("error",jsonData[3]);
                          } else {
                              $("#saveButton").attr('disabled', false);
                          }
                          
                      },
                      error: function (request,status, error) {
                          console.log(error);
                      }
                  });
              } else {
                  $.ajax({
                      url: '/'+ companyTeamName +'/contact/add',
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
          }
      });
      $("#cancel").click(function() {
          window.location = '/'+ companyTeamName +'/contact';
      });
  });