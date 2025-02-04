const a = ",hi my name is"

const b = 'hi my'

const c = "hi my name is"

const d = "hi. ,whats up?"

const e = d.replace(/\. ,/g, '. ')

console.log(a.includes(b))
console.log(a.includes(c))

console.log(e)
