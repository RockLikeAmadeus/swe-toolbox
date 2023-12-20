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
