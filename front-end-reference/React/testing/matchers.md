# List of Built-In Matchers:

https://jestjs.io/docs/expect

# Third-Party Matcher Libraries

Third-party matcher libraries can be installed as npm packages, and can be very useful. TODO: provide links to useful libraries...

# Matchers look like this:

```js
expect(document.body.textContent).toContain("Ashley");
```

`toContain()` is the matcher in the above snippet.

The argument to `expect()` should be the element(s) in question:

```js
document.querySelector("div#appointmentsDayView");
```

# Basic Matchers

```js
expect(document.body.textContent).toContain("Ashley");
```

```js
expect(document.querySelector("div#appointmentsDayView")).not.toBeNull();
```

```js
const listOfChildren = document.querySelectorAll("ol > li");
expect(listOfChildren).toHaveLength(2);
```

```js
expect(listOfChildren[0].textContent).toEqual("12:00");
```

# Creating Custom Matchers

## When to create a custom matcher

- When an expectation is length and/or doesn't read easily
- When expectation logic is duplicated across tests

## Setup

1. Create new directory `matchers` under the `test` directory.
2. Create a new file for the matcher (i.e. `toContainText.js`) and leave it empty
3. Create a new test file for the matcher, i.e. `toContainText.test.js`

## Writing the test

As usual, write these tests one at a time, writing the code to make each pass only after you've seen it fail.

```js
import { toContainText } from "./toContainText";

describe("toContainText matcher", () => {
  it("returns pass is true when text is found in the given DOM element", () => {
    const domElement = { textContent: "text to find" };
    const result = toContainText(domElement, "text to find");
    expect(result.pass).toBe(true);
  });
  it("returns pass is false when the text is not found in the given DOM element", () => {
    const domElement = { textContent: "" };
    const result = toContainText(domElement, "text to find");
    expect(result.pass).toBe(false);
  });
});
```
