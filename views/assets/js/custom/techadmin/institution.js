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

});

function editForm(event) {
    var form = document.forms['siteEditForm'];
    form.elements["Id"].value = event.id;
    form.elements["Name"].value = event.name;
    form.elements["Description"].value = event.description;
}