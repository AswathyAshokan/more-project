/*Author: Sarath
Date:01/02/2017*/
//Below line is for adding active class to layout side menu..
document.getElementById("nfc").className += " active";
var companyTeamName = vm.CompanyTeamName
//Fetching Key,Values from Database and Pushinng it into Main Array of Sub Arrays
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
    
    //Generate Datatabe from Main Array
    function dataTableManipulate(){
        table =  $("#nfc_details").DataTable({
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
            window.location = "/"+companyTeamName+"/nfc/add";
        });
        $('.table-wrapper .dataTables_filter').append(item);
    }
    if(vm.Values != null) {
        createDataArray(vm.Values, vm.Keys);
        console.log(vm.Values);
        console.log(vm.Keys);
    }
    dataTableManipulate(); 
    //Edit selected NFC Tag
    $('#nfc_details tbody').on( 'click', '#edit', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[4];
        //alert(data[4]);
        window.location = '/'+ companyTeamName +'/nfc/' + key + '/edit';
    });

    //Delete selcted NFC Tag from Datatable and Database
    $('#nfc_details tbody').on( 'click', '#delete', function () { 
        $("#myModal").modal();
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[4];
        console.log(data, key);
        $("#confirm").click(function(){
            $.ajax({
                type: "POST",
                url: "/" + companyTeamName  + "/nfc/"+data[4]+"/delete",
                data: {
                    Key : key
                },
                success: function(data){
                    if(data=="true"){
                        $('#nfc_details').dataTable().fnDestroy();
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


