
console.log(getTotal(1000,1,1.25));

function getTotal(invest, year, interest) {
   let total = invest;
   console.log(`Year 0 (เงินต้น) = ${total}`);
   for (let index = 0; index < year; index++) {
        total+=calInterest(total,interest)
        console.log(`Year ${index+1} (เงินต้น) = ${total}`);
   }
   return total;
}

function calInterest(principal,interest){
    return principal*interest/100
}