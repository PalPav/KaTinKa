$(document).ready(function () {

    let $add_tag_area = $("#add_tags_area")
    let $add_tag_button = $("#add_tag_btn")
    let $image_tags_list = $("#image_tags_list")
    let $tag_input = $("#tag_input")

    $add_tag_button.on("click", function (event) {
        let name = $tag_input.val()
        if (name.length < 2) {
            return
        }

        $.ajax({
            method: "PUT",
            url: "/tag",
            contentType: 'application/json',
            data: JSON.stringify({
                name: name,
            }),
            dataType: "JSON"
        }).done(function (data) {
            console.log("success", data)
            $tag_input.trigger("input")
        }).fail(function (data) {
            console.log("fail", data)
        }).always(function () {
            $add_tag_button.prop('disabled', false);
            console.log("finished")
        })
    })

    $tag_input.on('input', function (e) {
        let query = $tag_input.val()

        $add_tag_button.prop('disabled', query.length < 2);

        if (query.length < 2) {
            return
        }

        $.ajax({
            method: "GET",
            url: "/tag/search",
            data: {
                query: query
            },
            dataType: "JSON"
        }).done(function (data) {
            buildTagList(data.data)
            console.log("success", data)
        }).fail(function (data) {
            console.log("fail", data)
        }).always(function () {
            console.log("finished")
        })
    });

    $('#add_tags_area').on('click', '.assign-tag-btn', function (event) {
        let $btn = $(this)
        $.ajax({
            method: "POST",
            url: "/tag/assign",
            contentType: 'application/json',
            data: JSON.stringify({
                image_id: page_image_id,
                tag_id: $btn.parent().data("tag_id"),
            }),
            dataType: "JSON"
        }).done(function (data) {
            console.log("success", data)
            $image_tags_list.prepend('<li class="list-group-item">' +
                data.data.name +
                '<button type="button" data-tag_id="' + data.data.id + '" class="btn btn-danger btn-sm float-end">&times;</button></li>')
            $btn.parent().remove()

        }).fail(function (data) {
            console.log("fail", data)
        }).always(function () {
            console.log("finished")
        })
    })

    $image_tags_list.on('click', 'button', function (event) {
        let $btn = $(this)
        $btn.prop('disabled', true);
        $.ajax({
            method: "POST",
            url: "/tag/detach",
            contentType: 'application/json',
            data: JSON.stringify({
                image_id: page_image_id,
                tag_id: $btn.data("tag_id"),
            }),
            dataType: "JSON"
        }).done(function (data) {
            console.log("success", data)
            $btn.parent().remove()
            $tag_input.trigger("input")
        }).fail(function (data) {
            $btn.prop('disabled', false);
            console.log("fail", data)
        }).always(function () {
            console.log("finished")
        })
    })


    function buildTagList(data) {
        $add_tag_area.html("")
        $.each(data, function (index, item) {
            if ($image_tags_list.find("button[data-tag_id='" + item.id + "']").length) {
                return
            }

            $add_tag_area.append('<li class="list-group-item" data-tag_id="' + item.id + '">' +
                item.name +
                '<button type="button" class="assign-tag-btn btn btn-success btn-sm float-end">+</button></li>')
        })
    }
})