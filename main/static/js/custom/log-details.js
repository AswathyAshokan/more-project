
//Below line is for adding active class to layout side menu..
//document.getElementById("log").className += " active";
//console.log(vm.Values[1][5]);
console.log(vm);
$(function(){
    var mainArray = [];   
    var generalMainArray = [];
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
    var tableData = [];
    var lattitude;
    var longitude;
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
    // fuction for work log
    function createGeneralDataArray(values, keys){
        console.log("inside create");
        var subArray = [];
        for(i = 0; i < values.length; i++) {
            for(var propertyName in values[i]) {
                subArray.push(values[i][propertyName]);
            }
            subArray.push(keys[i])
            generalMainArray.push(subArray);
            subArray = [];
            
        }
    }
    
    function dataTableManipulate(mainArray){
        table =  $("#log-details").DataTable({
            data: mainArray,
            "searching": true,
            "info": false,
            "lengthChange":false,
           "columnDefs": [{
               "targets": [5],
                render : function (data, type, row) {
                       return '<button class="btn btn-primary btn-xs " id = "btnShow">Show Map</button>';
                }
           }]
        });
        $('#tbl_details_length').after($('.datepic-top'));
    }
    
    
    function generaldataTableManipulate(mainArray){
        table =  $("#log-details").DataTable({
            data: mainArray,
            "searching": true,
            "info": false,
            "lengthChange":false,
           "columnDefs": [{
                'visible': false, 'targets': [4] },
               
               {"targets": [5],
                render : function (data, type, row) {
                       return '<button class="btn btn-primary btn-xs " id = "btnShow">Show Map</button>';
                }
           }]
        });
        $('#tbl_details_length').after($('.datepic-top'));
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
            startDate = new Date(vm.Values[i][7]);
            unixStartDate = Date.parse(startDate)/1000;
           if( (unixFromDate <= unixStartDate && unixStartDate <= unixToDate) || (unixFromDate <= unixStartDate && unixStartDate <= unixToDate) || (unixFromDate >= startDate && unixStartDate >= unixToDate)) {
               
                tempArray.push(mainArray[i]);
           }
            $('#log-details').dataTable().fnDestroy();
            dataTableManipulate(tempArray);
        }
    } 
    
    
    $("#userlog").on('click',function(){
        completeTable = mainArray;
        $('#log-details').dataTable().fnDestroy();
       dataTableManipulate(completeTable);
    });
    
    $("#activitylog").on( 'click', function () {
         tableData = generalMainArray
        $('#log-details').dataTable().fnDestroy();
        generaldataTableManipulate(tableData);
    });
    
    
    if(vm.GeneralLogValues != null) {
        //$('#log-details').dataTable().fnDestroy();
        for( i=0;i<vm.GeneralLogValues.length;i++){
            console.log("hghghghg",vm.GeneralLogValues[i]);
            var utcTime = vm.GeneralLogValues[i][2];
            var dateFromDb = parseInt(utcTime)
            var d = new Date(dateFromDb * 1000);
            var dd = d.getDate();
            var mm = d.getMonth() + 1; //January is 0!
            var yyyy = d.getFullYear();
            var HH = d.getHours();
            var min = d.getMinutes();
            var sec = d.getSeconds();
            if (dd < 10) {
                dd = '0' + dd;
            }
            if (mm < 10) {
                mm = '0' + mm;
            }
            if (HH < 10) {
                HH = '0' + HH;
            }
            if (min < 10) {
                min = '0' + min;
            }
            if (sec < 10) {
                sec = '0' + sec;
            }
            var localTime = (HH + ':' + min);
            var localDate = (mm + '/' + dd + '/' + yyyy);
            vm.GeneralLogValues[i][2]= localDate;
            vm.GeneralLogValues[i][3] = localTime;
            lattitude = vm.GeneralLogValues[i][3];
            longitude= vm.GeneralLogValues[i][4];
             
            
        }
        createGeneralDataArray(vm.GeneralLogValues, vm.GeneralKey);
    }
    //generaldataTableManipulate(generalMainArray);
    
    

    
  $('#log-details').on( 'click', '#btnShow', function () {
      $("#myModal").modal();
      $('#myModal').on('shown.bs.modal', function(){
          var googleLocation = new google.maps.LatLng(lattitude, longitude);
          var mapOptions = {
          center: googleLocation,
          title: "Google Map",
          width: 50,
          height: 50,
          zoom: 15,
          mapTypeId: google.maps.MapTypeId.ROADMAP
          }
          var map = new google.maps.Map($("#dvMap")[0], mapOptions);
          var marker = new google.maps.Marker({
          position: googleLocation,
          });
          marker.setMap(map);
       });
  });
    
    if(vm.Values != null) {
        for( i=0;i<vm.Values.length;i++){
            var utcTime = vm.Values[i][3];
            //var datum = Date.parse(dateTime);
            var dateFromDb = parseInt(utcTime)
            var d = new Date(dateFromDb * 1000);
            var dd = d.getDate();
            var mm = d.getMonth() + 1; //January is 0!
            var yyyy = d.getFullYear();
            var HH = d.getHours();
            var min = d.getMinutes();
            var sec = d.getSeconds();
            if (dd < 10) {
                dd = '0' + dd;
            }
            if (mm < 10) {
                mm = '0' + mm;
            }
            if (HH < 10) {
                HH = '0' + HH;
            }
            if (min < 10) {
                min = '0' + min;
            }
            if (sec < 10) {
                sec = '0' + sec;
            }
            /*var utcInDateForm = new Date(dateFromDb);
            var datum = Date.parse(utcInDateForm);*/
            var localTime = (HH + ':' + min);
            var localDate = (mm + '/' + dd + '/' + yyyy);
            //var d = localDate.slice(0, 10).split('/');
            //var formatedDate = d[1] +'/'+ d[0] +'/'+ d[2];
           // console.log("formatedDate",formatedDate);
            /*var timeSplitArray = localTime.split(":");
            var hours = timeSplitArray[0];
            var minutes = timeSplitArray[1];
            var duration = vm.Values[i][2];
            var durationSplitArray = duration.split(":");
            var duartionHours = durationSplitArray[0];
            var durationMinutes = durationSplitArray[1];
            var localTimeInMinutes = parseFloat(hours)*60 + parseFloat(minutes);
            if (localTimeInMinutes>durationMinutes){
                var loggedTime = localTimeInMinutes - durationMinutes
                var loggedHours = window.parseInt(loggedTime/60);
                var loggedMins = loggedTime%60;
            }
            var actualloggedTime =loggedHours +   ":" +loggedMins
            var between = actualloggedTime + " &nbspto&nbsp" +hours +    ":"    +minutes;*/
            vm.Values[i][2]= localDate;
            vm.Values[i][3] = localTime;
            lattitude = vm.Values[i][5];
            longitude= vm.Values[i][6];
             //$('#log-details').dataTable().fnDestroy();
            
        }
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
    
    //checking plans
    
    if(vm.CompanyPlan == 'family' ){
        var parent = document.getElementById("menuItems");
        var contact = document.getElementById("contact");
        var job = document.getElementById("job");
        var crm = document.getElementById("crm");
        var leave = document.getElementById("leave");
        var timesheet  = document.getElementById("time-sheet");
        var consent = document.getElementById("consent")
        parent.removeChild(timesheet);
        parent.removeChild(consent);
        parent.removeChild(leave);
        parent.removeChild(contact);
        parent.removeChild(job);
        parent.removeChild(crm);
        
    } else if(vm.CompanyPlan == 'campus'){
            var parent = document.getElementById("menuItems");
            var contact = document.getElementById("contact");
            var job = document.getElementById("job");
            var crm = document.getElementById("crm");
            var leave = document.getElementById("leave");
            var timesheet  = document.getElementById("time-sheet");
            var consent = document.getElementById("consent")
            parent.removeChild(timesheet);
            parent.removeChild(consent);
            parent.removeChild(leave);
            parent.removeChild(contact);
            parent.removeChild(job);
            parent.removeChild(crm);
     }
    document.getElementById("username").textContent=vm.AdminFirstName;
    document.getElementById("imageId").src=vm.ProfilePicture;
    if (vm.ProfilePicture ==""){
        document.getElementById("imageId").src="/static/images/default.png"
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
    
    $('#log-details').on( 'click', '#activitylog', function () {
    
     window.location='/' + vm.CompanyTeamName +'/activityworkLog';
    
});
    
});





