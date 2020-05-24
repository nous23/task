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
        DoRequest(POST, "/register", r, true, callbackVerifyStatus)
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
        DoRequest(GET, "/login", l, true, callbackVerifyStatus)
    })
}

buttonEvent();