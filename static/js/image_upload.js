document.onpaste = function(event){

    var items = (event.clipboardData || event.originalEvent.clipboardData).items;

    for (var i = 0 ; i < items.length ; i++) {

        var item = items[i];

        if (item.type.indexOf("image") != -1) {
            var file = item.getAsFile();
            upload_file_with_ajax(file);
        }
    }
}

function upload_file_with_ajax(file){

    if (!confirm("Upload image from clipboard ?")) {
        return
    }

    let formData = new FormData();
    formData.append('upload_images', file);

    $.ajax('/upload' , {

        type: 'POST',
        contentType: false,
        processData: false,
        data: formData,
        error: function() {
            console.log("error");
        },
        success: function(res) {
            alert("File uploaded )")
            console.log("ok");
        }
    });
}