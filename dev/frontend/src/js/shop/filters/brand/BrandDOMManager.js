/**
 * Handles the creation, movement, and sorting of brand filter DOM elements.
 *
 * @description
 * Responsible for managing the visual representation of brand filters.
 * Moves elements between selected and suggestion holders, maintains alphabetical order,
 * and creates new filter elements with proper event bindings.
 *
 * Usage:
 * ```js
 * const domManager = new DOMManager(selectedsHolder, suggestionsHolder);
 * const element = domManager.createBrandFilterElement('CocaCola', handler);
 * domManager.select(element);
 * ```
 */
class BrandDOMManager{
    
    /**
     * @param {HTMLElement} selectedsHolder - Container element for selected brand filters.
     * @param {HTMLElement} suggestionsHolder - Container element for available brand suggestions.
     */
    constructor(selectedsHolder, suggestionsHolder) {
        this.selectedsHolder = selectedsHolder;
        this.suggestionsHolder = suggestionsHolder;
    }

    /**
     * Moves a brand element to the selected brands holder.
     *
     * @param {HTMLElement} brandElement - The <li> element representing the brand.
     */
    select(brandElement) {
        this.selectedsHolder.appendChild(brandElement);
    }

    /**
     * Moves a brand element back to the suggestions holder.
     *
     * @param {HTMLElement} brandElement - The <li> element representing the brand.
     */
    unselect(brandElement) {
        this.suggestionsHolder.appendChild(brandElement);
    }

    toggleElementVisibility(brandElement, force = null){
        brandElement.classList.toggle("hidden", force)
    }

    reset(){
        let selecteds = this.selectedsHolder.querySelectorAll("input:checked");
        selecteds.forEach(input => {
            input.checked = false;
            let li = input.closest('li');
            this.suggestionsHolder.appendChild(li);
        });

        this.reorderSuggestions();
    }

    /**
     * Sorts all brand elements in the suggestions holder alphabetically by text content.
     */
    reorderSuggestions() {
        [...this.suggestionsHolder.children]
            .sort((a, b) => a.textContent.trim().localeCompare(b.textContent.trim(), undefined, { sensitivity: 'base' }))
            .forEach(el => this.suggestionsHolder.appendChild(el));
    }

    /**
     * Creates a new brand filter <li> element with a checkbox and label.
     *
     * @param {string} brand - The brand name.
     * @param {function(HTMLInputElement):void} updateHandler - Function to call when the checkbox changes.
     * @returns {HTMLLIElement} The created <li> element.
     */
    createBrandFilterElement(brand, updateHandler) {
        let li = document.createElement("li");
        li.className = "checker";

        let label = document.createElement("label");

        let input = document.createElement("input");
        input.type = "checkbox";
        input.checked = true;
        input.name = "brand";
        input.value = brand;
        input.addEventListener('change', e => updateHandler(e.target));

        let span = document.createElement("span");
        span.className = "text";
        span.textContent = brand;

        label.appendChild(input);
        label.appendChild(span);
        li.appendChild(label);
        return li;
    }
}