/*Created By Farsana*/

console.log(vm)

var subArray = [];
var mainArray = [];
var keyArray = [];
var userLength = vm.Users.length;
for(var i = 0; i < userLength; i++) {
    keyArray.push(vm.Key[i])
  for(var propertyName in vm.Users[i]) {
      subArray.push(vm.Users[i][propertyName]);
  }
  //subArray.push(vm.Keys[i])
  mainArray.push(subArray);
  subArray = [];
}
console.log(keyArray)
  var i;
  var length;
  length=keyArray.length;
$(document).ready(function() {
            $('#example').DataTable( {
                data: mainArray,
                columns: [
                     { title: "FirstName" },
                     { title: "LastName" },
                     { title: "EmailId" },
                     { title: "UserType" },
                     { title: "Status"},
                     { "data": null,
                       "render": function ( data, type, full, meta ) {
                        for (i = 0; i<length;i++){

                           return '<div class="edit-wrapper">'+'<span class="icn">'+'<a href="/view-user/'+keyArray[i] +'"><i class="fa fa-eye" aria-hidden="true"></i></a>'+'    '+'<a href="/edit-user/'+keyArray[i] +'"><i class="fa fa-pencil-square-o" aria-hidden="true"></i></a>'+'    '+'<a href="/delete-user/'+keyArray[i] +'"><i class="fa fa-trash-o" aria-hidden="true"></i></a>'+'</span>'+'</div>'

                        }
                        }
                    },

                ]
    }) ;
} );