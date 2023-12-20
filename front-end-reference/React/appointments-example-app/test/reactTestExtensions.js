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
// Use like this:
//
// const button = document.querySelectorAll("button")[1];
// click(button);
export const click = (element) => act(() => element.click());
