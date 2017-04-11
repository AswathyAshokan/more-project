console.log(vm);


$(function(){
    var unixFromDate = 0;
    var unixToDate = 0;
    var mainArray = [];   
    var table = "";
    var selectedToDate;
    var actualToDate;
    var selectFromDate;
    var actualFromDate;
    var completeTable =[];
    
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
   
    completeTable = mainArray;
    $('#refreshButton').click(function(e) {
        $('#shareddocument-table').dataTable().fnDestroy();
        $('#fromDate').datepicker('setDate', null);
        $('#toDate').datepicker('setDate', null);
        dataTableManipulate(completeTable);
     });
    /*$('#fromDate').datepicker({
        
        minDate: new Date(2017, 4, 25),
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
            if(unixFromDate <= unixExpiryDate && unixToDate == 0){
                tempArray.push(mainArray[i]);
            
            } else if(unixFromDate ==0 && unixToDate >=unixExpiryDate){
                tempArray.push(mainArray[i]);
            
            }else if(unixFromDate <= unixExpiryDate && unixToDate >=unixExpiryDate ){
                tempArray.push(mainArray[i]);
            
            }
            
            $('#shareddocument-table').dataTable().fnDestroy();
            dataTableManipulate(tempArray);
        }
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
        
       
      
        selectFromDate = $('#fromDate').val();
         /*$('#toDate').datepicker({
            minDate: new Date(selectedToDate),
            maxDate: '+30Y',
            inline: true
        });*/
        actualFromDate = new Date(selectFromDate);
        actualFromDate.setHours(0);
        actualFromDate.setMinutes(0);
        actualFromDate.setSeconds(0);
        unixFromDate = Date.parse(actualFromDate)/1000;
        listSharedDocumentByDate(unixFromDate,unixToDate);
    });
    

    $('#toDate').change(function () {
        var output = selectFromDate.replace(/(\d\d)\/(\d\d)\/(\d{4})/, "$3-$1-$2");
        console.log("outt",output);
        $('#fromDate').datepicker({
            minDate:  output,
            maxDate: '+30Y',
            inline: true
        });
        selectedToDate = $('#toDate').val();
        actualToDate = new Date(selectedToDate);
        actualToDate.setHours(23);
        actualToDate.setMinutes(59);
        actualToDate.setSeconds(59);
        unixToDate = Date.parse(actualToDate)/1000;
        listSharedDocumentByDate(unixFromDate,unixToDate);
    });
    
});

