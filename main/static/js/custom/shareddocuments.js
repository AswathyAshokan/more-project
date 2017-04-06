console.log(vm);


$(function(){
    var selectFromDate;
        var selectedToDate;
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
    /*$('#dp1').datepicker();*/
    
    $('#fromDate-expiry').val('00-00-0000');
    $('#toDate-expiry').val('00-00-0000');
    $('#fromDate-expiry').change(function () {
        
        selectFromDate = $('#fromDate-expiry').val();
      
       
        alert(selectFromDate);
       
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
    
    
    $('#toDate-expiry').change(function () {
        selectedToDate = $('#toDate-expiry').val();
        alert(selectedToDate);
    });
    
    
    if(selectFromDate != null || selectedToDate != null){
         alert(selectFromDate);
    }
    
    
    
});

