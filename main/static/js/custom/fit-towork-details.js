document.getElementById("fitToWork").className += " active";
var companyTeamName = vm.CompanyTeamName;
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
    
    $('#fit-to-work-details tbody').on( 'click', '#edit', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var fitToWorkId = data[2];
        window.location = "/" + companyTeamName + "/fitToWork/"+fitToWorkId+"/edit";
        return false;
    });
    
     $('#fit-to-work-details tbody').on( 'click', '#delete', function () {
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