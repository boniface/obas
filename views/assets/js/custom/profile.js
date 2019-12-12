$(document).ready(function(){

    const BASE_URL = "http://localhost:4000/users/student/profile/";
    const LOCATION_RESTAPI = "http://localhost:4000/location/api/";

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
        location.href = BASE_URL + "address";
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
        location.href = BASE_URL + "contacts";
    });
    /** Contact ends here**/

    /** Distict starts here **/

    function getTownElement() {
        let townDropDown = $("#town");
        townDropDown.empty();
        townDropDown.append('<option value="" disabled selected>Select Town</option>');
        return townDropDown;
    }

    $("#province").change(function() {
        const provinceId = $(this).val();

        function getDistrictElement() {
            let districtDropDown = $('#district');
            districtDropDown.empty();
            districtDropDown.append('<option value="" disabled selected>Select District</option>');
            return districtDropDown;
        }

        let districtDropDown = getDistrictElement();
        getTownElement();
        if(provinceId != "") {
            const url = LOCATION_RESTAPI + "getforparent/" + provinceId;
            $.get(url, function(districts) {
                $.each(districts, function (key, value) {
                    let option = new Option(value.name, value.locationId);
                    districtDropDown.append(option);
                });
            });
        }
    });

    $('#district').change(function() {
        const districtId = $(this).val();
        let townDropDown = getTownElement();
        if (districtId != "") {
            const url = LOCATION_RESTAPI + "getforparent/" + districtId;
            $.get(url, function(towns) {
                $.each(towns, function (key, value) {
                    let option = new Option(value.name, value.locationId);
                    townDropDown.append(option);
                });
            });
        }
    });

    /** District ends here **/

    /** student profile ends here**/
});
