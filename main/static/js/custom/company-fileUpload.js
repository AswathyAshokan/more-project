console.log(vm);

var count =1;
var fileCount =1;
var fileNameCount =1;
var file ="";
var value ="";
var fileValue ="";
var fileUpload =false;
 var downloadURL="";
var formData="";
var companyTeamName = vm.CompanyTeamName;



$(function(){
    if (vm.PageType =="edit"){
        var folderNameReal ="folderName";
        var fileNameReal ="fileName";
        var fileUpload ="fileUploads";
        var spanId="mainSpan";
        var fileSpan ="fileSpan";
        
      document.getElementById('companyUploadButton').style.visibility = 'hidden';
        $("#container1").append("<div class='form-group clearfix upload-section' id='container'>"+"<div class='creat-folder' style='display:none;'>"+"Folder Name"+"  "+"<select id='folderName'  style='width: 219px;margin-left: 10px;' >"+"</select>"+"<span id="+spanId+" style='position: absolute;margin-top: 25px;'></span>"+"</div>"+
     " <div class='folder'>" +"<span class='folder-name' style='margin: 0px;'>"+"<span></span>"+"</span>"+"<div class='file-name' style='display: -webkit-box;'>"+"File Name  "+"<input class='form-control' id='fileName' type='text' placeholder='File Name' value="+ vm.FileName+"  style='width: 219px;margin-left: 28px;'>"+"<span class='dwnl-btn' style='margin-left: 20px; position: absolute;'><i class='fa fa-download fa-lg' aria-hidden='true' id='view'onclick='viewFile();' ></i>"+"  "+"<button class='btn btn-primary margin-left-15'>Upload file  <input type='file' id="+fileUpload+" onchange='uploadEditedFile("+folderNameReal+","+fileNameReal+","+fileUpload+","+spanId+","+fileSpan+");' >"+"</button>"+"</div>"+"<label  class='margin-left-15'></label>"+"<span id="+fileSpan+" style='position: absolute;margin-top: 25px;'></span>"+"</div>"+"<div class='button-new' style='margin-top: 58px;'>"+"<button class='btn btn-primary margin-left-15' id='save'style='margin-left: 114px;margin-top: -35px; width: 70px;' onclick='saveEdit();' >Save  "+"</button>"+"<button class='btn btn-primary margin-left-15' id='cancel'style='margin-left: 15px;margin-top: -35px; background-color: #13a016;' onclick='saveCancel();' >Cancel  "+"</button>"+"</div>"+"</div>"); 
        var ddlItems = document.getElementById("folderName");
           
            for (var i = 0; i < vm.FolderNameArray.length; i++) {
                var opt = vm.FolderNameArray[i];
                var el = document.createElement("option");
                el.textContent = opt;
                el.value = opt;
                ddlItems.appendChild(el);
            }
         $("#folderName").val(vm.FolderName);
        
        
    }else{
        var folderNameReal ="folderName";
        var fileNameReal ="fileName";
        var fileUpload ="fileUploads";
        var spanId="mainSpan";
        var fileSpan ="fileSpan";
        var statusBar ="statusBar";
        $("#container1").append("<div class='form-group clearfix upload-section' id='container'>"+"<div class='creat-folder' style='display:none;'>"+"<input class='form-control' id='folderName' type='text' placeholder='Folder Name' >"+"<button class='btn btn-primary margin-left-15' onClick='addFolder();'>Create folder</button>"+"<span id="+spanId+" style='position: absolute;margin-top: 25px;'></span>"+"</div>"+
     " <div class='folder'>" +"<span class='folder-name'>"+"<span></span>"+"</span>"+"<div class='file-name'>"+"<input class='form-control' id='fileName' type='text' placeholder='File Name'>"+"<button class='btn btn-primary margin-left-15'>Upload file  <input type='file'name='filename' id="+fileUpload+" onchange='displayFile("+folderNameReal+","+fileNameReal+","+fileUpload+","+spanId+","+fileSpan+","+statusBar+");' >"+"</button>"+"</div>"+"<label  class='margin-left-15'></label>"+"<button class='fa fa-plus-circle no-button-style' onclick='addFile("+folderNameReal+","+spanId+");'>"+"</button>"+"<span id="+fileSpan+" style='position: absolute;margin-top: 25px;'></span>"+"<div id='myProgress' style='display:none;'>"+"<div id="+statusBar+" style='height: 15px; background-color: #4CAF50; width:0%;'>"+"</div>"+"</div>"+"</div>"+"</div>"); 
    }
    
    
   
});


