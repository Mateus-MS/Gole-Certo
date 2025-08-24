function getpriceOrderFromURL(){
    let params = new URLSearchParams(window.location.search);
    let value = params.get("price-range")

    return value
}

function getpriceOrderFromDOM(){
    let element = document.querySelector(".filter-group-holder input[type='radio']:checked");
    return element.value
}

function setPriceOrder(priceOrderToActive){
    if(priceOrderToActive === "ascending"){
        document.getElementById("price-order-ascending").checked = true
    } else {
        document.getElementById("price-order-descending").checked = true
    }
}