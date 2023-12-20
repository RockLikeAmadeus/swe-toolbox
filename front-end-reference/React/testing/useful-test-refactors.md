# Extract Reusable Rendering Logic

todo

# General-Purpose React Testing Helper File

Copy this exact code into a new file in the `test` directory of your project. Name this file something like `reactTestExtensions.js`

```js
import ReactDOM from "react-dom/client";
import { act } from "react-dom/test-utils";

// Import these to simplify your beforeEach() block
export let container;
export const initializeReactContainer = () => {
  container = document.createElement("div");
  document.body.replaceChildren(container);
};

// Import this to simplify rendering your component in your test
export const render = (component) =>
  act(() => {
    ReactDOM.createRoot(container).render(component);
  });

// Import this to simplify click behavior and avoid using `act` in your test code
// Use like this:
//
// const button = document.querySelectorAll("button")[1];
// click(button);
export const click = (element) => act(() => element.click());

// Import this to replace calls to `describe` for react components.
// Simplifies test code by performing initialization within.
export const describeReactComponent = (componentName, fn) => {
  beforeEach(() => {
    initializeReactContainer();
  });

  fn();
};
```
