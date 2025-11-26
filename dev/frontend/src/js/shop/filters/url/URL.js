class URLManager{
    constructor(){
        this.url = new URLSearchParams(window.location.search); 
    }

    read(param){
        return this.url.get(param)
    }

    update(param, newValue){
        if (newValue === undefined){
            throw new Error("When updating the URL, the value passed must be non null")
        }

        this.url.set(param, newValue)
        this.#updatePath()
    }

    delete(param){
        this.url.delete(param)
        this.#updatePath()
    }

    toString(){
        return this.url.toString()
    }

    #updatePath(){
        history.replaceState({}, '', `/shop?${this.toString()}`)
    }
}