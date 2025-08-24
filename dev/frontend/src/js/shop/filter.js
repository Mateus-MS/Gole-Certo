const SUGGESTIONS_HOLDER = document.getElementById("suggestions-holder")
const SELECTEDS_HOLDER   = document.getElementById("selecteds-holder")

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
        createSelectedFiltersBrands(brandsArray)
    }

    let priceOrder = params.get("price-order");
    if(priceOrder !== null){
        setPriceOrder(priceOrder)
    }
}, {once: true})

function createSelectedFiltersBrands(brandsArray){
    for(let i = 0; i < brandsArray.length; i++){
        let li = document.createElement("li")
        li.classList.add("checker")
        let label = document.createElement("label")
        let input = document.createElement("input")
        input.type = "checkbox"
        input.checked = true
        input.name = "brand"
        input.value = brandsArray[i]
        input.onchange = function() {
            DeboucingFiltersConstUpdate(input);
        }
        
        let span = document.createElement("span")
        span.innerText = brandsArray[i]
        span.classList.add("text")

        label.appendChild(input)
        label.appendChild(span)

        li.appendChild(label)

        SELECTEDS_HOLDER.appendChild(li)
    }
}

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
    // TODO: Rename it to priceRange
    price: `${MIN_PRICE}-${MAX_PRICE}`,
    priceOrder: getpriceOrderFromURL(),
    brands: getFiltersBrandsSelectedsInRequestFormatFromURL()
};

function getFiltersBrandsSelectedsInRequestFormatFromURL(){
    let url = new URLSearchParams(window.location.search);
    let brandsRaw = url.get("brands")
    if(brandsRaw !== null){
        let brands = brandsRaw.split(":")

        let reqFormat = ""
        for(let i = 0; i < brands.length; i++){
            reqFormat += brands[i]
            
            if(i < brands.length){
                reqFormat += ":"
            }
        }

        return reqFormat
    }

    return ""
}

function getFiltersBrandsSelectedsInRequestFormatFromDOM(){
    let brands = document.getElementById("selecteds-holder")
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

function UnselectFiltersBrands(){
    let brandsElements = SELECTEDS_HOLDER.querySelectorAll("input[type=checkbox]")
    brandsElements.forEach(
        brand => {
            brand.checked = false
            handleBrandFilterToggle(brand.parentElement.parentElement)
        }
    )

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
    FILTERS.brands = getFiltersBrandsSelectedsInRequestFormatFromDOM()

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

    let temp = getpriceOrderFromDOM()
    if(temp !== null){
        params.set("price-order", temp)
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
        swap: 'outerHTML'
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

const debouncBrandFilterSuggestion = debounce(()=>{
    let brandsItems = SUGGESTIONS_HOLDER.children
    let searchQuery = document.getElementById("brand-search-input").value

    // Show all items
    if(searchQuery === "") {
        for(let i = 0; i < brandsItems.length; i++){
            brandsItems[i].style.display = "block"
        }
        return
    }

    // Hide all items that doesn't match
    let escapedSearch = searchQuery.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
    let regex = new RegExp(escapedSearch, 'i');

    for(let i = 0; i < brandsItems.length; i++){
        let item = brandsItems[i]
        let value = item.querySelector("input").value

        if(!regex.test(value)){
            item.style.display = "none"
        }
    }
}, 200);

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
function DeboucingFiltersConstUpdate(element) {
    FILTERS.page = 1

    debouncedUpdate();

    if(element !== undefined) {
        moveTheSelectedBrandFilterToSelectedHolder(element)
    }
}

function moveTheSelectedBrandFilterToSelectedHolder(element){
    let brandFilter = element.parentElement.parentElement

    handleBrandFilterToggle(brandFilter)
}

function handleBrandFilterToggle(brandFilter){
    // Try to select
    if(brandFilter.parentElement.id === SUGGESTIONS_HOLDER.id){
        SELECTEDS_HOLDER.appendChild(brandFilter)
        return
    }

    // Move it back to suggestions holder
    SUGGESTIONS_HOLDER.appendChild(brandFilter)

    // Re-order it alphabetically
    let items = Array.from(SUGGESTIONS_HOLDER.children)
    items.sort((a, b) => 
        a.textContent.trim().localeCompare(b.textContent.trim(), undefined, { sensitivity: 'base' })
    );

    // Re-append in sorted order
    items.forEach(item => SUGGESTIONS_HOLDER.appendChild(item));
}

function DebouncBrandFilterSuggestionCall(){
    debouncBrandFilterSuggestion()
}