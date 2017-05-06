/**
 *Convert the given string into lowercase and whitespaces with _(underscore)
 *
 */

function convertToLowerCase(str) {
    if (str !== undefined || str !== null) {
        return str.split(" ").map(function (item) {
            return item.split("").map(function (eachChar) {
                return eachChar.toLowerCase();
            }).join("");
        }).join("_");
    } else {
        return str;
    }
}

console.log(convertToLowerCase("rzr_rms"));
