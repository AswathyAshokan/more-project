$(document).ready(function() {

    var pageId = "#" + document.URL.substr(url.lastIndexOf('/') + 1);

    $(pageId).addClass('active').siblings().removeClass('active');

} );
