$(document).ready(function(){

    /** Matric form starts here **/

    function getDropDownElement(element, displayText) {
        let option = '<option value="" disabled selected>Select '+ displayText +'</option>';
        element.empty();
        element.append(option);
        return element;
    }

    $("form#matricInstitutionForm select#province").change(function() {
        const provinceId = $(this).val();
        let districtElement = $("form#matricInstitutionForm select#district");
        let townElement = $("form#matricInstitutionForm select#town");
        let districtDropDown = getDropDownElement(districtElement, 'District');
        getTownDropDown(townElement);
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

    $("form#matricInstitutionForm select#district").change(function() {
        const districtId = $(this).val();
        let townElement = $("form#matricInstitutionForm select#town");
        let townDropDown = getDropDownElement(townElement, 'Town');
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

    $("form#matricInstitutionForm select#institutionType").change(function() {
        const institutionTypeId = $(this).val();
        let institutionElement = $("form#matricInstitutionForm select#institution");
        let institutionDropDown = getDropDownElement(institutionElement, 'Institution');
        if (institutionTypeId != "") {
            const url = INSTUTITION_RESTAPI + "getInstitutionsInLocation/" + institutionTypeId;
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