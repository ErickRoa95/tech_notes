 const inputCriteria = {
  firstName: [
    value => value.length >= 2 
      ? '' 
      : "First name must be at least 2 characters long"
  ],
  lastName: [
    value => value.length >= 2 
      ? '' 
      : "Last name must be at least 2 characters long"
  ],
  zipCode:[
    value => value.length === 5 
      ? '' 
      : "Zip code must be 5 characters long",
  ],
  state: [
    value => value.length === 2 
      ? '' 
      : "state must be 2 characters long"
  ]
 }

const getErrorMessages = (inputs, criteria) =>{
  return Object.keys(inputs).reduce((acc, fieldName) =>[
    ...acc,
    ...criteria[fieldName].map(test => test(inputs[fieldName])).filter(msg => msg)
  ], []);

}

const currentInputValues = {
  firstName: 'J',
  lastName: 'Dr.1vce',
  zipCode: '123',
  state: 'NY'
}

console.log(getErrorMessages(currentInputValues, inputCriteria));