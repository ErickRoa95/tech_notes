const Person = ({name,age, job}) =>{
  var _name = name;
  var _age = age;
  var _job = job 

  return {
    getName: () => _name,
    getAge: () => _age, 
    getJob: () => _job,
    setJob: newJob => _job= newJob,
  }
} 

const me = Person({name: "Erick", age: 29, job: "Developer"});

console.log(me._name); // undefined, onlly accessabe trhough getters and setters. 
console.log(me.getJob()); // Only accesabe the private values from get/setters.
me.setJob("DB Administrator");
console.log(me.getJob());