const fs = require('fs');
const lines = fs.readFileSync('input.txt', {encoding: 'utf-8'}).split('\n');

let countValidStrings = 0;

function findPairs(str) {
    for (let i = 0; i < str.length - 1; i++) {
        let substr = str.slice(i, i+2);
        if (str.includes(substr, i+2)) {
            console.log(`Contains a pair: ${substr}`);
            return true;
        }
    }

    return false;
}

function findPattern(str) {
    for (let i = 0; i < str.length - 1; i++) {
        if (str[i] === str[i+2] && str[i] !== str[i+1]) {
            console.log(`Contains pattern: ${str.slice(i, i+3)}`);
            return true;
        }
    }
    
    return false;
}

lines.forEach(str => {
    if (findPairs(str) && findPattern(str)) countValidStrings++;
});

console.log(`Valid strings found = ${countValidStrings}`);
