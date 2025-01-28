# Basic Syntax

```js
let description = "";
let person = {
    fname:"Paul",
    lname:"Ken",
    age:18
};
for (let x in person){
    description += person[x] + " ";
}

let myPets = "";
let pets = ["cat", "dog", "hamster", "hedgehog"];
for (let pet of pets){
    myPets += pet + " ";
}

// JavaScript has function scope; functions get their own scope but other blocks
// do not.
if (true){
    var i = 5;
}
i; // = 5 - not undefined as you'd expect in a block-scoped language

// One of JavaScript's most powerful features is closures. If a function is
// defined inside another function, the inner function has access to all the
// outer function's variables, even after the outer function exits.
function sayHelloInFiveSeconds(name){
    var prompt = "Hello, " + name + "!";
    // Inner functions are put in the local scope by default, as if they were
    // declared with `var`.
    function inner(){
        alert(prompt);
    }
    setTimeout(inner, 5000);
    // setTimeout is asynchronous, so the sayHelloInFiveSeconds function will
    // exit immediately, and setTimeout will call inner afterwards. However,
    // because inner is "closed over" sayHelloInFiveSeconds, inner still has
    // access to the `prompt` variable when it is finally called.
}
sayHelloInFiveSeconds("Adam"); // will open a popup with "Hello, Adam!" in 5s
```

See more details syntax [here](https://learnxinyminutes.com/javascript/)

# Code Structure and Organization

Code structure and organization depends entirely on whether this is _browser-based_ JavaScript or _Node_ JavaScript.

## Browser-Based JavaScript

In the `<head>` or even the `<body>` section of your HTML page, you might put:

```html
<script src="assets/scripts/app.js" type="module"></script>
```

By specifying your JavaScript file as a `module`, you can use the modern `import` and `export` syntax.

# How it Works at a High Level
