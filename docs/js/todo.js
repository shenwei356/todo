
var url = "http://127.0.0.1:8080"

function newBtnClose(text, parent){
    var btn = $("<button/>").addClass("pull-right close ml-3");
    btn.html(text);
    btn.attr("id", "btn-close");
    btn.attr("title", "Close the task");
    btn.on("click", function () {
        var div = $(this).parents(parent);
        div.remove();

        id = div.attr("item-id")
        $.ajax({
            url: url+"/items/" + id,
            type: "DELETE",
            success: function(result) {
                alert("task #" + id + " deleted")
            }
        });
    });
    return btn
}

function newBtnDone(text, parent){
    var btn = $("<button/>").addClass("pull-right close ml-3");
    btn.html(text);
    btn.attr("id", "btn-done");
    btn.attr("title", "Mark as finished");
    btn.on("click", function () {
        var div = $(this).parents(parent);
        div.children("#btn-done").remove();
        div.appendTo($("#list-done"));

        id = div.attr("item-id")
        task = div.children("span").text()
        $.ajax({
            url: url+"/items/" + id, 
            type: "PUT",
            data: {done: "true", task: task}
            })
            .done(function(result) {
                alert("task #" + id + " updated")
            });
    });
    return btn
}

// $("li", "#list-todo").each(function (i, elm) {
//     newBtnDone("&#x2713;", "li").appendTo($(elm));
//     newBtnClose("&#215;", "li").appendTo($(elm));
// });

// $("li", "#list-done").each(function (i, elm) {
//     newBtnClose("&#215;", "li").appendTo($(elm));
// });


$.getJSON( url + "/items")
    .done(function(data) {
        $.each( data, function( key, val ) {
            var item = $("<li/>");
            item.addClass("nav-link");
            item.attr("item-id", val.id);
            item.append($("<span>" + val.task + "</span>"));

            newBtnClose("&#215;", "li").appendTo(item);

            if (val.done) {
                item.appendTo($("#list-done"));
            } else {
                newBtnDone("&#x2713;", "li").appendTo(item);
                item.appendTo($("#list-todo"));
            }
        });
    })
    .fail(function(data) {
        alert( "fail to fetch data from url:" + url );
    });

function addTodoItem() {
    var item = $("<li/>");
    item.addClass("nav-link")

    var inputValue = $("#ipt-todo-item").val();
    if (inputValue === '') {
        alert("Please enter task name!");
    } else {
        item.append($("<span>" + inputValue + "</span>"));
        newBtnDone("&#x2713;", "li").appendTo(item);
        newBtnClose("&#215;", "li").appendTo(item);

        $.post(url + "/items", {task: inputValue})
        .done(function(data) {
            val = JSON.parse(data)
            item.attr("item-id", val.id);
        });

        item.appendTo($("#list-todo"));
    }
    $("#ipt-todo-item").val("");
}

$("#btn-add-item").on("click", addTodoItem);
$("#ipt-todo-item").keypress(function(e){
    if (e.which == 13) {
        addTodoItem();
    }
});
