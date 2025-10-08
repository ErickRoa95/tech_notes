// Instead of returning an array, return true or false for the value. 

const number = [1,2,3,4,5,6,7,8,9,10];
const isOdd = x => x % 2 === 1;

// # Every example.
const everyOdd = number.every(isOdd);
console.log(everyOdd); // Return false, sisnce no every number is Odd. 

// # Some example.
const someOdd = number.some(isOdd);
console.log(someOdd); // Return false is at least one number is odd. 