
console.log(vm);
document.getElementById("accounts").className += " active";

$().ready(function() {
   /* if(vm.Pageatype == "add"){ */
        document.getElementById("superAdminName").value = vm.FirstName;
        document.getElementById("superadminEmail").value = vm.Email
        document.getElementById("superAdminPhone").value = vm.PhoneNo;
    /*} else {
        
        
        
        
        
    }*/
   
    
    
    $('#tbl_details').dataTable();

		// $(".edit-account input").prop( "disabled", false );
		$(".edit-account .dis-txt").prop( "disabled", true );


		$('#edit-txt').on('click', function () {

	        $(".edit-account input").prop( "disabled", false );
	        $(".edit-account input").toggleClass("dis-txt");	        

	        var btntxt = $("#edit-txt").text();

	        if (btntxt=='Edit') {
	        	$('#edit-txt').text("Save");
	        	$(".edit-account input").prop( "disabled", false );
	        }
	        else if (btntxt=='Save') {
	        	$('#edit-txt').text("Edit");
	        	$(".edit-account input").prop( "disabled", true );
	        }


	    });
    
    
     /*$("#editSuperAdminDetailsForm").validate({
        /*rules: {
            superAdminName:"required",
            superadminEmail:{
                required:true,
                email:true
            },
          superAdminPhone: "required"
        },
         messages: {
             superAdminName:"Please enter Customer Name ",
             superAdminPhone: "Please enter Phone Number",
             superadminEmail:"Please enter your Email id"
         },
        submitHandler: function(){//to pass all data of a form serial
            $("#saveButton").attr('disabled', true);
            if (vm.PageType == "edit"){
                var formData = $("#editSuperAdminDetailsForm").serialize();
                $.ajax({
                    url:'/accounts/edit',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/accounts';
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
        
    });*/
    
});