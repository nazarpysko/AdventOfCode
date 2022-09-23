const fs = require('fs');
const line = fs.readFileSync('input.txt', {encoding: 'utf-8'});

const MD5 = require("crypto-js/md5");

let nonce = 0; hash = '';

function findNonce(difficulty) {
    do {
        hash = MD5(line.concat(nonce++)).toString();
    } while(hash.slice(0, difficulty.length) !== difficulty);
    
    return --nonce;
}


console.log(`Nonce: ${findNonce('000000')}`);