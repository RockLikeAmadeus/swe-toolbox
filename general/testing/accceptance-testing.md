From [this video by Dave Farley](https://www.youtube.com/watch?v=JDD5EEJgpHU&ab_channel=ContinuousDelivery).

Acceptance tests are always written from the perspective of an expected user of the system. They exercise _what_ the system should do, not _how_ it does it. They're a great tool for driving development when written _before_ the code. If not written well, however, they can be very fragile, so keep the below in mind to write acceptance tests that are stable.

# Dave Farley's four-layer acceptance test architecture

At a high-level, these four layers are: **Specification->DSL->Driver->System**

**Step 1.** User story: describe the steps the user would take in an example use case.

**Step 2.** Convert this to an executable specification, e.g.

```java
@test
@channel(Amazon)
public void shouldBuyBookWithCreditCard
{
    shopping.goToStore();
    shopping.searchForBook(title: "Continuous Delivery");
    shopping.selectBook(author: "David Farley");
    shopping.addSelectedItemToShoppingBasket();
    shopping.checkOut(item: "Continuous Delivery");
    shopping.assertItemPurchased(item: "Continuous Delivery");
}
```

Notice that this specification could apply to any online book store, and indeed, any book store at all. If the actual system we were designing was a robot that could purchase the book for you in a physical store, the specification above would not have to change at all. What this means is that you can write the test specification once, and re-use it for your web interface and your mobile app alike.

Also from [Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests/testing-fundamentals/scaling-acceptance-tests):

> At this point, this level of ceremony to decouple our specification from implementation might make some people accuse us of "overly abstracting". **I promise you that acceptance tests that are too coupled to implementation become a real burden on engineering teams.** I am confident that most acceptance tests out in the wild are expensive to maintain due to this inappropriate coupling; rather than the reverse of being overly abstract.

It's helpful to use an internal DSL (Domain-Specific Language) to describe the steps, like in the Java example above. So, the next step, one layer down, is to

**Step 3.** Implement the DSL to be used for testing (in the language of the problem domain). Be sure to make it easy to design and capture our test cases. These should be designed with the expectation of sharing common functionality between tests.

```java
public void checkOut(String... args)
{
    Params params = new Params(args);
    String item = params.Optional(name: "item", defaultValue: "Continuous Delivery");
    String price = params.Optional(name: "price", defaultValue: "1.00");
    Card card = parseCard(params.Optional(name: "card", defaultValue: "1234123412341234"));

    driver.checkOut(item, price, card);
}
```

The next layer down is the protocol driver layer. The interface to the protocol driver is high-level and abstract, and should be described in the language of the _problem_ domain.

**Step 4.** Implement the protocol driver layer to translate from the language of the problem domain to the system under test.  For instance, if you're testing a GUI, you might make use of something like Selenium at this layer. If your application exposes an API, here is where you might define an HTTP client to make HTTP calls into the API. This is the only part of the test infrastructure that understands the _how_ of the system--clicking buttons in your UI occurs here, _not_ at the specification or DSL layers, which don't care about the actual structure of the interface to your system.

**Step 5.** Implement the system under test here, in a production-like environment.

