Sources: BDD in Action

# Learning progress

BDD in Action, Chapter 4

# High-Level Process

The process essentially has a 6-step cycle:

![alt text](BDD-activities.png)

1. [Speculate](#speculate-identifying-business-values-and-features): Identify high-level business goals and key features for achieving them
2. [Illustrate](): Get concrete about a specific feature using examples
3. [Formulate](): Transform examples into executable specifications
4. [Automate](): Transform executable specifications into automated acceptance tests
5. [Demonstrate](): Show that the feature does what it is supposed to do
6. [Validate](): Observe how the feature behaves in the real world

# Requirements Discovery and Specification

## Concrete examples and Gherkin

In a BDD workflow, features and their requirements are specified using concrete examples, partly because it is an effective way of communicating the intent of the system that leaves little room for ambiguity.

Most/many BDD tools rely on a format called _Gherkin_, which is designed to be easy to understand for non-technical stakeholders, and easy to automate using tools like _Cucumber_ and _SpecFlow_, so that the specification serves as both the definition for automated tests _and_ human-readable requirements documentation. _However, note that [David Farley](https://www.youtube.com/watch?v=YAZr3LsCzn0&ab_channel=ContinuousDelivery) prefers to use [his own custom DSL](/general/testing/accceptance-testing.md#dave-farleys-four-layer-acceptance-test-architecture) (domain-specific language) that he builds on top of his code, rather than Gherkin, so that's an option as well_.

Using the Gherkin format, the requirements for a particular feature are grouped into a single text file called a _feature file_, containing a short description of the feature, followed by a number of formalized examples for how it works in practice. 

```gherkin
Feature: Transferring money between accounts
  In order to manage my money more efficiently
  As a bank client
  I want to transfer funds between my accounts whenever I need to
  Scenario: Transferring money to a savings account
    Given Tess has a current account with $1000
    And a savings account with $2000.00
    When she transfers $500 from current to savings
    Then she should have $500 in her current account
    And she should have $2500 in her savings account
 
  Scenario: Transferring with insufficient funds
    Given Tess has a current account with $1000
    And a savings account with $2000.00
    When she transfers $1500 from current to savings
    Then she should receive an 'insufficient funds' error
    Then she should have $1000 in her current account
    And she should have $2000 in her Savings account
```

Each _scenario_ is made up of a number of _steps_ which each start with one of the following keywords: `Given`, `When`, `Then`, `And`, and `But`.

- `Given` described the pre-conditions for the scenario, and prepares the test environment
- `When` describes the actions being tested
- `Then` describes the expected outcomes from the test
- `And` and `But` can be used to join several `Given`, `When`, or `Then` steps in a more readable way.

### Simplifying an example with inline tables

The above example could be simplified to

```gherkin
Scenario: Transferring money between accounts within the bank
  Given Tess has the following accounts:
    | account | balance |
    | current | 1000    |
    | savings | 2000    |
  When she transfers 500.00 from current to savings
  Then her accounts should look like this:
    | account | balance |
    | current | 500     |
    | savings | 2500    |
```

### Tables of Examples

When appropriate, several related examples can be grouped together using a table of examples.

```gherkin
Scenario Outline: Earning interest
  Given Tess has a <account-type> account with $<initial-balance>
  And the interest rate for <account-type> accounts is <interest>
  When the monthly interest is calculated
  Then she should have earned $<earnings>
  And she should have $<new-balance> in her <account-type> account
  Examples:
  | initial-balance | account-type | interest | earnings | new-balance |
  | 10000           | Current      | 1.0      | 8.33     | 10008.33    |
  | 10000           | Savings      | 3.0      | 25       | 10025       |
  | 10000           | SuperSaver   | 5.0      | 41.67    | 10041.67    |
```

More about [Gherkin syntax](https://gist.github.com/dogoku/0c024c55ec124355f01472abc70550f5)

# Example

[High-Level end-to-end example](./high-level-end-to-end-bdd-example.md)