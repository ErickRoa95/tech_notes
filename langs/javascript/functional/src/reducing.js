const numbers = [1,2,3,4,5,6,7];


// acc is acumulated number. 
// element is the current value in the array from the looping. 
// {},0 => ACC initiated value in this scenario is 0.
const sum = numbers.reduce((acc, element) =>{
  return acc + element;
}, 0);

console.log(sum);