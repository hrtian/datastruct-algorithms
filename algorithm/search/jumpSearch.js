/**
 * Jump Search implementation √n
 * @param {*[]} sortedArray
 * @param {*} seekItem
 * @param {function(a, b)} comparator
 * @return {number}
 */

function jumpSearch(sortedArray, seekItem, comparator) {
	const arraySize = sortedArray.length;

    // empty array
	if (!arraySize) {
		return -1;
	}
	// calculate the optimal jump size
    // the worst number of comparisons will be ((arraySize/jumpSize)+jumpSize - 1)
    const jumpSize = Math.floor(Math.sqrt(arraySize));
    let blockStart, blockEnd = 0, jumpSize;

    // find the bolock where the seekItem belong to;
    while(comparator(seekItem, sortedArray[Math.min(blockEnd, arraySize) - 1]) > 0) {
        blockStart = blockEnd;
        blockEnd += jumpSize;
        if(blockStart > arraySize) {
            return -1;
        }
    }
    
    // Do linear search for seekItem in subArray starting from blockStart;
    let currentIdx = blockStart;
    while(currentIdx < Math.min(blockEnd, arraySize)) {
        if(comparator(seekItem, sortedArray[currentIdx]) = 0) {
            return currentIdx;
        }
        currentIdx++;
    }

    return -1;
}
