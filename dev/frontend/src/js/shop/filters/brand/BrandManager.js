/**
 * Manages the brand filter UI and its state.
 *
 * @description
 * Handles brand selection and deselection in the product filter UI.
 * Synchronizes the selected brands with the URL, updates the DOM accordingly,
 * and triggers a debounced product grid refresh when selections change.
 *
 * The class works with:
 * - `StateManager` for keeping track of selected brands and URL updates.
 * - `DOMManager` for creating, moving, and sorting brand filter elements.
 * 
 * @requires URLManager
 */
class BrandManager{
    constructor(urlManager){
        // External
        this.urlManager = urlManager;

        // Internal
        this.state = new BrandStateManager(urlManager);
        this.state.initFromURL();
        
        WhenDOMLoad(() => {
            this.#queryElements();
            this.searchInput.addEventListener("input", debounce(()=>{ this.onTypeSearch() }, 200));
            this.dom = new BrandDOMManager(this.selectedsHolder, this.suggestionsHolder);
            this.#activateBrandsFromURL();
        });
    }

    /**
     * Updates the brand filter selection when a brand input is toggled.
     *
     * @description
     * Triggered by the **Brand Filter Inputs**.
     * Determines whether the brand should be selected or unselected,
     * moves its `<li>` holder to the correct parent, updates the URL state,
     * and triggers a debounced product grid refresh.
     *
     * @param {HTMLInputElement} inputElement - The input element that triggered the update.
     */
    onClick(inputElement) {
        this.state.urlManager.update("page", "1");

        let action = this.state.brands.has(inputElement.value) ? "unselect" : "select";
        let holder = inputElement.closest('li');

        this.state.updateSet(inputElement.value, action);

        action === "select" ? this.dom.select(holder) : (this.dom.unselect(holder), this.dom.reorderSuggestions());

        this.state.syncWithURL();
        RefreshProductGrid_DEBOUNCED();
    }

    onTypeSearch(){
        let inputValue = this.searchInput.value
        let brandsElements = [...this.suggestionsHolder.children]

        // If the input search is empty, show all options
        if(inputValue === ""){
            brandsElements.forEach( element => {
                this.dom.toggleElementVisibility(element, false)
            })
            return
        }

        // Hide all items that doesn't match
        let escapedSearch = inputValue.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
        let regex = new RegExp(escapedSearch, 'i');

        brandsElements.forEach( element => {
            let value = element.querySelector("input").value

            if(!regex.test(value)){
                this.dom.toggleElementVisibility(element, true)
            }
        })
    }

    /**
     * Remove all selected filters
     */
    reset(){
        this.state.reset()
        this.dom.reset()

        RefreshProductGrid();
    }

    /**
     * @description
     * Returns all selected brands as a string, separated by "-".
     *
     * @returns {string} A string representing the selected brands for use in the URL.
     */
    getInURLFormat(){
        let brandsArray = Array.from(this.brands)
        return brandsArray.join("-")
    }

    /**
     * @description
     * Queries and stores references to the main brand filter DOM elements.
     * Should be run only after the DOM has fully loaded.
     *
     * @private
     */
    #queryElements() {
        this.searchInput = document.getElementById("brand-search-input");
        this.suggestionsHolder = document.getElementById("suggestions-holder");
        this.selectedsHolder = document.getElementById("selecteds-holder");
    }

    /**
     * Activates brand filters based on the current URL parameter.
     *
     * @description
     * Reads the brands from the URL, creates corresponding filter elements,
     * and appends them to the selected brands holder.
     *
     * @private
     */
    #activateBrandsFromURL() {
        let brands = this.state.urlManager.read(this.state.paramKey);
        if (!brands) return;

        brands.split("-").forEach(brand => {
            this.selectedsHolder.appendChild(
                this.dom.createBrandFilterElement(brand, this.update.bind(this))
            );
        });
    }
}