[Aspects of Software Engineering](general/aspects-of-software-engineering.md) <br>
[Design and Architecture](general/design-and-architecture/README.md) <br>
[Working with Legacy Code](general/legacy-code/legacy-code-tools.md) <br>
[Testing/TDD](general/testing/README.md) <br>
[Code-Review Checklist](general/processes/code-review-checklist.md) <br><br>

[I need to build myself a tool](./general/i-need-to-build-myself-a-tool.md)

###### Languages
[Language Reference](./language-reference/README.md) <br>

###### Stack-Specific
[Front End Reference](./front-end-reference/README.md) <br>

###### Applications Spaces
[Embedded Reference](./embedded-reference/README.md) <br>
[Game Dev Reference](./game-dev-reference/README.md) <br>

###### Full Application Stacks
[MERN Stack](./application-stacks/MERN/README.md)

###### More
[Network Programming](./network-programming/) <br>
[Protocols](./protocols/README.md) <br>
[Third-Party Tools](tools/README.md) <br>
[Other Resources](resources/README.md) <br>

<br><br><br>

---

# Building Software

At a high-level, the process of building software should look like this the process described [here](./general/design-and-architecture/behavior-driven-development/README.md#high-level-process).

Somewhere between steps 2 and 3 of this process, you additionally want to come up with a high-level system design, and then build up what [the GOOS book](https://www.amazon.com/Growing-Object-Oriented-Software-Guided-Tests/dp/0321503627) calls a walking skeleton based on that design. That way, as you start to build out acceptance tests, you already have the structure in place such that all that remains is a matter of implementing features. It's important to note that the design on which the skeleton is based doesn't need to be perfect or be the final design, but should support end-to-end use cases. From the book:

> The development of a "walking skeleton" is the moment when we start to make choices about the high-level structure of our application. We can't automate the build, deploy, and test cycle without _some_ idea of the overall project structure. We don't need much detail yet, just a broad-brush picture of what major system components will be needed to support the first planned release and how they will communicate. Our rule of thumb is that we should be able to draw the design for the "walking skeleton" in a few minutes on a white board.

If your project will have a user interface, then  that high-level design will need to include what the UI will look like. Follow [these steps](./front-end-reference/README.md#the-process-of-building-user-interfaces) when you get to that point.

Before building up that skeleton, you'll also need to think about the data models that will be created and passed around that represent the real-world entities mirroring your domain, as well as the separation of components and the APIs between those components. 

At a high-level, these steps then are

1. Determine high-level goals and features
2. Define specific examples
3. Create the high-level system and UI design
4. Build a walking skeleton
5. Write a failing acceptance test for the first feature
6. Implement the feature, using TDD recursively
7. Go back to step 5

##### Other resources on starting a new software project
https://quii.gitbook.io/learn-go-with-tests/testing-fundamentals/scaling-acceptance-tests

Also check out https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-template-repository