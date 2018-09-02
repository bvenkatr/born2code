let _ = require("lodash");
let pdf = require("html-pdf");


items = [
  {name: "macbook air laptop", "company": "apple", "model": "BMX1"},
  {name: "samsung on 8 mobile", "company": "samsung", "model": "BMX1"}
];

// temphtml = `<html><head>my pdf file</head><body>`;
// _.forEach((eachItem, index) => {
//   if (item)
//     temphtml = `<h4></h4><h4>contact</h4>`;
//     temphtml =
// });
//
let temphtml = `<html><head>my pdf file</head><body><h4>name</h4><h4>contact</h4></body></html>`;

pdf.create(temphtml, options).toFile('./MyPdf.pdf', function (err, res) {
  if (err) return console.log(err);
  console.log(res);
});
