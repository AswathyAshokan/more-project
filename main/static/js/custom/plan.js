var companyTeamName = vm.CompanyTeamName;
var sessionFlag = vm.SessionFlag;
var selectedCompanyPlan = vm.CompanyPlan
$().ready(function(){
    if (selectedCompanyPlan == "family"){
        var businessPlusLink = document.getElementById('businessPlus');
        businessPlusLink.classList.remove("selectPlanButton");
        businessPlusLink.style.cursor = null;
        
        var compusLink = document.getElementById('campus');
        compusLink.classList.remove("selectPlanButton");
        compusLink.style.cursor = null;
        
        var businessLink = document.getElementById('business');
        businessLink.classList.remove("selectPlanButton");
        businessLink.style.cursor = null;
    
    } else if(selectedCompanyPlan == "campus"){
         var businessLink = document.getElementById('business');
         businessLink.classList.remove("selectPlanButton");
         businessLink.style.cursor = null;
        
         var businessPlusLink = document.getElementById('businessPlus');
         businessPlusLink.classList.remove("selectPlanButton");
         businessPlusLink.style.cursor = null;
        
         var familyLink = document.getElementById('family');
         familyLink.classList.remove("selectPlanButton");
         familyLink.style.cursor = null;
     
    }else if(selectedCompanyPlan =="business"){
       var familyLink = document.getElementById('family');
         familyLink.classList.remove("selectPlanButton");
         familyLink.style.cursor = null;
         var compusLink = document.getElementById('campus');
         compusLink.classList.remove("selectPlanButton");
         compusLink.style.cursor = null;
         var businessPlusLink = document.getElementById('businessPlus');
         businessPlusLink.classList.remove("selectPlanButton");
         businessPlusLink.style.cursor = null;
         
     } else{
         var familyLink = document.getElementById('family');
         familyLink.classList.remove("family");
         familyLink.style.cursor = null;
         
         var compusLink = document.getElementById('campus');
         compusLink.classList.remove("campus");
         compusLink.style.cursor = null;
         
         var businessLink = document.getElementById('business');
         businessLink.classList.remove("business");
         businessLink.style.cursor = null;
     }
    
    $(".selectPlanButton").click(function(){
        if(sessionFlag == true){
            var companyPlan = $(this).attr('id');//to get the id of selected plan
            $.ajax({              
              url:'/plan/update',
              type:'post',
              datatype: 'json',
              data: {'companyPlan':companyPlan
                    },
              //call back or get response here
              success : function(data){
                var jsonData = JSON.parse(data)
                if(jsonData[0] == "true"){
                    window.location = '/'+ jsonData[1] +'/'+ jsonData[2] +'/payment';
                } else {
                      window.location = '/login';
                }
              }
            });
        } else {
            status = localStorage.setItem('loginStatus','false');
            window.location = '/login';
        }
    });
});









 /*if (selectedCompanyPlan == "Family"){
         console.log("haiii");
         var businessPlusLink = document.getElementById('businessPlus');
         businessPlusLink.classList.remove("selectPlanButton");
         //businessPlusLink.style.cursor = null;
         var compusLink = document.getElementById('campus');
         compusLink.classList.remove("campus");
         compusLink.style.cursor = null;
         var compusLink = document.getElementById('business');
         businessLink.classList.remove("business");
         businessLink.style.cursor = null;
     
     
     } else if(selectedCompanyPlan == "Campus"){
         var compusLink = document.getElementById('business');
         businessLink.classList.remove("business");
         businessLink.style.cursor = null;
         var businessPlusLink = document.getElementById('businessPlus');
         businessPlusLink.classList.remove("selectPlanButton");
         businessPlusLink.style.cursor = null;
         var familyLink = document.getElementById('family');
         familyLink.classList.remove("family");
         familyLink.style.cursor = null;
     
     } else if(selectedCompanyPlan =="Business"){
       var familyLink = document.getElementById('family');
         familyLink.classList.remove("family");
         familyLink.style.cursor = null;
         var compusLink = document.getElementById('campus');
         compusLink.classList.remove("campus");
         compusLink.style.cursor = null;
         var businessPlusLink = document.getElementById('businessPlus');
         businessPlusLink.classList.remove("selectPlanButton");
         businessPlusLink.style.cursor = null;
         
     } else{
          var familyLink = document.getElementById('family');
         familyLink.classList.remove("family");
         familyLink.style.cursor = null;
         var compusLink = document.getElementById('campus');
         compusLink.classList.remove("campus");
         compusLink.style.cursor = null;
          var compusLink = document.getElementById('business');
         businessLink.classList.remove("business");
         businessLink.style.cursor = null;
         
     }
*/



  