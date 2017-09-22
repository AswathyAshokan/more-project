
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
    
    
     document.getElementById("username").textContent=vm.AdminFirstName;
    document.getElementById("imageId").src=vm.ProfilePicture;
    if (vm.ProfilePicture ==""){
        document.getElementById("imageId").src="/static/images/default.png"
    }
if (vm.NotificationNumber !=0){
    document.getElementById("number").textContent=vm.NotificationArray.length;
}else{
    document.getElementById("number").textContent="";
}
   myNotification= function () {
        console.log("hiiii");
        document.getElementById("notificationDiv").innerHTML = "";
        var DynamicTaskListing="";
        if (vm.NotificationArray !=null){
            DynamicTaskListing ="<h5>"+"Notifications"+''<input type="button" id="notoficationClear">+"</h5>"+"<ul>";
        for(var i=0;i<vm.NotificationArray.length;i++){
            console.log("sp1");
            var timeDifference =moment(new Date(new Date(vm.NotificationArray[i][6]*1000)), "YYYYMMDD").fromNow();
            DynamicTaskListing += "<li>"+"User"+" "+vm.NotificationArray[i][2]+" "+vm.NotificationArray[i][3]+"  "+"delay to reach location"+" "+vm.NotificationArray[i][4]+" "+"for task"+" "+vm.NotificationArray[i][5]+" <span>"+timeDifference+"</span>"+"</li>";
            
            
        }
            $("#notificationDiv").prepend(DynamicTaskListing+"</ul>");
            document.getElementById("number").textContent="";
            $.ajax({
                url:'/'+ companyTeamName + '/notification/update',
                type: 'post',
                success : function(response) {
                    if (response == "true" ) {
                    } else {
                    }
                },
                error: function (request,status, error) {
                    console.log(error);
                }
            }); 
        }else{
            DynamicTaskListing ="<h5>"+" No New Notifications"+"</h5>";
            $("#notificationDiv").prepend(DynamicTaskListing);
            
        }
        
        }
     
     
     
     clearNotification= function () {
          document.getElementById("notificationDiv").innerHTML = "";
          $.ajax({
                url:'/'+ companyTeamName + '/notification/delete',
                type: 'post',
                success : function(response) {
                    if (response == "true" ) {
                    } else {
                    }
                },
                error: function (request,status, error) {
                    console.log(error);
                }
            }); 
         
         
         
     }
    
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
    
    
    
    
    //notification
      
    
    
    
    function listLogDetails(unixFromDate,unixToDate){
        var tempArray = [];
        var startDate =0;
        var unixStartDate = 0;
        for (i =0;i<vm.Values.length;i++){
            startDate = new Date(vm.Values[i][2]);
            console.log("date from date",vm.Values[i][2]);
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
        $('#log-details').on( 'click', '#btnShow', function () {
            var data = table.row( $(this).parents('tr') ).data();
              lattitude = data[5];
              longitude = data[6];
              console.log("data",data);
              $("#myModal").modal();
              $('#myModal').on('shown.bs.modal', function(){
                  console.log("lattitude  in general log:",lattitude);
                  console.log("longitude n general log:",longitude);
                  var googleLocation = new google.maps.LatLng(lattitude, longitude);
                  console.log("location",googleLocation);
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
                 var geocoder = new google.maps.Geocoder();
                  var latLng = new google.maps.LatLng(lattitude,longitude);
                  geocoder.geocode({       
                      latLng: latLng 
                  }, function(responses) { 
                      if (responses && responses.length > 0) {
                          $('#position').text(responses[0].formatted_address)
                      }
                  });
              });
        });
        
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
        
        
        
        
    });
    
    $("#activitylog").on( 'click', function () {
        console.log("generalMainArray",generalMainArray);
         tableData = generalMainArray
        $('#log-details').dataTable().fnDestroy();
        generaldataTableManipulate(tableData);
        $('#log-details').on( 'click', '#btnShow', function () {
              var data = table.row( $(this).parents('tr') ).data();
              lattitude = data[4];
              longitude = data[5];
              console.log("data",data);
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
                  var geocoder = new google.maps.Geocoder();
                  var latLng = new google.maps.LatLng(lattitude,longitude);
                  geocoder.geocode({  
                      latLng: latLng
                  }, function(responses) { 
                      if (responses && responses.length > 0) {
                          $('#position').text(responses[0].formatted_address)
                      }
                  });
              });
        });
        if(vm.GeneralLogValues != null) {
        console.log("vm.GeneralLogValues ########",vm.GeneralLogValues);
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
            /*lattitude = vm.GeneralLogValues[i][3];
            longitude= vm.GeneralLogValues[i][4];*/
            vm.GeneralLogValues[i][2]= localDate;
            vm.GeneralLogValues[i][3] = localTime;
            
            /*console.log("lattitude  in general log:",lattitude);
            console.log("longitude n general log:",longitude);*/
            
        }
        createGeneralDataArray(vm.GeneralLogValues, vm.GeneralKey);
    }
    });
    
    if( vm.WorkLocationValues !=null){
        for( i=0;i<vm.WorkLocationValues.length;i++){
            var utcTime = vm.WorkLocationValues[i][3];
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
            vm.WorkLocationValues[i][2]= localDate;
            vm.WorkLocationValues[i][3] = localTime;
        }
        createGeneralDataArray(vm.WorkLocationValues, vm.WorkLocationValues);
        
    }
    if(vm.GeneralLogValues != null) {
        console.log("vm.GeneralLogValues ########",vm.GeneralLogValues);
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
            /*lattitude = vm.GeneralLogValues[i][3];
            longitude= vm.GeneralLogValues[i][4];*/
            vm.GeneralLogValues[i][2]= localDate;
            vm.GeneralLogValues[i][3] = localTime;
            
            /*console.log("lattitude  in general log:",lattitude);
            console.log("longitude n general log:",longitude);*/
            
        }
        createGeneralDataArray(vm.GeneralLogValues, vm.GeneralKey);
    }
    //generaldataTableManipulate(generalMainArray);
    
    

    
  $('#log-details').on( 'click', '#btnShow', function () {
      var data = table.row( $(this).parents('tr') ).data();
      lattitude = data[5];
      longitude = data[6];
      console.log("data",data);
      $("#myModal").modal();
      $('#myModal').on('shown.bs.modal', function(){
          console.log("lattitude  in general log:",lattitude);
          console.log("longitude n general log:",longitude);
          var googleLocation = new google.maps.LatLng(lattitude, longitude);
          console.log("location",googleLocation);
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
          var geocoder = new google.maps.Geocoder();
          var latLng = new google.maps.LatLng(lattitude,longitude);
          geocoder.geocode({       
          latLng: latLng    
          }, function(responses) { 
              if (responses && responses.length > 0) { 
              $('#position').text(responses[0].formatted_address)
              }
          });
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
            var localTime = (HH + ':' + min);
            var localDate = (mm + '/' + dd + '/' + yyyy);
            vm.Values[i][2]= localDate;
            vm.Values[i][3] = localTime;
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
        var workLocation = document.getElementById("WorkLocation")
        parent.removeChild(workLocation)
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
            var workLocation = document.getElementById("WorkLocation")
            parent.removeChild(workLocation)
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
    
    $("#cancel").click(function() {
            window.location = '/'+companyTeamName+'/workLog';
    });
    
});





