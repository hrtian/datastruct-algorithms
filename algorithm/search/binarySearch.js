/**
 * Binary search implementation log(n)
 *
 * @name binarySearch
 * @param {*[]} sortedArray
 * @param {*} seekItem
 * @param {function(a, b)} comparator
 * @return {number}
 */

function binarySearch(sortedArray, seekItem, comparator) {
	let startIdx = 0;
	let endIdx = sortedArray.length - 1;
	while (startIdx <= endIdx) {
		const middleIdx = (startIdx + (endIdx - startIdx)) >> 1;
		if (comparator(seekItem, sortedArray[middleIdx]) == 0) {
			return middleIdx;
		}

		if (comparator(seekItem, sortedArray[middleIdx]) == 1) {
			startIdx = middleIdx + 1;
		} else {
			endIdx = middleIdx - 1;
		}
	}

	return -1;
}
