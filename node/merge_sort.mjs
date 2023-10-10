function merge(arr, l, meio, r) {

    const x = meio - l + 1;
    const y = r - meio;
    
    const leftArr = [];
    const rightArr = [];
    
    for (let i = 0; i < x; i++) {
        leftArr.push(arr[l + i]);
    }
    for (let i = 0; i < x; i++) {
        rightArr.push(arr[meio + 1 + i]);
    }
    let a = 0;
    let b = 0;
    let c = l;
    while (a < x && b < y) {
        if (leftArr[a] <= rightArr[b]) {
            arr[c] = leftArr[a];
            a++;
        } else {
            arr[c] = rightArr[b];
            b++;
        }
        c++;
    }

    while (a < x) {
        arr[c] = leftArr[a];
        c++;
        a++;
    }
    while (b < y) {
        arr[c] = rightArr[b];
        c++;
        b++;
    }
}



export function mergesort(arr, l, r) {
    if (l < r) {
        const meio = l + Math.floor((r - l) / 2);
        mergesort(arr, l, meio);
        mergesort(arr, meio + 1, r);
        merge(arr, l, meio, r);
    }   
}

