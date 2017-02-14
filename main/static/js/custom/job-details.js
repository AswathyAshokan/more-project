/* Author :Aswathy Ashok */

$(function(){  
    var table = "";
    var mainArray = [];  
    
    customerFilter = function(){
        var tempArray = [];
        var selectedCustomer = $("#customerDropdown").val();
        console.log(selectedCustomer);
        if (selectedCustomer == "All Customers") {
            $('#job-details').dataTable().fnDestroy();
            dataTableManipulate(mainArray); 
        } else {
            for(i = 0; i < mainArray.length; i++){
                if (mainArray[i][0] == selectedCustomer){
                    tempArray.push(mainArray[i]);
                }
            }
            $('#job-details').dataTable().fnDestroy();
            dataTableManipulate(tempArray);
            
            $("#customerDropdown").val(selectedCustomer);
        }         
    }
    
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
    
    function dataTableManipulate(dataArray){
        table =  $("#job-details").DataTable({
            data: dataArray,
            "columnDefs": [{
                       "targets": -1,
                       "width": "5%",
                       "data": null,
                       "defaultContent": '<div class="edit-wrapper"><span class="icn"><i class="fa fa-eye" aria-hidden="true" id="view"></i><i class="fa fa-pencil-square-o" aria-hidden="true" id="edit"></i><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
            }]
        });
        
        
        var dropdownItem = $('<div class="tbl-dropdown"><select class="form-control sprites-arrow-down" id="customerDropdown" onchange="customerFilter();"><option>All Customers</option></select></div>');
        
        var addItem = $('<span>+</span>');
        addItem.click(function() {
            window.location = "/job/add";
        });
        
        $('.table-wrapper .dataTables_filter').prepend(dropdownItem).append(addItem);
        
        var customerArray = vm.UniqueCustomerNames;
        
        for(i = 0; i < customerArray.length; i++){
            $("#customerDropdown").append("<option>"+customerArray[i]+"</option>");
        }
        
    }
    if(vm.Values != null) {
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate(mainArray); 

    $('#job-details tbody').on( 'click', '#edit', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[5];
        window.location = '/job/' + key + '/edit'
    });


    $('#job-details tbody').on( 'click', '#delete', function () {
        $("#myModal").modal();
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[5];
        
        $("#confirm").click(function(){
            $.ajax({
                type: "POST",
                url: '/job/' + key + '/delete',
                data: '',
                success: function(data){
                    if(data=="true"){
                        $('#job-details').dataTable().fnDestroy();
                        var index = "";
                        
                        for(var i = 0; i < mainArray.length; i++) {
                           index = mainArray[i].indexOf(key);
                           if(index != -1) {
                               console.log("dddd", i);
                             break;
                           }
                        }
                        mainArray.splice(i, 1);
                        dataTableManipulate(mainArray);   
                    }
                    else {
                        console.log("Removing Failed!");
                    }
                }

            });
        });
    });
    
});


