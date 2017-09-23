/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..

document.getElementById("job").className += " active";
var companyTeamName = vm.CompanyTeamName;
var DynamicNotification ="";
    if (vm.NotificationNumber !=0){
        document.getElementById("number").textContent=vm.NotificationNumber;
    }else{
        document.getElementById("number").textContent="";
    }
$().ready(function() {
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
    
     
    
    var pageType = vm.PageType;
    
    if(pageType == "edit") {
       var selectArray =vm.CustomerId;
        console.log("customer",selectArray);
        document.getElementById("jobName").value = vm.JobName;
        document.getElementById("orderNumber").value = vm.OrderNumber;
         document.getElementById("orderDate").value = vm.OrderDate;
        document.getElementById("jobNumber").value = vm.JobNumber;
        $("#customerId").val(selectArray);
        //$("#customerId option[text=selectArray]").attr("selected","selected");
        document.getElementById("jobHead").innerHTML = "Edit Job";
        
          $("#jobForm").validate({
        rules: {
            customerId:"required",
            jobName: {
                required: true,
                remote:{
                    url: "/isJobNameUsed/" + jobName+ "/" +vm.PageType+ "/" + vm.JobName,
                    type: "post"
                }
            },
            
            jobNumber: {
                required: true,
                remote:{
                    url: "/isJobNumberUsed/" + jobNumber+ "/" +vm.PageType+ "/" + vm.JobNumber,
                    type: "post"
                }
            }
        },
        messages: {
            jobName: {
                required: "Enter job name",
                remote: "Job name already exists!"
            },
            jobNumber:{
                required:"Enter job number",
                remote: "Job number already exists!",
            },
        },
        submitHandler: function() { 
            
             $("#saveButton").attr('disabled', true);
            var formData = $("#jobForm").serialize();
            var customerName = $('#customerId option:selected').text();
            formData = formData +"&customerName="+customerName;
            console.log(formData);
            var jobId = vm.JobId;
            if (pageType == "edit") {
                $.ajax({
                    url: '/' + companyTeamName  + '/job/'+ jobId +'/edit',
                    type: 'post',
                    datatype: 'json',
                    data: formData,
                    success : function(response) {
                        console.log(response);
                        if (response == "true") {
                            window.location = '/' + companyTeamName + '/job';
                        } else {
                            $("#saveButton").attr('disabled', false);
                        }
                    },
                    error: function (request,status, error) {
                        console.log(error);
                    }
                });
            } 
        }
    });
    }
    
     if(pageType == "add") {
         $("#jobForm").validate({
        rules: {
            customerId:"required",
            jobName: {
                required: true,
                remote:{
                    url: "/isJobNameUsed/" + jobName,
                    type: "post"
                }
            },
            
            jobNumber: {
                required: true,
                remote:{
                    url: "/isJobNumberUsed/" + jobNumber,
                    type: "post"
                }
            }
        },
        messages: {
            jobName: {
                required: "Enter job name",
                remote: "Job name already exists!"
            },
            jobNumber:{
                required:"Enter job number",
                remote: "Job number already exists!",
            },
        },
        submitHandler: function() { 
            
             $("#saveButton").attr('disabled', true);
            var formData = $("#jobForm").serialize();
            var customerName = $('#customerId option:selected').text();
            formData = formData +"&customerName="+customerName;
            console.log(formData);
            var jobId = vm.JobId;
             
                $.ajax({
                    url: '/' + companyTeamName + '/job/add',
                    type: 'post',
                    datatype: 'json',
                    data: formData,
                    success : function(response) {
                        if (response =="true") {
                            window.location ='/' + companyTeamName + '/job'
                        } else {
                            $("#saveButton").attr('disabled', false);
                        }
                    },
                    error: function (request,status, error) {
                        console.log(error);
                    }
                });
        }
    });
    }
    
    
    //adding alphanumeric validation 
    

  
    $("#cancel").click(function() {
        window.location = '/' +  companyTeamName  + '/job';
    });
});