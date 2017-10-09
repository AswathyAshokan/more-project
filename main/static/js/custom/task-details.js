/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..

document.getElementById("task").className += " active";
var companyTeamName = vm.CompanyTeamName
$(function(){
    console.log(vm.UserArray);
    console.log("array",vm);
    var ExposureArray =vm.ExposureArray;
    var mainArray = []; 
    var table = "";
    var selectedCustomer = "";
    var tempJobArray = [];
    var tempArray = [];
    var tempViewArray = [];
    var customerArray =[];
    var rowIndex ="";
    /*Function for Customer selection dropdown*/
    customerFilter = function(){
        tempArray = [];
        selectedCustomer = $("#customerDropdown").val();
        if (selectedCustomer == "All Customers") {
            $('#task-details').dataTable().fnDestroy();
            dataTableManipulate(mainArray); 
        } else {
            var tempSelectedCustomer = " (" + selectedCustomer + ")";
            console.log(tempSelectedCustomer);
            for(i = 0; i < mainArray.length; i++){                
                if (mainArray[i][1].indexOf(tempSelectedCustomer) != '-1'){
                    tempArray.push(mainArray[i]);
                }
            }
            $('#task-details').dataTable().fnDestroy();
            dataTableManipulate(tempArray);
            
            $("#customerDropdown").val(selectedCustomer);
            
            //filtering job dropdown
            tempJobArray = [];
            
            for(i = 0; i < tempArray.length; i++){                
                var tempCustomer = " (" + selectedCustomer + ")";
                var tempJob = tempArray[i][1].replace(tempCustomer, '');
                if (tempJobArray.indexOf(tempJob) == '-1') {
                    tempJobArray.push(tempJob);
                }
            }
            
            $("#jobDropdown").empty().append("<option>All Jobs</option>");
            
            for(i = 0; i < tempJobArray.length; i++){
                $("#jobDropdown").append("<option>"+tempJobArray[i]+"</option>");
            }      
        }         
    }
   
    /*Function for Customer selection dropdown*/
    jobFilter = function(){
        var selectedJob = $("#jobDropdown").val();
        selectedCustomer = $("#customerDropdown").val();
        if (selectedJob == "All Jobs") {
            if (selectedCustomer == "All Customers") {
                tempArray = mainArray;
            }
            $('#task-details').dataTable().fnDestroy();
            dataTableManipulate(tempArray);
        } else {        
            var tempJobTableArray = [];
            var tempSelectedJob = selectedJob + " (";
            for(i = 0; i < mainArray.length; i++){                
                if (mainArray[i][1].indexOf(tempSelectedJob) != '-1'){
                    tempJobTableArray.push(mainArray[i]);
                }
            }
            $('#task-details').dataTable().fnDestroy();
            dataTableManipulate(tempJobTableArray);            
        }
        if (selectedCustomer != "All Customers") {
            $("#jobDropdown").empty().append("<option>All Jobs</option>");
            for(i = 0; i < tempJobArray.length; i++){
                $("#jobDropdown").append("<option>"+tempJobArray[i]+"</option>");
            }
        }            
        $("#jobDropdown").val(selectedJob);
        $("#customerDropdown").val(selectedCustomer);
    }
    
     /*Function for setting task details of a particular job*/
    function taskAccordingToJob(){
        var tempArray = [];
        for(i = 0; i < mainArray.length; i++){
            if (mainArray[i][1].indexOf(vm.SelectedJob) != '-1'){
                tempArray.push(mainArray[i]);
            }
        }
        $('#task-details').dataTable().fnDestroy();
        dataTableManipulate(tempArray);
        $("#customerDropdown").val(vm.SelectedCustomerForJob);
        $("#jobDropdown").val(vm.SelectedJob);
    }

    //create data for datatable
    
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
    
    //function for place  data to datatable
    function dataTableManipulate(dataArray){
        console.log("inside data manipulatae");
        table =  $("#task-details").DataTable({
            data: dataArray,
            "paging": true,
            "columnDefs": [
                { className: "details-control" , "targets": [ 0 ] },
                {
                    "order": [[1, 'asc']]
                },
                {
                    "targets": 6,
                    "width": "10%",
                    "data": null,
                    "defaultContent": '<div class="edit-wrapper"><span class="icn"></i><i class="fa fa-pencil-square-o" aria-hidden="true" id="edit"></i><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
                }]
        });
        var addItem = $('<span>+</span>');
        addItem.click(function() {
            window.location = "/" + companyTeamName + "/task/add";
        });
        var customerDropdown = $('<div class="tbl-dropdown"><select class="form-control sprites-arrow-down" id="customerDropdown"  onchange="customerFilter();"><option>All Customers</option></select></div>');
        
        var jobDropdown = $('<div class="tbl-dropdown"><select class="form-control sprites-arrow-down" id="jobDropdown"  onchange="jobFilter();"><option>All Jobs</option></select></div>');       
        

        //function to show expanded row
        $('.table-wrapper .dataTables_filter').prepend(jobDropdown).prepend(customerDropdown).append(addItem);
       
        //......................................................
        customerArray = vm.UniqueCustomerNames;
        var newcustomerArray = new Array();
        for (var i = 0; i < customerArray.length; i++) {
            if (customerArray[i]) {
                newcustomerArray.push(customerArray[i]);
            }
        }
        for(i = 0; i < newcustomerArray.length; i++){
            $("#customerDropdown").append("<option>"+newcustomerArray[i]+"</option>");
        }
        var jobArray = vm.UniqueJobNames;
        
        var newjobArray = new Array();
        for (var i = 0; i < jobArray.length; i++) {
            if (jobArray[i]) {
                newjobArray.push(jobArray[i]);
            }
        }
        
        for(i = 0; i < newjobArray.length; i++){
            $("#jobDropdown").append("<option>"+newjobArray[i]+"</option>");
        }
            
        
        
    }
    
    
     $('#task-details tbody').on('click', 'td.details-control', function () {
         console.log("row expansion");
            var tr = $(this).closest('tr');
            var row = table.row( tr );
            if ( row.child.isShown() ) {
                row.child.hide();
                tr.removeClass('shown');
            }
            else {
                row.child( format(vm.UserArray,row.data(),vm.MinUserAndLoginTypeArray)).show();
                
                tr.addClass('shown');
            }
        } );
        
        
        //function to display data inside expanded area
        function format ( userDetailsArray, data,minUserArray ) {
        // `d` is the original data object for the row
            var taskID  = data[6];
            var result   ='<div class="pull-left dropdown-tbl" style="padding-right: 50px;">';
            result += "<table cellpadding='5' cellspacing='0'  style='padding-left:50px; border: 1px solid #dddddd !important;'>";
            result += '<th>User assigned</th>';
            result += '<th>Status</th>';
            result += "<tr>";
            for (var i = 0; i < userDetailsArray.length; i++) {
                if(userDetailsArray[i] != null && userDetailsArray[i][0].TaskId == taskID) {
                    for (var j=0; j<userDetailsArray[i].length ;j++){
                        result += "<td>"+userDetailsArray[i][j].Name+"</td>";
                        result += "<td>"+userDetailsArray[i][j].Status+"</td>";
                        result += "</tr>";
                    }
                }
            }
            result += "</table  >";
            result +="</div>";
            var minUser ='<div class="pull-left dropdown-tbl"  style="padding-right: 50px;">';
            minUser +="<table cellpadding='5' cellspacing='0' border='0' style=''>";
            minUser +='<tr>';
            for (var i=0; i<minUserArray.length; i++){
              
                if(minUserArray[i] != null && minUserArray[i][4] == taskID) {
                    minUser +='<td>Minimum no of users </td>';
                    minUser +='<td>'+minUserArray[i][0]+'</td>';
                    minUser +='</tr>';
                    minUser +='<tr>';
                    minUser +='<td>Login type </td>';
                    minUser +='<td>'+minUserArray[i][1]+'</td>';
                    minUser +='</tr>';
//                    minUser +='<tr>';
//                    minUser +='<td>Log Time In Minutes </td>';
//                    minUser +='<td>'+minUserArray[i][2]+'</td>';
//                    minUser +='</tr>';
                    minUser +='<tr>';
                    minUser +='<td>Fit To WorkName </td>';
                    minUser +='<td>'+minUserArray[i][3]+'</td>';
                    minUser +='</tr>';
                    
                }
            }
            minUser +="</table>";
            minUser +="</div>";
            minUser +="</table>";
            minUser +="</div>";
            
            //exposure
//            var exposure   ='<div class="pull-left dropdown-tbl" style="">';
//            exposure += "<table cellpadding='5' cellspacing='0' style='border: 1px solid #dddddd !important;'>";
//            exposure += '<th>Exposure Details</th>';
//            exposure += "<tr>";
//            for (var i = 0; i < ExposureArray.length; i++) {
//                 
//                if(ExposureArray[i] != null && ExposureArray[i][0].TaskId == taskID ) {
//                    console.log("task id exposure",ExposureArray[i][0].TaskId );
//                    console.log ("ggg",taskID);
//                     
//                    for (var j=0; j<ExposureArray[i].length ;j++){
//                        var Breakhours = Math.trunc(ExposureArray[i][j].BreakMinute/60);
//                        var Breakminutes = ExposureArray[i][j].BreakMinute % 60;
//                        var Workhours = Math.trunc(ExposureArray[i][j].WorkingHour/60);
//                        var Workminutes = ExposureArray[i][j].WorkingHour % 60;
//                        exposure += "<td>"+Breakhours +":"+ Breakminutes+" Minutes Break After"+Workhours +":"+ Workminutes+"Hours"+"</td>";
//                        exposure += "</tr>";
//                    }
//                }
//            }
//            exposure += "</table  >";
//            exposure +="</div>";
            
            return result+minUser;
        }
    //..................data table calling.......................
    if(vm.Values != null) {
        createDataArray(vm.Values, vm.Keys);
    }
    if(vm.SelectedJob == "" && vm.JobMatch == "true"){
        dataTableManipulate(mainArray);
    } else if(vm.JobMatch=="false" && vm.SelectedJob =="false"){
        dataTableManipulate(tempViewArray);
    }
    else {
        taskAccordingToJob();
    }
    //.....................editing..................
    $('#task-details tbody').on( 'click', '#edit', function () {
        console.log("edit");
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[6];
        $.ajax({
                type: "POST",
                url: '/'  +   companyTeamName + '/task/' + key + '/taskStatus',
                data: '',
                success: function(data){
                    if(data=="true"){
                           $("#myTaskStatus").modal();
                    }
                    else {
                        window.location = '/' + companyTeamName + '/task/' + key + '/edit'
                    }
                }
            });
        
        
        
        
//        window.location = '/' + companyTeamName + '/task/' + key + '/edit'
    });
//................deleting.........................
    $('#task-details tbody').on( 'click', '#delete', function () {
        console.log("delete");
        $("#myModal").modal();
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[6];
        
        $("#confirm").click(function(){
            $.ajax({
                type: "POST",
                url: '/'  +   companyTeamName + '/task/' + key + '/delete',
                data: '',
                success: function(data){
                    console.log("dddd",data);
                    if(data=="true"){
                        $('#task-details').dataTable().fnDestroy();
                        var index = "";
                        
                        for(var i = 0; i < mainArray.length; i++) {
                           index = mainArray[i].indexOf(key);
                           if(index != -1) {
                               console.log("inside delete");
                               break;
                           }
                        }
//                        mainArray.splice(i, 1);
//                        dataTableManipulate(mainArray);   
                         window.location = '/' + companyTeamName + '/task';
                    }
                    else {
                        console.log("Removing Failed!");
                    }
                }
            });
        });
    });
    
});