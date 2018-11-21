 
function newBtnClose(text, parent){
    var btn = $("<button/>").addClass("pull-right close ml-3");
    btn.html(text);
    btn.attr("id", "btn-close");
    btn.attr("title", "Close the task");
    btn.on("click", function () {
        var div = $(this).parents(parent);
        div.remove();
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
    });
    return btn
}

$("li", "#list-todo").each(function (i, elm) {
    newBtnDone("&#x2713;", "li").appendTo($(elm));
    newBtnClose("&#215;", "li").appendTo($(elm));
});

$("li", "#list-done").each(function (i, elm) {
    newBtnClose("&#215;", "li").appendTo($(elm));
});

function addTodoItem() {
    var lstTodoItem = $("<li/>");
    lstTodoItem.addClass("nav-link")

    var inputValue = $("#ipt-todo-item").val();
    if (inputValue === '') {
        alert("Please enter task name!");
    } else {
        lstTodoItem.append($("<span>" + inputValue + "</span>"));
        lstTodoItem.appendTo($("#list-todo"));
    }
    $("#ipt-todo-item").val("");

    newBtnDone("&#x2713;", "li").appendTo(lstTodoItem);
    newBtnClose("&#215;", "li").appendTo(lstTodoItem);
}

$("#btn-add-item").on("click", addTodoItem);
$("#ipt-todo-item").keypress(function(e){
    if (e.which == 13) {
        addTodoItem();
    }
});
