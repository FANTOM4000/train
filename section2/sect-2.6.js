function piramid() {
    let row = 5
    for (let i = 1; i <= row; i++) {
        let stars = ""

        for (let k = 0; k < row-i; k++) {
            stars+=" "
        }

        for (let k = 1; k <= i; k++) {
            stars+="*"
            if(k != i){
                stars+="."
            }
        }
        console.log(stars);

    }

    console.log();
    for (let i = row; i >= 1; i--) {
        let stars = ""

        for (let k = 0; k < row-i; k++) {
            stars+=" "
        }

        for (let k = 1; k <= i; k++) {
            stars+="*"
            if(k != i){
                stars+="."
            }
        }
        console.log(stars);

    }
}

piramid()