window.addEventListener("load", ()=>{
    const params = new URLSearchParams(window.location.search);

    let page = params.get("page")
    if(page !== null){
        FILTERS.page = page
    }
    
    let price = params.get("price");
    if(price !== null){
        let [min, max] = price.split("-");
        setValuesInPriceInInputs(min, max)
    }

    let brands = params.get("brands");
    if(brands !== null){
        let brandsArray = brands.split(":")
        selectFiltersBrands(brandsArray)
    }
}, {once: true})

/**
 * Represents the current state of filters based on user input.
 *
 * @constant {Object} FILTERS
 * @property {string} price - The selected price range in the format "min-max".
 * @property {string} brands - All brands names separated by ":".
 *
 * @file "/js/shop/filter.js"
 */
const FILTERS = {
    page: 1,
    price: `${MIN_PRICE}-${MAX_PRICE}`,
    brands: ""
};

function getFiltersBrandsSelectedsInRequestFormat(){
    let brands = document.getElementById("brands_filter")
    let selecteds = brands.querySelectorAll("input:checked")

    let filter = ""

    selecteds.forEach(selected => {
        if(filter == "") {
            filter += selected.value
        } else {
            filter += ":" + selected.value
        }
    })

    return filter
}

function selectFiltersBrands(brandsArray){
    let brandsElements = document.querySelectorAll("#brands_filter input[type=checkbox]")
    for(let i = 0; i < brandsElements.length; i++){
        for(let j = 0; j < brandsArray.length; j++){
            if(brandsElements[i].value === brandsArray[j]){
                // Mark the input box as selected
                brandsElements[i].checked = true;
                break
            }
        }
    }
}

function UnselectFiltersBrands(){
    let brandsElements = document.querySelectorAll("#brands_filter input[type=checkbox]")
    for(let i = 0; i < brandsElements.length; i++){
        if(brandsElements[i].checked === true){
            // Mark the input box as unselected
            brandsElements[i].checked = false;
        }
    }

    ProductsRefresh()
}

function getPriceInRequestFormat(){
    let min = document.getElementById("min_input").value.replace(",", ".")
    let max = document.getElementById("max_input").value.replace(",", ".")

    return `${min}-${max}`
}

function setValuesInPriceInInputs(min, max){
    min_range.updateThumb(PriceInput.ToRange(min))
    max_range.updateThumb(PriceInput.ToRange(max))
}

function UnselectFilterPrice(){
    min_range.updateThumb(PriceInput.ToRange(MIN_PRICE))
    max_range.updateThumb(PriceInput.ToRange(MAX_PRICE))

    ProductsRefresh()
}

function GetFiltersInRequestFormat(){
    // Update the values with the actual values from the inputs
    FILTERS.price = getPriceInRequestFormat()
    FILTERS.brands = getFiltersBrandsSelectedsInRequestFormat()

    // Build the URL parameters
    let params = new URLSearchParams();
    if(FILTERS.page !== ""){
        params.set("page", FILTERS.page);
    }

    // Not include the price parameter if the min and max prices are "default"
    if(FILTERS.price !== ""){
        let [min, max] = FILTERS.price.split("-")
        
        // If min and max are different from default
        if((parseFloat(min) !== MIN_PRICE) || (parseFloat(max) !== MAX_PRICE)){
            params.set("price", FILTERS.price);
        }
    }
    if(FILTERS.brands !== ""){
        params.set("brands", FILTERS.brands);
    }

    return params.toString()
}

function ClearAllFilters(){
    // TODO: refactor to re-use the code from: UnselectFilterPrice & UnselectFiltersBrands
    let brandsElements = document.querySelectorAll("#brands_filter input[type=checkbox]")
    for(let i = 0; i < brandsElements.length; i++){
        if(brandsElements[i].checked === true){
            // Mark the input box as unselected
            brandsElements[i].checked = false;
        }
    }

    min_range.updateThumb(PriceInput.ToRange(MIN_PRICE))
    max_range.updateThumb(PriceInput.ToRange(MAX_PRICE))

    ProductsRefresh()
}

function ProductsRefresh(){
    params = GetFiltersInRequestFormat()

    // Refresh the product list with HTMX
    htmx.ajax('GET', `/components/prodPage?${params}`, {
        target: '#pagination_container',
        swap: 'outerHTML',
        vals: {
            page: FILTERS.page,
            price: FILTERS.price,
            brands: FILTERS.brands
        }
    });

    // Update the history with the new URL parameters
    history.replaceState({}, '', `/shop?${params.toString()}`)
}

function UpdatePage(newPageIndex){
    FILTERS.page = newPageIndex
    window.scrollTo({ top: 0, behavior: 'smooth' })

    ProductsRefresh()
}

const debouncedUpdate = debounce(() => {
    ProductsRefresh()
}, 1000);

/**
 * Executes a debounced update to the `FILTERS` constant with a delay of 1000ms.
 *
 * This function relies on `debouncedUpdate`, which applies the debounce logic
 * to update the filter values based on the current state of the input elements.
 *
 * @returns {void}
 * 
 * @see FILTERS
 * @file /js/shop/filter.js
 * 
 */
function DeboucingFiltersConstUpdate() {
    FILTERS.page = 1

    debouncedUpdate();
}
