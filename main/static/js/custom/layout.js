console.log(vm);

$(document).ready(function() {
    
    //checking plans
    
     if(vm.CompanyPlan == 'family' ){
         $("#contact").remove();
         $("#crm").remove();
         $("#leave").remove();
         $("#fitToWork").remove();
         $("#time-sheet").remove();
         $("#consent").remove();
         $("#job").remove();
     } else if(vm.CompanyPlan == 'campus'){
         $("#contact").remove();
         $("#crm").remove();
         $("#leave").remove();
         $("#fitToWork").remove();
         $("#time-sheet").remove();
         $("#consent").remove();
         $("#job").remove();
    }
    document.getElementById("username").textContent=vm.AdminFirstName;
    document.getElementById("imageId").src=vm.ProfilePicture;
    if (vm.ProfilePicture ==""){
        document.getElementById("imageId").src="/static/images/default.png"
    }
    if(vm.CompanyPlan == "family")
    {
        $('#planChange').attr('data-target','#family');
    } else if (vm.CompanyPlan == "campus") {
        $('#planChange').attr('data-target','#campus');
    }else if (vm.CompanyPlan == "business") {
        $('#planChange').attr('data-target','#business');
    }else if (vm.CompanyPlan == "businessPlus") {
        $('#planChange').attr('data-target','#business-plus');
    }
    

} );
