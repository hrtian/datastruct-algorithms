/**
 * Merge sort implementation O(nLog(n))
 * @param {*[]} originalArray
 * @return {*[]}
 */

function mergeSort(originalArray) {
	if (originalArray.length <= 1) {
		return originalArray;
    }
    
    const middleIdx = originalArray.length >> 1;
    const leftArray = originalArray.slice(0, middleIdx);
    const rightArray = originalArray.slice(middleIdx);
    return merge(mergeSort(leftArray),mergeSort(rightArray));
}

/**
 * @param {*[]} leftArray
 * @param {*[]} rightArray
 * @return {*[]}
 */
function merge(leftArray, rightArray) {
    const help = [];
    while(leftArray.length > 0 && rightArray.length > 0) {
        help.push(leftArray[0] < rightArray[0] ? leftArray.shift() : rightArray.shift());
    }
    return help.concat(leftArray).concat(rightArray);
}
