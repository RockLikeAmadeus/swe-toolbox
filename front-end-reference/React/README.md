# Get Specific:

- [React Example](appointments-example-app)
- [Starting a new React project](starting%20a%20new%20project.md)
- [Testing](testing)

<br>

---

These notes and examples are compiled based on the following sources:

- react.dev/learn
- Mastering React Test-Driven Development by Daniel Irvine

# Concept Overview

React components are JavaScript functions that return markup. Because components have this structure, try to apply Uncle Bob's clean code rules for functions (short, descriptive, single-purpose, etc.).

```js
function MyButton() {
  return <button>I'm a button</button>;
}

export default function MyApp() {
  return (
    <div>
      <h1>Welcome to my app</h1>
      <MyButton />
    </div>
  );
}
```

You can always tell a React component apart from a regular HTML tag because components start with a capital letter.

You can't return multiple JSX tags. If you need to do something like this, wrap them into a shared parent like a `<div>` or an empty `<>` wrapper, like this:

```js
function AboutPage() {
  return (
    <>
      <h1>About</h1>
      <p>
        Hello there.
        <br />
        How do you do?
      </p>
    </>
  );
}
```

## Styling

Classes are specified withe `className` (rather than just `class` as in HTML):

```js
<img className="avatar" />
```

## Displaying Data

Curly braces (`{}`) let you "escape out of" markup and back into JavaScript:

```js
<img className="avatar" src={user.imageUrl} />
```

## Conditional Rendering

React doesn't have special syntax for writing conditions; regular JavaScript works. See the below examples (these should do essentially the same thing):

```js
let content;
if (isLoggedIn) {
  content = <AdminPanel />;
} else {
  content = <LoginForm />;
}
return <div>{content}</div>;
```

```js
<div>{isLoggedIn ? <AdminPanel /> : <LoginForm />}</div>
```

```js
<div>{isLoggedIn && <AdminPanel />}</div>
```

## Rendering Lists

React lists are rendered using standard JavaScript features like `for` loops and the array `map()` function, such as in:

```js
const products = [
  { title: "Cabbage", id: 1 },
  { title: "Garlic", id: 2 },
  { title: "Apple", id: 3 },
];

const listItems = products.map((product) => (
  <li key={product.id}>{product.title}</li>
));

return <ul>{listItems}</ul>;
```

Notice how <li> has a key attribute. For each item in a list, you should pass a string or a number that uniquely identifies that item among its siblings. Usually, a key should be coming from your data, such as a database ID. React uses your keys to know what happened if you later insert, delete, or reorder the items.

## Event Handling

```js
function MyButton() {
  function handleClick() {
    alert("You clicked me!");
  }

  return <button onClick={handleClick}>Click me</button>;
}
```

_Important: Do not call the event handler function: you only need to pass it down. Notice how `onClick={handleClick}` has no parentheses at the end._

## Managing Component State

First:

```js
import { useState } from "react";
```

Then call `useState() to initialize one or more _state variables_ in your component. The `useState()` function accepts the initial state as its argument, and returns the current state along with the function that you use to update the state.

```js
function MyButton() {
  const [count, setCount] = useState(0);

  function handleClick() {
    setCount(count + 1);
  }

  return <button onClick={handleClick}>Clicked {count} times</button>;
}
```

Any names are valid, but the convention is `something` and `setSomething`.

If a component is rendered multiple times, each instance gets its own state.

## Hooks

Hooks are functions that start with `use` like `useState()`. There are various built-in hooks and you can always write your own.

Hooks can only be called at the top level of components, or from other hooks (i.e. not in conditions or loops).

## Sharing Data Between Components

### Props

You can "lift the state up" to parent components.

```js
function MyButton({ count, onClick }) {
  return <button onClick={onClick}>Clicked {count} times</button>;
}

export default function MyApp() {
  const [count, setCount] = useState(0);

  function handleClick() {
    setCount(count + 1);
  }

  return (
    <div>
      <h1>Counters that update together</h1>
      <MyButton count={count} onClick={handleClick} />
      <MyButton count={count} onClick={handleClick} />
    </div>
  );
}
```

## Common Terms: JSX, Redux, Hooks, etc.

JSX is React's markup syntax. It is _technically_ optional.

# "Thinking in React"

Finish taking notes based on https://react.dev/learn/thinking-in-react

1. Mockup -> 2. Component Heirarchy -> 3. Static React version -> 4. Identify state ->

## 1. and 2.

See top level README "The Process of Building User Interfaces" (add link here).

## 3. Static React Version:

The most straightforward way to start is to build an initial version that renders the UI from your data model without any interactivity (don't use state at all for this version). The root-level component will take the data model as a prop, and pass it down the heirarchy via props as needed.

Implement components top down for simple interfaces, and bottom up for more complex ones. Remember, test-first.

## 4. Identify state

Identify the pieces of data in the static version that should be represented as state (that is, those data elements that change over time--data passed in from props or computed from other data are not state).

## 5.

## 6.

---

# Things to keep in mind

- You don't need a separate file for each component, particularly if the grouped components are small and closely-related, such as those that are composed of each other. However, ensure that the file itself is named after the root component of the heirarchy it contains.
