# Think "Spec" instead of "Test"

> generally, the term unit 'testing' is misleading. Testing something means that you want to know how that thing behaves. But with unit tests you know how it behaves (or should) and thus you are not interested in discovery.

> Unit tests are a way of writing your program's spec. You define (as in make up) rules for your system and now you need a way to lock that in. This is what you write the tests for. One test - one rule. E.g. you have a type `EvenI32`. Appropriate tests would be `EvenI32 can be created from i32`, `EvenI32 ctor return error on odd intput`, `EvenI32 and i32 addition gives correct result`, etc. You don't test functions or data, cause you can read the code and see how it works. What you need is a way to enforce rules on the system.

> You don't need to test for things that compiler can check for you. If you expect `f` to return `T` you will find it really quick that it now returns `U`. Based on this, you can have tests in strategic locations and use the compiler to connect the dots. This means that test coverage metric meaningless and will always be. Since the input for every piece of code is problem analysis, the output must be checked using problem analysis. Otherwise you have sort of the following chain: "Can a user withdraw from an account having a negative balance?" - "68% of the functions in the code base are test covered".

> Tests basically are a tool to enforce rules that the compiler can't.

> Also, some teams have opted not to use term 'tests' but switch to 'spec'. 

- [Source](https://www.reddit.com/r/rust/comments/zzqc9a/test_features_not_code_in_rust/)