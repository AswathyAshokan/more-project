//Created By Farsana
//Below line is for adding active class to layout side menu..
console.log("jdjdjd",vm.DocumentExpiryNotification);
document.getElementById("crm").className += " active";
console.log("vm.CustomerName",vm.CustomerName)
var companyTeamName = vm.CompanyTeamName;
var DynamicNotification ="";
    if (vm.NotificationNumber !=0){
        document.getElementById("number").textContent=vm.NotificationNumber;
    }else{
        document.getElementById("number").textContent="";
    }


$().ready(function() {
    
    
    
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
        console.log("new array", vm.NotificationArray);
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
                    if (reverseSorted[i][5]==""){
                        console.log("iam in first");
                        var utcTime =reverseSorted[i][3];
                        var dateFromDb = parseInt(utcTime);
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
                        var startDate = (mm + '/' + dd + '/' + yyyy);
                        var utcTime =reverseSorted[i][4];
                        var dateFromDb = parseInt(utcTime);
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
                            var endDate = (mm + '/' + dd + '/' + yyyy);
                            var timeDifferenceForLeave =moment(new Date(new Date(reverseSorted[i][6]*1000)), "YYYYMMDD").fromNow();
                             DynamicTaskListing += "<li>"+"User"+" "+reverseSorted[i][2]+" "+"Applied Leave For "+"  "+reverseSorted[i][7]+" "+"Days"+" "+"From"+" "+startDate+" "+"to"+" "+endDate+" <span>"+timeDifferenceForLeave+"</span>"+"</li>";
                    } else if (reverseSorted[i][5] =="Expiry111@@&&EEE"){
                        console.log("iam in second");
                        var timeDifferenceForLeave =moment(new Date(new Date(reverseSorted[i][6]*1000)), "YYYYMMDD").fromNow();
                         DynamicTaskListing += "<li>"+"User"+" "+reverseSorted[i][0]+"'s"+" "+"document of type"+" " +reverseSorted[i][4]+ " "+ "will be expired on" +" "+"<span>"+timeDifferenceForLeave+"</span>"+"</li>";
                    } else{
                        console.log("sp1");
                        var timeDifference =moment(new Date(new Date(reverseSorted[i][6]*1000)), "YYYYMMDD").fromNow();
                        DynamicTaskListing += "<li>"+"User"+" "+reverseSorted[i][2]+" "+reverseSorted[i][3]+"  "+"delay to reach location"+" "+reverseSorted[i][4]+" "+"for task"+" "+reverseSorted[i][5]+" <span>"+timeDifference+"</span>"+"</li>";
                    }
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
    
    

    
     //this is for notification of expiredDetails
        var expirycount = 0;
       
        var documentNotifyArry =  vm.DocumentExpiryNotification;
        if( vm.DocumentExpiryNotification!=null){
            for(i = 0;i<vm.DocumentExpiryNotification.length;i++){
                console.log("haiiiii");
                console.log("document values",documentNotifyArry[i]);
                var today = new Date();
                var dd = today.getDate();
                var mm = today.getMonth()+1; //January is 0!
                var yyyy = today.getFullYear();
                if(dd<10) {
                    dd = '0'+dd;
                } 
                if(mm<10) {
                    mm = '0'+mm;
                }
                var chechNotification
                var tempArray = [];
                var CurrentMonth = mm;
                var currentDay = dd;
                var currentYear = yyyy;
                var localToday = (mm + '/' + dd + '/' + yyyy);
                var dateParts = documentNotifyArry[i][3].split("/");
                var dateFromDb = (dateParts[1]+'/'+ dateParts[0]+'/'+ dateParts[2]);
                console.log("kkkkkkkkkkk",dateFromDb);
                if(CurrentMonth ==dateParts[1] && currentDay ==dateParts[0] &&currentYear ==  dateParts[2]){
                    console.log("iam in if loop");
                    if(documentNotifyArry[i][2] == "false"){
                       expirycount = expirycount+1;
                    }
                     
                    chechNotification = documentNotifyArry[i][0]+"111@@&&EEE";
                    tempArray.push(documentNotifyArry[i][5]);
                    tempArray.push(documentNotifyArry[i][1]);
                    tempArray.push(documentNotifyArry[i][2]);
                    tempArray.push(dateFromDb);
                    tempArray.push(documentNotifyArry[i][6]);
                    tempArray.push(chechNotification);
                    tempArray.push(documentNotifyArry[i][4]);
                    vm.NotificationArray.push(tempArray);
                    tempArray = [];
                    console.log("###########",vm.NotificationArray)
                    
                }
            }
        }
    vm.NotificationNumber = vm.NotificationNumber +expirycount;
    console.log("kkkkkgfgffggfgfgf",vm.NotificationArray);
    console.log("iddddd",vm.NotificationNumber);
    
    
    
    
    
    
    
    
   
    
    if(vm.PageType == "edit"){        
            
            document.getElementById("customername").value = vm.CustomerName;
            document.getElementById("contactperson").value = vm.ContactPerson;
            document.getElementById("country").value = vm.Country;
            document.getElementById("email").value = vm.Email;
            document.getElementById("phone").value = vm.Phone;
            document.getElementById("address").value = vm.Address;
            document.getElementById("state").value = vm.State;
            document.getElementById("zipcode").value = vm.ZipCode;
            document.getElementById("customerEdit").innerHTML = "Edit Customer"
    }
    $("#addcustomerForm").validate({
        rules: {
          customername:{
              required:true,
              remote:{
                  url: "/iscustomernameused/" + customername + "/" + vm.PageType + "/" + vm.CustomerName,
                  type: "post"
              }
              
          },
          contactperson:"required",
          email:{
              required:true,
              email:true
          },
            phone:"required",
            address:"required",
            country:"required",
            state: "required",
            zipcode: "required"
      },
        messages: {
            customername:{
                required: "Enter Customer Name ",
                remote: "Customer Name is already in use !"
                },
            contactperson:"Enter Contact Person",
            phone: {
                required:"Enter Phone Number"
            },
            address:"Enter your Address",
            state: "Enter your State",
            zipcode:"Enter zipcode  ",
            country:"Enter country name ",
            email:"Enter valid Email id"
    },
        submitHandler: function(){//to pass all data of a form serial
            $("#saveButton").attr('disabled', true);
            if (vm.PageType == "edit"){
                var formData = $("#addcustomerForm").serialize();
                var customerId = vm.CustomerId;
                $.ajax({
                    url:'/' + companyTeamName +'/customer/'+ customerId + '/edit',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/' + companyTeamName +'/customer';
                        }else {
                            $("#saveButton").attr('disabled', false);
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
            } else {
                var formData = $("#addcustomerForm").serialize();
                $.ajax({
                    url:'/' + companyTeamName +'/customer/add',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    //call back or get response here
                    success : function(response){
                        if(response == "true"){
                            window.location='/' + companyTeamName +'/customer';
                        }else {
                            $("#saveButton").attr('disabled', false);
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
            }
            return false;
        }
    });
    
    $("#cancel").click(function() {
            window.location = '/' + companyTeamName +'/customer';
    });
});