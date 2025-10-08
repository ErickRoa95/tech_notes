
const numbers = [8,3,5,1,2,8,10,9,6];

const ascending = (a,b) =>{
  if (a<b) return -1; // a should come before b
  if(a>b) return 1; // b should come before a
  return 0; //leave elemnts in whatever order we found them. 
}

const sortedNumbers = numbers.slice().sort(ascending);
console.log(numbers);
console.log(sortedNumbers);