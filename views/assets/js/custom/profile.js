$(document).ready(function(){

    $("#profileForm input").prop("disabled", true);

    $("#editButton").click(function () {
        $("#profileForm input").prop("disabled", false);
    });

    $("#clearButton").click(function () {
        $("#profileForm first").empty()
       // document.getElementById("myForm").reset();
    })

});