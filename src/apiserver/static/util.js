function newCompleteElement(id) {
    let input = document.createElement("input");
    input.type = "checkbox";
    id = "sub-task-complete-" + id
    input.id = id;
    input.classList.add("not-show");

    let label = document.createElement("label");
    label.for = id;
    label.setAttribute("aria-hidden", "true");

    let div = document.createElement("div");
    div.classList.add("complete");

    div.appendChild(input);
    div.appendChild(label);
    return div
}

function newSubTaskItemElement(subTask) {
    console.log(subTask);
    let id = subTask.id;
    let c = newCompleteElement(id);
    let text = document.createElement("textarea");
    text.classList.add("edit", "sub-task-edit");
    id = "sub-task-edit-" + id;
    text.id = id;
    text.setAttribute("rows", "1");
    text.value = subTask.title;
    let label = document.createElement("label");
    label.setAttribute("for", id);
    let delIcon = document.createElement("i");
    delIcon.classList.add("fa", "fa-times");
    let del = document.createElement("div");
    del.classList.add("del-sub-task");
    del.appendChild(delIcon);
    let e = document.createElement("div");
    e.classList.add("detail-item", "no-margin-bottom", "sub-task-item");
    e.appendChild(c);
    e.appendChild(text);
    e.appendChild(label);
    e.appendChild(del);
    return e;
}

function clearCreateSubTask() {
    document.getElementById("create-sub-task").value = "";
}
