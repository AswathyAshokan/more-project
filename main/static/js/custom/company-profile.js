
//Below line is for adding active class to layout side menu..
console.log(vm);
document.getElementById("user").className += " active";
console.log("keys",vm.Keys)

var companyTeamName = vm.CompanyTeamName;
/*Function for creating Data Array for data table*/
$(function(){ 
    
    var userResponse;
    var mainArray = [];   
    var table = "";
    function createDataArray(values, keys){
        var subArray = [];
        for(i = 0; i < values.length; i++) {
            console.log("keys of ",keys[i]);
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
        table =  $("#company-profile-table").DataTable({
            data: mainArray,
           "columnDefs": [
                
                { className: "details-control" , "targets": [0]},
                {
                    "order": [[1, 'asc']]
                },
                
            ]
        });
    }
         $('#company-profile-table tbody').on('click', 'td.details-control', function () {
            var tr = $(this).closest('tr');
            var row = table.row(tr);
            if ( row.child.isShown()){
                // This row is already open - close it
                row.child.hide();
                tr.removeClass('shown');
            }
            else {
                row.child( format(vm.UserExpand,row.data(),vm.NextOfKin) ).show();
                tr.addClass('shown');
            }
        });
        
        //function to display data inside expanded area
       function format ( userDetailsArray, data,nextOfKin ) {
           var userID  = data[4];
           var defaultImage ="/static/images/defaultImage.png";
           var result   ='<div class="user-info col-sm-7">';
           for (var i = 0; i < userDetailsArray.length; i++) {
                if(userDetailsArray[i] != null && userDetailsArray[i][8] == userID) {
                    
                    if (userDetailsArray[i][7] =="img/profile.png"){
                        result += '<div id ="imageId" class="info-img" style="background-image:url('+defaultImage+')"></div>';
                    }else{
                        console.log("kkkkkkkkkkk");
                        result += '<div id ="imageId" class="info-img" style="background-image:url('+userDetailsArray[i][7]+')"></div>';
								
                    }
                    result +='<div class="info-details">';
                    result +="Address "+" "+'<p>'+userDetailsArray[i][0]+" ,"+userDetailsArray[i][1]+" ,"+" ,"+userDetailsArray[i][2]+" ,"+userDetailsArray[i][3]+" ,"+userDetailsArray[i][4]+'</p>'; 
                    result +="Phone Number  "+'<p>'+userDetailsArray[i][5]+'</p>';
                   var  startDateInUnix = userDetailsArray[i][6];
                    if (startDateInUnix =="0" ||startDateInUnix =="" ){
                         result +="DOB "+'<p>'+""+'</p>';
                    }else{
                        var dateFromDb = parseInt(startDateInUnix);
                        var d = new Date(dateFromDb * 1000);
                        var dd = d.getDate();
                        var mm = d.getMonth() + 1; //January is 0!
                        var yyyy = d.getFullYear();
                        if (dd < 10) {
                            dd = '0' + dd;
                        }
                        if (mm < 10) {
                            mm = '0' + mm;
                        }
                        localDate = (mm + '/' + dd + '/' + yyyy);
                        result +="DOB "+'<p>'+localDate+'</p>';
                    }
                }
           }
           result += '</div>';
           var minUser ='<div class="col-sm-5 dropdown-tbl">';
           minUser +="<table cellpadding='5' cellspacing='0' style='border: 1px solid #dddddd !important;' class='pull-right'>";
           minUser += '<th colspan="2"> Next Of Kin Details</th>';
           for (var i = 0; i < nextOfKin.length; i++) {
                if(nextOfKin != null && nextOfKin[i][4] == userID ) {
                    minUser +='<tr>';
                    minUser +='<td>Email </td>';
                    minUser +='<td>'+nextOfKin[i][0]+'</td>';
                    minUser +='</tr>';
                    
                    minUser +='<tr>';
                    minUser +='<td>Name </td>';
                    minUser +='<td>'+nextOfKin[i][1]+'</td>';
                    minUser +='</tr>';
                    
                    minUser +='<tr>';
                    minUser +='<td>Phone Number </td>';
                    minUser +='<td>'+nextOfKin[i][2]+'</td>';
                    minUser +='</tr>';
                    
                    minUser +='<tr>';
                    minUser +='<td>Relation </td>';
                    minUser +='<td>'+nextOfKin[i][3]+'</td>';
                    minUser +='</tr>';
                }
           }
           minUser +='</table>';
           minUser +='</div>';
        result += '</div>';
           
           
           
           return result+minUser;
       }
/*---------------------------Initial data table calling---------------------------------------------------*/

    if(vm.Values != null) {
        createDataArray(vm.Values, vm.Keys);
    }
    dataTableManipulate();
/*--------------------------Ending Initial data table calling---------------------------------------------*/
    
   
    
       
    
});

