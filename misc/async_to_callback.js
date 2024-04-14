const util = require("util");

let fn = async() => {
	return "Hello World!";
}

//(async () => {
//	let res = await fn();
//	console.log(res);	
//})();

//const callbackFunction = util.callbackify(fn);

//callbackFunction((err, data) => {
//	if (err) return err;
//	console.log(data);
//});

fn().then((data) => {console.log(data)}).catch((error) => {console.log(error)})

