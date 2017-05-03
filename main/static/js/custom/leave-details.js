console.log(vm);
/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..
document.getElementById("leave").className += " active";
 var companyTeamName = vm.CompanyTeamName;
console.log("company team name",companyTeamName);
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
    console.log("main array",vm.Values);
    function dataTableManipulate(){
        table =  $("#leave_details").DataTable({
            data: mainArray,
            "searching": false,
            "paging": true, 
            "info": false,
            "lengthChange":false,
            "columnDefs": [{
                "targets": -1,
                "width": "5%",
                "data": null,
                "defaultContent": '<button class="btn btn-primary btn-xs " id ="accept">Accept</button>'+"  "+'<button class="btn btn-danger btn-xs " id="reject">Reject</button>'
            }]
        });
        
    }
    if(vm.Values != null) {
        console.log("inside array creation");
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate();
    
    //function when click on accept button
    $('#leave_details').on( 'click', '#accept', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var leaveKey = data[5];
        var userKey =data[6];
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
                               console.log("dddd", i);
                             break;
                           }
                        }
                        mainArray.splice(i, 1);
                        dataTableManipulate(mainArray);   
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
        var leaveKey = data[5];
        var userKey =data[6];
        //alert(data[4]);
       // window.location = '/'+ companyTeamName +'/leave/' + leaveKey +'/'+userKey+ '/edit';
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
                               console.log("dddd", i);
                             break;
                           }
                        }
                        mainArray.splice(i, 1);
                        dataTableManipulate(mainArray);   
                    }
                    else {
                        console.log("Updation Failed!");
                    }
                }

            });
    });


//    $('#leave_details tbody').on( 'click', '#delete', function () {
//        $("#myModal").modal();
//        var data = table.row( $(this).parents('tr') ).data();
//        var key = data[6];
//        
//        $("#confirm").click(function(){
//            $.ajax({
//                type: "POST",
//                url: '/' + companyTeamName + '/contact/' + key + '/delete',
//                data: '',
//                success: function(data){
//                    if(data=="true"){
//                        $('#leave_details').dataTable().fnDestroy();
//                        var index = "";
//                        
//                        for(var i = 0; i < mainArray.length; i++) {
//                           index = mainArray[i].indexOf(key);
//                           if(index != -1) {
//                             break;
//                           }
//                        }
//                        mainArray.splice(i, 1);
//                        dataTableManipulate();   
//                    }
//                    else {
//                        console.log("Removing Failed!");
//                    }
//                }
//
//            });
//        });
//    });
    
});


