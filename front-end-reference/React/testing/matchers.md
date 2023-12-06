Matchers look like this:

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
