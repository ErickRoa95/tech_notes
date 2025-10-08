const  Person = {
  name: "Erick",
  age: 29
};

const CareerData = {
  title: "Software Engineer",
  company: "Wizeline",
};

// with '...' operator, we spread all the atribues in Person and CareerData to the new object data. 
const PersonWithDataCareer = {
  ...Person,
  ...CareerData,
};

console.log(PersonWithDataCareer);
