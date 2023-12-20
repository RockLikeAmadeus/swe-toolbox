# Extract Reusable Rendering Logic

todo

# General-Purpose React Testing Helper File

Copy this exact code into a new file in the `test` directory of your project. Name this file something like `reactTestExtensions.js`

```js
// Import these to simplify your beforeEach() block
export let container;
export const initializeReactContainer = () => {
  container = document.createElement("div");
  document.body.replaceChildren(container);
};
```
