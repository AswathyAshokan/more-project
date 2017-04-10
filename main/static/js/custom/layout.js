console.log(vm);
$(document).ready(function() {
    
    //checking plans
    
    if(vm.CompanyPlan == 'family'){
       
        $('#group').bind('click', false);
        $('#contact').bind('click', false);
        $('#job').bind('click', false);
        $('#nfc').bind('click', false);
        $('#crm').bind('click', false);
        
    }
    
     if(vm.CompanyPlan == 'campus'){
         
       /* $('#group').attr('disabled', true);*/
        $('#group').bind('click', false);
        $('#contact').bind('click', false);
        $('#job').bind('click', false);
        $('#nfc').bind('click', false);
        $('#crm').bind('click', false);
     }
    
    
    if(vm.CompanyPlan == 'business'){
         
       /* $('#group').attr('disabled', true);*/
        $('#group').bind('click', false);
        $('#contact').bind('click', false);
        $('#job').bind('click', false);
        $('#nfc').bind('click', false);
    }
    
    document.getElementById("username").textContent=vm.AdminFirstName;
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
