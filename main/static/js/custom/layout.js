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
    
    
    
    
    
    
} );
