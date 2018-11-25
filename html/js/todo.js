
var url = "http://127.0.0.1:8080";

function info(atype, content) {
    var alert = $("<div/>").addClass("alert").addClass(atype).text(content);
    alert.delay(2000).slideUp(400, function() {
        $(this).alert('close');
    });
    return alert;
}

function newBtnClose(text, parent){
    var btn = $("<button/>").addClass("pull-right close ml-3");
    btn.html(text);
    btn.attr("id", "btn-close");
    btn.attr("title", "Close the task");
    btn.on("click", function () {
        var div = $(this).parents(parent);

        id = div.attr("item-id");
        $.ajax({
            url: url+"/items/" + id,
            type: "DELETE",
            success: function(result) {
                div.delay(100).slideUp(400, function() {
                    div.remove();
                });
            },
            fail: function(result) {
                $("#nav-tabContent").prepend(info("alert-danger", 
                    "fail to delete task #" + id + " from database"));
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

        id = div.attr("item-id");
        task = div.children("span").text();
        $.ajax({
            url: url+"/items/" + id, 
            type: "PUT",
            data: {done: "true", task: task}
            })
            .done(function(result) {
                div.delay(100).slideUp(400, function() {
                    div.children("#btn-done").remove();
                    div.show().appendTo($("#list-done"));
                });
            })
            .fail(function(result) {
                $("#nav-tabContent").prepend(info("alert-danger", 
                    "fail to update status of task #" + id + " in database"));
            });
    });
    return btn
}


$.getJSON( url + "/items")
    .done(function(data) {
        $("#loading-tasks").remove();

        $.each( data, function( key, val ) {
            var item = $("<li/>");
            item.addClass("nav-link");
            item.attr("item-id", val.id);
            item.append($("<span>" + val.task + "</span>"));

            if (val.done) {
                item.appendTo($("#list-done"));
            } else {
                newBtnDone("&#x2713;", "li").appendTo(item);
                item.appendTo($("#list-todo"));
            }

            newBtnClose("&#215;", "li").appendTo(item);
        });
    })
    .fail(function(data) {
        $("#loading-tasks").remove();

        $("#nav-tabContent").prepend(info("alert-danger", "Fail to fetch data from url: " + url));
    });

function addTodoItem() {
    var item = $("<li/>");
    item.addClass("nav-link");

    var inputValue = $("#ipt-todo-item").val();
    if (inputValue === '' || /^\s*$/.test(inputValue)) {
        alert("Please enter task name!");
    } else {
        item.append($("<span>" + inputValue + "</span>"));
        newBtnDone("&#x2713;", "li").appendTo(item);
        newBtnClose("&#215;", "li").appendTo(item);

        $.post(url + "/items", {task: inputValue})
        .done(function(data) {
            val = JSON.parse(data);
            item.attr("item-id", val.id);

            item.appendTo($("#list-todo")).hide().slideDown(400);
        })
        .fail(function(data) {
            alert("fail to save task to database: " + inputValue);
        });

    }
    $("#ipt-todo-item").val("");
}

$("#btn-add-item").on("click", addTodoItem);
$("#ipt-todo-item").keypress(function(e){
    if (e.which == 13) {
        addTodoItem();
    }
});
