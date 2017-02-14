/*Created By Farsana*/

$().ready(function() {
    if(vm.PageType == "edit"){        
            document.getElementById("customername").value = vm.CustomerName;
            document.getElementById("contactperson").value = vm.ContactPerson;
            document.getElementById("email").value = vm.Email;
            document.getElementById("phone").value = vm.Phone;
            document.getElementById("address").value = vm.Address;
            document.getElementById("state").value = vm.State;
            document.getElementById("zipcode").value = vm.ZipCode;
            document.getElementById("customerEdit").innerHTML = "Edit Customer"
    }
	$("#addcustomerForm").validate({
	  rules: {
          customername: "required",
          contactperson:"required",
          email:{
              required:true,
              email:true
          },
          phone:{
              required: true,
              minlength:10,
              maxlength:10
          },
          address:"required",
          state: "required",
          zipcode: "required"
      },
        messages: {
            customername:"please enter customer name ",
            contactperson:"please enter contact person",
            phone: {
                required:"please enter phone no",
                minlength:"enter 10 digit"
            },
            address:"please enter your address",
            state: "please enter your state",
            zipcode:"please enter zipcode  ",
            email:"please enter your email id"
    },
        submitHandler: function(){//to pass all data of a form serial
            if (vm.PageType == "edit"){
                var formData = $("#addcustomerForm").serialize();
                var customerId = vm.CustomerId;
                $.ajax({
                    url:'/customer/'+ customerId + '/edit',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/customer';
                        }else {
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
            } else {
                var formData = $("#addcustomerForm").serialize();
                $.ajax({
                    url:'/customer/add',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/customer';
                        }else {
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
            }
            return false;
        }
    });
});