function OpenFilterMenu(){
    let filterMenu = document.getElementById("filter-menu")

    if(!filterMenu.open){
        filterMenu.showModal()
        filterMenu.classList.add("show");
        document.body.classList.add("disable-scrolling")
    }
}

function CloseFilterMenu(event){
    let filterMenu = document.getElementById("filter-menu")

    let rect = filterMenu.getBoundingClientRect();
    let isInDialog = (
        event.clientX >= rect.left &&
        event.clientX <= rect.right &&
        event.clientY >= rect.top &&
        event.clientY <= rect.bottom
    );

    if (!isInDialog) {
        filterMenu.classList.remove("show");
        setTimeout(()=>{
            filterMenu.close();
            document.body.classList.remove("disable-scrolling")
        }, 200)
    }
}