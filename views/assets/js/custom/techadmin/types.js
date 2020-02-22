function editInstitutionTypeForm(app) {
    var form = document.forms['userInstitutionTypeEditForm'];
    form.elements["Id"].value = app.id;
    form.elements["Role"].value = app.name;
    form.elements["Description"].value = app.description;
}

function editDocumentnTypeForm(app) {
    var form = document.forms['userDocumentTypeEditForm'];
    form.elements["documentTypeId"].value = app.documentTypeId;
    form.elements["documetName"].value = app.documentTypename;
}

function editAddressTypeForm(app) {
    var form = document.forms['userAddressTypeEditForm'];
    form.elements["addressId"].value = app.addressTypeID;
    form.elements["addressName"].value = app.addressName;
}

function editApplicationTypeForm(app) {
    var form = document.forms['userApplicationTypeEditForm'];
    form.elements["applicationId"].value = app.id;
    form.elements["applicationName"].value = app.name;
    form.elements["applicationDescription"].value = app.description;
}
function editApplicantTypeForm(app) {
    var form = document.forms['userApplicantTypeEditForm'];
    form.elements["Id"].value = app.id;
    form.elements["Name"].value = app.name;
    form.elements["Description"].value = app.description;
}
function editLocationTypeForm(app) {
    var form = document.forms['locationTypeEditForm'];
    form.elements["locationId"].value = app.locationTypeId;
    form.elements["locationName"].value = app.name;
    form.elements["locationCode"].value = app.code;
}