
//Below line is for adding active class to layout side menu..
document.getElementById("group").className += " active";
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
        table =  $("#group-table").DataTable({
            data: mainArray,
            "columnDefs": [
                {
                    "targets": 2,
                    render : function(data, type, row) {
                        return '<div class="over-length">'+data+'</div>'
                    } 
                },
                {
                    "targets": -1,
                    "width": "5%",
                    "data": null,
                    "defaultContent": '<div class="edit-wrapper"><span class="icn"><i class="fa fa-pencil-square-o" aria-hidden="true" id="edit"></i><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
                }]
        });
/*Add a plus symbol in webpage for add new groups*/
        var item = $('<span>+</span>');
        item.click(function() {
            window.location = "/"+ companyTeamName +"/group/add";
        });
        $('.table-wrapper .dataTables_filter').append(item);
    }
    
 /*---------------------------Initial data table calling---------------------------------------------------*/
    if(vm.Values != null) {
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate(); 
 /*--------------------------Ending Initial data table calling---------------------------------------------*/


/* Edit group details when click on edit icon*/
    $('#group-table tbody').on( 'click', '#edit', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[3];
        window.location ='/' + companyTeamName + '/group/' + key + '/edit';
    });

/*Delete group details when click on delete icon*/
//    $('#group-table tbody').on( 'click', '#delete', function () {
//        $("#myModal").modal();
//        var data = table.row( $(this).parents('tr') ).data();
//        var key = data[3];
//        console.log(data, key);
//        $("#confirm").click(function(){
//            $.ajax({
//                type: "POST",
//                url: '/' + companyTeamName +'/group/'+ key + '/delete',
//                data: '',
//                success: function(data){
//                    if(data=="true"){
//                        $('#group-table').dataTable().fnDestroy();
//                        var index = "";
//                        for(var i = 0; i < mainArray.length; i++) {
//                            index = mainArray[i].indexOf(key);
//                            if(index != -1) {
//                                console.log("dddd", i);
//                                break;
//                            }
//                        }
//                        mainArray.splice(i, 1);
//                        console.log(mainArray);
//                        dataTableManipulate(); 
//                    }
//                    else {
//                        console.log("Removing Failed!");
//                    }
//                }
//            });
//        });
//    });
    $('#group-table tbody').on( 'click', '#delete', function () {
      
        var data = table.row( $(this).parents('tr') ).data();
       var key = data[3];
        $.ajax({
            type: "POST",
            url: '/' + companyTeamName +'/group/'+ key + '/delete',
            data: '',
            success: function(data){
                console.log("jjjj",data);
                if(data=="true"){
                    console.log("hdhhshhh");
                    $("#groupInTask").modal();
                    $("#deleteNotTask").click(function(){
                        $.ajax({
                            type: "POST",
                            url: '/' + companyTeamName +'/group/'+ key + '/deletionOfGroup',
                            data: '',
                            success: function(feedback){
                                console.log(feedback);
                                if(feedback=="true"){
                                    $('#group-table').dataTable().fnDestroy(); 
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
                else {
                   
                    $("#myGroupModal").modal();
                    $("#confirm").click(function(){
                        $.ajax({
                            type: "POST",
                            url: '/' + companyTeamName +'/group/'+ key + '/RemoveTask',
                            data: '',
                            success: function(response){
                                if(response=="true"){
                                    $('#group-table').dataTable().fnDestroy(); 
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
