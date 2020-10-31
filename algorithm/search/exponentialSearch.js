            /**
 * Exponential search implementation
 * @param {*[]} sortedArray
 * @param {*} seekItem
 * @param {number} n
 * @return {number}
 */

function exponentialSearch(sortedArray, seekItem) {
    const arraySize = sortedArray.length;
	if (sortedArray[0] === seekItem) {
		return 0; 
	}

	let n = 1;
	while (n < arraySize && sortedArray[n] <= seekItem) {
		n <<= 1;
	}

	let startIdx = n >> 1;
	let endIdx = Math.min(arraySize, n);

	while (startIdx <= endIdx) {
        const middleIdx = startIdx + (endIdx - startIdx) >> 1;
        
        if(seekItem === sortedArray[middleIdx]) {
            return middleIdx;
        }

        if(seekItem > sortedArray[middleIdx]) {
            startIdx = middleIdx + 1;
        } else {
            endIdx = middleIdx - 1;
        }
    }
    
    return -1;
}

console.log(exponentialSearch([1,2,3,4,5,6,7,8,9,10], 5))