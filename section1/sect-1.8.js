
console.log(get(5,8));

function get(width, height) {
    if(width>height){
        return "Landscape"
    }
    return "Portrait"
}