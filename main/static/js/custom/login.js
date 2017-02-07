/*Author: Sarath
Date:01/02/2017*/
$(function(){
    $("#signIn").click(function(){
        $.ajax({
            type    :   'POST',
            url     :   '/',
            data    : {
                'email'     :   $("#email").val(),
                'password'  :   $("#password").val()
            },
            success :   function(data){
                if(data=="true"){
                    window.location = '';
                }
                else{

                }
            }
        });
        return false;
    });
});

app = angular.module('app',['firebase']);
app.controller('AuthCtrl',[
    '$scope','$rootScope','$firebaseAuth',function($scope,$rootScope,$firebaseAuth){
        var ref = new Firebase('https://passporte.firebaseio.com');
        $rootScope.auth = $firebaseAuth(ref);
        $scope.signIn = function(){
            $rootScope.auth.signIn
        }
    }
]);