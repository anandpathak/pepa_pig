const URL = 'http://localhost:8080/fetch_countries'

$("#bylineInfo_feature_div").ready(function() {
    productDoc = $("#bylineInfo_feature_div");
    console.log($("#bylineInfo_feature_div")[0].innerText);
    productCreatedBy = productDoc[0].innerText;

    fetchCountriesNameForProduct(productDoc, productCreatedBy);
});

function fetchCountriesNameForProduct(productDoc, productCreatedBy){
    requestData = {
        source: 'amazon',
        companies:getCompaniesName(productCreatedBy)
    };
    fetch(URL, {
        method: 'POST',
        body: JSON.stringify(requestData),
        headers: {
            "Content-Type":"application/json"
        }
    }).then(resp => resp.json())
    .then(data => {
        console.log(data.data);
        c = Object.values(data.data);
        country_name = c.find(val => val.length > 0);
        country_name && (productDoc[0].innerHTML += ` country: ${country_name}`);
    });

}
function getCompaniesName(companiesAsString) {
    companiesAsString = companiesAsString.toLowerCase();
    list = companiesAsString.split(' ');
    result = [];
    list.map((elem, pos) => {

        result.push({name: elem});
        list[pos+1] && result.push({
            name: `${elem}_${list[pos+1]}`
        });
        list[pos+2] && result.push({
            name:`${elem}_${list[pos+1]}_${list[pos+2]}`
        });
    });
    return result;
}