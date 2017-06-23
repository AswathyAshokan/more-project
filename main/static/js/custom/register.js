/*Author: Sarath
Date:01/02/2017*/
console.log(vm.CountryAllData);
$(function(){
    //Enabling Register button when if I Agree check box is chechked
    $("#agree").click(function() {
      $("#register").attr("disabled", !this.checked);
    });
    window.onload = function() {
     function test(){
        vm.CountryName.sort();
        console.log( vm.CountryName);
        
    }
    }
   
/*---------------------for display heading of each webpage----------------------*/
   $('#planType').on('change', function() {
       if( document.getElementById("planType").value =="family"){
           
           document.getElementById("companyNumberLabelId").style.display = "none";
           document.getElementById("number").style.display = "none";
           document.getElementById("companyNameLabel").innerHTML = "Family Name";
       
       } else if(document.getElementById("planType").value =="campus"){
           
           document.getElementById("companyNumberLabelId").style.display = "none";
           document.getElementById("number").style.display = "none";
           document.getElementById("companyNumberLabelId").innerHTML == "Campus Number";
           document.getElementById("companyNameLabel").innerHTML = "Campus Name";
       
       }else if (document.getElementById("planType").value =="business"){
            document.getElementById("companyNumberLabelId").innerHTML == "Company Number";
           document.getElementById("companyNameLabel").innerHTML = "Company Name"; 
      
       }
   });
    

    $('#country').on('change', function() {
        document.getElementById('state').options.length = 0;
        
         var countryName = $(this).val();
         
         /*-----------for fill dialcode------------------*/
         for(var i=0; i <vm.CountryAllData.length; i++){
             if (vm.CountryAllData[i][0] == countryName ) {
                 $("#dialCode").val(vm.CountryAllData[i][1]);
                  $("#dialCode").attr('disabled', true);
                 var countryCode = vm.CountryAllData[i][2]
                 
            }
         }
        
        /*-----------for fill states----------------------*/
        var formData = "&countryName="+countryName+"&countryCode="+countryCode
         $.ajax({
             type    : 'POST',
             url     : '/register/getstate',
             data    : formData,
             success : function(data){
                 
                 var jsonData = JSON.parse(data);
                 if(jsonData[0] == "true" ){
                     jsonData[1].sort();
                     for(i = 0;i< jsonData[1].length;i++){
                         var select = document.getElementById("state"),
                         opt = document.createElement("option");
                         opt.textContent =jsonData[1][i];
                         select.appendChild(opt);
                     }
                 }
                 else{
                     //$("#register").attr('disabled', false);
                 }
             },
             error: function (request,status, error) {
                 console.log(error);
             }
         });
    });     
    //Validate and Register Company Admin
    $("#companyRegisterForm").validate({
                    
                    rules: {
                        firstName :{
                            minlength: 3,
                            required: true,
                        },
                        lastName :{
                            required: true
                        },
                        emailId:{
                            required: true,
                            email: true,
                            remote:{
                                url: "/isEmailUsed/" + emailId,
                                type: "post"
                            }
                        },
                        password:{
                            required: true,
                            minlength: 8,
                        },
                        confirmPassword:{
                            required: true,
                            equalTo: "#password",
                        },
                        companyName:{
                            required: true,
                        },
                        teamName:{
                            required: true,
                            teamRegEx: true,
                        },
                        phoneNo:"required",
                        state : "required",
                        country :"required",
                        number :"required",
                        planType :"required",
                        country : "required",
                        
                        
                    },
                    messages: {
                        firstName:{
                            required: "Please enter your First name!",
                            minlength: "First Name atleast have 3 characters!",
                        },
                        lastName:{
                            required: "Please enter your Last name!"
                        },
                        emailId:{
                            required: "Please enter your Email address!",
                            email: "Please enter a valid Email address!",
                            remote: "The Email you have entered is already in use!"
                        },
                        password:{
                            required: "Please enter a Password!",
                            minlength: "Password atleast have 8 characters!"
                        },
                        confirmPassword:{
                            required: "Please re-type your Password!",
                            equalTo: "Password doesnot match!",
                        },
                        companyName: "Please enter your Company Name!",
                        teamName:{
                            required: "Please enter your Passporte Team Name!",
                            teamRegEx: "Only lower case alphanumeric characters and '-' is allowed!",
                        },
                        number :"Please fill this column",
                        phoneNo:"Please fill this column",
                        state:"Please select state!",
                        country : "Please select country!",
                        planType: "Please select Plan Type!",
                        country : "please fill this column",
                        
                        
                    },
        
    	            submitHandler: function() {
                        var formData = $("#companyRegisterForm").serialize();
                                         
    				    $.ajax({
                            type    : 'POST',
                            url     : '/register',
                            data    : formData,
                            success : function(data){
                                console.log("haiii",data);
                                if(data=="true"){
                                    window.location = '/login';
                                }
                                else{
                                    $("#register").attr('disabled', false);
                                }
                            },
                            error: function (request,status, error) {
                                console.log(error);
                            }
                        });
                    }
    });
    $.validator.addMethod("teamRegEx", function(value, element){
            return this.optional(element) || /^[a-z0-9\-]+$/.test(value);
        },
            "Team Name must contain only letters, numbers, dashes"
    );
    
});