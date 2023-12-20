# Extract Reusable Rendering Logic

todo

# Useful General Purpose React Testing Helper File

Name this file something like `reactTestExtensions.js`

```js
// Import these to simplify your beforeEach() block
export let container;
export const initializeReactContainer = () => {
  container = document.createElement("div");
  document.body.replaceChildren(container);
};
```
