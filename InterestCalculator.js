let moment = require("moment");

var interestCalculator = (principal, interest, startDate, endDate, byMonth, byDay) => {
	let finalResult = {};
	let interestPerDay = ((principal*interest)/100)/30;
    console.log(interestPerDay)
	let totalDays = moment(endDate, "YYYY-MM-DD").diff(moment(startDate, "YYYY-MM-DD"), "days") + 1;
    console.log(totalDays)
	let totalInterestByDay = totalDays * interestPerDay;
	console.log(totalInterestByDay)
	
	let totalMonths = moment(endDate, "YYYY-MM-DD").diff(moment(startDate, "YYYY-MM-DD"), "months") + 1;
	console.log(totalMonths)
	let totalInterestByMonth = ((principal*interest*totalMonths)/100)
	console.log(totalInterestByMonth)
	finalResult.byDay = totalInterestByDay;
	finalResult.byMonth = totalInterestByMonth;
	return finalResult;
}

console.log(process.argv);
let res = interestCalculator(process.argv[2], process.argv[3], process.argv[4], process.argv[5])
console.log("final Result is :", res)
