// Mapping allow us to modify the values of the target array based on an instructions. 

const numbers = [1,2,4,5,6];

// double function that will be pass to map function. 
const double = x => x * 2;

const doubleNumbers = numbers.map(double);

console.log(doubleNumbers);