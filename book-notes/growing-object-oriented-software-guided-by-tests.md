# Chapter 1

Chapter 1 discusses the ideas of incremental and iterative development and the importance of feedback, given how little is typically known going into a project and therefore the importance of a process of discovery. It then talks at a high level about TDD and refactoring, and discusses testing at multiple layers (a unit-testing cycle inside of a larger integration testing cycle inside of an acceptance-testing cycle), and mentions the importance of testing not just the system, but the processes of building and deploying the system as well. There are many parallels in this section to the ideas in David Farley's Modern Software Engineering.

It ends with a discussion of internal vs external quality (the quality of the system as seen by users vs the quality of the code that makes it work), and how automated testing works for the benefit of both (followed by a quick sidebar about coupling and cohesion).

# Chapter 2

Chapter 2 is very abstract and I'm hesitant to try and summarize it until I (hopefully) understand better by working through the examples in later chapters. Primarily, though, I think it's about the importance of relying on abstract interfaces rather than concrete implementations, and how that relates to mock objects (I assume, though don't know for sure, that these are referring to all test doubles in general rather than mocks specifically) in unit tests.

# Chapter 3

Chapter 3 is just a minimalist explanation of the tools being used in the later examples to illustrate the concepts. The idea is to explain what the tools are and how they're used in just enough detail to understand the examples, as you'll likely use different/equivalent tools when writing your own code.