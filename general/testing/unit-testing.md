# Naming unit tests to encourage decoupling from implementation

Source: BDD in Action

Think in terms of what the class, method, or function "should" do. When written this way, unit tests read more like specifications than tests, which is much more useful, maintainable, and encourages better design.

```java
public class WhenTransferringInternationalFunds {
    @Test
    public void should_transfer_funds_to_a_local_account() {...}
    @Test
    public void should_transfer_funds_to_a_different_bank() {...}
    ...
    @Test
    public void should_deduct_fees_as_a_separate_transaction() {...}
    ...
}
```