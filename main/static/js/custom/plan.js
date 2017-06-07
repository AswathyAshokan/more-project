var companyTeamName = vm.CompanyTeamName;
var sessionFlag = vm.SessionFlag;

$().ready(function(){
     if (localStorage.getItem('planType') == "Family"){
         console.log("haiii");
          $("#campus").attr('disabled', false);
         $("#business").attr('disabled', false);
         $("#businessPlus").attr('disabled', false);
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
  