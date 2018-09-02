var fs = require('fs');
var CsvReadableStream = require('csv-reader');

let finalData = [];
var inputStream = fs.createReadStream('reserve_20180423.csv', 'utf8');
let first = 1;
inputStream
    .pipe(CsvReadableStream({ parseNumbers: true, parseBooleans: true, trim: true }))
    .on('data', function (row) {
        if (first === 1) {
            // console.log('A row arrived: ', row);
            finalData.push(row);
            first++
        }
    })
    .on('end', function (data) {
        console.log("from analyze file");
        console.log(finalData);
    });
