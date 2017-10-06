console.log(vm);
var companyTeamName = vm.CompanyTeamName;
$(function(){
    console.log("number",vm.NotificationNumber);
    if (vm.NotificationNumber !=0){
        console.log("kkk");
        document.getElementById("number").textContent=vm.NotificationNumber;
    }else{
        document.getElementById("number").textContent="";
    }
    document.getElementById("username").textContent=vm.AdminFirstName;
    document.getElementById("imageId").src=vm.ProfilePicture;
    if (vm.ProfilePicture ==""){
        document.getElementById("imageId").src="/static/images/default.png"
    }
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
         console.log("inside",values);
        var subArray = [];
        for(i = 0; i < values.length; i++) {
            for(var propertyName in values[i]) {
                subArray.push(values[i][propertyName]);
            }
            subArray.push(keys[i])
            mainArray.push(subArray);
            subArray = [];
        }
          console.log("main array",mainArray);
    }
    function dataTableManipulate(dataArray){
        console.log("manipulate");
       table =  $("#company-document-table").DataTable({
           data: dataArray,
           "columnDefs": [{
               "targets": -1,
               "width": "13%",
               "data": null,
               "defaultContent": '<div class="edit-wrapper"><span class="icn"><span class="dwnl-btn"><i class="fa fa-download fa-lg" aria-hidden="true" id="view"></i>'+" "+'</span><i class="fa fa-pencil-square-o" aria-hidden="true" id="edit"></i><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
           }],
       
       });
         var item = $('<span>+</span>');
        item.click(function() {
            console.log("teamname",companyTeamName)
            window.location ="/" + companyTeamName + "/companyFileUpload/add";
        });
        $('.table-wrapper .dataTables_filter').append(item);

   }
    
    /*Add a plus symbol in webpage for add new groups*/
       
    

/*----------------------------------Initialize Datatable--------------------------------------------------*/
   if(vm.Values != null) {
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate(mainArray);
/*--------------------------------Download-------------------------------------------------------------*/

    $('#company-document-table tbody').on( 'click', '#view', function () {
        var data = table.row( $(this).parents('tr') ).data();
        if(data[1] !=""){
            window.location =   data[1];
        } else{
            $("#noDocument").modal();
        }
        
        return false;
    });
/*------------------------------------------------------------------------------------------------------*/
//click on delete button
      $('#company-document-table tbody').on( 'click', '#delete', function () {
           $("#myModal").modal();
          console.log("indide delete");
         var data = table.row( $(this).parents('tr') ).data();
         var  documentId = data[2];
          $("#confirm").click(function(){
          $.ajax({
              type: "POST",
              url: '/' + companyTeamName +'/companyFileUpload/'+ documentId + '/delete',
              data: '',
              success: function(feedback){
                  console.log(feedback);
                  if(feedback=="true"){  
                      window.location ="/" + companyTeamName + "/companyFileUpload";
                      
                      
                  }
                  else {
                  }
              }
          });
          });
      });
      
    //click on edit button 
    
     $('#company-document-table tbody').on( 'click', '#edit', function () {
          console.log("indide edit");
         var data = table.row( $(this).parents('tr') ).data();
         var  documentId = data[2];
         window.location = "/" + companyTeamName + "/companyFileUpload/"+documentId+"/edit";
         return false;
     });
    console.log(vm);
});
    
$(document).ready(function() {
    if(vm.CompanyPlan == 'family' ){
        var parent = document.getElementById("menuItems");
        var contact = document.getElementById("contact");
        var job = document.getElementById("job");
        var crm = document.getElementById("crm");
        var leave = document.getElementById("leave");
        var timesheet  = document.getElementById("time-sheet");
        var consent = document.getElementById("consent");
        var workLocation = document.getElementById("workLocation");
        var leave = document.getElementById("leave");
        var log =  document.getElementById("log");
        var timesheet =document.getElementById("time-sheet");
        var fitToWork = document.getElementById("fitToWork");
        var dashBoard = document.getElementById("dashBoard");
        parent.removeChild(workLocation);
        parent.removeChild(timesheet);
        parent.removeChild(consent);
        parent.removeChild(leave);
        parent.removeChild(contact);
        parent.removeChild(job);
        parent.removeChild(crm);
        parent.removeChild(leave);
        parent.removeChild(log);
        parent.removeChild(timesheet);
        parent.removeChild(fitToWork);
        parent.removeChild(dashBoard);
        
        
    } else if(vm.CompanyPlan == 'campus'){
            var parent = document.getElementById("menuItems");
        var contact = document.getElementById("contact");
        var job = document.getElementById("job");
        var crm = document.getElementById("crm");
        var leave = document.getElementById("leave");
        var timesheet  = document.getElementById("time-sheet");
        var consent = document.getElementById("consent");
        var workLocation = document.getElementById("workLocation");
        var leave = document.getElementById("leave");
        var log =  document.getElementById("log");
        var timesheet =document.getElementById("time-sheet");
        var fitToWork = document.getElementById("fitToWork");
        //var dashBoard = document.getElementById("dashBoard");
        parent.removeChild(workLocation);
        parent.removeChild(timesheet);
        parent.removeChild(consent);
        parent.removeChild(leave);
        parent.removeChild(contact);
        parent.removeChild(job);
        parent.removeChild(crm);
        parent.removeChild(leave);
        parent.removeChild(log);
        parent.removeChild(timesheet);
        parent.removeChild(fitToWork);
       // parent.removeChild(dashBoard);
     }
    document.getElementById("username").textContent=vm.AdminFirstName;
    document.getElementById("imageId").src=vm.ProfilePicture;
    if (vm.ProfilePicture ==""){
        document.getElementById("imageId").src="/static/images/default.png"
    }
    if(vm.CompanyPlan == "family")
    {
        $('#planChange').attr('data-target','#family');
    } else if (vm.CompanyPlan == "campus") {
        $('#planChange').attr('data-target','#campus');
    }else if (vm.CompanyPlan == "business") {
        $('#planChange').attr('data-target','#business');
    }else if (vm.CompanyPlan == "businessPlus") {
        $('#planChange').attr('data-target','#business-plus');
    }
    
} );

    
    


