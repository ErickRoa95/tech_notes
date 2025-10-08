// 1 1 2 3 5 8 13 21
const fibonacci = n =>{
  if (n <= 2) {
    return 1;
  }

  return fibonacci(n-1) + fibonacci(n-2);
}

console.log(fibonacci(6)) // 8
console.log(fibonacci(7)) // 13
console.log(fibonacci(8)) // 21