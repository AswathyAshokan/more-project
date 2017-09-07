document.getElementById("crm").className += " active";

var companyTeamName = vm.CompanyTeamName;

/*Function for creating Data Array for data table*/
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
        table =  $("#workLocation-table").DataTable({
            data: mainArray,
            "columnDefs": [{
                "targets": -1,
                "width": "10%",
                "data": null,
                "defaultContent": '<div class="edit-wrapper"><span class="icn"><i class="fa fa-pencil-square-o" aria-hidden="true" id="edit"></i><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
            }]
        });
        
/*Add a plus symbol in webpage for add new groups*/
        var item = $('<span>+</span>');
        item.click(function() {
            window.location ='/' + companyTeamName + '/worklocation/add';
        });
        
        $('.table-wrapper .dataTables_filter').append(item);
    }
/*---------------------------Initial data table calling---------------------------------------------------*/
    var tempArry = [];
    if(vm.Values != null) {
        for( i=0;i<vm.Values.length;i++){
            for( j=0;j<vm.Users.length;j++){
                for(k=0;k<vm.Users[j].length;k++){
                    console.log("vm.Users[j][k].Name",vm.Users[j][k].Name);
                    if(vm.Values[i][1] == vm.Users[j][k].UserKey){
                        if(vm.Users[j][k].Name != null){
                            
                            console.log("kkk",vm.Values[j][0])
                            vm.Values[i][0] = vm.Values[j][0];
                        }
                        tempArry.push(vm.Users[j][k].Name);
                    }
                }
            }
            console.log("tempArry");
            vm.Values[i][1] = tempArry;
            tempArry = [];
        }
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate(); 
 /*--------------------------Ending Initial data table calling---------------------------------------------*/


    /*Edit customer details when click on edit icon*/
    $('#workLocation-table tbody').on( 'click', '#edit', function () {
        var data = table.row( $(this).parents('tr') ).data();
        console.log("data",data)
        var workLocationId = data[2];
        window.location = '/' + companyTeamName +'/worklocation/'+ workLocationId + '/edit';
        return false;
    });
    
    $('#workLocation-table tbody').on( 'click', '#delete', function () {
         $("#myGroupModal").modal();
        var data = table.row( $(this).parents('tr') ).data();
        console.log("data",data[2]);
        var workLocationId = data[2];
        $("#confirm").click(function(){
            $.ajax({
                type: "POST",
                url: '/' + companyTeamName +'/worklocation/'+ workLocationId + '/delete',
                data:'',
                success: function(response){
                    if(response=="true"){
                        $('#workLocation-table').dataTable().fnDestroy();
                        var index = "";
                        
                        for(var i = 0; i < mainArray.length; i++) {
                           index = mainArray[i].indexOf(workLocationId);
                           if(index != -1) {
                               console.log("dddd", i);
                             break;
                           }
                        }
                        mainArray.splice(i, 1);
                        dataTableManipulate() 
                    }
                    else {
                        console.log("Removing Failed!");
                    }
                }

            });
        });
    });
});

