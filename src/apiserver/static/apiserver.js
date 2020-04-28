const GET = "GET"
const POST = "POST"
const PUT = "PUT"
const DELETE = "DELETE"

function DoRequest(method, url, body, async, callback) {
    let x = new XMLHttpRequest();
    x.method = method;
    x.url = url;
    x.onload = callback;
    x.onerror = function () {
        console.error(`do http request failed, method: ${method}, url: ${url}`)
    };
    x.open(method, url, async);
    if (body !== null) {
        x.send(JSON.stringify(body));
    } else {
        x.send();
    }
}

function callbackVerifyStatus() {
    let okCode;
    switch (this.method) {
        case POST:
            okCode = 201;
            break;
        default:
            okCode = 200;
    }
    if (this.status !== okCode) {
        console.error(`request ${this.method}:${this.url} failed, status: ${this.status}, response: ${this.responseText}`);
    }
}
