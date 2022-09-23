const fs = require('fs');
const line = fs.readFileSync('input.txt', {encoding: 'utf-8'});

const MD5 = require("crypto-js/md5");

let nonce = 0; hash = '';

do {
    hash = MD5(line.concat(nonce++)).toString();
} while(hash.slice(0, 5) !== '00000');

console.log(`Nonce: ${--nonce}`);