import ReactDOM from "react-dom/client";
import { act } from "react-dom/test-utils";

// Import these to simplify your beforeEach() block
export let container;
export const initializeReactContainer = () => {
  container = document.createElement("div");
  document.body.replaceChildren(container);
};

// Import this to replace calls to `describe` for react components.
// Simplifies test code by performing initialization within.
export const describeReactComponent = (componentName, fn) => {
  fnWrapper = () => {
    beforeEach(() => {
      initializeReactContainer();
    });
    fn();
  };

  describe(componentName, fnWrapper);
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

// Wraps document.querySelector()
export const getElement = (selector) => document.querySelector(selector);

// Wraps document.querySelectorAll() and returns an array
export const getElements = (selector) =>
  Array.from(document.querySelectorAll(selector));

// Maps an array of elements to an array of their types
export const typesOf = (elements) => elements.map((element) => element.type);

// Maps an array of elements to an array of their text contents
export const textOf = (elements) =>
  elements.map((element) => element.textContent);
