/* Author :Aswathy Ashok */
console.log(vm);


var subArray = [];
var keyArray= [];
var mainArray = [];
for(i = 0; i < vm.User.length; i++) {
   for(var propertyName in vm.User[i]) {
       subArray.push(vm.User[i][propertyName]);
   }
   //subArray.push(vm.Keys[i])
   mainArray.push(subArray);
   keyArray.push(vm.Key[i])
   subArray = [];
}
//var dataSet = [mainArray];
Key=keyArray
console.log(Key)

    $(document).ready(function() {
        $('#example').DataTable( {
            data: mainArray,
            columns: [
                { title:"Name"},
                { title: "Address" },
                { title: "State" },
                { title: "Zipcode" },
                { title: "Email" },
                { title: "Phone Number"},
                {
                  data:null,
                 mRender: function (data, type, row) {
                 for(i = 0; i < vm.User.length; i++) {
                     return '<div class="edit-wrapper"><span class="icn">'+'<a href="editContact/'+ Key[i] + '"><i class="fa fa-pencil-square-o" aria-hidden="true"></i></a>'+"   "+'<a href="deleteContact/'+ Key[i] + '"><i class="fa fa-trash-o" aria-hidden="true"></i></a>'+'</span>'+'</div>'
                     }
                }

                },



            ]
        } );


    } );