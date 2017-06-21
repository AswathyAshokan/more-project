console.log(vm);
document.getElementById("customer-management").className += " active";
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
    
/*Function for assigning data array into data table*/
    function dataTableManipulate(){
        table =  $("#customerManagement-table").DataTable({
            data: mainArray,
            "columnDefs": [{
                       "targets": -1,
                       "width": "5%",
                       "data": null,
                       "defaultContent": '<div class="edit-wrapper"><span class="icn"><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
            }]
        });
        
/*Add a plus symbol in webpage for add new groups*/
        /*var item = $('<span>+</span>');
        item.click(function() {
            window.location = "/"+ companyTeamName +"/group/add";
        });
        $('.table-wrapper .dataTables_filter').append(item);*/
    }
    
 /*---------------------------Initial data table calling---------------------------------------------------*/
    if(vm.Values != null) {
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate(); 
 /*--------------------------Ending Initial data table calling---------------------------------------------*/


/* Edit group details when click on edit icon*/
    /*$('#customerManagement-table tbody').on( 'click', '#edit', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[3];
        window.location ='/' + companyTeamName + '/group/' + key + '/edit';
    });
*/
/*Delete group details when click on delete icon*/
    $('#customerManagement-table tbody').on( 'click', '#delete', function () {
        $("#myModal").modal();
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[7];
        console.log(data, key);
        $("#confirm").click(function(){
            $.ajax({
                type: "POST",
                url: '/customer-management/'+ key + '/delete',
                data: '',
                success: function(data){
                    if(data=="true"){
                        $('#customerManagement-table').dataTable().fnDestroy();
                        var index = "";
                        for(var i = 0; i < mainArray.length; i++) {
                            index = mainArray[i].indexOf(key);
                            if(index != -1) {
                                console.log("dddd", i);
                                break;
                            }
                        }
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
