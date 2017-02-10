/*Created By Farsana*/

$().ready(function() {
    if(vm.PageType == "edit"){        
            document.getElementById("groupName").value = vm.GroupName;
            document.getElementById("addUser").value = vm.GroupMembers;
            /*var array1 = new Array();
            var counter = 0;
            array1 = vm.GroupMembers;
            for (var i = 0; i < array1.length; i++) {
                foo = array1[i].split(" ");
                for (var j = 0; j < foo.length; j++) {
                    document.getElementById('addUser' + counter).value = foo[j];
                    counter++;

                }
            }*/
               
    }

	$("#addgroupForm").validate({

	  rules: {
		        	groupName: "required",
                    addUser : "required"
	},
	messages: {
		            groupName:"please enter group name ",
                    addUser:"please fill this column"

	},
	submitHandler: function(){//to pass all data of a form serial
         if (vm.PageType == "edit"){

	           var formData = $("#addgroupForm").serialize() 
               var groupId = GroupId;
                    $.ajax({

                		  url:'/group/'+ groupId +'/edit',
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
                 var formData = $("#addgroupForm").serialize();
                 //var values = $('#addUser').val();
                 
                 console.log(formData);

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




























