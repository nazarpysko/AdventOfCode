const fs = require('fs');
const line = fs.readFileSync('input.txt', {encoding: 'utf-8'});

let currentLocation = {
    x: 0,
    y: 0
};

let visitedHouses = new Set();
visitedHouses.add(`${currentLocation.x}x${currentLocation.y}`);
    
function move(nextMove) {
    switch(nextMove) {
        case '^':
            currentLocation.y++;
            break;
        case 'v':
            currentLocation.y--;
            break;
        case '>':
            currentLocation.x++;
            break;
        case '<':
            currentLocation.x--;
            break;
    }
}

line.split('').forEach(nextMove => {
    move(nextMove);
    let newLocation = `${currentLocation.x}x${currentLocation.y}`;

    if (!visitedHouses.has(newLocation)) {
        visitedHouses.add(newLocation);
    }
});

console.log(`Total houses visited: ${visitedHouses.size}`);