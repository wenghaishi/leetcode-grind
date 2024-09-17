
class Calculator {

    constructor (value) {
        this.result = value;
    }

    add(value) {
        this.result += value;
        return this;
    }

    subtract(value) {
        this.result -= value;
        return this;
    }
}

const a = new Calculator(1)

a.add(1)
console.log(a)