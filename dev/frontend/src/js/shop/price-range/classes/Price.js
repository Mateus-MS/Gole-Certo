// TODO: Heavly depedent on external consts outside this file
class PriceInput{

    constructor(element){
        this.element = element;

        if(element.id.includes("min")){
            this.isMin = true
            this.element.value = MIN_PRICE
        } else {
            this.isMin = false
            this.element.value = MAX_PRICE
        }

        this.registerEvents()
    }

    get rawValue(){
        return parseFloat(this.element.value)
    }

    registerEvents(){
        this.element.addEventListener("focusout", this.onConfirmation.bind(this))
        this.element.addEventListener("keypress", (e)=>{
            if(e.key === "Enter"){
                e.stopPropagation();     
                e.preventDefault();
                e.target.blur();
            }
        })
    }

    onConfirmation(){
        if(this.isMin){
            let att = EnforceMinDifference(PriceInput.ToRange(this.rawValue), max_range.rawValue, "min") 
            min_range.updateThumb(att)
            this.element.value = RangeInput.ToPrice(min_range.rawValue)
        } else {
            let att = EnforceMinDifference(min_range.rawValue, PriceInput.ToRange(this.rawValue), "max") 
            max_range.updateThumb(att)
            this.element.value = RangeInput.ToPrice(max_range.rawValue)
        }

        DeboucingFiltersConstUpdate()
    }

    // Converts this price object value into a range object value
    static ToRange(value){
        return (((value - MIN_PRICE) / (MAX_PRICE - MIN_PRICE)) * 100).toFixed(2);
    }
}