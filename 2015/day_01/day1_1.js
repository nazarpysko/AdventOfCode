const fs = require('fs');
const line = fs.readFileSync('input.txt', {encoding: 'utf-8'});

let floor = 0;

line.split('').forEach((char)=> {
    if (char === '(')
        floor++;
    else
        floor--;
})

console.log(`Floors: ${floor}`);
