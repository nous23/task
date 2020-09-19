
const LeftBarOptionActive = "active";
// timezone offset in milliseconds
const timezoneOffset = new Date().getTimezoneOffset() * 60000;



function leftBarEvent() {
    let es = document.getElementsByClassName("left-bar-option");
    let i;
    for (i = 0; i < es.length; i++) {
        let e = es[i];
        e.onmouseover = function () {
            this.classList.add("mouseover");
        };
        e.onmouseleave = function () {
            this.classList.remove("mouseover");
        };
        e.onclick = function () {
            let i;
            for (i = 0; i < es.length; i++) {
                es[i].classList.remove(LeftBarOptionActive);
            }
            this.classList.add(LeftBarOptionActive);
            document.getElementById("task-top-name").innerText = this.getElementsByClassName("left-bar-option-text")[0].innerText;
        }
    }
}




function taskListEvent() {
    let es = document.getElementsByClassName("task-item");
    let i;
    for (i = 0; i < es.length; i++) {
        let e = es[i];
        e.addEventListener("click", function () {
            for (i = 0; i < es.length; i++) {
                es[i].classList.remove(LeftBarOptionActive);
            }
            this.classList.add(LeftBarOptionActive);
            let id = this.id.split("-")[1];
            DoRequest(GET, "/task/" + id, null, true, callbackShowRightBar);
        });

        e.addEventListener("mouseover", function () {
            this.classList.add("task-item-mouseover");
        });
        e.addEventListener("mouseleave", function () {
            this.classList.remove("task-item-mouseover");
        });

        let completeId = "complete-" + e.id;
        let input = document.getElementById(completeId);
        let taskid = e.id.split("-")[1];
        input.addEventListener("change", function () {
            updateTaskCompleted(input.checked, taskid);
        });
    }
}

function callbackShowTaskList() {
    if (this.status !== 200) {
        console.error(`list task failed: ${this.responseText}`)
        return
    }
    let tasks = JSON.parse(this.responseText);
    let taskList = document.createElement("div");
    taskList.classList.add("task-list");
    taskList.id = "task-list";
    for (let i = 0; i < tasks.length; i++) {
        let task = tasks[i];

        // input
        let input = document.createElement("input");
        input.type = "checkbox";
        input.classList.add("not-show");
        input.id = "complete-taskid-" + task.id;
        if (task.completed) {
            input.checked = true;
        }
        // label
        let label = document.createElement("label");
        label.setAttribute("aria-hidden", "true");
        label.setAttribute("for", input.id);
        // complete div
        let completeDiv = document.createElement("div");
        completeDiv.classList.add("complete");
        completeDiv.appendChild(input);
        completeDiv.appendChild(label);

        // title div
        let titleDiv = document.createElement("div");
        titleDiv.classList.add("title");
        titleDiv.innerHTML = task.title;

        // task item div
        let taskDiv = document.createElement("div");
        taskDiv.classList.add("task-item");
        taskDiv.id = "taskid-" + task.id;
        taskDiv.appendChild(completeDiv);
        taskDiv.appendChild(titleDiv);

        taskList.appendChild(taskDiv);
    }

    let cte = document.getElementById("create-task-edit");
    cte.value = "";
    let createTask = document.getElementById("create-task");
    let existingTaskList = document.getElementById("task-list");
    if (existingTaskList != null) {
        existingTaskList.remove();
    }
    document.getElementById("list").insertBefore(taskList, createTask);
    taskListEvent();
}

function callbackShowRightBar() {
    if (this.status !== 200) {
        console.error("get task failed: " + this.responseText);
        return;
    }
    let task = JSON.parse(this.responseText);
    let titleEdit = document.getElementById("title-edit");
    if (titleEdit != null) {
        titleEdit.value = task.title;
    }
    let tc = document.getElementById("task-complete");
    tc.checked = task.completed;

    let contentEdit = document.getElementById("content-edit");
    contentEdit.value = task.detail;

    let deadline = document.getElementById("input-date");
    let d = new Date(task.deadline);
    let time = new Date(d.getTime() - timezoneOffset).toISOString();
    time = time.slice(0, time.lastIndexOf(":"));
    deadline.value = time;

    let taskTypeEdit = document.getElementById("task-type");
    taskTypeEdit.value = task.type;

    let describe = document.getElementById("describe");
    let message;
    if (task.completed) {
        let completeTime = new Date(task.end_time);
        message = "完成于 " + completeTime.toLocaleString();
    } else {
        let startTime = new Date(task.start_time);
        message = "开始于 " + startTime.toLocaleString();
    }
    describe.innerHTML = message;

    let rightBar = document.getElementById("right-bar");
    rightBar.setAttribute("data-taskid", task.id);

    DoRequest(GET, `/sub_task/${task.id}`, null, true, callbackShowSubTaskOnPage);
    let cst = document.getElementById("create-sub-task");
    cst.value = "";

    rightBar.classList.remove("not-show");

    let ess = document.getElementsByClassName("edit");
    for (let i = 0; i < ess.length; i++) {
        let e = ess[i];
        e.style.height = "auto";
        e.style.height = e.scrollHeight + 'px';
    }
}

