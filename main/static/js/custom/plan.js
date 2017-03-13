var companyTeamName = vm.CompanyTeamName;
var sessionFlag = vm.SessionFlag;

$().ready(function(){
    
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
                    window.location = '/'+ jsonData[1] +'/invite';
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
  