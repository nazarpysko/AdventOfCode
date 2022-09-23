const fs = require('fs');

const line = fs.readFileSync('input.txt', {encoding: 'utf-8'});

let floor = 0, position, basementVisited = false;

line.split('').forEach((char, index)=> {
    if (char === '(')
        floor++;
    else
        floor--;

    if (floor === -1 && !basementVisited) {
        position = index + 1;
        basementVisited = true;
    } 
})

console.log(`Floors: ${floor} and first position to basement: ${position}`);
