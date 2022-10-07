$(document).ready(function() {
    let $search_select = $('#search_select')
    let $image_container = $('#images_container')
    let $search_button = $('#search_button')

    $search_select.select2({
        width: 'resolve' // need to override the changed default
    });

    $search_button.on('click', function (event) {
        let tags_ids = $search_select.val().map(Number);
        $image_container.html("")

        if (tags_ids.length < 1) {
            return;
        }

        $.ajax({
            method: "GET",
            url: "/image/search",
            data: {
                tags_ids: tags_ids.join(','),
            },
            dataType: "JSON"
        }).done(function (data) {
            buildImagesList(data.data)
        }).fail(function (data) {
            console.log("fail", data)
        }).always(function () {
            console.log("finished")
        })
    })

    function buildImagesList(data) {
        console.log(data)
        $.each(data, function (index, item) {
            let img_element = '<div class="image">'
            if (item.extension === ".webm") {
                img_element += '<video autoPlay loop muted src="/image/'+ item.id +'"></video>'
            } else {
                img_element += '<img src="/image/'+ item.id +'" alt="'+ item.id +'">'
            }
            img_element += '</div>'

            $image_container.append(img_element)
        })
    }
});