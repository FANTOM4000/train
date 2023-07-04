// โจทย์: เขียนโปรแกรมเพื่อเรียงลำดับตัวเลขในอาร์เรย์จากน้อยไปหามาก
// ตัวอย่าง Input: [5, 3, 8, 2, 1]
// ตัวอย่าง Output: [1, 2, 3, 5, 8]

function sortAsc(arr) {
    for (let index = 0; index < arr.length; index++) {
        for (let k = 0; k < arr.length; k++) {
            if(arr[index] > arr[k]) {
                let tmp = arr[index];
                arr[index] = arr[k];
                arr[k] = tmp
            }
        }
    }
    return arr;
}