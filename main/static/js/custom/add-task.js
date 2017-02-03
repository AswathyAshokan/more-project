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
                    //alert(val);
                    // console.log("fgsg",val);
                    //var form_data = $("#taskDoneForm").serialize();
                    //console.log(form_data);
                    //form_data.append("loginType", val)
                    //console.log(form_data);
                    $.ajax({
                                   url: '/task',
                                   type: 'post',
                                   datatype: 'json',


                            //name: $('#name').val(),
                            //phoneNumber: $('#phoneNumber').val(),
                            //emailAddress: $('#emailAddress').val(),
                            //address: $('#address').val(),
                            //state: $('#state').val(),
                            //zipcode: $('#zipcode').val(),
                            //request: serverName
                          data: $("#taskDoneForm").serialize() + "&loginType=" + val,
                          success : function(response) {
                            },
                         error: function (request,status, error) {
                             console.log(error);
                        }
                    });


            }

       });

});
