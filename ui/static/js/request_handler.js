


function sendJSON(OperationType) {
    // с помощью jQuery обращаемся к элементам на странице по их именам
    let OriginalMessage = document.querySelector('#OriginalMessage');
    let ConvertedMessage = document.querySelector('#ConvertedMessage');
    let Key = document.querySelector('#Key');

    // а вот сюда мы поместим ответ от сервера
    //let result = document.querySelector('#ConvertedMessage');
    // создаём новый экземпляр запроса XHR
    let xhr = new XMLHttpRequest();
    // адрес, куда мы отправим нашу JSON-строку

    const title = document.querySelector('#title')
    let url

    switch (title) {
        case 'Caesar Cipher':
            url = "http://localhost:8080/caesar/";
        
        case 'Affine Cipher':
            url = "http://localhost:8080/affine/";

        case 'Vigenere Cipher':
            url = "http://localhost:8080/vigenere/";

        case 'Simple Substitution Cipher':
            url = "http://localhost:8080/caesar/";

    }


    // открываем соединение
    xhr.open("POST", url, true);
    // устанавливаем заголовок — выбираем тип контента, который отправится на сервер, в нашем случае мы явно пишем, что это JSON
    xhr.setRequestHeader("Content-Type", "application/json");

    // когда придёт ответ на наше обращение к серверу, мы его обработаем здесь
    xhr.onreadystatechange = function () {
        // если запрос принят и сервер ответил, что всё в порядке
        if (xhr.readyState === 4 && xhr.status === 200) {
        // выводим то, что ответил нам сервер — так мы убедимся, что данные он получил правильно
            if (OperationType === 'GetRandomKey') {
                ConvertedMessage.value = ''
                Key.value = JSON.parse(this.responseText).key;

            } 
            else {
                ConvertedMessage.value = JSON.parse(this.responseText).converted;
            }
        }
    };
   
    // преобразуем наши данные JSON в строку
    // OperationType    string `json:"operation_type"`
    // OriginalMessage  string `json:"original"`
    // ConvertedMessage string `json:"converted"`
    // Key              string `json:"key"`
    // Language         string `json:"language"`

    var data = JSON.stringify({ "operation_type": OperationType,
     "original": OriginalMessage.value, "converted": "", 
    "key": Key.value, "language": "eng"});

    // когда всё готово, отправляем JSON на сервер
    xhr.send(data);
}

// document.querySelector("#GetRandomKey").onclick = function() {
//     sendJSON('GetRandomKey')
//     sendJSON('Encrypt')
// }

// document.querySelector("#EncryptButton").onclick = function() {
//     sendJSON('Encrypt')
// }

// document.querySelector("#DecryptButton").onclick = function() {
//     sendJSON('Decrypt')
// }




function sendRequest() {
    let OriginLabel = document.querySelector('#originLabel');

    if (OriginLabel.value === "Original message") {
        sendJSON('Encrypt')
    }
    else {
        sendJSON('Decrypt')
    }
}


document.querySelector("#swap").onclick = function() {
    let OriginalMessage = document.querySelector('#OriginalMessage');
    let ConvertedMessage = document.querySelector('#ConvertedMessage');
    let OriginLabel = document.querySelector('#originLabel');
    let ConvertedLabel = document.querySelector('#convertedLabel');

    let temp = OriginLabel.value
    OriginLabel.value = ConvertedLabel.value
    ConvertedLabel.value = temp

    temp = OriginalMessage.value
    OriginalMessage.value = ConvertedMessage.value;
    ConvertedMessage.value = temp;
}