/* Author :Aswathy Ashok */


console.log(vm);

var subArray = [];
var keyArray= [];
var mainArray = [];
for(i = 0; i < vm.Task.length; i++) {
   for(var propertyName in vm.Task[i]) {
       subArray.push(vm.Task[i][propertyName]);
   }

   mainArray.push(subArray);
   keyArray.push(vm.Key[i])
   subArray = [];
}

Key=keyArray

    $(document).ready(function() {
        $('#example').DataTable( {
            data: mainArray,
            columns: [
                { title:"Project Name"},
                { title: "Task Name" },
                { title: "Location" },
                { title: "Start Date" },
                { title: "End Date" },
                { title: "Login Type"},
                 { title: "Status"},
                 {
                     data:null,
                     mRender: function (data, type, row) {
                    for(i = 0; i < vm.Task.length; i++) {
                                      return '<div class="edit-wrapper"><span class="icn">'+'<a href="task/'+ Key[i] + '/edit"><i class="fa fa-pencil-square-o" aria-hidden="true"></i></a>'+"   "+'<a href="task/'+ Key[i] + '/delete"><i class="fa fa-trash-o" aria-hidden="true"></i></a>'+'</span>'+'</div>'
                     }
                    }

                 },



            ]
        });
    });