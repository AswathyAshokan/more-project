/*created by Aswathy Ashok*/
$(function () {
    //add text box dynamically
     $("#btnAdd").bind("click", function () {
        var div = $("<div class='plus'/>");
        div.html(GetDynamicTextBox(""));
        $("#TextBoxContainer").prepend(div);
    });
    $("body").on("click", ".delete-decl", function () {
        $(this).closest("div").remove();
    });
    
    
});
function GetDynamicTextBox(value) {
    return ' <input class="form-control"  name = "DynamicTextBox"  id=  "DynamicTextBox"  type="text" value = "" />&nbsp;' +
            '<button    class="add-decl">+</button>'
}
