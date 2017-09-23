document.getElementById("consent").className += " active";
var companyTeamName = vm.CompanyTeamName;
var DynamicNotification ="";
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
    
/*Function for assigning data array into data table*/
    function dataTableManipulate(){
        table =  $("#consent-details").DataTable({
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
    
     
    
    
/*Add a plus symbol in webpage for add new groups*/
        var item = $('<span>+</span>');
        item.click(function() {
            console.log("teamname",companyTeamName)
            window.location ="/" + companyTeamName + "/consent/add";
        });
        $('.table-wrapper .dataTables_filter').append(item);
    }
    
    $('#consent-details tbody').on('click', 'td.details-control', function () {
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
            for (var i=0; i<InnerContent.length;i++){
                if(InnerContent[i].InstructionKey ==userId){
                    result += '<th>Instructions</th>'
                    result += "<td><div class ='over-length'>"+InnerContent[i].Description+"</div></td>";
                    result += "<tr>";
                    result += '<th>Accepted Users</th>';
                    if (InnerContent[i].AcceptedUsers != null){
                        result += "<td><div class ='over-length'><span>"+InnerContent[i].AcceptedUsers+"</div></span></td>";
                    } else{
                        result += "<td><div class ='over-length'><span>"+"Nill"+"</div></span></td>";
                        
                    }
                    
                    result += "</tr>";
                    result += "<tr>";
                    result += '<th>Rejected Users</th>';
                    if (InnerContent[i].RejectedUsers !=null){
                        result += "<td><div class ='over-length'><span>"+InnerContent[i].RejectedUsers+"</div></span></td>";
                   } else{
                        result += "<td><div class ='over-length'><span>"+"Nill"+"</div></span></td>";
                        
                    }
                    result += "</tr>";
                    result += "<tr>";
                    result +='<th>Pending Users</th>'
                    if (InnerContent[i].PendingUsers !=null){
                        result += "<td><div class ='over-length'><span>"+InnerContent[i].PendingUsers+"</div</span</td>";
                    }else{
                        result += "<td><div class ='over-length'><span>"+"Nill"+"</div></span></td>";
                        
                    }
                    result += "</tr>";
                }
            }
            result += "</tr>";
            result += "</table>";
            result +="</div>";
            return result;
        }
    /*---------------------------Initial data table calling---------------------------------------------------*/

    if(vm.Values != null) {
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate(); 
    
    $('#consent-details tbody').on( 'click', '#edit', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var consentId = data[2];
        window.location = "/" + companyTeamName + "/consent/"+consentId+"/edit";
        return false;
    });
    
     $('#consent-details tbody').on( 'click', '#delete', function () {
        $("#myModal").modal();
        var data = table.row( $(this).parents('tr') ).data();
        var  consentId = data[2];
        console.log(data, consentId);
        $("#confirm").click(function(){
            console.log("cp1");
            $.ajax({
                type: "POST",
                url: '/' + companyTeamName +'/consent/'+ consentId + '/delete',
                data: '',
                success: function(data){
                    if(data=="true"){
                        $('#consent-details').dataTable().fnDestroy();
                        var index = "";
                        for(var i = 0; i < mainArray.length; i++) {
                            index = mainArray[i].indexOf(consentId);
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
    });
    
});