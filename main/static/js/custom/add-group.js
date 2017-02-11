/*Created By Farsana*/
console.log(vm.GroupName);
console.log(vm.PageType);
console.log("group id",vm.GroupId);
$().ready(function() {
    function getMemberName(sel) {
        alert(sel.options[sel.selectedIndex].text);
    }
    
    if(vm.PageType == "edit"){        
            document.getElementById("groupName").value = vm.GroupNameToEdit;
            document.getElementById("addUser").value = vm.GroupMembers;
            //document.getElementById("addUser").value = vm.GroupKey;
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
                    selectedUserIds : "required"
	},
	messages: {
		            groupName:"please enter group name ",
                    selectedUserIds:"please fill this column"

	},
	submitHandler: function(){//to pass all data of a form serial
         var formData = $("#addgroupForm").serialize();
        if (vm.PageType == "edit"){

	                
                    var groupId = vm.GroupId;
                	         $.ajax({

                		    	url:'/group/'+ groupId  +'/edit',
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
                var selectedUserIds = $("#selectedUserIds").val();
                var selectedUsersNames = [];
                
                $("#selectedUserIds option:selected").each(function () {
                    var $this = $(this);
                    if ($this.length) {
                        var selectedUsersName = $this.text();
                        selectedUsersNames.push(selectedUsersName);
                    }
                });
                
                console.log(selectedUsersNames);
                
               
                
                
               
                
                
                /*var data = $('#addgroupForm').serializeArray();
                data.push({name: 'selectedUsersNames', value: selectedUsersNames});*/
                
                 console.log(formData);
                for(i = 0; i < selectedUsersNames.length; i++) {
                    formData = formData+"&selectedUserNames="+selectedUsersNames[i];
                }
                
                
                console.log(typeof(formData));

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




























