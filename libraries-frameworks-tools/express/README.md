New Express.js project:

```bash
$ npm init
$ npm install --save express body-parser
$ npm install --save-dev nodemon
```

Add the following script to the `scripts` section in `package.json`

```
    "start": "nodemon server.js"
```

And then add `server.js`

```js
const express = require('express');
const bodyParser = require('body-parser');

const app = express();

// middleware here

app.listen(5000);
```

TODO: Add how to set this up for TypeScript