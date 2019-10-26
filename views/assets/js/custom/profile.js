$(document).ready(function(){

    const BASE_URL = "http://localhost:4000/";

    /** For all student profile pages starts here**/
    $("#profileForm input").prop("disabled", true);
    $("#profileForm select").prop("disabled", true);
    $("#updateButton, #clearButton, #cancelButton").hide();

    $("#editButton").click(function () {
        $("#profileForm input").prop("disabled", false);
        $("#profileForm select").prop("disabled", false);
        $("#updateButton, #clearButton, #cancelButton").show(1000);
        $(this).hide(500);
    });

    $("#cancelButton").click(function () {
        $("#profileForm").trigger("reset");
        $("#profileForm input").prop("disabled", true);
        $("#profileForm select").prop("disabled", true);
        $("#updateButton, #clearButton, #cancelButton").hide();
        $("#editButton").show(500);
    });
    /** profile ends here**/

    /** Address starts here**/
    $("#addressForm input").prop("disabled", true);
    $("#addressUpdateBtn, #addressClearBtn").hide();

    $("#addressEditBtn").click(function () {
        $("#addressForm input").prop("disabled", false);
        $("#addressUpdateBtn, #addressClearBtn").show(1000);
        $(this).hide(500);
    });

    $("#addressCancelBtn").click(function () {
        location.href = BASE_URL + "users/student/profile/address";
    });
    /** Address ends here**/

    /** Contact starts here**/
    $("#contactForm input").prop("disabled", true);
    $("#contactUpdateBtn, #contactClearBtn").hide();

    $("#contactEditBtn").click(function () {
        $("#contactForm input").prop("disabled", false);
        $("#contactUpdateBtn, #contactClearBtn").show(1000);
        $(this).hide(500);
    });

    $("#contactCancelBtn").click(function () {
        location.href = BASE_URL + "users/student/profile/contacts";
    });
    /** Contact ends here**/


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
