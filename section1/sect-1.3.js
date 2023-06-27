
console.log(f(8));

function f(n) {
    if(n==0) return 1

    return 2*f(n-1) 
}