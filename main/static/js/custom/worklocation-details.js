document.getElementById("WorkLocation").className += " active";
console.log("uuuuuuuuuuuuuuu",vm.Values);
var companyTeamName = vm.CompanyTeamName;
var ExposureArray =vm.ExposureArray;


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
            "columnDefs": [
                { className: "details-control" , "targets": [0]},
                {
                    "order": [[1, 'asc']]
                },
                {
                    "targets": 2,
                    render : function(data, type, row) {
                        return '<div class="over-length">'+data+'</div>'
                    } 
                },
                
                {
                "targets": -1,
                "width": "10%",
                "data": null,
                "defaultContent": '<div class="edit-wrapper"><span class="icn"><i class="fa fa-pencil-square-o" aria-hidden="true" id="edit"></i><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
            }]
        });
        
        
        $('#workLocation-table tbody').on('click', 'td.details-control', function () {
        var tr = $(this).closest('tr');
        var row = table.row( tr );
 
        if ( row.child.isShown() ) {
            // This row is already open - close it
            row.child.hide();
            tr.removeClass('shown');
        }
        else {
            // Open this row
            row.child( format(row.data(),vm.MinUserAndLoginTypeArray) ).show();
            tr.addClass('shown');
        }
    } );
        
        
        function format (data,minUserArray ) {
            
             var workLocationID  = data[5];
            var minUser ='<div class="pull-left dropdown-tbl"  style="padding-right: 50px;">';
            minUser +="<table cellpadding='5' cellspacing='0' border='0' style=''>";
            minUser +='<tr>';
            for (var i=0; i<minUserArray.length; i++){
              
                if(minUserArray[i] != null && minUserArray[i][3] == workLocationID) {
                    minUser +='<tr>';
                    minUser +='<td>Login type </td>';
                    minUser +='<td>'+minUserArray[i][0]+'</td>';
                    minUser +='</tr>';
                    minUser +='<tr>';
                    minUser +='<td>Log Time In Minutes </td>';
                    minUser +='<td>'+minUserArray[i][1]+'</td>';
                    minUser +='</tr>';
                    minUser +='<tr>';
                    minUser +='<td>Fit To WorkName </td>';
                    minUser +='<td>'+minUserArray[i][2]+'</td>';
                    minUser +='</tr>';
                    
                }
            }
            minUser +="</table>";
            minUser +="</div>";
            minUser +="</table>";
            minUser +="</div>";
            
            //exposure
            var exposure   ='<div class="pull-left dropdown-tbl" style="margin-left: 25px;">';
            exposure += "<table cellpadding='5' cellspacing='0' style='border: 1px solid #dddddd !important;'>";
            exposure += '<tr><th>Exposure Details</th></tr>';
           
            for (var i = 0; i < ExposureArray.length; i++) {
                 
                if(ExposureArray[i] != null && ExposureArray[i][0].TaskId == workLocationID ) {
                    console.log("task id exposure",ExposureArray[i][0].TaskId );
                    for (var j=0; j<ExposureArray[i].length ;j++){
                        var Breakhours = Math.trunc(ExposureArray[i][j].BreakMinute/60);
                        var Breakminutes = ExposureArray[i][j].BreakMinute % 60;
                        var Workhours = Math.trunc(ExposureArray[i][j].WorkingHour/60);
                        var Workminutes = ExposureArray[i][j].WorkingHour % 60;
                         exposure += "<tr>";
                        exposure += "<td>"+Breakhours +":"+ Breakminutes+" Minutes Break After    "+Workhours +":"+ Workminutes+"Minutes"+"</td>";
                        exposure += "</tr>";
                    }
                }
            }
            exposure += "</table  >";
            exposure +="</div>";
            
            
            //fit To Work
            
            var fitToWork   ='<div class="pull-left dropdown-tbl" style="margin-left:25px;">';
            fitToWork += "<table cellpadding='5' cellspacing='0' style='border: 1px solid #dddddd !important;'>";
//            fitToWork += '<th>Fit To Work Details </th>';
             if ( vm.FitToWorkDetailsDisplayArray !=null)
                {
                    fitToWork += "<th>"+"Fit To Work  Detail"+"</th>";
                    fitToWork += '<th> </th>';
                    fitToWork += "<tr>";
                    fitToWork += '<th>Accepted User </th>';
                    fitToWork += '<th>Date And Time </th>';
                    fitToWork += "</tr>";
                    fitToWork += "<tr>";
           
                    for (var i = 0; i < vm.FitToWorkDetailsDisplayArray.length; i++) {
                 
                if(vm.FitToWorkDetailsDisplayArray[i] != null && vm.FitToWorkDetailsDisplayArray[i][1] == workLocationID ) {
                    console.log("task id exposure");
                    fitToWork += "<td>"+vm.FitToWorkDetailsDisplayArray[i][2]+"</td>";
                    var utcTime = vm.FitToWorkDetailsDisplayArray[i][3];
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
                    var fitTime = (HH + ':' + min);
                    var fitDate = (mm + '/' + dd + '/' + yyyy);
                     fitToWork += "<td>"+fitDate+" ("+fitTime+")"+"</td>";
                     fitToWork += "</tr>";
                }
            }
                }
            
            fitToWork += "</table  >";
            fitToWork +="</div>";
            return minUser+exposure+fitToWork;
         }
        
        
        
