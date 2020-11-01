/**
 * Shell sort implementation. O(Log(n))
 * @param {*[]} originalArray
 * @return {*[]}
 */

function shellSort(originalArray) {
	const array = [...originalArray];

	// define a gap distance.
	let gap = array.length >> 1;

	while (gap > 0) {
		for (let i = 0; i < array.length - gap; i++) {
			let currentIdx = i;
			let gapShiftIdx = i + gap;

			while (currentIdx >= 0) {
				if (array[gapShiftIdx] < array[currentIdx]) {
					[array[gapShiftIdx], array[currentIdx]] = 
                    [array[currentIdx], array[gapShiftIdx]];
                }
                gapShiftIdx = currentIdx;
                currentIdx -= gap;
			}
        }
        gap = gap >> 1;
    }
    
    return array;
}
console.log(shellSort([1, 5, 1, 2, 4, 9, 51, 48, 23, 84, 0, 6]));

