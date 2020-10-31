/**
 * Interpolation search implementation.
 * @param {*[]} sortedArray - sorted array with uniform distributed values;
 * @param {*} seekItem
 * @return {number}
 */

function interpolateSearch(sortedArray, seekItem) {
	let leftIdx = 0;
	let rightIdx = sortedArray.length - 1;

	while (leftIdx <= rightIdx) {
		const idxDelta = rightIdx - leftIdx;
		const valueDelta = seekItem - sortedArray[leftIdx];
		const rangeDelta = sortedArray[rightIdx] - sortedArray[leftIdx];

        // if valueDelta is less than zero, it means that the minimum of sortedArray is more than the seekItem;
		if (valueDelta < 0) {
			return -1;
		}

        // 
		if (!rangeDelta) {
			return sortedArray[leftIdx] === seekItem ? leftIdx : -1;
		}
        // Do interpolation of the middle index.
		const middleIdx = leftIdx + Math.floor((valueDelta * idxDelta) / rangeDelta);
		if (sortedArray[middleIdx] === seekItem) {
			return middleIdx;
		}

		if (sortedArray[middleIdx] < seekItem) {
			leftIdx = middleIdx + 1;
		} else {
			rightIdx = middleIdx - 1;
		}
	}
	return -1;
}

