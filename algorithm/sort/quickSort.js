/**
 * Quick sort implementation by no in-place. O(nLog(n))
 * @param {*[]} originalArray
 * @return {*[]}
 */

function quickSort(originalArray) {
	const array = [...originalArray];

	if (array.length <= 1) {
		return array;
	}

	const lowArray = [];
	const highArray = [];

	const pivotElement = array.shift();
	const centerArray = [pivotElement];

	while (array.length) {
		const currentElement = array.shift();

		if (currentElement == pivotElement) {
			centerArray.push(currentElement);
		} else if (currentElement < pivotElement) {
			lowArray.push(currentElement);
		} else {
			highArray.push(currentElement);
		}
	}

	const lowArraySorted = quickSort(lowArray);
	const highArraySorted = quickSort(highArray);
	return lowArraySorted.concat(centerArray).concat(highArraySorted);
}
