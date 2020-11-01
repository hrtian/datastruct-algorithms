/**
 * @param {number[]} originalArray
 * @param {number} [smallestElement]
 * @param {number} [biggestElement]
 */

function count(
	originalArray,
	smallestElement = undefined,
	biggestElement = undefined
) {
	let detectedSmallestElement = smallestElement || 0;
	let detectedBiggestElement = biggestElement || 0;
	if (smallestElement === undefined || biggestElement === undefined) {
		detectedSmallestElement = Math.min(...originalArray);
		detectedBiggestElement = Math.max(...originalArray);
	}

	const buckets = Array(
		detectedBiggestElement - detectedSmallestElement + 1
	).fill(0);

	originalArray.forEach((element) => {
		buckets[element - detectedSmallestElement]++;
	});

	for (let bucketIdx = 1; bucketIdx < buckets.length; bucketIdx++) {
		buckets[bucketIdx] += buckets[bucketIdx - 1];
	}
	buckets.pop();
	buckets.unshift(0);

	const sortedArray = Array(originalArray.length).fill(null);
	for (
		let elementIndex = 0;
		elementIndex < originalArray.length;
		elementIndex += 1
	) {
		const element = originalArray[elementIndex];
		const elementSortedPosition =
			buckets[element - detectedSmallestElement];
		sortedArray[elementSortedPosition] = element;
		buckets[element - detectedSmallestElement] += 1;
	}
	return sortedArray;
}
