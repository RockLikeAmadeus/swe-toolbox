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
4. Create a file (if it doesn't yet exist) to plug the test into jest, called something like `domMatchers.js`

## Writing the test

As usual, write these tests one at a time, writing the code to make each pass only after you've seen it fail (the following tests could use some refactoring to remove duplication).

```js
import { toContainText } from "./toContainText";

const stripTerminalColor = (text) => text.replace(/\x1B\[\d+m/g, "");

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
  it("returns a message that contains the source line if no match", () => {
    const domElement = { textContent: "" };
    const result = toContainText(domElement, "text to find");
    expect(stripTerminalColor(result.message())).toContain(
      `expect(element).toContainText("text to find")`
    );
  });
  it("returns a message that contains the source line if negated match", () => {
    const domElement = { textContent: "text to find" };
    const result = toContainText(domElement, "text to find");
    expect(stripTerminalColor(result.message())).toContain(
      `expect(element).not.toContainText("text to find")`
    );
  });
  it("returns a message that contains the actual text", () => {
    const domElement = { textContent: "text to find" };
    const result = toContainText(domElement, "text to find");
    expect(stripTerminalColor(result.message())).toContain(
      `Actual text: "text to find"`
    );
  });
});
```

Note about the `message`: it should provide a helpful string to display when the expectation fails, but is itself a _function that returns a string_ for performance reasons. `message` won't be used if `pass` is `true`, _unless_ it gets used alongside a `.not` qualifier.

Note about `stripTerminalColor`: a quick helper function to remove the ASCII escape codes that Jest uses to print results in different colors.

## Writing the matcher

```js
import { matcherHint, printExpected, printReceived } from "jest-matcher-utils";

export const toContainText = (received, expectedText) => {
  const pass = received.textContent.includes(expectedText);
  const sourceHint = () =>
    matcherHint("toContainText", "element", printExpected(expectedText), {
      isNot: pass,
    });
  const actualTextHint = () =>
    "Actual text: " + printReceived(received.textContent);
  const message = () => [sourceHint(), actualTextHint()].join("\n\n");
  return { pass, message };
};
```

## Plugging it in

In `domMatchers.js`

```js
import { toContainText } from "./toContainText";

expect.extend({ toContainText });
```

Then in `package.json`

```json
  ...
  "jest": {
    ...
    "setupFilesAfterEnv": ["./test/matchers/domMatchers.js"]
  }
```

If your tests are set to automatically detect changes and re-run, you need to Ctrl+C and then re-run the test manually to pick up these changes.

## Using your new matcher

```js
expect(document.body).toContainText("Ashley");
```

## Other concerns

The above works just fine when we are only using our custom matchers internally, but if you intend on exporting them for use in other projects, they should be more robust. For instance, they should check to be sure that the received value is, in fact, an HTML element.
