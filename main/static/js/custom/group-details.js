/*Created By Farsana*/

console.log(vm)

var subArray = [];
var mainArray = [];
var keyArray = [];
var userLength = vm.Groups.length;
for(var i = 0; i < userLength; i++) {
    keyArray.push(vm.Key[i])
  for(var propertyName in vm.Groups[i]) {
      subArray.push(vm.Groups[i][propertyName]);
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
                     { title: "GroupName" },
                     { title: "GroupMembers" },
                     { "data": null,
                       "render": function ( data, type, full, meta ) {
                        for (i = 0; i<length;i++){

                           return '<div class="edit-wrapper">'+'<span class="icn">'+'<a href="/view-group/'+keyArray[i] +'"><i class="fa fa-eye" aria-hidden="true"></i></a>'+'    '+'<a href="/edit-group/'+keyArray[i] +'"><i class="fa fa-pencil-square-o" aria-hidden="true"></i></a>'+'    '+'<a href="/delete-group/'+keyArray[i] +'"><i class="fa fa-trash-o" aria-hidden="true"></i></a>'+'</span>'+'</div>'

                        }
                        }
                    },

                ]
    }) ;
} );