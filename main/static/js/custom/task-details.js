/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..
document.getElementById("task").className += " active";
var companyTeamName = vm.CompanyTeamName
$(function(){ 
    
    var mainArray = [];   
    var table = "";
    var selectedCustomer = "";
    var tempJobArray = [];
    var tempArray = [];
    var tempViewArray = [];
    
    /*Function for Customer selection dropdown*/
    customerFilter = function(){
        tempArray = [];
        selectedCustomer = $("#customerDropdown").val();
        if (selectedCustomer == "All Customers") {
            $('#task-details').dataTable().fnDestroy();
            dataTableManipulate(mainArray); 
        } else {
            var tempSelectedCustomer = " (" + selectedCustomer + ")";
            for(i = 0; i < mainArray.length; i++){                
                if (mainArray[i][0].indexOf(tempSelectedCustomer) != '-1'){
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
                var tempJob = tempArray[i][0].replace(tempCustomer, '');
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
                if (mainArray[i][0].indexOf(tempSelectedJob) != '-1'){
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
            if (mainArray[i][0].indexOf(vm.SelectedJob) != '-1'){
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
        table =  $("#task-details").DataTable({
            data: dataArray,
            "paging": true,
            "columnDefs": [
                { "width": "15%", "targets": 0 },
                { "width": "15%", "targets": 1 },
                {   "width": "15%","targets": 2 },
                {   "width": "15%","targets": 3 },
                { "width": "15%", "targets": 4 },
                { "width": "15%", "targets": 5 },
                {
                    
                    "targets": 6,
                    "width": "20%",
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
        
        
        
        $('.table-wrapper .dataTables_filter').prepend(jobDropdown).prepend(customerDropdown).append(addItem).prepend( $('#example_wrapper .row:first-child > div.col-sm-6').addClass("gutter-padding")).prepend($('#example_wrapper .row:first-child > div.col-sm-6').removeClass("col-sm-6"));
        
        var customerArray = vm.UniqueCustomerNames;
        
        for(i = 0; i < customerArray.length; i++){
            $("#customerDropdown").append("<option>"+customerArray[i]+"</option>");
        }
        
        var jobArray = vm.UniqueJobNames;
        
        for(i = 0; i < jobArray.length; i++){
            $("#jobDropdown").append("<option>"+jobArray[i]+"</option>");
        }
    }
    
    
    
    
    //..................data table calling.......................
    if(vm.Values != null) {
        createDataArray(vm.Values, vm.Keys);
    }
    if(vm.SelectedJob == "" && vm.JobMatch == "true"){
        console.log("set1");
        dataTableManipulate(mainArray);
    } else if(vm.JobMatch=="false" && vm.SelectedJob =="false"){
        console.log("set2");
        dataTableManipulate(tempViewArray);
    }
    else {
        console.log("set3")
       taskAccordingToJob();
    }
    
    
    
    //.....................editing..................
    $('#task-details tbody').on( 'click', '#edit', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[11];
        window.location = '/' + companyTeamName + '/task/' + key + '/edit'
    });
//................deleting.........................
    $('#task-details tbody').on( 'click', '#delete', function () {
        $("#myModal").modal();
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[11];
        
        $("#confirm").click(function(){
            $.ajax({
                type: "POST",
                url: '/'  +   companyTeamName + '/task/' + key + '/delete',
                data: '',
                success: function(data){
                    if(data=="true"){
                        $('#task-details').dataTable().fnDestroy();
                        var index = "";
                        
                        for(var i = 0; i < mainArray.length; i++) {
                           index = mainArray[i].indexOf(key);
                           if(index != -1) {
                               break;
                           }
                        }
                        mainArray.splice(i, 1);
                        dataTableManipulate(mainArray);   
                    }
                    else {
                        console.log("Removing Failed!");
                    }
                }
            });
        });
    });
    
});