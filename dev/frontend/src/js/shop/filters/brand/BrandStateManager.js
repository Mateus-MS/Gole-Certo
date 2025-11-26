/**
 * Manages the state of selected brands and synchronizes it with the URL.
 *
 * @description
 * Keeps track of the currently selected brands using a `Set`.
 * Provides methods to initialize from the URL, update the selection,
 * and synchronize the current state back to the URL.
 *
 * This class is typically used in conjunction with a DOM manager
 * that handles the visual representation of the selected/unselected brands.
 */
class BrandStateManager{

    /**
     * @param {URLManager} urlManager - An instance of URLManager to read/write URL parameters.
     * @param {string} [paramKey="brands"] - The URL parameter key for storing the brands selection.
     */
    constructor(urlManager, paramKey = "brands"){
        this.urlManager = urlManager;
        this.paramKey = paramKey;
        this.brands = new Set();
    }

    /**
     * Initializes the selected brands from the URL parameter.
     */
    initFromURL() {
        let raw = this.urlManager.read(this.paramKey);
        if (!raw) return;
        raw.split("-").forEach(brand => this.brands.add(brand));
    }

     /**
     * Updates the selection state of a single brand.
     *
     * @param {string} value - The brand name.
     * @param {"select"|"unselect"} action - Whether to add or remove the brand from the selection.
     */
    updateSet(value, action) {
        action === "select" ? this.brands.add(value) : this.brands.delete(value);
    }

    reset(){
        this.brands.clear()
        this.syncWithURL()
    }

    /**
     * Synchronizes the current set of selected brands with the URL.
     * If no brands are selected, the parameter is removed from the URL.
     */
    syncWithURL(){
        if(this.brands.size === 0) {
            this.urlManager.delete(this.paramKey)
        } else {
            this.urlManager.update(this.paramKey, [...this.brands].join("-"));
        }
    }
}