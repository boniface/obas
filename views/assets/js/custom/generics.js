const LOCATION_RESTAPI = "http://localhost:4000/location/api/";
const INSTITUTION_RESTAPI = "http://localhost:4000/institution/api/";

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
