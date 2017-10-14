$(document).ready(function() {


    $( function() {
       $('.datepicker').datepicker({
           changeYear: true,
           yearRange: '1900:' + '2050',
       });
     } );


    $(".multiple-drop").select2();




    //  $('.select2-selection').on('click', function () {

    //     var total=$(this).find("ul > li.select2-selection__choice").length;

    //     if (total > 1) {
    //        $(this).addClass("no-background")
    //     }
    //     else{
    //         $(this).removeClass("no-background")
    //     }
    // });

    // $(".edit-account .dis-txt").prop( "disabled", true );


});