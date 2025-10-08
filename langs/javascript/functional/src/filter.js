// Filtering, allow us to extract specific values from an array based on an instruction. 

// # Example 1: Get evenNumbers
const numbers = [1, 2,3,4,5,6,7,8,9,10];

const isEven = x => x % 2 === 0;

const evenNumbers = numbers.filter(isEven);
console.log(evenNumbers);

// #Example2: Get words longer than 5. 

const words = [
  "Hello","World", "EstadosUnidos", "Chocolate", "Lira", "Rojo"
];

const longerThan5 = words.filter(w => w.length > 5);
console.log(longerThan5); 