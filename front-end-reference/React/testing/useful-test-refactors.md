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

// import this to simplify rendering your component in your test
export const render = (component) =>
  act(() => {
    ReactDOM.createRoot(container).render(component);
  });

// import this to simplify click behavior and avoid using `act` in your test code
export const click = (element) => act(() => element.click());
```

You could also building a wrapper function for `describe()` called `describeReactComponent`. It could automatically includes a `beforeEach` block and builds a `container` variable that's accessible within the scope of the `describe` block.
