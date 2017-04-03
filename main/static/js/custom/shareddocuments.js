

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
        table =  $("#shareddocument-table").DataTable({
            data: mainArray,
            "columnDefs": [{
                "targets": -1,
                "width": "10%",
                "data": null,
                "defaultContent": '<span class="dwnl-btn" ><i class="fa fa-download fa-lg"  aria-hidden="true"></i></span>'
            }],
            "searching": false,
            "paging": true, 
            "info": false,
            "lengthChange":false
        });
        $('#tbl_details_length').after($('.datepic-top'));
    }
    if(vm.Values != null) {
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate();
});

