


const BudgetURL = "155.238.30.101:4000/admin/budget/give-budget";

var table = document.getElementById("myTable");
function GetDocuments(userId,applicationId) {
    var i = 0;
    var row;
    var dType;
    var dView;
    var dDownload;
    var ddate;
    var stats;
    for (var i = table.rows.length - 1; i > 0; i--) {
        table.deleteRow(i);
    }

    const url = "155.238.30.101:4000/admin/applicant/application/" + userId+"/"+applicationId;
    $.get(url, function(data) {
        console.log(data)
        $.each(data, function (key, value) {

            row = table.insertRow(i + 1);

            dType = row.insertCell(0);
            dView = row.insertCell(1);
            dDownload = row.insertCell(2);
            ddate = row.insertCell(3);
            stats = row.insertCell(4);

            //dType.innerHTML = value.Id;
            dType.innerHTML = value.DocumentType;
            dView.setAttribute("class","fa fa-image col-md-3");
            dView.setAttribute("href",value.Doc.url);

            dDownload.setAttribute("class","col-md-5 nc-icon nc-cloud-download-93 icon-bold ");
            dDownload.setAttribute("href",value.Doc.url);

            ddate.innerHTML = value.Doc.date;
            stats.innerHTML = value.Status;

        });
    });
}
function GetDocumentStat(DocumentStatu) {
    document.getElementById("ModifiedBy").value=DocumentStatu.modifiedBy;
    document.getElementById("Comment").value=DocumentStatu.comment;
    document.getElementById("DateTime").value=DocumentStatu.dateTime;
}
function GetApplicationStat(modifier,timeDate,comment) {
    document.getElementById("ModifiedBy").value=modifier;
    document.getElementById("Comment").value=comment;
    document.getElementById("DateTime").value=timeDate;
}
function budgetAllocation() {
    swal({
        title: "Are you sure?",
        text: 'You will not be able to recover this Action.  You are about to allocate   R '+document.getElementById("amount").value+'.00  to Espoir',
        type: "warning",
        showCancelButton: true,
        confirmButtonText: "Yes, proceed !",
        cancelButtonText: "No, cancel !",
        closeOnConfirm: false,
        closeOnCancel: false
    },function (isConfirm) {
        if (isConfirm) {
            const url =BudgetURL+"/"+applicationId+"/"+amount
            $.get(url, function(data){
                if(data==true){
                    swal("Allocated!", "You have successfully made allocation.", "success");
                }else {
                    swal("Error!", "Something went wrong.", "error");
                }
            });

        } else {
            swal("Cancelled", "You have successfully cancelled allocation :)", "error");
        }
    });
}
