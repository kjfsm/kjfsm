var a = [1, 2, 3, 4, 5];
var b = [2, 4];

var d = a.filter((x, i, self) => !b.includes(x));

console.log(d); // ほしい [1.3.5]
