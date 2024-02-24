#!/usr/bin/node

const values = ['lamaie', 'mere', 'pere', 'inpanere', 'lamai', 'chiftele', 'legume', 'dulai']
const users = ['mircea', 'alin', 'sorin', 'elena']

function get_random_string(){
  const randomIndexValues = Math.floor(Math.random() * values.length)
  const randomIndexUsers = Math.floor(Math.random() * users.length)
  return [users[randomIndexUsers], values[randomIndexValues]]
}

setInterval(() => {
  let [usr, vals] = get_random_string()
  console.log(`${usr}:${vals}`);
}, 1000);