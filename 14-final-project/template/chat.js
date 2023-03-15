var app = {}
app.ws = undefined
app.container = undefined

app.print = function (message) {
    var el = document.createElement("p")
    el.innerHTML = message
    app.container.append(el)
}

app.doSendMessage = function () {
    var messageRaw = document.querySelector('.input-message').value
    app.ws.send(JSON.stringify({
        Message: messageRaw
    }));

    var message = new Date().toLocaleTimeString() + ' <b>me</b>: ' + messageRaw
    app.print(message)

    document.querySelector('.input-message').value = ''
}

app.init = function () {
    if (!(window.WebSocket)) {
        alert('Your browser does not support WebSocket')
        return
    }

    var name = prompt('Enter your name please:') || "No name"
    document.querySelector('.username').innerText = name

    app.container = document.querySelector('.container')

    app.ws = new WebSocket("ws://localhost:8080/ws?name=" + name)

    app.ws.onopen = function () {
        var message = new Date().toLocaleTimeString() + ' <b>me</b>: connected'
        app.print(message)
    }

    app.ws.onmessage = function (event) {
        var res = JSON.parse(event.data)

        var messsage = ''
        if (res.Type === 'NEW_USER') {
            message = new Date().toLocaleTimeString() + ' <b>' + res.From + '</b>: connected'
        } else if (res.Type === 'DISCONNECT') {
            message = new Date().toLocaleTimeString() + ' <b>' + res.From + '</b>: disconnected'
        } else {
            message = new Date().toLocaleTimeString() + ' <b>' + res.From + '</b>: ' + res.Message
        }

        app.print(message)
    }

    app.ws.onclose = function () {
        var message = new Date().toLocaleTimeString() + ' <b>me</b>: disconnected'
        app.print(message)
    }
}

window.onload = app.init