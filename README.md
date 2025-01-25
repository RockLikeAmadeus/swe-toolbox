[Aspects of Software Engineering](general/aspects-of-software-engineering.md) <br>
[Design and Architecture](general/design-and-architecture/README.md) <br>
[Working with Legacy Code](general/legacy-code/legacy-code-tools.md) <br>
[Testing/TDD](general/testing/README.md) <br>
[Code-Review Checklist](general/processes/code-review-checklist.md) <br>

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
[Protocols](./protocols/README.md) <br>
[Third-Party Tools](tools/README.md) <br>
[Other Resources](resources/README.md) <br>

<br><br><br>

---

# Building Software

At a high-level, the process of building software should look like this the process described [here](./general/design-and-architecture/behavior-driven-development/README.md#high-level-process).

Somewhere between steps 2 and 3 of this process, you additionally want to come up with a high-level system design, and then build up what [the GOOS book] calls a walking skeleton based on that design. That way, as you start to build out acceptance tests, you already have the structure in place such that all that remains is a matter of implementing features. It's important to note that the design on which the skeleton is based doesn't need to be perfect or be the final design, but should support end-to-end use cases.

If your project will have a user interface, then for each feature you implement during the BDD process you want to think about what that UI will look like. Follow [these steps](./front-end-reference/README.md#the-process-of-building-user-interfaces) when you get to that point.

##### Other resources on starting a new software project
https://quii.gitbook.io/learn-go-with-tests/testing-fundamentals/scaling-acceptance-tests

Also check out https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-template-repository