var companyTeamName = vm.CompanyTeamName;
var sessionFlag = vm.SessionFlag;
var selectedCompanyPlan = vm.CompanyPlan
$().ready(function(){
    console.log(selectedCompanyPlan);
    if (selectedCompanyPlan == "family"){
        
        document.getElementById('businessDiv').className += ' disable'
        document.getElementById('campusDiv').className += ' disable'
        
        document.getElementById("selectPlanButton")
        var compusLink = document.getElementById('campus');
        compusLink.classList.remove("selectPlanButton");
        compusLink.style.cursor = null;
        
        var businessLink = document.getElementById('business');
        businessLink.classList.remove("selectPlanButton");
        businessLink.style.cursor = null;
    
    } else if(selectedCompanyPlan == "campus"){
         document.getElementById('businessDiv').className += ' disable'
         document.getElementById('familyDiv').className += ' disable'
        
         var businessLink = document.getElementById('business');
         businessLink.classList.remove("selectPlanButton");
         businessLink.style.cursor = null;
        var familyLink = document.getElementById('family');
         familyLink.classList.remove("selectPlanButton");
         familyLink.style.cursor = null;
     
    }else if (selectedCompanyPlan =="business"){
        
        document.getElementById('campusDiv').className += ' disable'
        document.getElementById('familyDiv').className += ' disable'
        
        var familyLink = document.getElementById('family');
         familyLink.classList.remove("selectPlanButton");
         familyLink.style.cursor = null;
         var compusLink = document.getElementById('campus');
         compusLink.classList.remove("selectPlanButton");
         compusLink.style.cursor = null;
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





  