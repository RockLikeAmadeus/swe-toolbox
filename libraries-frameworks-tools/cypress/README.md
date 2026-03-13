- [Finding page elements](./finding-page-elements.md)
- [Simulating user interaction](./simulating-user-interaction.md)
- [Writing assertions](./writing-assertions.md)

# Project Setup

(I think) Cypress gets added to a web project, such as a React project.

```bash
$ npm install cypress
$ npx cypress open
```

Select E2E Testing in the application that launches, then click Continue. At this point, you may run into an error. You can fix this by removing this line from `package.json`: `"type": "module",`.

Next, choose your preferred browser, then in the browser that opens, choose "Create new spec"

# Run Tests

```bash
$ npx cypress open
```