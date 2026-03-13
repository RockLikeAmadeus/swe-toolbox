# Get Specific:

- [HTML](HTML)
- [CSS](CSS)
- [JavaScript](Frontend%20JavaScript)
- [React](./React/README.md)

<br>

---

# The Process of Building User Interfaces

## 1. Map out requirements/stories

Think about what the system should do and how that would work. Think about what information you expect the user to provide to it, and what the outputs would be, and how it should communicate the output to the user (should it be continuously available, or on demand, or by push notifications?). These requirements should be first defined as use cases, and then those use cases can be "refined into one or more [UML] sequence diagrams" (see https://developer.ibm.com/articles/the-sequence-diagram/)

## 2. Use the requirements to break out the different views/tabs/windows that wil be necessary

## 3. Use the results of 1. and 2. to mock up the UI

Use a (ideally low-fidelity) tool for this (try out Pencil Project, there might be multiple versions, but we want to try the open-source Electron-based one) (or, check out this link https://www.zenframe.com/post/top-open-source-wireframe-tools) (or just draw with a tablet)

## 4. Use the mockup to map out a component tree

Can Pencil Project, or a different diagramming tool, help with visualizing this? Realistically, a tabbed-and-bulleted list should work plenty well.

## 5. Build each component, using the TDD process (red, green, refactor).

Should you start with the "leaf" components and work up? Or start with the "root" components and work down? I think trial and error has the answer to this question, but I initially lean root-first. According to the react website, "in simpler examples, it’s usually easier to go top-down, and on larger projects, it’s easier to go bottom-up."
