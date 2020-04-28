function newCompleteElement(id, completed) {
    let input = document.createElement("input");
    input.type = "checkbox";
    let elementId = "sub-task-complete-" + id
    input.id = elementId;
    input.classList.add("not-show");
    if (completed) {
        input.checked = true;
    }
    input.addEventListener("change", function () {
        let update = {id: id};
        update.completed = input.checked;
        DoRequest(PUT, `/sub_task/${id}`, update, true, callbackVerifyStatus);
    })

    let label = document.createElement("label");
    label.setAttribute("aria-hidden", "true");
    label.setAttribute("for", elementId);

    let div = document.createElement("div");
    div.classList.add("complete");

    div.appendChild(input);
    div.appendChild(label);
    return div
}

function newSubTaskItemElement(subTask) {
    let id = subTask.id;
    let completed = subTask.completed;
    let c = newCompleteElement(id, completed);
    let text = document.createElement("textarea");
    text.classList.add("edit", "sub-task-edit");
    id = "sub-task-edit-" + id;
    text.id = id;
    text.setAttribute("rows", "1");
    text.value = subTask.title;
    let label = document.createElement("label");
    label.setAttribute("for", id);
    let delIcon = document.createElement("i");
    delIcon.classList.add("del-sub-task-icon", "fa", "fa-times");
    delIcon.addEventListener("click", function () {
        DoRequest(DELETE, `/sub_task/${subTask.id}`, null, false, callbackVerifyStatus);
        DoRequest(GET, `/sub_task/${subTask.task_id}`, null, true, callbackShowSubTaskOnPage);
    });
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

function getRightBarTaskId() {
    return document.getElementById("right-bar").getAttribute("data-taskid");
}

function updateTaskCompleted(completed, id) {
    let update = {completed: completed};
    if (completed) {
        let utc = new Date();
        let local = new Date(utc.getTime() - timezoneOffset).toISOString();
        update.end_time = local.split(".")[0];
    }
    DoRequest(PUT, `/task/${id}`, update, false, callbackVerifyStatus);
}

function hideRightBar() {
    let rb = document.getElementById("right-bar");
    rb.classList.add("not-show");
}
