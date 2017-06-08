
/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..
//document.getElementById("log").className += " active";

$(function(){ 
    var mainArray = [];   
    var table = "";
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
        console.log("inside create");
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
        $('#log-details').dataTable().fnDestroy();
        $('#fromDate').datepicker('setDate', null);
        $('#toDate').datepicker('setDate', null);
        dataTableManipulate(completeTable);
     });
    function listLogDetails(unixFromDate,unixToDate){
        var tempArray = [];
        var startDate =0;
        var unixStartDate = 0;
        for (i =0;i<vm.Values.length;i++){
            startDate = new Date(vm.Values[i][5]);
            unixStartDate = Date.parse(startDate)/1000;
           if( (unixFromDate <= unixStartDate && unixStartDate <= unixToDate) || (unixFromDate <= unixStartDate && unixStartDate <= unixToDate) || (unixFromDate >= startDate && unixStartDate >= unixToDate)) {
               
                tempArray.push(mainArray[i]);
           }
            
            $('#log-details').dataTable().fnDestroy();
            dataTableManipulate(tempArray);
        }
    } 
    
    
    function dataTableManipulate(mainArray){
        table =  $("#log-details").DataTable({
            data: mainArray,
            "searching": false,
            "info": false,
            "lengthChange":false
            
            
        });
        $('#tbl_details_length').after($('.datepic-top'));
        
    }
    if(vm.Values != null) {
        console.log("inside if");
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate(mainArray);   
    
    
    
    $('#fromDate').change(function () {
        selectFromDate = $('#fromDate').val();
        var fromYear = selectFromDate.substring(6, 10);
        var fromDay = selectFromDate.substring(3, 5);
        var fromMonth = selectFromDate.substring(0, 2);
        $('#toDate').datepicker("option", "minDate", new Date(fromYear, fromMonth-1, fromDay));
        actualFromDate = new Date(selectFromDate);
        actualFromDate.setHours(0);
        actualFromDate.setMinutes(0);
        actualFromDate.setSeconds(0);
        unixFromDate = Date.parse(actualFromDate)/1000;
        listLogDetails(unixFromDate,unixToDate);
    });
    
    $('#toDate').change(function () {
        selectedToDate = $('#toDate').val();
        var year = selectedToDate.substring(6, 10);
        var day = selectedToDate.substring(3, 5);
        var month = selectedToDate.substring(0, 2);
        $('#fromDate').datepicker("option", "maxDate", new Date(year, month-1, day));
        actualToDate = new Date(selectedToDate);
        actualToDate.setHours(23);
        actualToDate.setMinutes(59);
        actualToDate.setSeconds(59);
        unixToDate = Date.parse(actualToDate)/1000;
        listLogDetails(unixFromDate,unixToDate);
    });
    
    
    if(vm.CompanyPlan == 'family'){
       
        $('#group').bind('click', false);
        $('#contact').bind('click', false);
        $('#job').bind('click', false);
        $('#nfc').bind('click', false);
        $('#crm').bind('click', false);
        
    }
    
     if(vm.CompanyPlan == 'campus'){
         
       /* $('#group').attr('disabled', true);*/
        $('#group').bind('click', false);
        $('#contact').bind('click', false);
        $('#job').bind('click', false);
        $('#nfc').bind('click', false);
        $('#crm').bind('click', false);
     }
    
    
    if(vm.CompanyPlan == 'business'){
         
       /* $('#group').attr('disabled', true);*/
        $('#group').bind('click', false);
        $('#contact').bind('click', false);
        $('#job').bind('click', false);
        $('#nfc').bind('click', false);
    }
    
    document.getElementById("username").textContent=vm.AdminFirstName;
    document.getElementById("imageId").src=vm.ProfilePicture;
    if (vm.ProfilePicture ==""){
        document.getElementById("imageId").src="https://firebasestorage.googleapis.com/v0/b/passporte-b9070.appspot.com/o/profilePicturesOfAdmin%2Fdefault.png?alt=media&token=7444c8f3-2254-4494-9588-a41ecee96b01"
    }
    if(vm.CompanyPlan == "family")
    {
        $('#planChange').attr('data-target','#family');
    } else if (vm.CompanyPlan == "campus") {
        $('#planChange').attr('data-target','#campus');
    }else if (vm.CompanyPlan == "business") {
        $('#planChange').attr('data-target','#business');
    }else if (vm.CompanyPlan == "businessPlus") {
        $('#planChange').attr('data-target','#business-plus');
    }
    
});


