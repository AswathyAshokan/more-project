var companyTeamName = vm.CompanyTeamName;
var thumbnail = "";
var file ="";
var img ="";
var unixDateTime ="";
var thumbUrl = ""; 
var profileUrl = "";
var tempProfilePicture ="";
var tempThumbPicture ="";
 
//function for displaying image
function displayImage() {

    file    = document.querySelector('input[type=file]').files[0];
    var reader  = new FileReader();
    reader.onloadend = function () {
        document.getElementById("imageUpload").src = reader.result;
    }
    console.log("newww",document.getElementById("imageUpload").src);
    if (file) {
        reader.readAsDataURL(file);
    } else {
        document.getElementById("imageUpload").src = "";
    }
        var btntxt = $("#edit-txt").text();
    if (btntxt == 'Edit') {
        $(".edit-account input").prop( "disabled", false );
        $(".edit-account input").toggleClass("dis-txt");
        $('#edit-txt').text("Save");
        $('#edit-txt').attr('type', 'submit');
        return false;
    }
}
//uploading image
var config = {
            apiKey: "AIzaSyDME5QGEf2AZd0eJGf5NAzOqKui7RtH4qc",
            authDomain: "passporte-b9070.firebaseapp.com",
            databaseURL: "https://passporte-b9070.firebaseio.com",
            projectId: "passporte-b9070",
            storageBucket: "passporte-b9070.appspot.com",
            messagingSenderId: "196354561117"
        };
firebase.initializeApp(config);


function resizeImg() {
    console.log("inside");
    img  = document.querySelector('input[type=file]').files[0];
    img.height = 200;
    img.width = 100;
}
$().ready(function() {
    
    document.getElementById("name").value = vm.FirstName;
    document.getElementById("emailId").value = vm.Email;
    document.getElementById("phoneNumber").value = vm.PhoneNo;
    
    //to check the plan and load modal according to plan
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
    
    
    //function for editing form
   $('#edit-txt').on('click', function() {
        var btntxt = $("#edit-txt").text();
        if (btntxt == 'Edit') {
            $(".edit-account input").prop( "disabled", false );
            $(".edit-account input").toggleClass("dis-txt");	
            $('#edit-txt').text("Save");
            $('#edit-txt').attr('type', 'submit');
            return false;
        }
        $("#adminAccountDetail").validate({
            rules: {
                name:"required",
                emailId:{
                    required:true,
                    email:true
                },
                phoneNumber: "required"
            },
            messages: {
                name:"Please enter your Name ",
                emailId: "Please enter Phone Number",
                phoneNumber:"Please enter your Email id"
            },
            
            submitHandler: function(){//to pass all data of a form serial
                
                //image upload
                //var imageNumber =0;
                 var now = new Date();
                var datetime = now.getFullYear()+'/'+(now.getMonth()+1)+'/'+now.getDate(); 
                datetime += ' '+now.getHours()+':'+now.getMinutes()+':'+now.getSeconds();
                unixDateTime = Date.parse(datetime)/1000;
//                tempProfilePicture  =file.name(' ').join('_');
//                console.log("temp",tempProfilePicture);
                 var tempProfilePicture = file.name.replace(/\s/g, '');
                var storageRef = firebase.storage().ref('profilePicturesOfAdmin/original/'+unixDateTime+tempProfilePicture);
                storageRef.put(file);
                function error(err) {
                    console.log("error",err);
                }
                
                //retrieve image from firebase storage
                //tempThumbPicture  =img.name(' ').join('_');
                var tempThumbPicture =img.name.replace(/\s/g, '');
                var storageRef = firebase.storage().ref('profilePicturesOfAdmin/thumbnail/'+unixDateTime+tempThumbPicture);
                storageRef.put(img);
                function error(err) {
                    console.log("error",err);
                }
                
                var displayThumbRef = firebase.storage().ref('profilePicturesOfAdmin/thumbnail/'+unixDateTime+tempThumbPicture);
                setTimeout(function() { displayThumbRef.getDownloadURL().then(function(url) {
                    // Get the download URL for 'images/stars.jpg'
                    // This can be inserted into an <img> tag
                    thumbUrl=url;
                   
                }).catch(function(error) {
                    console.error(error);
                }); }, 3000);
                 var displayProfileRef = firebase.storage().ref('profilePicturesOfAdmin/original/'+unixDateTime+tempProfilePicture);
                setTimeout(function() { displayProfileRef.getDownloadURL().then(function(url) {
                    // Get the download URL for 'images/stars.jpg'
                    // This can be inserted into an <img> tag
                    profileUrl=url;
                    
                }).catch(function(error) {
                    console.error(error);
                }); }, 2000);
                setTimeout(function(){var formData = $("#adminAccountDetail").serialize()+ "&profilePicture=" + profileUrl+"&profilePicturePath=" + file+"&thumbPicture=" + thumbUrl;
                                      console.log("thumb",thumbUrl);
                                      console.log("profile",profileUrl);
                                      $.ajax({
                                          url:'/'+ companyTeamName + '/editProfile',
                                          type:'post',
                                          datatype: 'json',
                                          data: formData,
                    //call back or get response here
                                          success : function(response){
                                              if(response == "true"){
                                                   window.location =  '/login';
//                                                  $('#edit-txt').text("Edit");
//                                                  $(".edit-account input").prop( "disabled", true );
//                                                  $('#edit-txt').attr('type', 'button');
                                              } else {
                                                  $('#edit-txt').text("Edit");
                                              }
                                          },
                                          error: function (request,status, error) {
                                          }
                                      });
                                      return false;},5000);
            }
        });
   });

    //function to change password
    
     $('#updateAdminPassword').on('click', function() {
        $("#adminPasswordChangeModal").validate({
            rules: {
                newPassword:"required",
                confirmpassword:{
                    equalTo : "#newPassword"
                } ,
                oldPassword: {
                required: true,
                remote:{
                    url: '/'+ companyTeamName +"/isOldAdminPasswordCorrect/" + oldPassword,
                    type: "post"
                }
            },
            },
            messages: {
                 oldPassword:{
                     required: "Please enter Old Password ",
                     remote: "The password entered is not correct !!!"
                 },
                newPassword: "Please enter New Password",
                confirmpassword:"Retype password is incorrect"
            },
            submitHandler: function(){//to pass all data of a form serial
                var formData = $("#adminPasswordChangeModal").serialize();
                $.ajax({
                    url:'/'+ companyTeamName +'/changePassword',
                    type:'post',
                    datatype: 'json',
                    data: formData,
                    success : function(response){
                        if(response == "true"){
                            window.location =  '/login';
                        } else {
                            alert("password incorrect");
                        }
                    },
                    error: function (request,status, error) {
                    }
                });
                return false;
            }
        });
    });
});