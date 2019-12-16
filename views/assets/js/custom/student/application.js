$(document).ready(function(){

    /** Matric form starts here **/

    $("form#matricInstitutionForm select#province").change(function() {
        const provinceId = $(this).val();
        let districtElement = $("form#matricInstitutionForm select#district");
        let townElement = $("form#matricInstitutionForm select#town");
        getDropDownElement(townElement, 'Town');

        populateDropDown(districtElement, provinceId);
    });

    $("form#matricInstitutionForm select#district").change(function() {
        const districtId = $(this).val();
        let townElement = $("form#matricInstitutionForm select#town");
        populateDropDown(townElement, districtId);
    });

    $("form#matricInstitutionForm select#institutionType").change(function() {
        const institutionTypeId = $(this).val();
        let institutionElement = $("form#matricInstitutionForm select#institution");
        let institutionDropDown = getDropDownElement(institutionElement, 'Institution');
        if (institutionTypeId) {
            const url = INSTITUTION_RESTAPI + "getInstitutionsInLocation/" + institutionTypeId;
            $.get(url, function(institutions) {
                $.each(institutions, function (key, value) {
                    let option = new Option(value.name, value.id);
                    institutionDropDown.append(option);
                });
            });
        }
    });

    /** Matric form ends here **/

});