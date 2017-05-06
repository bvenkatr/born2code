/**
 * https://www.freecodecamp.com/challenges/return-largest-numbers-in-arrays
 * @param arr
 * @returns {Array}
 */
function largestOfFour(arr) {
    var res = [];
    arr.forEach(function(item){
        var lNum = 0;
        item.forEach(function(num){
            if (lNum < num){
                lNum = num;
            }
        });
        res.push(lNum);
    });
    return res;
}

largestOfFour([[4, 5, 1, 3], [13, 27, 18, 26], [32, 35, 37, 39], [1000, 1001, 857, 1]]);
