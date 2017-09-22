$(function(){
 var companyTeamName =vm.CompanyTeamName;
 var NumberOfUsers =vm.NumberOfUsers;
 successFunction= function () {
        console.log("hiiii");
        $.ajax({
                url:'/'+companyTeamName+'/invite'+'/'+NumberOfUsers+'/AddExtraUserByUpgradePlan',
                type: 'post',
                success : function(response) {
                    if (response == "true" ) {
                    } else {
                    }
                },
                error: function (request,status, error) {
                    console.log(error);
                }
            });
             }
});