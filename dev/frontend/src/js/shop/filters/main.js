const URL_MANAGER = new URLManager()
URL_MANAGER.update("page", "1")

const BRAND_MANAGER = new BrandManager(URL_MANAGER)
const PRICE_MANAGER = new PriceManager(URL_MANAGER)
const PRICE_ORDER_MANAGER = new PriceOrderManager(URL_MANAGER)

const DEBOUNCED_REFRESH = debounce(() => {
    RefreshProductGrid()
}, 1000);

function RefreshProductGrid(){
    htmx.ajax('GET', `/components/prodPage?${URL_MANAGER.toString()}`, {
        target: '#pagination_container',
        swap: 'outerHTML'
    });
}

function RefreshProductGrid_DEBOUNCED(){
    DEBOUNCED_REFRESH();
}

function ClearAllFilters(){
    URL_MANAGER.update("page", "1")

    BRAND_MANAGER.state.reset()
    BRAND_MANAGER.dom.reset()

    PRICE_MANAGER.state.reset()
    PRICE_MANAGER.dom.reset()

    RefreshProductGrid()
}

function UpdatePage(newPage){
    URL_MANAGER.update("page", newPage)
    RefreshProductGrid()
}

function OpenProductPage(productID){
    window.location.href = `/productpage?id=${encodeURIComponent(productID)}`
}