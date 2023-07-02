function factorial(n) {
    if(n === 0) return 1
    return n*factorial(n-1);
}

function callFacttorial(n) {
    console.log(factorial(n));
    console.log(factorial(n+1));
}

callFacttorial(4)