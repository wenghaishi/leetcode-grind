
function processData(arr) {
    const map = new Map()

    for (const str of arr) {
        const arrSplit= str.split(";")
        
        const fields = new Map()

        let name = ''

        for (const e of arrSplit) {
            const [key, value] = e.split("=")
            if (key === "Name") {
                name = value
            } else {
                fields.set(key, value)
            }
        }

        if (!map.has(name)) {
            map.set(name, fields)
        } else {
            const existingData = peopleMap.get(name)
            
        }



    }
}