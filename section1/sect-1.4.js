
console.log(f(2));

function f(n) {
    if(n%2==0){
        // if(n==0)return 0

        return n/2 + f(n/2)
    }
    return 1 + f(n-1)
}