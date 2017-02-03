/* Author :Aswathy Ashok */


console.log(vm);

var subArray = [];
var mainArray = [];
var keyArray= [];
for(i = 0; i < vm.Project.length; i++) {
   for(var propertyName in vm.Project[i]) {
       subArray.push(vm.Project[i][propertyName]);
   }

   mainArray.push(subArray);
   keyArray.push(vm.Key[i])
   subArray = [];
}
Key=keyArray
console.log(Key)

    $(document).ready(function() {
        $('#example').DataTable( {
            data: mainArray,
            columns: [
                { title:"Customer Name"},
                { title: "Project Name" },
                { title: "Project Number" },
                { title: "Number Of Tasks" },
                { title: "Status" },
                {
                                 data:null,
                                 mRender: function (data, type, row) {
                                 for(i = 0; i < vm.Project.length; i++) {
                                     return '<div class="edit-wrapper"><span class="icn">'+'<a href="project/'+ Key[i] + '/edit"><i class="fa fa-eye" aria-hidden="true"></i></a>'+"   "+'<a href="/'+ Key[i] +'"><i class="fa fa-pencil-square-o" aria-hidden="true"></i></a>'+ "  "+'<a href="/project/'+ Key[i] + '/delete"><i class="fa fa-trash-o" aria-hidden="true"></i></a>'+'</span>'+'</div>'
                                     }
                                 }

                },





            ]
        });
    });