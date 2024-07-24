let button = document.querySelector("#mainForm > button")

if (button) {
    button.onclick = function (e) {

        let inputs = document.querySelectorAll("#mainForm > input")
        let block = document.querySelector("#mainForm > .block")
        let data = {}

        console.log(inputs)

        inputs.forEach(input => {
            data[input.name] = input.value
        })

        console.log(data)

        let xhr = new XMLHttpRequest()
        xhr.open("POST", "/messages")
        xhr.setRequestHeader("Content-Type", "application/json")

        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response)
            if (response && "Error" in response) {
                if (response.Error == null) {
                    console.log("fff")
                    console.log("Операция прошла успешноdd")
                    block.innerText = "Сообщение успешно отправлено"
                } else {
                    console.log("Возникла ошибка", response.Error)
                }
            } else {
                console.log("Некорректные данные")
            }
        };
        xhr.send(JSON.stringify(data))
    }
}