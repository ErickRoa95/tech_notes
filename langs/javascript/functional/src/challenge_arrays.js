const electionVotes = [
  'Erick', 
  'Erick',
  'Ivan',
  'Antony',
  'Erick',
  'Antony',
  'Ivan',
  'Erick',
  'Erick',
  'Antony',
];

const tallyVotes = votes =>{
  return votes.reduce((acc, vote) => ({
    ...acc, // spread the current values in the accumulator object.
    [vote]: acc[vote] ?  acc[vote] + 1 : 1,
  }), {});
}

console.log(tallyVotes(electionVotes));