function buttonEvent() {
    let e;
    e = document.getElementById("register")
    e.addEventListener("click", function () {
        let ele = document.getElementById("confirm-password");
        ele.classList.remove("hide");

        ele = document.getElementById("confirm-register");
        ele.classList.remove("hide");

        ele = document.getElementById("login");
        ele.classList.add("hide");
    })

    e = document.getElementById("confirm-register");
    e.addEventListener("click", function () {
        let username = document.getElementById("username-input").value;
        let password = document.getElementById("password-input").value;
        let passwordConfirm = document.getElementById("password-confirm").value;
        if (password !== passwordConfirm) {
            alert("密码不一致")
            return
        }
        let r = {
            username: username,
            password: password
        };
        console.log(JSON.stringify(r))
        DoRequest(POST, "/register", r, true, confirmRegisterCallback)
    })

    e = document.getElementById("login");
    e.addEventListener("click", function () {
        let username = document.getElementById("username-input").value;
        let password = document.getElementById("password-input").value;
        let l = {
            username: username,
            password: password
        };
        console.log(JSON.stringify(l));
        DoRequest(POST, "/login", l, true, redirect)
    })
}

function confirmRegisterCallback() {
    if (this.status !== 201) {
        alert(this.responseText);
    } else {
        window.location.assign("http://127.0.0.1:8080/");
    }
}

function redirect() {
    if (this.status === 201) {
        console.log("login success");
        window.location.assign("http://127.0.0.1:8080/task");
    } else {
        alert(this.responseText);
    }
}

buttonEvent();
