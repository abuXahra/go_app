var myName = prompt('What is your name?');

function greeting(myName) {
    var result = 'Hello' + " " + myName;
    console.log(result);
}

greeting(myName);


function sumNumbers(num1, num2) {
    var result = num1 + num2;
    console.log(result);
}




var num1 = prompt('enter num1');
var num2 = prompt('enter num2');

sumNumbers(num1, num2)

var myString = "I am really hungry for some";

console.log(myString)

var myUpperString = myString.toUpperCase();

console.log(myUpperString);


var foods = ["grapes", "cheese", "bread", "butter"]


// console.log(`${myString} ${foods[0]}`);

// for (const foodItem of foods) {
//     console.log(`${myString} ${foodItem}`);
// }

// for (var i = 0; i < foods.length; i++) {
//     if (foods[i] % 2 === 0) {
//         console.log(`${myString} ${foods[i].toUpperCase()}`);
//     }
// }

var numArray = [1, 2, 3, 4, 5, 6, 7, 8, 9];

function evenNumber(arr) {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] % 2 === 0) {
            console.log(arr[i]);
        }
    }
}
evenNumber(numArray);



var foodss = ["grapes", "cheese", "bread", "butter"];
function evenUppcase(arr) {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] % 2 === 0) {
            console.log(`${arr[i].toUpperCase()}`);
        }
    }
}
evenNumber(foodss);