/* Author :Aswathy Ashok */
//Below line is for adding active class to layout side menu..

document.getElementById("contact").className += " active";
var companyTeamName = vm.CompanyTeamName;
console.log("company name",vm.CompanyTeamName);
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
    
 
    function dataTableManipulate(){
        table =  $("#contact-details").DataTable({
            data: mainArray,
            "columnDefs": [{
                    "targets": 1,
                    render : function(data, type, row) {
                        return '<div class="over-length min-150">'+data+'</div>'
                    } 
                },
                {
                       "targets": -1,
                       "width": "5%",
                       "data": null,
                       "defaultContent": '<div class="edit-wrapper"><span class="icn"><i class="fa fa-pencil-square-o" aria-hidden="true" id="edit"></i><i class="fa fa-trash-o" aria-hidden="true" id="delete"></i></span></div>'
            }]
        });
        
        var item = $('<span>+</span>');
        item.click(function() {
            window.location = "/" + companyTeamName + "/contact/add";
        });
        
        $('.table-wrapper .dataTables_filter').append(item);
    }
    if(vm.Values != null) {
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate(); 

    $('#contact-details tbody').on( 'click', '#edit', function () {
        var data = table.row( $(this).parents('tr') ).data();
        var key = data[7];
        window.location = "/" + companyTeamName + "/contact/" + key + "/edit";
    });



    //...................
    $('#contact-details tbody').on( 'click', '#delete', function () {
      
        var data = table.row( $(this).parents('tr') ).data();
       var key = data[7];
        $.ajax({
            type: "POST",
            url: '/' + companyTeamName +'/contact/'+ key + '/delete',
            data: '',
            success: function(data){
                console.log("jjjj",data);
                if(data=="true"){
                    console.log("hdhhshhh");
                    $("#myContactModel").modal();
                    $("#deleteNotTask").click(function(){
                        $.ajax({
                            type: "POST",
                            url: '/' + companyTeamName +'/contact/'+ key + '/deletionOfContact',
                            data: '',
                            success: function(feedback){
                                console.log(feedback);
                                if(feedback=="true"){
                                    $('#contact-details').dataTable().fnDestroy(); 
                                    var index = "";
                                    for(var i = 0; i < mainArray.length; i++) {
                                    index = mainArray[i].indexOf(key);
                                    if(index != -1) {
                                        console.log("dddd", i);
                                        break;
                                    }
                                }
                                mainArray.splice(i, 1);
                                dataTableManipulate()
                                }
                                else {
                                }
                            }
                        });
                    });
                    
                }
                else {
                   
                    $("#myContactDelete").modal();
                    $("#confirm").click(function(){
                        $.ajax({
                            type: "POST",
                            url: '/' + companyTeamName +'/contact/'+ key + '/RemoveTask',
                            data: '',
                            success: function(response){
                                if(response=="true"){
                                    $('#contact-details').dataTable().fnDestroy(); 
                                    var index = "";
                                    for(var i = 0; i < mainArray.length; i++) {
                                    index = mainArray[i].indexOf(key);
                                    if(index != -1) {
                                        console.log("dddd", i);
                                        break;
                                    }
                                }
                                mainArray.splice(i, 1);
                                dataTableManipulate()
                                }
                                else {
                                }
                            }
                        });
                    });
                }
            }
        });
    });
    
});


