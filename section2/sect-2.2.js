
function calCash(n){
    let banknote10 = Math.floor(n/10)

    n = n-(10*banknote10)

    console.log(`$10 x ${banknote10}`);

    let banknote5 = Math.floor(n/5)

    console.log(`$5 x ${banknote5}`);

    n = n-(5*banknote5)

    console.log(`$1 x ${n}`);

}

calCash(28)