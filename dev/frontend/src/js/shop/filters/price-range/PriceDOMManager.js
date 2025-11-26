class PriceDOMManager{
    constructor(stateManager){
        this.stateManager = stateManager

        WhenDOMLoad(()=>{
            this.#queryElements()

            // Before the page loads, the PriceStateManager automatically starts with the price setted in the URL data
            // This just update when the page loads, to reflect this data
            this.updateThumb("min")
            this.updateThumb("max")

            this.updateInput("min")
            this.updateInput("max")
        });
    }

    /**
     * @param {"input" | "range"} source  
     * @param {"min" | "max"} mode
     */
    getValue(source, mode){
        if(source === "input"){
            return this.inputElements.get(mode)
        }       

        if(source === "range"){
            return this.rangeElements.get(mode)
        }

        throw new Error("Invalid object. Should be 'input' or 'range', received: ", object)
    }

    update(mode){
        this.updateThumb(mode)
        this.updateInput(mode)
    }

    /**
     * This is called after the StateManager resets
     * This just to reflect the reseted data of the StateManager
     */
    reset(){
        this.update("min")
        this.update("max")
    }

    /**
     * Updates the position and value of a range thumb element.
     *
     * Retrieves the corresponding value from `PriceStateManager` based on the given
     * mode (`"min"` or `"max"`) and updates:
     * - The thumb's value (for accessibility and internal state)
     * - The CSS custom property (`--start` or `--end`) to visually position the thumb
     *
     * @param {"min"|"max"} mode - Determines which thumb to update:
     *  - `"min"` updates the starting (minimum) thumb.
     *  - `"max"` updates the ending (maximum) thumb.
     */
    updateThumb(mode){
        let element = this.rangeElements.get(mode)
        let propertyName = mode === "min" ? "--start" : "--end"
        let newValue = mode === "min" ? this.stateManager.min : this.stateManager.max

        // Update thumb position
        element.value = newValue
        element.parentElement.style.setProperty(propertyName, `${newValue}%`)
    }
 
    /**
     * Updates the corresponding price input field based on the current state.
     *
     * Retrieves the value from `PriceStateManager` depending on the specified mode
     * (`"min"` or `"max"`) and sets it in the associated input element.
     *
     * @param {"min"|"max"} mode - Determines which input to update:
     *  - `"min"` updates the minimum price input.
     *  - `"max"` updates the maximum price input.
     */
    updateInput(mode){
        let element = this.inputElements.get(mode)
        let value = mode === "min" ? this.stateManager.min : this.stateManager.max
        element.value = rangeToInput(value)
    }

    /**
     * @description
     * Queries and stores references.
     * Should be run only after the DOM has fully loaded.
     *
     * @private
     */
    #queryElements() {
        this.rangeElements = new Map()
        this.rangeElements.set("min", document.getElementById("min_range"))
        this.rangeElements.set("max", document.getElementById("max_range"))

        this.inputElements = new Map()
        this.inputElements.set("min", document.getElementById("min_input"))
        this.inputElements.set("max", document.getElementById("max_input"))
    }
}