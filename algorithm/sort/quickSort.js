/**
 * Quick sort implementation by no in-place. O(nLog(n))
 * @param {*[]} originalArray
 * @return {*[]}
 */

function quickSort(originalArray) {
    const array = [...originalArray];

    if(array.length <= 1) {
        return array;
    }

    const lowArray = [];
    const highArray = [];

    const pivot = array.shift();
    const centerArray = [pivot];

    while (array.length) {
        const currentItem = array.shift();
  
        if (currentItem == pivot) {
          centerArray.push(currentItem);
        } else if (currentItem < pivot) {
          lowArray.push(currentItem);
        } else {
          highArray.push(currentItem);
        }
      }

    const lowArraySorted = quickSort(lowArray);
    const highArraySorted = quickSort(highArray);
    return lowArraySorted.concat(centerArray).concat(highArraySorted);
}
console.log(quickSort([1, 5, 1, 2, 4, 9, 51, 48, 23, 84, 0, 6]));