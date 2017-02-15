/*Created By Farsana*/
//Below line is for adding active class to layout side menu..
document.getElementById("crm").className += " active";

//push data to data table
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
    
// Add a empty column to datatable and fill with edit delete and list icons
    function dataTableManipulate(){
        table =  $("#customer-table").DataTable({
            data: mainArray,
            "columnDefs": [{
                "targets": -1,
                "width": "10%",
                "data": null,
                "defaultContent": '<div class="edit-wrapper"><span class="icn"><i class="fa fa-eye" aria-hidden="true" id="view"></i><i class="fa fa-pencil-square-o" aria-hidden="true" id="edit"></i><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
            }]
        });
        
// Add a plus symbol in webpage for add new groups
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
    
// list job details of each customer when click on list icon
    $('#customer-table tbody').on( 'click', '#view', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var cusomerId = data[7];
        window.location = '/customer/'+ cusomerId + '/job';
        return false;
    });

// Edit customer details when click on edit icon
    $('#customer-table tbody').on( 'click', '#edit', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[7];
        window.location = '/customer/'+ key + '/edit';
        return false;
    });
    
// Delete customer details when click on delete icon
    $('#customer-table tbody').on( 'click', '#delete', function () {
        $("#myModal").modal();
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[7];
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
                        mainArray.splice(i, 1);
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

