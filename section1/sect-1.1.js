
console.log(f(3));

function f(n) {
    if(n==0) return 1

    return n*f(n-1)
}