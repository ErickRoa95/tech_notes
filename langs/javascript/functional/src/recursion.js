const countdown = x => {
  if (x < 0) return;
  console.log(x);
  countdown(x - 1);
}

const countup = (x,max) =>{
  if (x > max) return;
  console.log(x);
  countup(x + 1, max);
}

countdown(10);
console.log("-----");
countup(1,10);