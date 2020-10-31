/**
 * Linear Search implementation O(n) n
 * 
 * @param {*[]} array
 * @param {*} seekElement
 * @param {function(a, b)} comparator
 * @return {*[]}
 */

function linearSearch(array, seekElement, comparator) {
	const foundIdxs = [];
	array.forEach((item, index) => {
		if (comparator(item, seekElement)) {
			foundIdxs.push(index);
		}
    });
    return foundIdxs;
}
