/**
 *Convert the given string into lowercase and whitespaces with _(underscore)
 *
 */

/**
 * Level 1 answer
 */
function convertToLowerCase1(str) {
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

console.time("Level 1 answer took");
console.log(convertToLowerCase1("Market Pace"));
console.timeEnd("Level 1 answer took");

/**
 * Level 2 answer
 * I found this answer with kesav help
 */
function convertToLowerCase2(str) {
    return str ? str.toLowerCase().replace(" ", "_") : str;
}

console.time("Level 2 answer took");
console.log(convertToLowerCase2("Market Pace"));
console.timeEnd("Level 2 answer took");
