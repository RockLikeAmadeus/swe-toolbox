MERN is a common application stack used for building web applications. It utilizes MongoDB, Express.js, React.js, and Node.js.

![alt text](mern-stack-overview.png)

The connection between the client and the server can take many forms, and which technique is most appropriate [depends on your use case](../../general/design-and-architecture/communicating-between-services.md/over-the-web.md), although since we're using React, [GraphQL is likely a good candidate](https://academind.com/tutorials/graphql-with-node-react-full-app).

The backend could be split up any number of ways, for instance with a single server hosting both the react app _and_ the app's backend API, or we can have each of these handled by distinct web servers.