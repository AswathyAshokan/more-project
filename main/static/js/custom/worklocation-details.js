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
                "defaultContent": '<div class="edit-wrapper"><span class="icn"><i class="fa fa-eye" aria-hidden="true" id="view"></i><i class="fa fa-pencil-square-o" aria-hidden="true" id="edit"></i><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
            }]
        });
        
/*Add a plus symbol in webpage for add new groups*/
        var item = $('<span>+</span>');
        item.click(function() {
            console.log("temname",companyTeamName);
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
                    console.log("vm.Values[i][1]",vm.Values[i][1]);
                    if(vm.Values[i][1] == vm.Users[j][k].UserKey){
                        tempArry.push(vm.Users[j][k].Name);
                        vm.Values[i][0] = vm.Values[j][0];
                    }
                }
            }
            console.log("tempArry",tempArry);
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
        var workLocationId = data[2];
        window.location = '/' + companyTeamName +'/worklocation/'+ workLocationId + '/edit';
        return false;
    });
    
    $('#customer-table tbody').on( 'click', '#delete', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[8];
        $.ajax({
            type: "POST",
            url: '/' + companyTeamName +'/customer/'+ key + '/deletionOfCustomerIfUsedForJob',
            data: '',
            success: function(response){
                console.log("dhfg",response)
                
                if(response=="true"){
                   
                     $("#myCustomerJobModel").modal();
                    $("#deleteNotJob").click(function(){
                         $.ajax({
                            type: "POST",
                            url: '/' + companyTeamName +'/customer/'+ key + '/deletionOfCustomerFromJob',
                             data: '',
                            success: function(feedback){
                                console.log(feedback);
                                if(feedback=="true"){
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
                                dataTableManipulate()
                                }
                                else {
                                }
                            }
                         });
                    });
                } else {
                    console.log("inside else part");
                   
                    $("#myModal").modal();
                    $("#confirm").click(function(){
                        $.ajax({
                            type: "POST",
                            url: '/' + companyTeamName +'/customer/'+ key + '/RemoveTask',
                            data: '',
                            success: function(response){
                                if(response=="true"){
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
                                dataTableManipulate()
                                }
                                else {
                                }
                            }
                        });
                    });
                }
            }
        });
    });
});

