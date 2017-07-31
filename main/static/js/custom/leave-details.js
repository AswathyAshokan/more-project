console.log(vm.Values);
/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..
document.getElementById("leave").className += " active";
 var companyTeamName = vm.CompanyTeamName;
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
        $('#leave_details').dataTable().fnDestroy();
        $('#fromDate').datepicker('setDate', null);
        $('#toDate').datepicker('setDate', null);
        dataTableManipulate(completeTable);
     });
    function listSharedDocumentByDate(unixFromDate,unixToDate){
        var tempArray = [];
        var startDate =0;
        var unixEndDate =0;
        var unixStartDate = 0;
        for (i =0;i<vm.Values.length;i++){
            startDate = new Date(vm.Values[i][1]);
            endDate = new Date(vm.Values[i][2]);
            unixStartDate = Date.parse(startDate)/1000;
            unixEndDate = Date.parse(endDate)/1000;
           if( (unixFromDate <= unixStartDate && unixStartDate <= unixToDate) || (unixFromDate <= unixEndDate && unixEndDate <= unixToDate) || (unixFromDate >= startDate && unixEndDate >= unixToDate)) {

                tempArray.push(mainArray[i]);
           }

            $('#leave_details').dataTable().fnDestroy();
            dataTableManipulate(tempArray);
        }
    }

    function dataTableManipulate(mainArray){
        table =  $("#leave_details").DataTable({
            data: mainArray,
            "paging": true,
            "info": false,
            "lengthChange":false,
            "columnDefs": [{
                "targets": [5],
                 render : function (data, type, row) {
                     switch(data) {
                         case 'Accepted' : return '<button class="btn btn-primary btn-xs " >Accepted</button>'; break;
                         case 'Rejected' : return '<button class="btn btn-danger btn-xs " >Rejected</button>'; break;
                         case 'Pending' : return '<button class="btn btn-primary btn-xs " id ="accept">Accept</button>'+"  "+'<button class="btn btn-danger btn-xs " id="reject">Reject</button>'; break;
                         case 'Subcontractor':return '<button class="btn btn-primary btn-xs " >Leave Applied</button>'; break;

                         default  : return 'N/A';
                     }
                 }
            }]
        });
        $('#tbl_details_length').after($('.datepic-top'));

    }
    if(vm.Values != null) {
        for(i = 0;i<vm.Values.length;i++){
            var startUtcDate = vm.Values[i][1];
            var startUtcInDateForm = new Date(startUtcDate);
            var startLocalDate = (startUtcInDateForm.toLocaleDateString());
            var startDate = startLocalDate.slice(0, 10).split('/');
            var formatedStartDate = startDate[1] +'/'+ startDate[0] +'/'+startDate[2];

            var endUtcDate = vm.Values[i][2];
            var endUtcInDateForm = new Date(endUtcDate);
            var endLocalDate = (endUtcInDateForm.toLocaleDateString());
            var d = endLocalDate.slice(0, 10).split('/');
            var formatedEndDate = d[1] +'/'+ d[0] +'/'+ d[2];

            vm.Values[i][1] = formatedStartDate;
            vm.Values[i][2] = formatedEndDate;
        }
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate(mainArray);

    //function when click on accept button
    $('#leave_details').on( 'click', '#accept', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var leaveKey = data[6];
        var userKey =data[7];
        //alert(data[4]);
       // window.location = '/'+ companyTeamName +'/leave/' + leaveKey +'/'+userKey+ '/edit';
        $.ajax({
                type: "GET",
                url: '/'+ companyTeamName +'/leave/' + leaveKey +'/'+userKey+ '/accept',
                data: '',
                success: function(data){
                    if(data=="true"){
                        $('#leave_details').dataTable().fnDestroy();
                        var index = "";

                        for(var i = 0; i < mainArray.length; i++) {
                           index = mainArray[i].indexOf(leaveKey);
                           if(index != -1) {

                             break;
                           }
                        }
                        mainArray.splice(i, 1);
                        dataTableManipulate(mainArray);
                        window.location =  '/'+ companyTeamName +'/leave';
                    }
                    else {
                        console.log("Updation Failed!");
                    }
                }

            });
    });

    //function when click on reject button
    $('#leave_details').on( 'click', '#reject', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var leaveKey = data[6];
        var userKey =data[7];
        $.ajax({
                type: "GET",
                url: '/'+ companyTeamName +'/leave/' + leaveKey +'/'+userKey+ '/reject',
                data: '',
                success: function(data){
                    if(data=="true"){
                        $('#leave_details').dataTable().fnDestroy();
                        var index = "";

                        for(var i = 0; i < mainArray.length; i++) {
                           index = mainArray[i].indexOf(leaveKey);
                           if(index != -1) {

                             break;
                           }
                        }
                        mainArray.splice(i, 1);
                        dataTableManipulate(mainArray);
                        window.location =  '/'+ companyTeamName +'/leave';
                    }
                    else {
                        console.log("Updation Failed!");
                    }
                }

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
        console.log("unixFromDate",unixFromDate);
         console.log("unixToDate",unixToDate);
        listSharedDocumentByDate(unixFromDate,unixToDate);
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
        listSharedDocumentByDate(unixFromDate,unixToDate);
    });
});

