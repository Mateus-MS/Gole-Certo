class PriceStateManager{
    constructor(urlManager, paramKey = "price-range"){
        this.urlManager = urlManager;
        this.paramKey = paramKey;

        this.min_difference = 30
        this.min = 0
        this.max = 100

        this.initFromURL()
    }

    initFromURL(){
        let raw = this.urlManager.read(this.paramKey);
        if(!raw) return;
        let [min, max] = raw.split("-")

        this.min = inputToRange(min)
        this.max = inputToRange(max)
    }

    reset(){
        this.min = 0
        this.max = 100
        this.syncWithURL()
    }

    /**
     * Updates the minimum or maximum value while enforcing the minimum difference.
     *
     * @param {number} min
     * @param {number} max
     * @param {"min"|"max"} mode - Determines which value to update:
     */
    update(min, max, mode){
        let calculated = this.enforceMinDifference(min, max, mode) 
        
        if(mode === "min"){
            this.min = calculated
            return
        }

        if(mode === "max"){
            this.max = calculated
            return
        }
    }

    /**
     * Ensures that the difference between `min` and `max` values is not less than `min_difference`.
     *
     * If the provided value would violate the minimum difference, it is adjusted to maintain
     * the constraint.
     *
     * @param {number} min - Proposed minimum value.
     * @param {number} max - Proposed maximum value.
     * @param {"min"|"max"} mode - Determines which value is being adjusted.
     * @returns {number} - The adjusted value that satisfies the minimum difference rule.
     * @throws {Error} If `mode` is not `"min"` or `"max"`.
     */
    enforceMinDifference(min, max, mode){
        if(mode === "min"){
            return (max - this.min_difference <= min) ? max - this.min_difference : min
        }

        if(mode === "max"){
            return (min + this.min_difference >= max) ? min + this.min_difference : max
        }

        throw new Error("Must be min or max. Received: ", mode)
    }

    syncWithURL(){
        if((parseFloat(this.min) !== 0) || (parseFloat(this.max) !== 100)){
            this.urlManager.update(this.paramKey, `${rangeToInput(this.min)}-${rangeToInput(this.max)}`)
            return
        }
        this.urlManager.delete(this.paramKey)
    }
}