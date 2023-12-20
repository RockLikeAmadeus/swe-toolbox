# Get Specific:

- [Matchers](matchers.md)

<br>

---

# Test Basics

Test-driven react development using Jest looks something like this:

```js
describe("Appointment", () => {
  it("renders the customer first name", () => {
    expect(document.body.textContent).toContain("Ashley");
  });
});
```

Run your tests with `$ npm test`

- The test files go in a directory which is separate from the `src` directory, since our tests should not have a one-to-one relationship with our application structure (as that would tightly-couple the two).
- The `describe()` function defines a _test suite_, and the first argument is the name of the **unit** you are testing.
- The `it()` function is the preferred test method; there are equivalents with different names, but `it` reads best.
  - The first argument to `it()` should be a string which is a present-tense phrase that describes the behavior we expect, as this reads very nicely in the Jest test results (assuming the name of the suite is the name of the unit, as in the above example).
- The `toContain()` function is an example of a _matcher_, and you can and should write your own matchers that are specific to your project.
- In order for `document` to be defined as we expect, we need to use the `jsdom` **test environment** (a piece of code that performs setup and teardown before and after your test runs, respectively), which gives us a headless DOM we can access in the Node runtime for stuff like this. We install and setup `jsdom` in the project setup steps (2 and 4, to be exact).

# Defining what usage of your component will look like (get to red)

Since this is TDD, you write something like the following _before creating your component_.

```js
import React from "react";
import ReactDOM from "react-dom/client";
import { act } from "react-dom/test-utils";
// `Appointment` is not the default export, which is intentional
import { Appointment } from "../src/Appointment";

describe("Appointment", () => {
  it("renders the customer first name", () => {
    const customer = {
      firstName: "Ashley",
    };
    const component = <Appointment customer={customer} />;
    const container = document.createElement("div");
    document.body.replaceChildren(container); // appendChild isn't actually recommended
    act(() => ReactDOM.createRoot(container).render(component));
    expect(document.body.textContent).toContain("Ashley");
  });
});
```

# Creating the component (get to green)

Create an empty file

```
$ mkdir src
$ touch src/Appointment.js
```

And then define an empty component in the new file to _do the smallest thing to fix the reported failure (in this case, the component being undefined)_:

```js
export const Appointment = () => {};
```

Then run your test to find the next thing to do:

```js
export const Appointment = () => "Ashley";
```

# Triangulate to remove hard coding

Triangulation is adding more specific tests, which requires more general production code, in a cycle. In this case, the first cycle would involve duplicating our test, but passing in (and then verifying) a different customer name.

Note (move this elsewhere): If you already have multiple tests and you want to work on one in isolation, we can skip a test by adding `.skip` like this:

```js
it.skip("renders another customer first name", () => {
```

To get both to pass, rewrite the component as:

```js
import React from "react";

export const Appointment = ({ customer }) => <div>{customer.firstName}</div>;
```

# Refactor

Refactoring applies to tests just as much as production code. Remember to DRY up your test code often. One way to do this is with the `beforeEach` block:

```js
describe("Appointment", () => {
  let container;
  beforeEach(() => {
    container = container = document.createElement("div");
  });

  const render = (component) => {
    document.body.replaceChildren(container);
    act(() => ReactDOM.createRoot(container).render(component));
  };

  it("renders the customer first name", () => {
    const customer = {
      firstName: "Ashley",
    };
    render(<Appointment customer={customer} />);
    expect(document.body.textContent).toContain("Ashley");
  });

  it("renders another customer first name", () => {
    const customer = {
      firstName: "Jordan",
    };
    render(<Appointment customer={customer} />);
    expect(document.body.textContent).toContain("Jordan");
  });
});
```

Be careful with using `let` in the `describe` scope--use the following rule: any variable declared in the `describe` scope should be assigned a new value in the `beforeEach` block (or in the first part of each test). See https://reacttdd.com/use-of-let.