function callbackShowSubTaskOnPage() {
    if (this.status !== 200) {
        console.error(`list sub task failed: ${this.responseText}`);
        return;
    }
    let subTasks = JSON.parse(this.responseText);
    let stElement = document.getElementById("sub_tasks");
    stElement.innerHTML = '';
    if (subTasks == null) {
        return
    }
    for (let i = 0; i < subTasks.length; i++) {
        let st = newSubTaskItemElement(subTasks[i]);
        stElement.appendChild(st)
    }
    autoResize(stElement);
}

function rightBarEvent() {
    let titleEdit = document.getElementById("title-edit");
    if (titleEdit != null) {
        titleEdit.addEventListener("change", function () {
            let rightBar = document.getElementById("right-bar");
            let taskid = rightBar.getAttribute("data-taskid");
            let update = {title: this.value};
            DoRequest(PUT, `/task/${taskid}`, update, true, listIncompleteTask);
        });
    }

    let contentEdit = document.getElementById("content-edit");
    contentEdit.addEventListener("change", function () {
        let rightBar = document.getElementById("right-bar");
        let taskid = rightBar.getAttribute("data-taskid");
        let update = {detail: this.value};
        DoRequest(PUT, `/task/${taskid}`, update, true, callbackVerifyStatus);
    });

    let deadline = document.getElementById("input-date");
    deadline.addEventListener("change", function () {
        let rightBar = document.getElementById("right-bar");
        let taskid = rightBar.getAttribute("data-taskid");
        let update = {deadline: this.value};
        DoRequest(PUT, `/task/${taskid}`, update, true, callbackVerifyStatus);
    });

    let taskTypeEdit = document.getElementById("task-type");
    taskTypeEdit.addEventListener("change", function () {
        let rightBar = document.getElementById("right-bar");
        let taskid = rightBar.getAttribute("data-taskid");
        let update = {type: this.value};
        DoRequest(PUT, `/task/${taskid}`, update, true, callbackVerifyStatus);
    });

    let footerIcons = document.getElementsByClassName("detail-footer-icon");
    for (let i = 0; i < footerIcons.length; i++) {
        footerIcons[i].addEventListener("mouseover", function () {
            this.classList.add("mouseover");
        });
        footerIcons[i].addEventListener("mouseleave", function () {
            this.classList.remove("mouseover");
        });
    }

    let hide = document.getElementById("hide");
    hide.addEventListener("click", hideRightBar);

    let cst = document.getElementById("create-sub-task");
    cst.addEventListener("change", function () {
        let rightBar = document.getElementById("right-bar");
        let taskid = rightBar.getAttribute("data-taskid");
        let subTask = {task_id: taskid, title: this.value}
        DoRequest(POST, "/sub_task", subTask, false, callbackVerifyStatus);
        DoRequest(GET, `/sub_task/${taskid}`, null, true, callbackShowSubTaskOnPage);
        clearCreateSubTask();
    })

    let tc = document.getElementById("task-complete");
    tc.addEventListener("change", function () {
        let taskId = getRightBarTaskId();
        updateTaskCompleted(this.checked, taskId);
        listIncompleteTask();
    })

    autoResize(document.getElementById("right-bar"));
}

function autoResize(root) {
    let es = root.getElementsByClassName("edit");
    for (let i = 0; i < es.length; i++) {
        let e = es[i];
        e.addEventListener('input', function () {
            this.style.height = "auto";
            this.style.height = e.scrollHeight + 'px';
        });
    }
}

function createTaskEvent() {
    let createEdit = document.getElementById("create-task-edit");
    createEdit.addEventListener("change", function () {
        let startTime = new Date();
        let deadline = new Date(startTime.getTime() + 86400000);
        let task = {
            completed: 'false',
            title: this.value,
            detail: "",
            type: "未分类",
            start_time: getLocalTimeString(startTime),
            deadline: getLocalTimeString(deadline),
        }
        DoRequest(POST, "/task", task, false, callbackVerifyStatus);
        listIncompleteTask();
    });
}

function getLocalTimeString(utc) {
    let local = new Date(utc.getTime() - timezoneOffset).toISOString();
    local = local.split(".")[0];
    local.replace("T", " ");
    return local;
}

function deleteTaskEvent() {
    let dlt = document.getElementById("delete");
    dlt.addEventListener("click", function () {
        let rightBar = document.getElementById("right-bar");
        let id = rightBar.getAttribute("data-taskid");
        DoRequest(DELETE, `/task/${id}`, null, false, callbackVerifyStatus);
        listIncompleteTask();
        hideRightBar();
    });
}

function listAllTask() {
    DoRequest(GET, "/tasks", null, true, callbackShowTaskList)
}

function listIncompleteTask() {
    DoRequest(GET, "/tasks?complete=false", null, true, callbackShowTaskList)
}

let o = {
    addEvent: function () {},
    show: function () {},
    init: function () {
        this.addEvent();
        this.show();
    }
};

let leftBarElement = Object.create(o);
leftBarElement.addEvent = leftBarEvent;
let taskListElement = Object.create(o);
taskListElement.show = listIncompleteTask;
let rightBarElement = Object.create(o);
rightBarElement.addEvent = rightBarEvent;
let createTaskElement = Object.create(o);
createTaskElement.addEvent = createTaskEvent;
let deleteTaskElement = Object.create(o);
deleteTaskElement.addEvent = deleteTaskEvent;


let items = [leftBarElement, taskListElement, createTaskElement, rightBarElement, deleteTaskElement];
for (let j = 0; j < items.length; j++) {
    items[j].init();
}

