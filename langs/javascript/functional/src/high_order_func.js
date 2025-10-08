const divide = (x,y) => x/y; // based Function. 


const secondArgumentIsntZero = func => 
  (...args) => { // arguments of the func.
      if (args[1] === 0) { // second argument. 
        console.log('Error: Divding by zero.');
        return null
      }

      return func(...args); // no issues returnt he original func with arguments. 
  }

const divideSafe = secondArgumentIsntZero(divide)

console.log(divideSafe(7,0));
console.log(divideSafe(10,5));