var companyTeamName = vm.CompanyTeamName;
var sessionFlag = vm.SessionFlag;
console.log("flaggggg",sessionFlag);
$().ready(function(){
    console.log(companyTeamName);




    redirect = function(){
         window.location = '/'+companyTeamName+'/dashBoard';
    }

    reply_click= function(plan){
    var selectedCompanyPlan=plan;
    if(sessionFlag == true){
                $.ajax({
                  url:'/plan/update',
                  type:'post',
                  datatype: 'json',
                  data: {'companyPlan':selectedCompanyPlan
                        },
                  //call back or get response here
                  success : function(data){
                    var jsonData = JSON.parse(data)
                    if(jsonData[0] == "true"){;
                        window.location ='/' + jsonData[1] +'/invite';
                       /* window.location = '/'+ jsonData[1] +'/'+ jsonData[2] +'/payment';*/
                    } else {
                        console.log("haiiii");
                          window.location = '/login';
                    }
                  }
                });
          }
            else {
                $("#plan-confirm").modal();
                status = localStorage.setItem('loginStatus','false');
                //window.location = '/login';
            }


    }
});





