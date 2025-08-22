const min_range = new RangeInput(document.getElementById("min_range"))
const max_range = new RangeInput(document.getElementById("max_range"))

const min_price = new PriceInput(document.getElementById("min_input"))
const max_price = new PriceInput(document.getElementById("max_input"))

function EnforceMinDifference(min, max, mode){
    if(mode === "min"){
        return (max - MIN_DIFFERENCE <= min) ? max - MIN_DIFFERENCE : min
    }

    if(mode === "max"){
        return (min + MIN_DIFFERENCE >= max) ? min + MIN_DIFFERENCE : max
    }

    throw new Error("Must be min or max. Received: ", mode)
}