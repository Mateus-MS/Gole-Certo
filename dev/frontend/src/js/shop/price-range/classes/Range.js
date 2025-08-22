// TODO: Heavly depedent on external consts outside this file

class RangeInput{

    constructor(element){
        this.element = element;

        this.propertyName = element.id.includes("min") ? "--start" : "--end"
        this.isMin = element.id.includes("min")

        this.registerEvents()
    }

    get rawValue(){
        return parseFloat(this.element.value)
    }

    updateThumb(newThumbValue){
        this.element.value = newThumbValue
        this.element.parentElement.style.setProperty(this.propertyName, `${newThumbValue}%`)

        if(this.isMin){
            min_price.element.value = RangeInput.ToPrice(this.rawValue)
        } else {
            max_price.element.value = RangeInput.ToPrice(this.rawValue)
        }
    }

    registerEvents(){
        this.element.addEventListener("input", this.onInputChange.bind(this))
    }

    //Enforces min-max difference when user moves slider thumb.
    onInputChange(){
        this.updateThumb(EnforceMinDifference(min_range.rawValue, max_range.rawValue, this.isMin ? "min" : "max"))
    
        DeboucingFiltersConstUpdate()
    }

    // Converts this range object value into a price object value
    static ToPrice(value){
        return (MIN_PRICE + (parseFloat(value) / 100) * (MAX_PRICE - MIN_PRICE)).toFixed(2);
    }

}