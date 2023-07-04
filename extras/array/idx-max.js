// โจทย์: เขียนโปรแกรมเพื่อหาตำแหน่งแรกของตัวเลขในอาร์เรย์ที่มีค่ามากกว่าหรือเท่ากับที่กำหนด
// ตัวอย่าง Input: [5, 10, 15, 20, 25], ค่าที่กำหนด: 15
// ตัวอย่าง Output: 2

function findIdxOfMaxVal(arr) {

    if(arr.length==0)return -1

    let idx = 0
    let maxVal = arr[0]
    for (let index = 0; index < arr.length; index++) {
        const element = arr[index];
        if(element > maxVal) {
            idx = index
            maxVal = element
        }
    }
    return idx
}

console.log(findIdxOfMaxVal([5, 10, 15, 20, 25]));