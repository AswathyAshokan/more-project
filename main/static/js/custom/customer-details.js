/*Created By Farsana*/

/*console.log(vm)

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
} );*/

console.log(vm)
$(function(){ 
    var mainArray = [];   
    var table = "";
    function createDataArray(values, keys){
        var subArray = [];
        for(i = 0; i < values.length; i++) {
            for(var propertyName in values[i]) {
                subArray.push(values[i][propertyName]);
            }
            subArray.push(keys[i])
            mainArray.push(subArray);
            subArray = [];
        }
    }
    
    function dataTableManipulate(){
        table =  $("#customer-table").DataTable({
            data: mainArray,
            "columnDefs": [{
                       "targets": -1,
                       "width": "5%",
                       "data": null,
                       "defaultContent": '<div class="edit-wrapper"><span class="icn"><i class="fa fa-pencil-square-o" aria-hidden="true" id="edit"></i><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
            }]
        });
        var item = $('<span>+</span>');
        item.click(function() {
            window.location = "/customer/add";
        });
        $('.table-wrapper .dataTables_filter').append(item);
    }
    if(vm.Values != null) {
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate();   
    console.log(mainArray);
    
    $('#customer-table tbody').on( 'click', '#edit', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[7];
        window.location = '/customer/'+ key + '/edit';
        return false;
    });


    $('#customer-table tbody').on( 'click', '#delete', function () {
        $("#myModal").modal();
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[7];
        console.log(data, key);
        $("#confirm").click(function(){
            $.ajax({
                type: "POST",
                url: '/customer/'+ key + '/delete',
                data: '',
                success: function(data){
                    if(data=="true"){
                        $('#customer-table').dataTable().fnDestroy();
                        var index = "";
                        
                        for(var i = 0; i < mainArray.length; i++) {
                           index = mainArray[i].indexOf(key);
                           if(index != -1) {
                               console.log("dddd", i);
                             break;
                           }
                        }
                        
                        console.log(i);
                        //var index = mainArray.indexOf(key);
                        console.log(index);
                        mainArray.splice(i, 1);
                        console.log(mainArray);
                        dataTableManipulate();   
                    }
                    else {
                        console.log("Removing Failed!");
                    }
                }

            });
        });
    });
});

