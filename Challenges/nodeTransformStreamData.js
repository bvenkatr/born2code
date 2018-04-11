let {createReadStream} = require("fs");
let {Transform} = require("stream");

class ToObject extends Transform {
    constructor(options) {
        super(options);
    }

    _transform(chunk, enc, cb) {
        cb(null, JSON.parse(chunk.toString()));
    }
}

const stream = createReadStream("../data/testdata.json")
    .pipe(new ToObject({objectMode: true}));

stream.on("error", err => console.error(err));
stream.on("readable", () => {
    while (null !== (chunk = stream.read())) {
        console.log(chunk);
    }
});
