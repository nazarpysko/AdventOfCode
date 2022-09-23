const fs = require('fs');
const lines = fs.readFileSync('input.txt', {encoding: 'utf-8'}).split('\n');

let totalWrappingPaper = 0, totalRibbon = 0;

lines.forEach(dimensions => {
    let dim = dimensions.split('x');
    dim.sort(function(a, b){return a - b});
    
    let lxw = dim[0]*dim[1];
    let wxh = dim[1]*dim[2];
    let lxh = dim[0]*dim[2];
    
    let slack = Math.min(lxw, wxh, lxh);

    totalWrappingPaper += 2*lxw + 2*wxh + 2*lxh + slack;

    totalRibbon += dim.reduce((total, dim) => total *= dim);
    totalRibbon += 2*dim[0] + 2*dim[1];
});

console.log(`Total wrapping paper: ${totalWrappingPaper} sqf\nTotal ribbon: ${totalRibbon} ft`);
