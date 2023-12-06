# Creating a new React application

React applications should be built test-first, just like everything else. That means, according to YAGNI, we don't want anything in our applications that doesn't meet a needed specification, so we won't typically want to use `create-react-app` to avoid all of the extra stuff it includes.

## Pre-requisites

We need to have NPM, Jest, React, and Babel installed

### NPM

Make sure you have the latest NPM with

```
$ npm install npm@latest -g
```

### Jest

Alternatives exist, but Jest is fine for now.

### Babel

We need Babel for a few things. It gets installed with Jest, so we really just need to add presets and plugins (a preset is a set of plugins). Each plugin enables a specific feature of the ECMAScript standards, or a preprocessor, such as JSX.

## Steps

### 1. Create and navigate to a new directory for the project. Initialize it as a git repository and push it to a remote store, i.e. GitHub.

### 2. Install dependencies

`$ npm init` (the defaults are OK _except_ for the `test command` question: enter `jest`)

```
$ npm install --save-dev jest
$ npm install --save react react-dom
$ npm install --save-dev @babel/preset-env @babel/preset-react
$ npm install --save-dev @babel/plugin-transform-runtime
$ npm install --save @babel/runtime
$ npm install --save-dev jest-environment-jsdom
$ npm install --save-dev webpack webpack-cli babel-loader
$ echo "node_modules/" >> .gitignore
```

### 3. Enable the packages we've just installed by creating a new file `.babelrc` with the contents

```json
{
  "presets": ["@babel/env", "@babel/react"],
  "plugins": ["@babel/transform-runtime"]
}
```

### 4. Add the test environment we just installed to the project by opening `package.json` and adding the following section at the bottom:

```json
{
    ...,
    "jest": {
        "testEnvironment": "jsdom",
        "globals": {
          "IS_REACT_ACT_ENVIRONMENT": true
        }
    }
}
```

Note: this enables the React `act()` function, which pauses until asynchronous rendering has completed. It's useful, but not requierd for React testing. See https://reacttdd.com/understanding-act.

### 5. Add a `build` script using webpack. Also, if you'd like, add the `watchAll` flag to the `test` command to automatically re-run tests when changes are detected:

```json
  ...
  "scripts": {
    "test": "jest --watchAll",
    "build": "webpack"
  },
  ...
```

### 6. Create a place for initial sample data, for manual testing.

```
$ touch src/sampleData.js
```

It's probably easier to populate this file after you've created some initial components and have a better idea of the structure of data they'll expect, but you could also use this as an opportunity to define that structure.

See the example [here](appointments-example-app/src/sampleData.js).

### 7. Add an application entry point.

```
$  mkdir src
$ touch src/index.js
```

React apps are heirarchies of components, and the entry point should render the root component (note, we typically don't want to test-drive root components). Keep the entry point brief, and only use it to instantiate dependencies and to call `render()`:

```js
import React from "react";
import ReactDOM from "react-dom";
import { MyRootComponent } from "./MyRootComponent";
import { mySampleData } from "./sampleData";
ReactDOM.createRoot(document.getElementById("root")).render(
  <MyRootComponent data={mySampleData} />
);
```

### 8. Configure webpack in development mode (see webpack docs for setting up for production builds).

```
$ touch webpack.config.js
```

in webpack.config.js:

```js
const path = require("path");
const webpack = require("webpack");
module.exports = {
  mode: "development",
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        loader: "babel-loader",
      },
    ],
  },
};
```

### 9. Add the `index.html` file

```
mkdir dist
touch dist/index.html
```

In `index.html`:

```html
<!DOCTYPE html>
<html>
  <head>
    <title>My Application</title>
  </head>
  <body>
    <div id="root"></div>
    <script src="main.js"></script>
  </body>
</html>
```

### 10. Build and run the application

```
$ npm run build
```

Ignore main.js (since it's generated) with:

```
$ echo "dist/main.js" >> .gitignore
```

Then view it at `dist/index.html` in your browser.
