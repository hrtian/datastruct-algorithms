/**
 * Insert sort implementation O(n^2)
 * @param {*[]} originalArray
 * @param {function(a, b)} comparator
 * @retturn {*[]}
 */

 function insertSort(originalArray, comparator) {
     const array = [...originalArray];

     for (let i = 1; i < array.length; i++) {
         for(let j = i - 1; j >= 0; j--) {
             if(comparator(array[j],array[j+1]) > 0) {
                 [array[j], array[j+1]] = [array[j+1], array[j]]
             }
         }
     }
     return array
 }