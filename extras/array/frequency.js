// โจทย์: เขียนโปรแกรมเพื่อหาจำนวนครั้งที่ตัวเลขปรากฏในอาร์เรย์
// ตัวอย่าง Input: [1, 2, 3, 2, 1, 2, 3, 4, 2]
// ตัวอย่าง Output: {1: 2, 2: 4, 3: 2, 4: 1}

function frequency(arr) {
    let freq = {}
    for (let index = 0; index < arr.length; index++) {
        if(!freq[arr[index]]){
            freq[arr[index]] = 0
        }

        freq[arr[index]]++
    }
    return freq
}