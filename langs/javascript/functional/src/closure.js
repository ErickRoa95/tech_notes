const createPrinter = () => {
  const favoriteNumber = 42; // variable will be capture by the closure of the inner function
  return () => console.log(`My Favorite number is ${favoriteNumber}`) ;
}

createPrinter()(); // ()() invokes the returned function immediately
const printer = createPrinter();
printer(); // My Favorite number is 42