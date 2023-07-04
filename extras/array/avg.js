// โจทย์: เขียนโปรแกรมเพื่อหาค่าเฉลี่ยของตัวเลขที่อยู่ในอาร์เรย์
// ตัวอย่าง Input: [10, 20, 30, 40, 50]
// ตัวอย่าง Output: 30

function avg(arr) {
    let sum = 0
    for (let index = 0; index < arr.length; index++) {
        const element = arr[index];
        sum += element
    }
    return sum / arr.length
}

console.log(avg([10, 20, 30, 40, 50]));