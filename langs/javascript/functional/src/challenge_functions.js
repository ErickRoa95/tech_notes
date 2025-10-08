// Implement a function 'map' that takes an array and a function as arguments.
// array: [1,2,3]
// function: x => x * 2
// return a new array with all values modified by the function.
// expected output: [2,4,6]
const map = (arr, func) => {
  let newArray = []
  for (let i=0; i < arr.length; i++){
    const result = func(arr[i])
    newArray.push(result);
  }
  return newArray;
}

const double = x => x * 2 ;
const negative = x => -x ;
const upperCase = x => x.toUpperCase() ;

// return array with all values doubled.
console.log(map([1,2,3], double));

// return array with all values negative.
console.log(map([1,2,3],  negative));

// return array with all values in upperCase.
console.log(map(['a','b','c'], upperCase))

