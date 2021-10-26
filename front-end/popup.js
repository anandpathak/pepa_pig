let submitButton = document.getElementById('submit');

const URL = 'http://localhost:8080/add_country';

const ALERT_LEVEL = {
    'warn': 'orange',
    'error': 'red',
    'success': 'green'
}
chrome.storage.sync.get('color', function (data) {


});

submitButton.onclick = function (element) {
    requestData = {
        "source": "amazon",
        "company": document.getElementById('company').value,
        "country": document.getElementById('country').value
    }
    fetch(URL, {
            method: "POST",
            body: JSON.stringify(requestData),
            headers: {
                "Content-Type": "application/json"
            }
        }).then(resp => {
            if (resp.ok)
                return resp.json();
            throw new Error("Failed to add in review");
            })
        .then(data => {
            Alert(data.data, 'success');
        }).catch(err => {
            Alert(err, 'error')
        });
};

// PushCountryForReview(companyName, countryName) {
//     fetch({

//     })
// }

function Alert(msg, type) {

    let err = document.getElementById('error');
    console.log(err);
    err.innerHTML = msg;
    err.style.display = 'block';
    err.style.backgroundColor = ALERT_LEVEL[type];
    setTimeout(function () {
        err.style.display = 'none';
    }, 3000);
}