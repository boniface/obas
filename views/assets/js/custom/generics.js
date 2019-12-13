const LOCATION_RESTAPI = "http://localhost:4000/location/api/";
const INSTITUTION_RESTAPI = "http://localhost:4000/institution/api/";
/**institution location**/
const INSTITIONLOCATION_RESTAPI =""

let populateDropDown = function(element, locationId) {
    let dropDownElement = getDropDownElement(element, 'Town');
    if (locationId) {
        const url = LOCATION_RESTAPI + "getforparent/" + locationId;
        $.get(url, function(data) {
            $.each(data, function (key, value) {
                let option = new Option(value.name, value.locationId);
                dropDownElement.append(option);
            });
        });
    }
};

let getDropDownElement = function(element, displayText) {
    let option = '<option value="" disabled selected>Select '+ displayText +'</option>';
    element.empty();
    element.append(option);
    return element;
};

/**Institution Location**/
let popullateInstitutionDrop=function (element,institutionId) {
    let dropDownElement = getDropDownElement(element, 'Institution');
    if(institutionId){
            const url = INSTITUTION_RESTAPI+"getinstitutions/"+institutionId;
        $.get(url, function(data) {
            $.each(data, function (key, value) {
                let option = new Option(value.name, value.locationId);
                dropDownElement.append(option);
            });
        });

    }

}