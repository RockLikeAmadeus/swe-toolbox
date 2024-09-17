[Source: BDD in Action, Chapter 3](https://livebook.manning.com/book/bdd-in-action-second-edition/chapter-3/)

# High-Level End-to-End BDD Example

This example describes the process of building an application for a public transport department which provides train timetable data and rel-time updates about delays, track work, etc.

It's broken down based on the steps specified [here](./_index.md#high-level-process).

## 1. Speculate: Identifying business values and features

### 1. Identify business objectives

Identify the business problem we're solving. In this case, at a high level, we want to build an application that makes it easier for commuters to plan their commutes. But we want to get more specific than that, so we'll state the project vision as such:

"The application will help to reduce average travel time for regular commuters by 10% within a year, by allowing them to plan their journeys more effectively.”

### 2. Discover capabilities and features

Now, we focus on ideation for features that will directly support our stated business objectives. The technique of Impact Mapping can be a useful exercise to discover and prioritize high-level capabilities and features through the lens of four questions: why, who, how, and what.
1. Why (are we doing this)?
2. Who (is this for, are the key actors)?
3. How (might their behavior change with our solution)?
4. What (features might support this behavior change)?

![alt text](./impact-mapping-example.png)

### 3. Describe each feature

The following format is common and helps ensure that each feature actively contributes to a business goal. It is especially appropriate for higher-level features:

```
In order to <achieve a business goal or deliver business value>
As a <stakeholder>
I want <to be able to do something>
```

For example:

```
In order to plan my trips more effectively
As a commuter
I want to know the optimal itinerary between two stations
```

Another valid format, perhaps more appropriate for the detailed user stories within a feature:


```
As a <stakeholder>
I want <something>
So that <I can achieve some business goal>
```

For example:

```
As a commuter
I want to know the best way to travel between two stations
So that I can get to my destination quickly
```

Or even:

```
As a commuter
I want to be able to easily find the optimal route between two stations
So that I can get to my destination quickly
Whereas currently I need to look up the timetable in a paper booklet
```

## 2. Illustrate: Exploring features with examples

Thinking only about abstractions can lead to flawed assumptions and other problems. Here, we want to talk through concrete examples, using a technique such as [Example Mapping](internal-link-here).

![alt text](./example-mapping-example.png)

### Slice the feature into user stories

_Direct Connections_

As a commuter traveling between two stations on the same line

I want to know what time the next trains for my destination will leave

So that I can spend less time waiting at the station

_Alternative Routes_

As a commuter traveling between two stations on different lines

I want to know what time the fastest train to my destination will leave

So that I can spend less time waiting at the station or for connecting trains

## 3. Formulate: From examples to executable specifications

In BDD, the following notation is often used to express examples:

```
Given <a context>
When <something happens>
Then <you expect some outcome>
```

When expressed this way, these are called _scenarios_, and they are intentionally a bit less flexible in order to use these as the basis for automated tests later on. With concrete scenarios defined in this way, you'll be able to more easily notice logical errors, missing context, or situations that are more complex than first thought:

```
Scenario: Display the next train going to the requested destination
  Given the T1 train to Chatswood leaves Hornsby at 8:02, 8:15, 8:21
  When Travis wants to travel from Hornsby to Chatswood at 8:00        
  Then he should be told about the trains at: 8:02, 8:15
```

## 4. Automate: From executable specifications to automated tests

Depending on your tech stack, there are many specialized BDD tools that you can use to automate your acceptance criteria, such as Cucumber, SpecFlow, and Behave. These aren't absolutely necessary, but can make it easier to express automated tests in a structured form (given, when, then).

There are also tools for other aspects of this process, such as the open source [Serenity BDD](https://serenity-bdd.info/), and [Selenium](https://www.selenium.dev/).