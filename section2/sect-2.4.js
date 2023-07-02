function findAllInt(n) {
    for (let index = n; index > 0; index--) {
        // console.log(index);
        if(n%index == 0) {
            console.log(index);
        }
    }
}

findAllInt(20)