/*Created By Farsana*/

console.log(vm)

var subArray = [];
var mainArray = [];
var GroupKeyArray = [];
var userLength = vm.Groups.length;
for(var i = 0; i < userLength; i++) {
    GroupKeyArray.push(vm.GroupKey[i])
    for(var propertyName in vm.Groups[i]) {
        subArray.push(vm.Groups[i][propertyName]);
    }
    //subArray.push(vm.Keys[i])
    mainArray.push(subArray);
    subArray = [];
}
console.log(GroupKeyArray)
var i;
var length;
length=GroupKeyArray.length;
$(document).ready(function() {
            $('#example').DataTable( {
                data: mainArray,
                columns: [
                     { title: "GroupName" },
                     { title: "GroupMembers" },
                     { "data": null,
                       "render": function ( data, type, full, meta ) {
                            for (i = 0; i<length;i++){

                                 return '<div class="edit-wrapper">'+'<span class="icn">'+'   '+'<a href="/group/'+GroupKeyArray[i] +'/edit"><i class="fa fa-pencil-square-o" aria-hidden="true"></i></a>'+'    '+'<a href="/group/'+GroupKeyArray[i] +'/delete"><i class="fa fa-trash-o" aria-hidden="true"></i></a>'+'</span>'+'</div>'

                            }
                        }
                    },

                ]
             }) ;
} );