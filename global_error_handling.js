new Promise((resolve, reject) => {
	throw new Error("This is from promise error thrown at line 2");
}).then((res) => {
	console.log("In in then");
	console.log(res);
}).catch((err) => {
	console.log("In catch");
	console.error(err);
});

process.on('unhandledRejection', (err) => {
	console.log("In unhandledRejection");
	console.log(err);
})
