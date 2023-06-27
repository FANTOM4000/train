
console.log(getSquareArea(5,8));

function getSquareArea(width, height) {
    if(width<height){
        return Math.pow(width,2)
    }
    return Math.pow(height,2)
}