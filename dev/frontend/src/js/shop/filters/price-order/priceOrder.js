class PriceOrderManager{
    constructor(urlManager, paramKey = "price-order"){
        this.urlManager = urlManager
        this.paramKey = paramKey

        WhenDOMLoad(()=>{
            this.#queryElements()
            this.#setFromURL()
        })
    }

    updatePriceOrder(active){
        this.elements.get(active).checked = true
        this.order = active
        this.syncWithURL()

        RefreshProductGrid_DEBOUNCED()
    }

    syncWithURL(){
        this.urlManager.update(this.paramKey, this.order)
    }

    /**
     * Sets the sorting order based on the value stored in the URL.
     *
     * If the URL contains a value for `paramKey`, it is used as the active order.
     * Otherwise, defaults to `"descending"`. Updates both the corresponding
     * DOM element's checked state and the internal `order` property.
     *
     * @private
     */
    #setFromURL(){
        // If there is something in URL
        // Set it as the active one
        // If nothing on URL, set the ascending as default
        let urlValue = this.urlManager.read(this.paramKey)
        urlValue = urlValue ? urlValue : "ascending"

        this.elements.get(urlValue).checked = true
        this.order = urlValue
    }

    #queryElements(){
        this.elements = new Map()

        let ascending = document.getElementById("price-order-ascending")
        this.elements.set("ascending", ascending)
        ascending.addEventListener("click", ()=>{this.updatePriceOrder("ascending")})
        
        let descending = document.getElementById("price-order-descending")
        this.elements.set("descending", descending)
        descending.addEventListener("click", ()=>{this.updatePriceOrder("descending")})
    }
}