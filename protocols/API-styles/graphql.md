Sources:
https://graphql.org/learn/

# Overview

GraphQL is a query language for your API, and a server-side runtime for executing queries using a type system you define for your data.

For example, a GraphQL service that tells you who the logged in user is (me) as well as that user’s name might look like this:
```graphql
type Query {
  me: User
}
 
type User {
  id: ID
  name: String
}

function Query_me(request) {
  return request.auth.user
}
 
function User_name(user) {
  return user.getName()
}
```

The first type defined above is called `Query` because it's the _Query_ type, also known as the _Root_ type, which is the top level and entry point of every GraphQL API.

[Find tools for GraphQL in any language](https://graphql.org/community/tools-and-libraries/)
(but there's also [this](https://github.com/Khan/genqlient) for Go)

# Querying a GraphQL API

Based on the above types and fields, a query might look like this:

```graphql
{
  me {
    name
  }
}
```

and might provide a response that looks like:

```json
{
  "me": {
    "name": "Luke Skywalker"
  }
}
```

Notice that the query has the exact same shape as the result.