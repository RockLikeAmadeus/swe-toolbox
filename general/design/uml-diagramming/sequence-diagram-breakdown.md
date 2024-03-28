Compile, with images, from https://developer.ibm.com/articles/the-sequence-diagram/

A sequence diagram is less about the messages and more about the order in which messages are passed, and to/from whom. The vertical dimension shows time sequence of messages/calls as they occur (top-down). The horizontal dimension shows the object _instances_ between which messages are exchanged.

Objects are specified as "lifelines", the syntax for which is "instanceName : ClassName". It is perfectly valid to have only the instance name (during up-front design, when class breakdowns aren't fleshed out) or only the class name (when we want to be specific about roles without specifying who plays those roles). The lifeline then has a vertical line extending down from the name box--this line represents the object itself for message passing.

**To show an object  sending a message to another object**, a solid, horizontal line is drawn from the sending object (line) to the receiving object (line). A solid arrowhead (--|>) is used for synchronous calls and a stick arrowhead (-->) for asynchronous messages. The message or method name is then placed above the arrow line.

When a method name is used, the name refers to a method on the receiving object (a diagram helps describe this). e.g.

```
[A]				 [B]
 |  methodOfB()   |
 |--------------|>|
 |				  |
```

In this example, A calls synchronous method `methodOfB` on B.

Dotted lines under these method calls that indicate "return messages" are sometimes included, but optional. They can be useful when a value is returned, for ease of interpretation.

It is also valid to show an object "sending a message to itself" (i.e. calling its own method). This is of course common, but should only be included in the diagram when it adds clarity or when we want to draw attention to a specific internal method call. (include image here)

# More Advanced Stuff

## Guards

## Alternatives

## Options

## Loops

## Gates

## Break

## Parallel