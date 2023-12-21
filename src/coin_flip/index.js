const start = +process.argv[2] || 69
const end = +process.argv[3] || 89

const rand = Math.random()

function random() {
  return Math.floor(rand*(end - start) + start)
}

console.log(random())