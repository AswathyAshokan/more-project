/* Author :Aswathy Ashok */
var subArray = [];
var mainArray = [];
var keyArray= [];
for(i = 0; i < vm.Job.length; i++) {
   for(var propertyName in vm.Job[i]) {
       subArray.push(vm.Job[i][propertyName]);
   }

   mainArray.push(subArray);
   keyArray.push(vm.Key[i])
   subArray = [];
}
key = keyArray;

    $(document).ready(function() {
        $('#example').DataTable( {
            data: mainArray,
            columns: [
                { title:"Customer Name"},
                { title: "Job Name" },
                { title: "Job Number" },
                { title: "Number Of Tasks" },
                { title: "Status" },
                {
                                 data:null,
                                 mRender: function (data, type, row) {
                                 for(i = 0; i < vm.Job.length; i++) {
                                     return '<div class="edit-wrapper"><span class="icn">'+'<a href="job/'+ key[i] + '/edit"><i class="fa fa-eye" aria-hidden="true"></i></a>'+"   "+'<a href="/job/'+ key[i] +'/edit"><i class="fa fa-pencil-square-o" aria-hidden="true"></i></a>'+ "  "+'<a href="/job/'+ key[i] + '/delete"><i class="fa fa-trash-o" aria-hidden="true"></i></a>'+'</span>'+'</div>'
                                     }
                                 }

                },





            ]
        });
    });