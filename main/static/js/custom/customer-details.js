/*Created By Farsana*/

console.log(vm)

var subArray = [];
var mainArray = [];
var customerKeyArray = [];
var userLength = vm.Customers.length;
for(var i = 0; i < userLength; i++) {
    customerKeyArray.push(vm.CustomerKey[i])
  for(var propertyName in vm.Customers[i]) {
      subArray.push(vm.Customers[i][propertyName]);
  }
  //subArray.push(vm.Keys[i])
  mainArray.push(subArray);
  subArray = [];
}
console.log(customerKeyArray)
  var i;
  var length;
  length=customerKeyArray.length;
$(document).ready(function() {
            $('#example').DataTable( {
                data: mainArray,
                columns: [
                     { title: "CustomerName" },
                     { title: "ContactPerson" },
                     { title: "Address" },
                     { title: "Phone" },
                     { title: "Email"},
                     { title: "State"},
                     { title: "ZipCode"},
                     { "data": null,
                       "render": function ( data, type, full, meta ) {
                        for (i = 0; i<length;i++){

                                return '<div class="edit-wrapper">'+'<span class="icn">'+'<a href="/view-user/'+customerKeyArray[i] +'"><i class="fa fa-eye" aria-hidden="true"></i></a>'+' '+'<a href="/customer/'+customerKeyArray[i] +'/edit"><i class="fa fa-pencil-square-o" aria-hidden="true"></i></a>'+' '+'<a href="/customer/'+customerKeyArray[i] +'/delete"><i class="fa fa-trash-o" aria-hidden="true"></i></a>'+'</span>'+'</div>'


                        }
                        }
                    },

                ]
    }) ;
} );