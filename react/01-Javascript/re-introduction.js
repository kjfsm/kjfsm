console.log('Hello'.length); // 5

console.log('hello'.charAt(0)); // "h"
console.log('hello, world'.replace('hello', 'goodbye')); // "goodbye, world"
console.log('hello'.toUpperCase()); // "HELLO"

// let はブロックスコープ

// const は変更できない

// var はブロックスコープではなく、変更できる

const array = [0, 1, 2, 3, 4, 5];

for (let value of array) {
  // 値に関する処理
  console.log(value);
}

for (let property in array) {
  // オブジェクトのプロパティに関する処理
  console.log(property);
}

// forinはArray に新しいプロパティを追加していたらそれもループでくりかえされてしまう
// できるだけforofを使う

// switch はフォールスルーする

switch (1) {
  case 1:
  case 2:
    console.log("1か2です");
    break;
  default:
    console.log("1や2じゃないです")
}

// オブジェクトを作る
// どっちでもいい

var obj = new Object();
var obj = {};
