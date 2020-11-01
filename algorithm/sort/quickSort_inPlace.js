/**
 * Quick sort implementation by in-place. O(nLog(n))
 * @param {*[]} originalArray
 * @param {number} lowIdx
 * @param {number} highIdx
 * @return {*[]}
 */
function quickSortInPlace(
	originalArray,
	lowIdx = 0,
	highIdx = originalArray.length - 1
) {
	if (lowIdx < highIdx) {
		const partitionIdx = partitionArray(originalArray, lowIdx, highIdx);
		quickSortInPlace(originalArray, lowIdx, partitionIdx - 1);
		quickSortInPlace(originalArray, partitionIdx + 1, highIdx);
	}
	return originalArray;
}

/**
 *
 * @param {*[]} array;
 * @param{number} lowIdx;
 * @param{number} highIdx;
 * @return {number}
 */

function partitionArray(array, lowIdx, highIdx) {
	const swap = (leftIdx, rightIdx) => {
		[array[leftIdx], array[rightIdx]] = [array[rightIdx], array[leftIdx]];
	};

    // take a random element as pivot
    const pivot = array[highIdx];
    // swap the elements less than the pivot to the front of the array;
    let partitionIdx = lowIdx;
	for (let currentIdx = lowIdx; currentIdx < highIdx; currentIdx++) {
		if (array[currentIdx] < pivot) {
			partitionIdx != currentIdx && swap(partitionIdx, currentIdx);
			partitionIdx++;
		}
    }
    // swap partitionIdx to the array boundary;
	swap(partitionIdx, highIdx);
	return partitionIdx;
}