/*Add a plus symbol in webpage for add new groups*/
        var item = $('<span>+</span>');
        item.click(function() {
            window.location ='/' + companyTeamName + '/worklocation/add';
        });
        
        $('.table-wrapper .dataTables_filter').append(item);
    }
/*---------------------------Initial data table calling---------------------------------------------------*/
    var tempArry = [];
    if(vm.Values != null) {
        if(vm.Users !=null){
        for( i=0;i<vm.Values.length;i++){
            for( j=0;j<vm.Users.length;j++){
                if(vm.Users[j] !=null){
                for(k=0;k<vm.Users[j].length;k++){
                    console.log("keyyyy",vm.Values[i][4]);
                    if(vm.Values[i][4] == vm.Users[j][k].UserKey){
                        if(vm.Users[j][k].Name != null){
                            
                            console.log("kkk",vm.Values[j][0])
                            vm.Values[i][1] = vm.Values[i][1];
                        }
                        tempArry.push( " "+vm.Users[j][k].Name+" ");
                      
                    }
                }
                }
            }
            var startDateInUnix = vm.Values[i][2];
            var dateFromDb = parseInt(startDateInUnix)
            var d = new Date(dateFromDb * 1000);
            var dd = d.getDate();
            var mm = d.getMonth() + 1; //January is 0!
            var yyyy = d.getFullYear();
            if (dd < 10) {
                dd = '0' + dd;
            }
            if (mm < 10) {
                mm = '0' + mm;
            }

            var localstartDate = (mm + '/' + dd + '/' + yyyy);

            var endDateInUnix = vm.Values[i][3];
            var dateEndFromDb = parseInt(endDateInUnix)
            var d = new Date(dateEndFromDb * 1000);
            var dd = d.getDate();
            var mm = d.getMonth() + 1; //January is 0!
            var yyyy = d.getFullYear();
            if (dd < 10) {
                dd = '0' + dd;
            }
            if (mm < 10) {
                mm = '0' + mm;
            }
            var localEndDate = (mm + '/' + dd + '/' + yyyy);
            vm.Values[i][3] = localstartDate;
            vm.Values[i][4] = localEndDate;
            vm.Values[i][2] = tempArry;
            tempArry = [];
        }
        createDataArray(vm.Values, vm.Keys);
    }
    }
    dataTableManipulate(); 
 /*--------------------------Ending Initial data table calling---------------------------------------------*/


    /*Edit customer details when click on edit icon*/
    $('#workLocation-table tbody').on( 'click', '#edit', function () {
        var data = table.row( $(this).parents('tr') ).data();
        console.log("data",data)
        var workLocationId = data[5];
        window.location = '/' + companyTeamName +'/worklocation/'+ workLocationId + '/edit';
        return false;
    });
    
    $('#workLocation-table tbody').on( 'click', '#delete', function () {
         
        var data = table.row( $(this).parents('tr') ).data();
        console.log("full data",data);
        console.log("data id",data[5]);
        var workLocationId = data[5];
         $.ajax({
              type: "POST",
             url: '/' + companyTeamName +'/worklocation/'+ workLocationId + '/checkbeforedelete',
             data: '',
             success: function(data){
                 var jsonData = JSON.parse(data)
                 if(jsonData[0]=="true"){
                     
                     var endDateFromDb = jsonData[1];
                     console.log("json data",endDateFromDb)
                     var today = new Date();
                     var dd = today.getDate();
                     var mm = today.getMonth()+1; //January is 0!
                     var yyyy = today.getFullYear();
                     if(dd<10) {
                         dd = '0'+dd
                     } 
                     if(mm<10) {
                        mm = '0'+mm
                    }
                     var actualCurrentDate = (yyyy+'-'+ mm+'-'+dd);
                     var endDateInt = parseInt(endDateFromDb)
                     var endDateParse = new Date(endDateInt * 1000);
                     var ddfromDb = endDateParse.getDate();
                    var mmfromDb = endDateParse.getMonth() + 1; //January is 0!
                    var yyyyfromDb = endDateParse.getFullYear();
                    if (dd < 10) {
                         dd = '0' + dd;
                    }
                    if (mm < 10) {
                         mm = '0' + mm;
                     }
                     var actualDateFormDb = (yyyyfromDb+'-'+mmfromDb+'-'+ddfromDb);
                     console.log("actualDateFormDb",actualDateFormDb);
                     
                     var currentDt  = new Date(actualCurrentDate).setHours(0,0,0,0);
                     var dbDt  = new Date(actualDateFormDb).setHours(0,0,0,0);
                     console.log("first11",currentDt);
                     console.log("second 11",dbDt)
                     if (currentDt > dbDt){
                         console.log("inside if condition");
                         $("#myGroupModal").modal();
                         $("#confirm").click(function(){
                             $.ajax({
                                type: "POST",
                                url: '/' + companyTeamName +'/worklocation/'+ workLocationId + '/delete',
                                data:'',
                                success: function(response){
                                    if(response=="true"){
                                        $('#workLocation-table').dataTable().fnDestroy();
                                        var index = "";

                                        for(var i = 0; i < mainArray.length; i++) {
                                           index = mainArray[i].indexOf(workLocationId);
                                           if(index != -1) {
                                               console.log("dddd", i);
                                             break;
                                           }
                                        }
                                        mainArray.splice(i, 1);
                                        dataTableManipulate() 
                                    }
                                    else {
                                        console.log("Removing Failed!");
                                    }
                                }
                             });
                         });
                     } else{
                         $("#myWorkLocatinDeleteStatus").modal();
                     }
                 }else if(data =="false"){
                     console.log("iam in else part........");
                     $("#myGroupModal").modal();
                     $("#confirm").click(function(){
                        $.ajax({
                            type: "POST",
                            url: '/' + companyTeamName +'/worklocation/'+ workLocationId + '/delete',
                            data:'',
                            success: function(response){
                                if(response=="true"){
                                    $('#workLocation-table').dataTable().fnDestroy();
                                    var index = "";

                                    for(var i = 0; i < mainArray.length; i++) {
                                       index = mainArray[i].indexOf(workLocationId);
                                       if(index != -1) {
                                           console.log("dddd", i);
                                         break;
                                       }
                                    }
                                    mainArray.splice(i, 1);
                                    dataTableManipulate() 
                                }
                                else {
                                    console.log("Removing Failed!");
                                }
                            }

                        });
                    });
                 }
             }
         });
                     
    });
});

