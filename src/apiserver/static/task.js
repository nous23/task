
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
            console.log("you are fucking clicking task " + id);
            showDetail(id);
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
            if (input.checked) {
                let utc = new Date();
                let local = new Date(utc.getTime() - timezoneOffset).toISOString();
                updateTask(taskid, {completed: "1", end_time: local.split(".")[0]});
            } else {
                updateTask(taskid, {completed: "0"});
            }
        });
    }
}


function showTaskList() {
    let x = new XMLHttpRequest();
    x.onreadystatechange = function () {
        if (x.readyState === 4 && x.status === 200) {
            console.log("query tasks success");
            console.log(x.responseText);
            let tasks = JSON.parse(x.responseText);
            createTaskList(tasks);
        }
    };
    x.open("GET", "/tasks", true);
    x.send();
}

function createTaskList(tasks) {
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




function showDetail(id) {
    let xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function () {
        if (xmlHttp.readyState === 4 && xmlHttp.status === 200) {
            console.log(xmlHttp.responseText);
            let task = JSON.parse(xmlHttp.responseText);
            showRightBar(task);
        }
    };
    xmlHttp.open("GET", "/task/" + id, true); // false for synchronous request
    xmlHttp.send(null);
}

function updateTask(id, update) {
    console.log("update task " + id);
    console.log(update);
    let xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function () {
        if (xmlHttp.readyState === 4 && xmlHttp.status === 200) {
            if (update.hasOwnProperty("title")) {
                updateTaskTitleOnPage(id, update.title);
            }
        }
    };
    xmlHttp.open("put", "/task/" + id, true);
    xmlHttp.send(JSON.stringify(update));
}

function updateTaskTitleOnPage(id, title) {
    let task = document.getElementById("taskid-" + id);
    if (task == null) {
        return;
    }
    let es = task.getElementsByClassName("title");
    for (let i = 0; i < es.length; i++) {
        es[i].innerHTML = title;
    }
}



function showRightBar(task) {
    let titleEdit = document.getElementById("title-edit");
    titleEdit.value = task.title;

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
    rightBar.classList.remove("not-show");

    let ess = document.getElementsByClassName("edit");
    for (let i = 0; i < ess.length; i++) {
        let e = ess[i];
        e.style.height = "auto";
        e.style.height = e.scrollHeight + 'px';
    }
}

function rightBarEvent() {
    let titleEdit = document.getElementById("title-edit");
    titleEdit.addEventListener("change", function () {
        let rightBar = document.getElementById("right-bar");
        let taskid = rightBar.getAttribute("data-taskid");
        updateTask(taskid, {title: this.value});
    });

    let contentEdit = document.getElementById("content-edit");
    contentEdit.addEventListener("change", function () {
        console.log("you are fucking changing content");
        let rightBar = document.getElementById("right-bar");
        let taskid = rightBar.getAttribute("data-taskid");
        updateTask(taskid, {detail: this.value});
    });

    let deadline = document.getElementById("input-date");
    deadline.addEventListener("change", function () {
        let rightBar = document.getElementById("right-bar");
        let taskid = rightBar.getAttribute("data-taskid");
        updateTask(taskid, {deadline: this.value})
    });

    let taskTypeEdit = document.getElementById("task-type");
    taskTypeEdit.addEventListener("change", function () {
        let rightBar = document.getElementById("right-bar");
        let taskid = rightBar.getAttribute("data-taskid");
        updateTask(taskid, {type: this.value})
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
    hide.addEventListener("click", function () {
        let rightBar = document.getElementById("right-bar");
        rightBar.classList.add("not-show");
    });

    autoResize()
}



function autoResize() {
    let es = document.getElementsByClassName("edit");
    for (let i = 0; i < es.length; i++) {
        let e = es[i];
        e.addEventListener('input', function () {
            this.style.height = "auto";
            this.style.height = e.scrollHeight + 'px';
        });
        e.addEventListener('load', function () {
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
        createTask({
            completed: 'false',
            title: this.value,
            detail: "",
            type: "未分类",
            start_time: getLocalTimeString(startTime),
            deadline: getLocalTimeString(deadline),
        });
    });
}

function createTask(task) {
    console.log("create task" + JSON.stringify(task));
    let x = new XMLHttpRequest();
    x.onreadystatechange = function () {
        if (x.readyState === 4 && x.status === 201) {
            console.log("create task success");
            showTaskList();
        } else {
            console.log("create task failed: " + x.responseText);
        }
    };
    x.open("POST", "/task");
    x.send(JSON.stringify(task));
}



function getLocalTimeString(utc) {
    let local = new Date(utc.getTime() - timezoneOffset).toISOString();
    local = local.split(".")[0];
    local.replace("T", " ");
    return local;
}

function deleteTask(id) {
    console.log("delete task " + id);
    let x = new XMLHttpRequest();
    x.onreadystatechange = function () {
        if (x.readyState === 4 && x.status === 200) {
            console.log("delete task success");
            showTaskList();
        } else {
            console.log("delete task failed: " + x.responseText);
        }
    };
    x.open("DELETE", "/task/" + id);
    x.send();
}

function deleteTaskEvent() {
    let dlt = document.getElementById("delete");
    dlt.addEventListener("click", function () {
        let rightBar = document.getElementById("right-bar");
        let id = rightBar.getAttribute("data-taskid");
        deleteTask(id);
    });
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
taskListElement.show = showTaskList;
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

