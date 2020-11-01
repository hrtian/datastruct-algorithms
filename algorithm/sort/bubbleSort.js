/**
 * Bubble sort implementation O(n^2)
 * @param {*[]} originalArray
 * @param {function(a, b)} comparator
 * @retturn {*[]}
 */

function bubbleSort(originalArray, comparator) {
	const array = [...originalArray];
	let swapped = false;
	for (let i = 1; i < array.length && swapped; i++) {
		swapped = false;
		for (let j = 0; j < array.length - i; j++) {
			if (comparator(array[j], array[j + 1]) > 0) {
				[array[j], array[j + 1]] = [array[j + 1], array[j]];
				swapped = true;
			}
		}

		if (!swapped) {
			return array;
		}
	}

	return array;
}
