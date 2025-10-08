const add = (a, b, c) => a + b + c; // normal function

const addPartial = a => // Partial application
  (b,c) => add(a,b,c);

// To call fully addPartial we need to call it twice. 
// addPartial(5)(2,3) => add(5,2,3)

// We can create a new function by fixing the first argument of addPartial
const add5 = addPartial(5); // b,c => add(5,b,c)

console.log(add5(2,3)); // 10
console.log(add5(10,20)); // 35
