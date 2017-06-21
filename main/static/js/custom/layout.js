console.log(vm);
$(document).ready(function() {
    
    //checking plans
    
    if(vm.CompanyPlan == 'family'){
        var parent = document.getElementById("menuItems");
        var group = document.getElementById("group");
        var contact = document.getElementById("contact");
        var job = document.getElementById("job");
        var nfc = document.getElementById("nfc");
        var crm = document.getElementById("crm");
        parent.removeChild(group);
        parent.removeChild(contact);
        parent.removeChild(job);
        parent.removeChild(nfc);
        parent.removeChild(crm);
        
    }
    
     if(vm.CompanyPlan == 'campus'){
         
       /* $('#group').attr('disabled', true);*/
        var parent = document.getElementById("menuItems");
        var group = document.getElementById("group");
        var contact = document.getElementById("contact");
        var job = document.getElementById("job");
        var nfc = document.getElementById("nfc");
        var crm = document.getElementById("crm");
        parent.removeChild(group);
        parent.removeChild(contact);
        parent.removeChild(job);
        parent.removeChild(nfc);
        parent.removeChild(crm);
     }
    
    
    if(vm.CompanyPlan == 'business'){
         
       var parent = document.getElementById("menuItems");
        var group = document.getElementById("group");
        var contact = document.getElementById("contact");
        var job = document.getElementById("job");
        var nfc = document.getElementById("nfc");
        var crm = document.getElementById("crm");
        parent.removeChild(group);
        parent.removeChild(contact);
        parent.removeChild(job);
        parent.removeChild(nfc);
        parent.removeChild(crm);
    }
    
    document.getElementById("username").textContent=vm.AdminFirstName;
    document.getElementById("imageId").src=vm.ProfilePicture;
    if (vm.ProfilePicture ==""){
        document.getElementById("imageId").src="https://firebasestorage.googleapis.com/v0/b/passporte-b9070.appspot.com/o/profilePicturesOfAdmin%2Fdefault.png?alt=media&token=7444c8f3-2254-4494-9588-a41ecee96b01"
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
