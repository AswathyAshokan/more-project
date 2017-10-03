document.getElementById("WorkLocation").className += " active";
console.log("uuuuuuuuuuuuuuu",vm.Values);
var companyTeamName = vm.CompanyTeamName;
var ExposureArray =vm.ExposureArray;
if (vm.NotificationNumber !=0){
    document.getElementById("number").textContent=vm.NotificationNumber;
}else{
    document.getElementById("number").textContent="";
}

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
    
   //notification
    var notificationSorted =[[]];
    function sortByCol(arr, colIndex){
    notificationSorted=arr.sort(sortFunction);
    function sortFunction(a, b) {
        a = a[colIndex]
        b = b[colIndex]
        return (a === b) ? 0 : (a < b) ? -1 : 1
    }
}

    
     myNotification= function () {
         if (vm.NotificationArray !=null){
        console.log("hiiii");
         sortByCol(vm.NotificationArray, 6);
         console.log("jjjjj",notificationSorted);
         var reverseSorted =[[]];
         reverseSorted=notificationSorted.reverse();

        document.getElementById("notificationDiv").innerHTML = "";
        var DynamicTaskListing="";
        if (reverseSorted !=null){
            DynamicTaskListing ="<h5>"+"Notifications"+ "<button class='no-button-style' method='post' onclick='clearNotification()'>"+"clear all"+"</button>"+"</h5>"+"<ul>";
        for(var i=0;i<reverseSorted.length;i++){
            console.log("sp1");
            var timeDifference =moment(new Date(new Date(reverseSorted[i][6]*1000)), "YYYYMMDD").fromNow();
            DynamicTaskListing += "<li>"+"User"+" "+reverseSorted[i][2]+" "+reverseSorted[i][3]+"  "+"delay to reach location"+" "+reverseSorted[i][4]+" "+"for task"+" "+reverseSorted[i][5]+" <span>"+timeDifference+"</span>"+"</li>";
            
            
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
             document.getElementById("notificationDiv").innerHTML = "";
            DynamicTaskListing ="<h5>"+" No New Notifications"+"</h5>";
                        $("#notificationDiv").prepend(DynamicTaskListing);
            
        }
        
        }else{
             document.getElementById("notificationDiv").innerHTML = "";
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
                        DynamicTaskListing ="<h5>"+" No New Notifications"+"</h5>";
                        $("#notificationDiv").prepend(DynamicTaskListing);
                    } else {
                    }
                },
                error: function (request,status, error) {
                    console.log(error);
                }
            }); 
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
            var exposure   ='<div class="pull-left dropdown-tbl" style="">';
            exposure += "<table cellpadding='5' cellspacing='0' style='border: 1px solid #dddddd !important;'>";
            exposure += '<th>Exposure Details</th>';
            exposure += "<tr>";
            for (var i = 0; i < ExposureArray.length; i++) {
                 
                if(ExposureArray[i] != null && ExposureArray[i][0].TaskId == workLocationID ) {
                    console.log("task id exposure",ExposureArray[i][0].TaskId );
                    for (var j=0; j<ExposureArray[i].length ;j++){
                        var Breakhours = Math.trunc(ExposureArray[i][j].BreakMinute/60);
                        var Breakminutes = ExposureArray[i][j].BreakMinute % 60;
                        var Workhours = Math.trunc(ExposureArray[i][j].WorkingHour/60);
                        var Workminutes = ExposureArray[i][j].WorkingHour % 60;
                        exposure += "<td>"+Breakhours +":"+ Breakminutes+" Minutes Break After    "+Workhours +":"+ Workminutes+"Hours"+"</td>";
                        exposure += "</tr>";
                    }
                }
            }
            exposure += "</table  >";
            exposure +="</div>";
            
            return minUser+exposure;
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
                        tempArry.push(vm.Users[j][k].Name);
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
         $("#myGroupModal").modal();
        var data = table.row( $(this).parents('tr') ).data();
        console.log("full data",data);
        console.log("data id",data[5]);
        var workLocationId = data[5];
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
    });
});

