$(document).ready(function(){

    const BASE_URL = "http://localhost:4000/";

    /** Personal starts here**/
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
    /** Personal ends here**/

    /** Address starts here**/
    $("#addressForm input").prop("disabled", true);
    $("#addressUpdateBtn, #addressClearBtn").attr("disabled", true);

    $("#addressEditBtn").click(function () {
        $("#addressForm input").prop("disabled", false);
        $("#addressUpdateBtn, #addressClearBtn").attr("disabled", false);
        $(this).attr("disabled", true);
    });

    $("#addressCancelBtn").click(function () {
        location.href = BASE_URL + "users/student/profile/address";
    });
    /** Address ends here**/


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
