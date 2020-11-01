/**
 * Count sort implementation. O(n+max)
 *
 * 1. 当数字变化范围小的时候效果好；
 * 2. 仅适用于整数数组；
 *
 * @param {number[]} originalArray
 * @param {number} [smallestItem]
 * @param {number} [biggestItem]
 */

function countSort(
	originalArray,
	smallestItem = undefined,
	biggestItem = undefined
) {
	let detectedSmallestItem = smallestItem || 0;
	let detectedBiggestItem = biggestItem || 0;
	if (smallestItem === undefined || biggestItem === undefined) {
		detectedSmallestItem = Math.min(...originalArray);
		detectedBiggestItem = Math.max(...originalArray);
	}

	// The buckets will ho;d frequency of each number from originalArray;
	const buckets = Array(detectedBiggestItem - detectedSmallestItem + 1).fill(
		0
	);
	originalArray.forEach((item) => {
		buckets[item - detectedSmallestItem]++;
	});

	// calculate how many elements which are less than or equals for the given index.
	for (let bucketIdx = 1; bucketIdx < buckets.length; bucketIdx++) {
		buckets[bucketIdx] += buckets[bucketIdx - 1];
	}
	buckets.pop();
	buckets.unshift(0);

	const sortedArray = Array(originalArray.length).fill(null);
	for (let itemIdx = 0; itemIdx < originalArray.length; itemIdx++) {
		const item = originalArray[itemIdx];
		const itemSortedPosition = buckets[item - detectedSmallestItem];
		sortedArray[itemSortedPosition] = item;
		buckets[item - detectedSmallestItem]++;
	}
	return sortedArray;
}