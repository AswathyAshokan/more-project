var companyTeamName = vm.CompanyTeamName;
var thumbnail = "";
var file ="";
var img ="";
var unixDateTime ="";
var thumbUrl = ""; 
var profileUrl = "";
var tempProfilePicture ="";
var tempThumbPicture ="";
var pictureUploded ="";

var originalUploaded=false;
var thumbUploaded=false;
console.log("profile picture",vm.ProfilePicture);
if (vm.ProfilePicture !=""){
    console.log("inside pic");
    document.getElementById("imageUpload").src = vm.ProfilePicture;
}
//var image = document.getElementsByClassName("imageUpload");
//image.src = vm.ProfilePicture;
//function for displaying image

function displayImage() {
    
//     var filePath = $(this).val();
//    console.log("kkkkkkk",filePath);
//            
    file    = document.querySelector('input[type=file]').files[0];
    var reader  = new FileReader();
    reader.onloadend = function () {
//        document.getElementById("imageUpload").src = reader.result;
//        console.log("d",reader.result);
        document.getElementById('imageUpload').style.backgroundImage="url(reader.result)"; // specify the image path here

    }
    console.log("newww",document.getElementById("imageUpload").src);
    
    if (file) {
        reader.readAsDataURL(file);
    } else {
        console.log("ooooo",vm.ProfilePicture);
        document.getElementById("imageUpload").src = vm.ProfilePicture;
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
    img.height = 100;
    img.width = 100;
}
$().ready(function() {
    
    $('#fileButton').on('change',function ()
        {
        
            var filePath = $(this).val();
            console.log(filePath);
//        document.getElementById('imageUpload').style.backgroundImage="url(filePath)";
            
//                document.getElementById('imageUpload').style.backgroundImage="url(filePath)"; // specify the image path here

        });
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
            document.getElementById("name").removeAttribute('readonly');
            document.getElementById("emailId").removeAttribute('readonly');
            document.getElementById("phoneNumber").removeAttribute('readonly');
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
                emailId: "Please enter Email Id ",
                phoneNumber:"Please enter Phone Number"
            },
            
            submitHandler: function(){//to pass all data of a form serial
                $("#edit-txt").attr('disabled', true);
                var now = new Date();
                var datetime = now.getFullYear()+'/'+(now.getMonth()+1)+'/'+now.getDate(); 
                datetime += ' '+now.getHours()+':'+now.getMinutes()+':'+now.getSeconds();
                unixDateTime = Date.parse(datetime)/1000;
                var tempProfilePicture = file.name.replace(/\s/g, '');
                var uploadTaskOriginal = firebase.storage().ref().child('profilePicturesOfAdmin/original/'+unixDateTime+tempProfilePicture).put(img);
                uploadTaskOriginal.on('state_changed', function(snapshot){
                    var progress = (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
                      console.log('Upload is ' + progress + '% done');
                }, function(error) {
                      // Handle unsuccessful uploads
                }, function() {
                      // Handle successful uploads on complete
                      // For instance, get the download URL: https://firebasestorage.googleapis.com/...
                    var downloadURL = uploadTaskOriginal.snapshot.downloadURL;
                    profileUrl=downloadURL;
                    originalUploaded=true;
                });
                var uploadTaskThumb = firebase.storage().ref().child('profilePicturesOfAdmin/thumbnail/'+unixDateTime+tempThumbPicture).put(img);
                    uploadTaskThumb.on('state_changed', function(snapshot){
                        var progress = (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
                        console.log('Upload is ' + progress + '% done');
                    }, function(error) {
                          // Handle unsuccessful uploads
                    }, function() {
                        var downloadURL1 = uploadTaskThumb.snapshot.downloadURL;
                        thumbUrl=downloadURL1;
                        thumbUploaded=true;
                    });
                var editProfile=  setInterval(function(){
                    console.log("originalUploaded",originalUploaded)
                   console.log("thumbUploaded",thumbUploaded)
                   if(originalUploaded && thumbUploaded ){
                       var formData = $("#adminAccountDetail").serialize()+ "&profilePicture=" + profileUrl+"&profilePicturePath=" + file+"&thumbPicture=" + thumbUrl;
                       console.log("thumb",thumbUrl);
                       console.log("profile",profileUrl);
                       $.ajax({
                           url:'/'+ companyTeamName + '/editProfile',
                           type:'post',
                           datatype: 'json',
                           data: formData,
                           success : function(response){
                               if(response == "true"){
                                   window.location =  '/'+companyTeamName+'/dashBoard';
                               } else {
                                   $('#edit-txt').text("Edit");
                               }
                           },
                           error: function (request,status, error) {
                           }
                       });
                       clearInterval(editProfile);
                   }
                },3000);
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
                 $("#updateAdminPassword").attr('disabled', true);
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