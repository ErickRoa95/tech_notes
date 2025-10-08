// Arrow Functions showcases

const add = (a, b) => a + b;
const square = x => x * x;
const sayHello = () => console.log("Hello, World!");
const getPersonalData = () => ({ name: "John", age: 30 });

// Multiple Functions on a const. 
const double = x => x * 2;
const substractOne = x => x -1;
const triple = x => x * 3;
const add5 = x => x + 5;

const funcArray = [
  double,
  substractOne,
  triple,
  add5
];

var number = 10
funcArray.forEach(func => number = func(number));
console.log(number); 


// Return function from another function
const createMultipler = y => x => x * y;
/** This equivalent to
 * 
 * function createMultipler(y){
 *  return function(x){
 *    return x * y;
 *   }
 * }
 * 
*/

const multiplyBy2 = createMultipler(2);
const multiplyBy3 = createMultipler(3);

console.log(multiplyBy2(5)); // 10
console.log(multiplyBy3(5)); // 15