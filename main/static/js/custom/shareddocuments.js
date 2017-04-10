console.log(vm);


$(function(){
    var unixFromDate = 0;
    var unixToDate = 0;
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
    
    
    /*$('#fromDate-expiry').datepicker({
        
        minDate: new Date(2017, 1 -0, 25),
        maxDate: '+30Y',
        inline: true
    });*/
    
    function listSharedDocumentByDate(unixFromDate,unixToDate){
        var tempArray = [];
        var expiryDate =0;
        var unixExpiryDate = 0;
        for (i =0;i<vm.Values.length;i++){
            expiryDate = new Date(vm.Values[i][1]);
            unixExpiryDate = Date.parse(expiryDate)/1000;
            console.log(unixFromDate);
            console.log(unixExpiryDate);
            console.log(unixToDate);
            if(unixFromDate <= unixExpiryDate && unixToDate == 0){
                console.log("main",mainArray[i]);
                tempArray.push(mainArray[i]);
            } else if(unixFromDate ==0 && unixToDate >=unixExpiryDate){
                console.log("to");
                tempArray.push(mainArray[i]);
            }else if(unixFromDate <= unixExpiryDate && unixToDate >=unixExpiryDate ){
                console.log("both");
                tempArray.push(mainArray[i]);
            }else if(unixToDate == 0 && unixFromDate ==0) {
                dataTableManipulate(mainArray);
            }
        }
            $('#shareddocument-table').dataTable().fnDestroy();
            dataTableManipulate(tempArray);
            
        
    }
    
    
   function dataTableManipulate(dataArray){
       table =  $("#shareddocument-table").DataTable({
           data: dataArray,
           "columnDefs": [{
               "targets": -1,
               "width": "3%",
               "data": null,
               "defaultContent": '<span class="dwnl-btn"><i class="fa fa-download fa-lg" aria-hidden="true" id="view"></i></span>'
           }],
           "searching": false,
           "paging": true,
           "info": false,
           "lengthChange":false
       });
       $('#tbl_details_length').after($('.datepic-top'));
   }

/*----------------------------------Initialize Datatable--------------------------------------------------*/
    if(vm.Values != null) {
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate(mainArray);

/*--------------------------------Download-------------------------------------------------------------*/

    $('#shareddocument-table tbody').on( 'click', '#view', function () {
        var data = table.row( $(this).parents('tr') ).data();
        window.location =   data[2];
        return false;
    });
/*-------------------------------------------------------------------------------------------------------------------*/

    $('#fromDate').change(function () {
        var selectFromDate;
        var actualFromDate;
        selectFromDate = $('#fromDate').val();
        console.log("onclick",selectFromDate)
        actualFromDate = new Date(selectFromDate);
        actualFromDate.setHours(0);
        actualFromDate.setMinutes(0);
        actualFromDate.setSeconds(0);
        unixFromDate = Date.parse(actualFromDate)/1000;
        listSharedDocumentByDate(unixFromDate,unixToDate);        
       /* if(data[1]>=selectFromDate && data[1]<=selectedToDate){*/
             /*var tempArray = [];
        for(i = 0; i < mainArray.length; i++){
             var data = table.row( $(this).parents('tr') ).data();
             if(mainArray[i][1].indexOf(vm.ExpirationDate)>=selectFromDate){
                 alert("haiai");
                tempArray.push(mainArray[i]);
            }
            $('#shareddocument-table').dataTable().fnDestroy();
            dataTableManipulate(tempArray);   
            
        }*/
             
       /* }*/

    });
    

    $('#toDate').change(function () {
        var selectedToDate;
        var actualToDate;
        selectedToDate = $('#toDate').val();
        actualToDate = new Date(selectedToDate);
        actualToDate.setHours(23);
        actualToDate.setMinutes(59);
        actualToDate.setSeconds(59);
        unixToDate = Date.parse(actualToDate)/1000;
        console.log(unixToDate);
        listSharedDocumentByDate(unixFromDate,unixToDate);   
    });
    
    
});

