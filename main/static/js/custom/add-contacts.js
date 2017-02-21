
/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..
//add contact.js

document.getElementById("contact").className += " active";
var pageType = array.PageType;
$(function () {
    if( pageType==  "edit") {
            
                document.getElementById("name").value = array.Name;
                document.getElementById("address").value = array.Address;
                document.getElementById("state").value = array.State;
                document.getElementById("zipcode").value = array.ZipCode;
                document.getElementById("emailAddress").value = array.Email;
                document.getElementById("phoneNumber").value = array.PhoneNumber;
                document.getElementById("contactHead").innerHTML = "Edit Contact";
                
                }
});
  $().ready(function() {
      $("#contactForm").validate({
          rules: {
              name: "required",
              emailAddress: {
                  required: true,
                  email: true
              },
              phoneNumber: {
                  required: true,
                  minlength : 10
              }
          },
          messages: {
              firstName: "Please enter your firstName",
              lastName: "Please enter your lastName",
              phoneNumber:{
                  required:"please provide a phone number",
                  minlength:"your phone number at least 10 digit long"
              },
              emailAddress: "Please enter a valid email address"
          },
          submitHandler: function() {
              var form_data = $("#contactForm").serialize();
              var contactId = array.ContactId
              if(pageType ==  "edit"){
                  $.ajax({
                      url: '/contact/'+contactId+'/edit',
                      type: 'post',
                      datatype: 'json',
                      data: form_data,
                      success : function(response) {
                          if (response =="true") {
                              window.location = '/contact';
                          } else {
                          }
                      },
                      error: function (request,status, error) {
                          console.log(error);
                      }
                  });
              } else {
                  $.ajax({
                      url: '/contact/add',
                      type: 'post',
                      datatype: 'json',
                      data: form_data,
                      success : function(response) {
                          if (response =="true") {
                              window.location = '/contact';
                          } else {
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
          window.location = '/contact';
      });
  });