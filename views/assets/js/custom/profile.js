$(document).ready(function(){

    $("#profileForm input").prop("disabled", true);
    $("#updateButton, #clearButton, #cancelButton").attr("disabled", true);



    $("#editButton").click(function () {
        $("#profileForm input").prop("disabled", false);
        $("#updateButton, #clearButton, #cancelButton").attr("disabled", false);
        $(this).attr("disabled", true);
    });

    $("#cancelButton").click(function () {
        $("#profileForm").trigger("reset");
        $("#profileForm input").prop("disabled", true);
        $("#updateButton, #clearButton, #cancelButton").attr("disabled", true);
        $("#editButton").attr("disabled", false);
    });


   // <!-- guardianForm-->
    $("#guardianForm input").prop("disabled",true);
    $("#guardianupdateButton, #guardianclearButton, #guardiancancelButton").attr("disabled",true);


    $("#guardianEditButton").click(function () {
        $("#guardianForm input").prop("disabled",false);
        $("#guardianupdateButton,#guardianclearButton,#guardiancancelButton").attr("disabled",false);
        $("#guardianEditButton").attr("disabled",false);
    });
    // <!-- guardianForm-->
});
