/*Author: Sarath
Date:01/02/2017*/
var content = '<div class="edit-wrapper"><span class="icn"><i class="fa fa-pencil-square-o" aria-hidden="true" id="edit"></i><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
var subArray = [];
var mainArray = [];
//alert(vm.Values);
for(i = 0; i < vm.Values.length; i++) {
    for(var propertyName in vm.Values[i]) {
        subArray.push(vm.Values[i][propertyName]);
    }
    subArray.push(vm.Keys[i])
    mainArray.push(subArray);
    subArray = [];
}
console.log(mainArray);


$(function(){

  var table =  $("#nfc_details").DataTable({
        data: mainArray,
        "columnDefs": [ {
                   "targets": -1,
                   "width": "5%",
                   "data": null,
                   "defaultContent": '<div class="edit-wrapper"><span class="icn"><i class="fa fa-pencil-square-o" aria-hidden="true" id="edit"></i><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
               } ]
           } );
    /*var table =  $("#nfc_details").DataTable({
                       "processing": true,
                       "serverSide": true,
                       "ajax": {
                                   "url": "/datatable",

                                   "dataSrc": function(data){
                                            var subArray = [];
                                            var mainArray = [];
                                            alert(vm.Values);
                                            alert(data);
                                            for(i = 0; i < data.length; i++) {
                                                for(var propertyName in data[i]) {
                                                    subArray.push(data[i][propertyName]);
                                                }
                                                subArray.push(vm.Keys[i])
                                                mainArray.push(subArray);
                                                subArray = [];
                                            }
                                            alert(mainArray);
                                            return mainArray;
                                            }
                                  },

                                   "columnDefs": [ {
                                                      "targets": -1,
                                                      "width": "5%",
                                                      "data": null,
                                                      "defaultContent": '<div class="edit-wrapper"><span class="icn"><i class="fa fa-pencil-square-o" aria-hidden="true" id="edit"></i><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
                                                  } ]
           } );*/

           $('#nfc_details tbody').on( 'click', '#edit', function () {
                   var data = table.row( $(this).parents('tr') ).data();
                   alert( data[0] +"'s key is: "+ data[4] );
            } );


           $('#nfc_details tbody').on( 'click', '#delete', function () {
                          $("#myModal").modal();
                           var data = table.row( $(this).parents('tr') ).data();
                           var key = data[4];

                          $("#confirm").click(function(){

                              $.ajax({
                                type: "POST",
                                url: "/nfc/delete",
                                data: {
                                    Key : key
                                },
                                success: function(data){
                                    if(data=="true"){
                                        //$('#nfc_details').DataTable().ajax.reload();
                                        //$(this).closest('tr').remove();
                                         window.location ='/nfc';
                                    }
                                    else{
                                        console.log("Removing Failed!");
                                    }
                                }

                              });
                          });

            });
    });
