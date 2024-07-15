// Write a function argumentsLength that returns the count of arguments passed to it.


var argumentsLength = function(...args) {
    let counter = 0;
    args.forEach((e)=> {
        counter++;
    })

    return counter;
};