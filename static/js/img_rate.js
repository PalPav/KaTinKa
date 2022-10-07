$(document).ready(function(){
    $(".hit-btn").click(function (event) {
        let $button = $(this)
        let id = $button.data('image-id')
        $button.prop('disabled', true)
        $button.html("...")
        $.ajax({
            method: "POST",
            url:"/image/" + id + "/hit",
        }).done(function (data) {
            console.log("success", data)
            $button.html("&check;")
        }).fail(function (data) {
            console.log("fail", data)
            $button.html("...")
        }).always(function () {
            console.log("finished")
        })

    });
})