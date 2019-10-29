$(document).ready(function(){

    const BASE_URL = "http://localhost:4000/";

    /** student profile pages starts here**/
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
        $("#updateButton, #clearButton, #cancelButton").hide(1000);
        $("#editButton").show(500);
    });

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

    /** District starts here **/
    // $("#provinceForm select").prop("disabled", true);
    // $("#districtForm select").prop("disabled", true);
    // $("#townForm select").prop("disabled", true);
    // $("#townUpdateBtn, #townClearBtn, #townCancelBtn, #districtBtn, #townBtn").hide();
    // $("#townEditBtn").click(function() {
    //     $("#provinceForm select").prop("disabled", false);
    //     $("#districtForm select").prop("disabled", false);
    //     $("#townForm select").prop("disabled", false);
    //     $(this).hide();
    //     $("#townUpdateBtn, #townClearBtn, #townCancelBtn, #districtBtn").show();
    // });
    // $("#townCancelBtn").click(function () {
    //     $("#provinceForm, #districtForm, #townForm").trigger("reset");
    //     $("#provinceForm select").prop("disabled", true);
    //     $("#districtForm select").prop("disabled", true);
    //     $("#townForm select").prop("disabled", true);
    //     $("#townUpdateBtn, #townClearBtn, #townCancelBtn, #districtBtn, #townBtn").hide();
    //     $("#townEditBtn").show();
    // });
    /** District ends here **/

    /** student profile ends here**/
});
