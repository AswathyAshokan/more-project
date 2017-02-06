/* Author :Aswathy Ashok */

$().ready(function() {

       var val;
       $(".radio-inline").change(function () {

            val = $('.radio-inline:checked').val();


        });

       $("#taskDoneForm").validate({
         rules: {
            taskName: "required",
            emailAddress: {
                required: true,
                email: true
            },
            phoneNumber: {
                required: true,
                minlength : 10
            },
            password: {
                required: true,
                minlength: 8
            },
            confirmpassword: {
                required: true,
                equalTo :"#password"
            }
          },

         submitHandler: function() {
            $.ajax({
              url: '/task/add',
              type: 'post',
              datatype: 'json',
              data: $("#taskDoneForm").serialize() + "&loginType=" + val,
              success : function(response) {
                if (response =="true") {
                     window.location = '/task';
                } else {

                }
              },
             error: function (request,status, error) {

                   console.log(error);
            }
          });

         }

       });

});