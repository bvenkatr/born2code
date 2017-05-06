/**
 * Find the longest word in a string
 * @param str
 * @returns {number}
 */
function findLongestWord(str) {
    var tempArray = str.split(" ");
    var tempVar = 0;
    tempArray.forEach(function(item){
        if(tempVar < item.length) {
            tempVar = item.length;
        }
    });
    return tempVar;
}

findLongestWord("The quick brown fox jumped over the lazy dog");


/**
 * 2) Convert givne string to titleCase
 *
 */
function titleCase(str) {
    var strArray = str.split(" ");
    return strArray.map(function(item){
        var firstEle = item[0];
        return item.split("").map(function(eachCharacter, i){
            if (i === 0) {
                return eachCharacter.toUpperCase();
            } else {
                return eachCharacter.toLowerCase();
            }
        }).join("");
    }).join(" ");
}

titleCase("I'm a little tea pot");
