// Given a function fn, an array of arguments args, and an interval time t, return a cancel function cancelFn.

var cancellable = function(fn, args, t) {

    fn(...args);
    const interval = setInterval(()=> {
        fn(...args);
    }, t)

    const cancelFn = () => {
        clearTimeout(interval);
    }

    return cancelFn;

};