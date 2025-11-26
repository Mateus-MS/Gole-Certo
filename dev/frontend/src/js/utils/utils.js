function WhenDOMLoad(func){
    window.addEventListener("load", ()=>{
        func()
    }, {once: true})
}