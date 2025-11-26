function inputToRange(value){
    return (((value - MIN_PRICE) / (MAX_PRICE - MIN_PRICE)) * 100).toFixed(2);
}

function rangeToInput(value){
    return (MIN_PRICE + (parseFloat(value) / 100) * (MAX_PRICE - MIN_PRICE)).toFixed(2);
}

class PriceManager{
    constructor(urlManager){
        this.state = new PriceStateManager(urlManager)
        this.dom = new PriceDOMManager(this.state)

        WhenDOMLoad(()=>{
            this.#initiateListeners()
        })
    }

    reset(){
        // Reset the data
        this.state.reset()
        // Reset the DOM
        this.dom.reset()

        // Ask for new product with the reseted data
        RefreshProductGrid()
    }

    /**
     * @param {"input" | "range"} source  
     * @param {"min" | "max"} mode  
     */
    update(source, mode){
        // Get the value from DOM
        let minValue = parseFloat(this.dom.getValue(source, "min").value)
        let maxValue = parseFloat(this.dom.getValue(source, "max").value)

        // The values should be in range format
        if(source === "input"){
            minValue = inputToRange(minValue)
            maxValue = inputToRange(maxValue)
        }

        // Update the state
        // here the value will be enforced in boundaries and keep the min_difference
        this.state.update(minValue, maxValue, mode)

        // Reflect the calculated state on DOM
        this.dom.update(mode)

        // Synchronize the URL with the new state
        this.state.syncWithURL()

        // Update the products grid
        RefreshProductGrid_DEBOUNCED()
    }

    #initiateListeners(){
        this.dom.rangeElements.get("min").addEventListener("input", ()=>{this.update("range", "min")})
        this.dom.rangeElements.get("max").addEventListener("input", ()=>{this.update("range", "max")})

        this.dom.inputElements.get("min").addEventListener("input", ()=>{this.update("input", "min")})
        this.dom.inputElements.get("max").addEventListener("input", ()=>{this.update("input", "max")})
    }
}