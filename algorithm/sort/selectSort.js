/**
 * Select sort implementation O(n^2)
 * @param {*[]} originalArray
 * @param {function(a, b)} comparator
 * @retturn {*[]}
 */
function selectSort(originalArray, comparator) {
	const array = [...originalArray];

	for (let i = 0; i < array.length - 1; i++) {
		let minIdx = i;
		for (let j = i + 1; j < array.length; j++) {
			if (comparator(array[minIdx], array[j]) > 0) {
				minIdx = j;
			}
		}

		if (minIdx != i) {
			[array[i], array[minIdx]] = [array[minIdx], array[i]];
		}
	}

	return array;
}