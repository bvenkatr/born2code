var luxon1 = require("luxon1");
var luxon2 = require("luxon2");

luxon1.Settings.defaultZoneName = 'Pacific/Honolulu';

console.log(luxon1.DateTime.local().zoneName);
console.log(luxon1.DateTime.fromISO('2018-05-01'));
console.log(luxon2.DateTime.local().zoneName);
console.log(luxon2.DateTime.fromISO('2018-05-01'));
