document.getElementById("fitToWork").className += " active";
var companyTeamName = vm.CompanyTeamName;
/*Function for creating Data Array for data table*/

//if (vm.NotificationArray.length !=0){
//        document.getElementById("number").textContent=vm.NotificationArray.length;
//
//    }else{
//        document.getElementById("number").textContent="";
//    }

var DynamicNotification ="";
    if (vm.NotificationNumber !=0){
        document.getElementById("number").textContent=vm.NotificationNumber;
    }else{
        document.getElementById("number").textContent="";
    }
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
        table =  $("#fit-to-work-details").DataTable({
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
        $('#fit-to-work-details tbody').on('click', 'td.details-control', function () {
            var tr = $(this).closest('tr');
            var row = table.row(tr);
            if ( row.child.isShown()){
                // This row is already open - close it
                row.child.hide();
                tr.removeClass('shown');
            }
            else {
                row.child( format(vm.InnerContent,row.data())).show();
                tr.addClass('shown');
            }
        });
        
        //function to display data inside expanded area
        function format ( InnerContent,data) {
            var userId  = data[2];
            var result  ='<div class="pull-left dropdown-tbl" style="padding-right: 50px;">';
             result +=  "<table cellpadding='5' cellspacing='0' border='0' style='padding-left:50px;' class='drp-tbl-wrp'>";
            
           
            
            /*result += '<th>Instructions</th>';*/
           
            result += "<tr>";
            result += '<th>Instructions</th>'
            for (var i=0; i<InnerContent.length;i++){
                if(InnerContent[i].InstructionKey ==userId){
                    result += "<tr>";
                    result += "<td><div class ='over-length'>"+InnerContent[i].Description+"</div></td>";
                    result += "</tr>";
                }
            }
            result += "</tr>";
            result += "</table>";
            result +="</div>";
            return result;
        }
        
/*Add a plus symbol in webpage for add new groups*/
        var item = $('<span>+</span>');
        item.click(function() {
            console.log("teamname",companyTeamName)
            window.location ="/" + companyTeamName + "/fitToWork/add";
        });
        $('.table-wrapper .dataTables_filter').append(item);
    }
    /*---------------------------Initial data table calling---------------------------------------------------*/

    if(vm.Values != null) {
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate();
    
    //notification
//     myNotification= function () {
//       console.log("hiiii");
//       var DynamicTaskListing="";
//        DynamicTaskListing ="<h5>"+"Notifications"+"</h5>"+"<ul>";
//       for(var i=0;i<vm.NotificationArray.length;i++){
//            console.log("sp1");
//           DynamicTaskListing += "<li>"+"User"+" "+vm.NotificationArray[i][2]+" "+vm.NotificationArray[i][3]+"  "+"delay to reach location"+" "+vm.NotificationArray[i][4]+" "+"for task"+" "+vm.NotificationArray[i][5]+"</li>";
//       }
//        $("#notificationDiv").prepend(DynamicTaskListing+"</ul>");
//        document.getElementById("number").textContent="";
//       $.ajax({
//           url:'/'+ companyTeamName + '/notification/update',
//           type: 'post',
////           datatype: 'json',
////           data: formData,
//           success : function(response) {
//               if (response == "true" ) {
////                   window.location = '/' + companyTeamName + '/task';
//               } else {
//               }
//           },
//           error: function (request,status, error) {
//               console.log(error);
//           }
//       });
//   }
    
    $('#fit-to-work-details tbody').on( 'click', '#edit', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var fitToWorkId = data[2];
        window.location = "/" + companyTeamName + "/fitToWork/"+fitToWorkId+"/edit";
        return false;
    });
    
     //notification
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
    
     
    
    
    
     $('#fit-to-work-details tbody').on( 'click', '#delete', function () {
         var data = table.row( $(this).parents('tr') ).data();
         var  fitToWorkId = data[2];
          $.ajax({
            type: "POST",
            url: '/' + companyTeamName +'/fitToWork/'+ fitToWorkId + '/deletionOfFitToWorkIfUsedForTask',
              
            data: '',
            success: function(response){
                console.log("dhfg",response)
                
                if(response=="true"){
                   
                     $("#myFitWorkModel").modal();
                }else{
                    $("#myModal").modal();
                    var data = table.row( $(this).parents('tr') ).data();
                    
                    $("#confirm").click(function(){
                        console.log("cp1");
                        $.ajax({
                            type: "POST",
                            url: '/' + companyTeamName +'/fitToWork/'+ fitToWorkId + '/delete',
                            data: '',
                            success: function(data){
                                if(data=="true"){
                                    $('#fit-to-work-details').dataTable().fnDestroy();
                                    var index = "";
                                    for(var i = 0; i < mainArray.length; i++) {
                                        index = mainArray[i].indexOf(fitToWorkId);
                                        if(index != -1) {
                                            console.log("dddd", i);
                                            break;
                                        }
                                    }
                                    mainArray.splice(i, 1);
                                    console.log(mainArray);
                                    dataTableManipulate();
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