const debounce = function (fn, time) {
    let timer;

    const timeout = setTimeout(()=> {
        fn()
    }, time)

    clearTimeout()
};
