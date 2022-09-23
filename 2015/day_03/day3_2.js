const fs = require('fs');
const line = fs.readFileSync('input.txt', {encoding: 'utf-8'});

let currentLocationSanta = {
    x: 0,
    y: 0
};

let currentLocationRoboSanta = {
    x: 0,
    y: 0
};

let visitedHouses = new Set();
visitedHouses.add(`${currentLocationSanta.x}x${currentLocationSanta.y}`);
    
function move(nextMove, currentLocation) {
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

line.split('').forEach((nextMove, step) => {
    let newLocation;

    if (step % 2 !== 0) {
        move(nextMove, currentLocationSanta);
        newLocation = `${currentLocationSanta.x}x${currentLocationSanta.y}`;
    } else {
        move(nextMove, currentLocationRoboSanta);
        newLocation = `${currentLocationRoboSanta.x}x${currentLocationRoboSanta.y}`;  
    }

    if (!visitedHouses.has(newLocation)) {
        visitedHouses.add(newLocation);
    }
});

console.log(`Total houses visited: ${visitedHouses.size}`);