//db configuration
    var config = {
        apiKey: "AIzaSyDME5QGEf2AZd0eJGf5NAzOqKui7RtH4qc",
        authDomain: "passporte-b9070.firebaseapp.com",
        databaseURL: "https://passporte-b9070.firebaseio.com",
        projectId: "passporte-b9070",
        storageBucket: "passporte-b9070.appspot.com",
        messagingSenderId: "196354561117"
    };
    firebase.initializeApp(config);

/*Function for creating Data Array for data table*/
$().ready(function() {
    saveCancel=function() {
        window.location ="/" + companyTeamName + "/companyFileUpload";
    }
    
    saveUpdateNew=function() {
        $("#companyUploadButton").attr('disabled', false);
        window.location ="/" + companyTeamName + "/companyFileUpload";
    }
    
    //file upload on edit 
    
    
     uploadEditedFile=function(textboxId,fileTextBox,fileName,folderSpan,fileSpan) {
         
         console.log("display file");
         console.log("text1",textboxId.id);
         console.log("text2",fileTextBox.id);
         console.log("file id",fileName.id);
         console.log("span folder",folderSpan.id);
         console.log("span file",fileSpan.id);

         var folderId=textboxId.id;
         var fileId =fileTextBox.id;
         var uploadedfile=fileName.id;
         var folderSpanId=folderSpan.id;
         var fileSpanId =fileSpan.id;
        file = $('#'+uploadedfile).prop('files')[0];
//         file    = document.querySelector('input[type=file]').files[0];
         var reader  = new FileReader();
         reader.onloadend = function () {
         }
         if (file) {
             reader.readAsDataURL(file);
         } else {
         }
         var e = document.getElementById(folderId);
         var strOptions = e.options[e.selectedIndex].value;
         var value =strOptions;
         fileValue =document.getElementById(fileId).value;
         document.getElementById(folderId).required = true;
         document.getElementById(fileId).required = true;
         if (value.length !=0 &&fileValue.length !=0 ){
             document.getElementById(folderSpanId).innerHTML = "";
             document.getElementById(fileSpanId).innerHTML = "";
             console.log("value",value);
             console.log("file value",fileValue);
             console.log("u1");
             console.log("uploaded file",file);
             var now = new Date();
             var datetime = now.getFullYear()+'/'+(now.getMonth()+1)+'/'+now.getDate();
             datetime += ' '+now.getHours()+':'+now.getMinutes()+':'+now.getSeconds();
             unixDateTime = Date.parse(datetime)/1000;
             var tempFileName = file.name.replace(/\s/g, '');
             var uploadDocumentOriginal =
             firebase.storage().ref().child('CompanyDocuments/'+value+'/'+fileValue+'/'+tempFileName+unixDateTime).put(file);
         uploadDocumentOriginal.on('state_changed', function(snapshot){
             var progress = (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
             console.log('Upload is ' + progress + '% done');
         }, function(error) {
             // Handle unsuccessful uploads
         }, function() {
             // Handle successful uploads on complete
             // For instance, get the download URL: https://firebasestorage.googleapis.com/...
              downloadURL = uploadDocumentOriginal.snapshot.downloadURL;
             console.log("download url",downloadURL);
             fileUpload =true;
              if(fileUpload ){
                 
                  console.log("insideeee");
                   formData = $("#CompanyFileUpload").serialize()+"&fileName=" + fileValue+"&folderName=" + value+"&downloadUrl=" + downloadURL;
                  console.log("file name",fileValue);
                  console.log("folder name",value);
                  console.log("download url",downloadURL);
                  console.log("form data",formData);
                  if (downloadURL.length !=0){
                      alert("successfully uploaded");
//                      
                  }
              }
         });
         
     }else{
         if(value.length =="0" &&fileValue.length =="0"){
             document.getElementById(folderSpanId).innerHTML = " folder name required";
             document.getElementById(fileSpanId).innerHTML = " file name required";
         }
         else {
             if (value.length =="0"){
                 document.getElementById(folderSpanId).innerHTML = " folder name required";
             }else{
                 document.getElementById(fileSpanId).innerHTML = " file name required";
             }
         }
     }
     }
    
    
    

    //file upload
     displayFile=function(textboxId,fileTextBox,fileName,folderSpan,fileSpan,statusBarId) {
          var status = document.getElementById("myProgress");
         status.style.display = "block";
         console.log("display file");
         console.log("text1",textboxId.id);
         console.log("text2",fileTextBox.id);
         console.log("file id",fileName.id);
         console.log("span folder",folderSpan.id);
         console.log("span file",fileSpan.id);

         var folderId=textboxId.id;
         var fileId =fileTextBox.id;
         var uploadedfile=fileName.id;
         var folderSpanId=folderSpan.id;
         var fileSpanId =fileSpan.id;
         var statusBarId =statusBarId.id;
        file = $('#'+uploadedfile).prop('files')[0];
//         file    = document.querySelector('input[type=file]').files[0];
         var reader  = new FileReader();
         reader.onloadend = function () {
         }
         if (file) {
             reader.readAsDataURL(file);
         } else {
         }
         value =document.getElementById(folderId).value;
         value="test";
         fileValue =document.getElementById(fileId).value;
         document.getElementById(folderId).required = true;
         document.getElementById(fileId).required = true;
         if (value.length !=0 &&fileValue.length !=0 ){
                console.log("value",value);
         console.log("file value",fileValue);
          console.log("u1");
         console.log("uploaded file",file);
         var now = new Date();
         var datetime = now.getFullYear()+'/'+(now.getMonth()+1)+'/'+now.getDate();
         datetime += ' '+now.getHours()+':'+now.getMinutes()+':'+now.getSeconds();
         unixDateTime = Date.parse(datetime)/1000;
         var tempFileName = file.name.replace(/\s/g, '');
         var uploadDocumentOriginal =
             firebase.storage().ref().child('CompanyDocuments/'+value+'/'+fileValue+'/'+tempFileName+unixDateTime).put(file);
         uploadDocumentOriginal.on('state_changed', function(snapshot){
             var progress = (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
             console.log('Upload is ' + progress + '% done');
             var elem = document.getElementById(statusBarId); 
             var width = 1;
             var id = setInterval(frame, 10);
             function frame() {
                 if (width >= 100) {
                     clearInterval(id);
                 } else {
                     width++; 
                     elem.style.width = progress + '%';
                 }
             }
            
         }, function(error) {
             // Handle unsuccessful uploads
         }, function() {
             // Handle successful uploads on complete
             // For instance, get the download URL: https://firebasestorage.googleapis.com/...
              downloadURL = uploadDocumentOriginal.snapshot.downloadURL;
             console.log("download url",downloadURL);
             fileUpload =true;
              if(fileUpload ){
                  document.getElementById(folderSpanId).innerHTML = "";
                  document.getElementById(fileSpanId).innerHTML = "";
                  console.log("insideeee");
                  formData = $("#CompanyFileUpload").serialize()+"&fileName=" + fileValue+"&folderName=" + value+"&downloadUrl=" + downloadURL;
                  console.log("file name",fileValue);
                  console.log("folder name",value);
                  console.log("download url",downloadURL);
                  console.log("form data",formData);
                  if (downloadURL.length !=0){
                      $.ajax({
                          url:'/'+ companyTeamName + '/companyFileUpload/add',
                          type:'post',
                          datatype: 'json',
                          data: formData,
                          success : function(response){
                              if(response == "true"){
                                  alert("successfully uploaded");
                              } else {
                              }
                          },
                          error: function (request,status, error) {
                              console.log(error);
                          }
                      });
                  }
              }
         });
         
     }else{
         if(value.length =="0" &&fileValue.length =="0"){
             //document.getElementById(folderSpanId).innerHTML = " folder name required";
             document.getElementById(fileSpanId).innerHTML = " file name required";
         }
         else {
             if (value.length =="0"){
                 document.getElementById(folderSpanId).innerHTML = " folder name required";
             }else{
                 document.getElementById(fileSpanId).innerHTML = " file name required";
             }
         }
     }
     }
   
   viewFile= function () {
       if(vm.DownloadUrl !=""){
           window.location =   vm.DownloadUrl;
       } else{
           $("#noDocument").modal();
       }
       
   }
   saveEdit= function () {
       if (formData ==""){
           var e = document.getElementById("folderName");
           var strOptions = e.options[e.selectedIndex].value;
           var value =strOptions;
           var fileValue =document.getElementById("fileName").value;
           var downloadLink="";
           formData = $("#CompanyFileUpload").serialize()+"&fileName=" + fileValue+"&folderName=" + value+"&downloadUrl=" + downloadLink;;
       }
       
       $.ajax({
           url:'/'+ companyTeamName +'/'+vm.DocumentId+ '/companyFileUpload/EditWithoutChange',
           type:'post',
           datatype: 'json',
           data: formData,
           success : function(response){
               if(response == "true"){
                    window.location ="/" + companyTeamName + "/companyFileUpload";
               } else {
               }
           },
           error: function (request,status, error) {
               console.log(error);
           }
       });
       
   }
    
    
    
    addFile= function (folderName,folderSpanId){
        console.log("add file");
        fileNameCount=fileNameCount+1;
        var fileNewName ="file"+fileNameCount;
        var uploadedFile="uploaded"+fileNameCount;
        var folderNewName =folderName.id;
        var spanFolderName =folderSpanId.id;
        var spanFileId="span"+fileNameCount;
        var statusBar ="status"+fileNameCount;
        var mainStatus ="mainStatus"+fileNameCount;
        $("#container").append(
     " <div class='folder'>" +"<span class='folder-name'>"+"<span></span>"+"</span>"+"<div class='file-name'>"+"<input class='form-control' id="+fileNewName+" type='text' placeholder='File Name' style='width: 198px;'>"+"<button class='btn btn-primary margin-left-15'>Upload file  <input type='file' id="+uploadedFile+" onchange='displayFile("+folderNewName+","+fileNewName+","+uploadedFile+","+spanFolderName+","+spanFileId+","+statusBar+");' >"+"</button>"+"</div>"+"<span id="+spanFileId+" style='position: absolute;margin-top: 25px;'></span>"+"<div id='myProgress'>"+"<div id="+statusBar+" style='height: 15px; background-color: #4CAF50; width:0%;'>"+"</div>"+"</div>"+"</div>"+"<label  class='margin-left-15'></label>"); 
        
    }  
   addFileDup= function (id,folderName,folderSpan){
       console.log("hhhhh",id.id);
       var name =id.id;
      
       console.log("folder name",folderName);
       fileCount =fileCount+1;
       var uplodedFolder =folderName.id;
       var fileId ='fileNamei'+fileCount;
       var uploadedFile='upload'+fileCount;
       var folderSpanId=folderSpan.id;
       var fileSpanId ='spanNewi'+fileCount;
       var statusBar='statusbari'+fileCount;
         $("#"+name).append(
     " <div class='folder'>" +"<span class='folder-name'>"+"<span></span>"+"</span>"+"<div class='file-name'>"+"<input class='form-control' id="+fileId+" type='text' placeholder='File Name'>"+"<button class='btn btn-primary margin-left-15'>Upload file<input type='file' id="+uploadedFile+" onchange='displayFile("+uplodedFolder+","+fileId+","+uploadedFile+","+folderSpanId+","+fileSpanId+","+statusBar+");'></button>"+"<span id="+fileSpanId+" style='position: absolute;margin-top: 25px;'></span>"+"<div id='myProgress'>"+"<div id="+statusBar+" style='height: 15px; background-color: #4CAF50; width:0%;'>"+"</div>"+"</div>"+"<label  class='margin-left-15'></label>"); 
   }  
   
   addFolder= function (){
        count = count + 1;
       var containerId = 'container'+count;
       var folderName ='folderName'+count;
       var fileName ='fileName'+count;
       var uploaded ='uploadedNewFile'+count;
       var spanFolderId ="spanIdFolder"+count;
       var spanFileId ="spanIdFile"+count;
       $("#container1").append("<div class='form-group clearfix upload-section' id="+containerId+">"+ "<div class='creat-folder' style='display:none;'>"+"<input class='form-control' id="+folderName+" type='text' placeholder='Folder Name'>"+"<span id="+spanFolderId+" style='position: absolute;margin-top: 25px;'></span>"+"</div>"+
                                 
     " <div class='folder'>" +"<span class='folder-name'>"+"<span></span>"+"</span>"+"<div class='file-name'>"+"<input class='form-control' id="+fileName+" type='text' placeholder='File Name'>"+"<button class='btn btn-primary margin-left-15'>Upload file  <input type='file' id="+uploaded+" onchange='displayFile("+folderName+","+fileName+","+uploaded+","+spanFolderId+","+spanFileId+");'>"+"</button>"+"<span id="+spanFileId+" style='position: absolute;margin-top: 25px;'></span>"+"</div>"+"<label  class='margin-left-15'></label>"+"<button class='fa fa-plus-circle no-button-style' onclick='addFileDup("+containerId+","+folderName+","+spanFolderId+");'></button>"+
                                " </div>"+"</div>"); 
       console.log("iddddddddd",'container'+count);
   }    
});



