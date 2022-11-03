const fs = require('fs');
const lines = fs.readFileSync('input.txt', {encoding: 'utf-8'}).split('\n');

const disallowedSubstrings = ['ab', 'cd', 'pq', 'xy'];
let countValidStrings = 0;

function findThreeVowels(str) {
    const vowels = 'aeiou';
    let strCopy = str.split('');
    let count = 0;

    count += strCopy.filter(char => /[aeiou]/.test(char)).length;

    if (count >= 3) {
        // console.log('Contains >=3 vowels');
        return true;
    }

    return false;
}

function findLetterTwiceInARow(str) {
    for (let i = 0; i < str.length - 1; i++) {
        if (str[i] === str[i+1]) {
            // console.log(`Contains twice in a row: ${str.slice(i, i+2)} (at position ${i})`);
            return true;
        }
    }

    return false;
}

function findDisallowedSubstring(str) {
    for (let substr of disallowedSubstrings) {
        if (str.includes(substr))
            return true;
    }

    // console.log('Not included disallowed strings');
    return false;
}

lines.forEach(str => {
    if (findThreeVowels(str) && findLetterTwiceInARow(str) && !findDisallowedSubstring(str)) countValidStrings++;
});

console.log(`Valid strings found = ${countValidStrings}`);
