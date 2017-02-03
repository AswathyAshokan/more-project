/*Created By Farsana*/

$().ready(function() {

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
		var formData = $("#addcustomerForm").serialize();
	         $.ajax({
			url:'/add-customer',
			type:'post',
			datatype: 'json',
			data: formData,
			//call back or get response here
			success : function(response){
				console.log(response);
				 window.location ='/customer-details';

			},
		error: function (request,status, error) {
				}	
		
	
		});
	return false;
     	}
	});
	
});