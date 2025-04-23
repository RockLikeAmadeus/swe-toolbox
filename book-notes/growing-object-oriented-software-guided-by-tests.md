# Part 1

Part 1 is an overview of the concepts of TDD and why it's important, as well as the importance of test doubles. It also takes a moment to describe the tools the book will use for the later examples.

## Individual Chapters

### Chapter 1

Chapter 1 discusses the ideas of incremental and iterative development and the importance of feedback, given how little is typically known going into a project and therefore the importance of a process of discovery. It then talks at a high level about TDD and refactoring, and discusses testing at multiple layers (a unit-testing cycle inside of a larger integration testing cycle inside of an acceptance-testing cycle), and mentions the importance of testing not just the system, but the processes of building and deploying the system as well. There are many parallels in this section to the ideas in David Farley's Modern Software Engineering.

It ends with a discussion of internal vs external quality (the quality of the system as seen by users vs the quality of the code that makes it work), and how automated testing works for the benefit of both (followed by a quick sidebar about coupling and cohesion).

### Chapter 2

Chapter 2 is very abstract and I'm hesitant to try and summarize it until I (hopefully) understand better by working through the examples in later chapters. Primarily, though, I think it's about the importance of relying on abstract interfaces rather than concrete implementations, and how that relates to mock objects (I assume, though don't know for sure, that these are referring to all test doubles in general rather than mocks specifically) in unit tests.

### Chapter 3

Chapter 3 is just a minimalist explanation of the tools being used in the later examples to illustrate the concepts (JUnit, Hamcrest, and jMock), in addition to some a couple of the related concepts and terms. The idea is to explain what the tools are and how they're used in just enough detail to understand the examples, as you'll likely use different/equivalent tools when writing your own code.

JUnit is a pretty standard testing framework with tests that are fairly easy to read without an deep understanding of the framework itself. The book touches on the concept of _Test Fixtures_ which it defines as the fixed state that exists at the start of a test run, ensuring that the test is independently repeatable (a fixture is _set up_ and ultimately _torn down_), and then mentions how that looks in a JUnit test.

Hamcrest is a framework for writing match criteria, and defines the `Matcher<T>` type and the `.matches()` method. Matchers are designed so that when they're combined, both the code and the failure messages are self-explanatory. It's basically a way to simplify writing assertions without explicitly defining conditional expressions.

Finally, jMock creates mock objects dynamically so that you don't have to write your own implementations of the types you want to mock, and provides a high-level API for specifying how the object under test should invoke the mock objects and how the mocks should behave in response. The book describes here the jMock concepts of the _mockery_ (the context of the object under test, which is its neighboring objects) and _expectations_, and provides an example.

# Part 2

Part 2 goes into the concepts specific to the book's approach, which boil down to the two core principles of incremental development and expressive code.

## Individual Chapters

### Chapter 4