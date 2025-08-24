// Funtions that doesn't relly on DOM are here

function GetBrandsSuggestionsItemsQuantity(){
    let parent = document.getElementById("suggestions-holder")
    if(parent !== undefined){
        return parent.children.length - 1
    }
    return 0
}