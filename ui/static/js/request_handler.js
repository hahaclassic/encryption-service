


function sendJSON(OperationType) {
    let OriginalMessage = document.querySelector('#OriginalMessage');
    let ConvertedMessage = document.querySelector('#ConvertedMessage');
    let Key = document.querySelector('#Key');

    let xhr = new XMLHttpRequest();

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

    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/json");

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            if (OperationType === 'GetRandomKey') {
                ConvertedMessage.value = ''
                Key.value = JSON.parse(this.responseText).key;

            } 
            else {
                ConvertedMessage.value = JSON.parse(this.responseText).converted;
            }
        }
    };
   
    // OperationType    string `json:"operation_type"`
    // OriginalMessage  string `json:"original"`
    // ConvertedMessage string `json:"converted"`
    // Key              string `json:"key"`
    // Language         string `json:"language"`

    var data = JSON.stringify({ "operation_type": OperationType,
     "original": OriginalMessage.value, "converted": "", 
    "key": Key.value, "language": "eng"});

    xhr.send(data);
}

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
