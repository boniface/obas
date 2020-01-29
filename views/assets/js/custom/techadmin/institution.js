$(document).ready(function () {

    /** Institution Location start here **/

    $("#province").change(function() {
        const provinceId = $(this).val();
        let districtElement = $('#district');
        let townElement = $('#town');
        getDropDownElement(townElement, 'Town');
        let districtDropDown = getDropDownElement(districtElement, "District");
        populateLocationDropDown(districtDropDown, provinceId);
    });

    $('#district').change(function() {
        const districtId = $(this).val();
        let townElement = $('#town');
        //let townDropDown = getDropDownElement(institutionElement, "Town");
        let townDropDown = getDropDownElement(townElement, "Town");
        populateLocationDropDown(townDropDown, districtId);
    });


    /** Institution types ends here **/

    $("#institutionTypeAddressDrop").change(function() {
        const institutionTypeAddress = $(this).val();
        let institutionAddressElement = $('#institutionAddressDrop');
        let institutionAddressDropDown = getDropDownElement(institutionAddressElement, "Institution");
        populateInstitutionDropDownByType(institutionAddressDropDown, institutionTypeAddress);
    });

    $("#institutionTypeCourseDrop").change(function() {
        const institutionTypeAddress = $(this).val();
        let institutionAddressElement = $('#institutionCourseDrop');
        let institutionAddressDropDown = getDropDownElement(institutionAddressElement, "Institution");
        populateInstitutionDropDownByType(institutionAddressDropDown, institutionTypeAddress);
    });
    //location
    $("#institutionTypeLocationDrop").change(function() {
        const institutionTypeAddress = $(this).val();
        let institutionAddressElement = $('#institutionLocationDrop');
        let institutionAddressDropDown = getDropDownElement(institutionAddressElement, "Institution");
        populateInstitutionDropDownByType(institutionAddressDropDown, institutionTypeAddress);
    });
});



function editCourseForm(event) {
    var form = document.forms['siteEditForm'];
    form.elements["Id"].value = event.id;
    form.elements["Name"].value = event.name;
    form.elements["Description"].value = event.description;
}

function editLocationForm(app) {
    var form = document.forms['locationEditForm'];
    var myArea = document.getElementById("courseSubjectContent");
    form.elements["LocationId"].value = app.Location.locationId;
    form.elements["Name"].value = app.Location.name;
    form.elements["Longitude"].value = app.Location.longitude;
    form.elements["Latitude"].value = app.Location.latitude;
    myArea.innerHTML="<p> Location Type: "+app.LocationType.name+"<br/> Parent: "+app.ParentLocation.name+"<p>";
}
