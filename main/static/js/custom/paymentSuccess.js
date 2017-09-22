
 var companyTeamName =vm.CompanyTeamName;
 var NumberOfUsers =vm.NumberOfUsers;
    console.log("number of users",NumberOfUsers);
    console.log("")
 window.onload = function () {
        console.log("hiiii");
     
     window.location='/'+companyTeamName+'/invite'+'/'+urlParams+'/AddExtraUserByUpgradePlan';
//        $.ajax({
//                url:'/'+companyTeamName+'/invite'+'/'+urlParams+'/AddExtraUserByUpgradePlan',
//                type: 'post',
//                success : function(response) {
//                    if (response == "true" ) {
//                    } else {
//                    }
//                },
//                error: function (request,status, error) {
//                    console.log(error);
//                }
//            });
             }