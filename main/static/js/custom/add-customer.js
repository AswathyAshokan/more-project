//Created By Farsana
//Below line is for adding active class to layout side menu..
console.log(vm);
document.getElementById("crm").className += " active";
console.log("vm.CustomerName",vm.CustomerName)
var companyTeamName = vm.CompanyTeamName;


$().ready(function() {
    if(vm.PageType == "edit"){        
            
            document.getElementById("customername").value = vm.CustomerName;
            document.getElementById("contactperson").value = vm.ContactPerson;
            document.getElementById("country").value = vm.Country;
            document.getElementById("email").value = vm.Email;
            document.getElementById("phone").value = vm.Phone;
            document.getElementById("address").value = vm.Address;
            document.getElementById("state").value = vm.State;
            document.getElementById("zipcode").value = vm.ZipCode;
            document.getElementById("customerEdit").innerHTML = "Edit Customer"
    }
    $("#addcustomerForm").validate({
        rules: {
          customername:{
              required:true,
              remote:{
                  url: "/iscustomernameused/" + customername + "/" + vm.PageType + "/" + vm.CustomerName,
                  type: "post"
              }
              
          },
          contactperson:"required",
          email:{
              required:true,
              email:true
          },
            phone:"required",
            address:"required",
            country:"required",
            state: "required",
            zipcode: "required"
      },
        messages: {
            customername:{
                required: "Enter Customer Name ",
                remote: "Customer Name is already in use !"
                },
            contactperson:"Enter Contact Person",
            phone: {
                required:"Enter Phone Number"
            },
            address:"Enter your Address",
            state: "Enter your State",
            zipcode:"Enter zipcode  ",
            country:"Enter country name ",
            email:"Enter valid Email id"
    },
        submitHandler: function(){//to pass all data of a form serial
            $("#saveButton").attr('disabled', true);
            if (vm.PageType == "edit"){
                var formData = $("#addcustomerForm").serialize();
                var customerId = vm.CustomerId;
                $.ajax({
                    url:'/' + companyTeamName +'/customer/'+ customerId + '/edit',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/' + companyTeamName +'/customer';
                        }else {
                            $("#saveButton").attr('disabled', false);
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
            } else {
                var formData = $("#addcustomerForm").serialize();
                $.ajax({
                    url:'/' + companyTeamName +'/customer/add',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/' + companyTeamName +'/customer';
                        }else {
                            $("#saveButton").attr('disabled', false);
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
            }
            return false;
        }
    });
    
    $("#cancel").click(function() {
            window.location = '/' + companyTeamName +'/customer';
    });
});