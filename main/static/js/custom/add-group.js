/*Created By Farsana*/
 allValues = new Array();
valzz = new Array();
function getAllSelectUser()
{
 var value;
  var x=document.getElementById("addUser");
      for (var i = 0; i < x.options.length; i++) {
            if(x.options[i].selected){
                 value = x.options[i].value;

                allValues.push("value");
                alert("valuess",valzz)
             }
      }
}


$().ready(function() {
    if(vm.PageType == "2"){        
            document.getElementById("groupName").value = vm.GroupName;
            document.getElementById("addUser").value = vm.GroupMembers;
               
    }

	$("#addgroupForm").validate({

	  rules: {
		        	groupName: "required",
                   
	},
	messages: {
		            groupName:"please enter group name ",
                    //groupMember:"please fill this column"

	},
	submitHandler: function(){//to pass all data of a form serial
         if (vm.PageType == "edit"){

	           var formData = $("#addgroupForm").serialize() + "&addUser="+allValues
                    $.ajax({

                		  url:'/group/:groupkey/edit',
                          type:'post',
                          datatype: 'json',
                          data: formData,
                          //call back or get response here
                          success : function(response){
                			  if(response == "true"){
                                    window.location='/group';
                              }else {
                              }

                          },
                		  error: function (request,status, error) {
                          }


                   });
	        } else {
                 var formData = $("#addgroupForm").serialize() + "&addUser="+allValues
                 //var values = $('#addUser').val();

                     $.ajax({
                            url:'/group/add',
                            type:'post',
                            datatype: 'json',
                            data: formData,
                            //call back or get response here
                            success : function(response){
                                if(response == "true"){

                                     window.location='/group';
                                }else {
                                }
                            },
                            error: function (request,status, error) {
                            }


                     });
	     return false;
        }
    }

	});

});




























