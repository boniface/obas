$(document).ready(function () {

    /** Distict starts here **/

    $("#province").change(function() {
        const provinceId = $(this).val();
        let districtElement = $('#district');
        let townElement = $('#town');
        getDropDownElement(townElement, 'Town');
        populateDropDown(districtElement, provinceId);
    });

    $('#district').change(function() {
        const districtId = $(this).val();
        let townElement = $('#town');
        populateDropDown(townElement, districtId);
    });

    /** District ends here **/

    /**institution Location start here**/
    $("#myInstitutionType").change(function() {
        const myInstitutionTypeId = $(this).val();
        let myInstitutionType = $('#myInstitutionType');
        let myInstitutionElement = $('#myInstitution');
        getDropDownElement(myInstitutionType, 'Institutions');
        popullateInstitutionDrop(myInstitutionElement, myInstitutionTypeId);

    });

});

function editForm(event) {
    var form = document.forms['siteEditForm'];
    form.elements["Id"].value = event.id;
    form.elements["Name"].value = event.name;
    form.elements["Description"].value = event.description;
}
function editLocationForm(app) {

    var form = document.forms['locationEditForm'];
    form.elements["LocationId"].value = app.locationId;
    form.elements["Name"].value = app.name;
    form.elements["Longitude"].value = app.longitude;
    form.elements["Latitude"].value = app.latitude;
}