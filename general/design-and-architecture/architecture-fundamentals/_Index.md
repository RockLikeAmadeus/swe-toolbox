Sources: Fundamentals of Software Architecture; Richards & Ford

Architecture, like every other part of software development, should be developed iteratively.

At a high-level, I'm boiling the process down into these steps:

1. Determine the desired architecture characteristics.
2. Decompose the system into components, with those characteristics in mind.
3. Decide on an architecture style.


# [Software Architecture Characteristics](architecture-characteristics)

The first step in defining the architecture of a new application or system is to decide on which architecture characteristics are most important to support. Supporting architecture characteristics is a matter of making trade-offs, which means all systems cannot support all characteristics. As a good starting point, attempt to determine the top 3 most important characteristics that the system must support in order for the project to be successful.

# Component Identification/Decomposition

After prioritizing architecture characteristics, one of the first concerns is how to divide the system up into discrete components (libraries, services, etc.). This should always be an iterative process, with the following five steps in a general case (though modifications to this process can apply for certain specialized domains, such as those with complex security requirements):

1. Identify your initial components
   - Base this on top-level partitioning, i.e.
     - Technical partitioning, as in layered architectures like MVC, or
     - Domain partitioning, like a conceptually decomposed structure for a microservices approach.
   - Outside of top-level partitioning, just try and think of a reasonable structure.
   - Finally, map domain functionality to specific components, to see where specific behavior should reside.
   - The likelihood that this is your final component structure is quite low.
2. Assign requirements (user stories) to your components
   - This may involve
     - Creating new components.
     - Consolidating existing components.
     - Breaking components apart because they have too much responsibility.
3. Analyze the roles and responsibilities of your components
   - This is really a sub-step of Step 2.
   - Thinking about both the roles and behaviors that the system must support allows the architect to align the component and domain granularity.
   - One of the greatest challenges for architects is discovering the correct [granularity for components](#component-granularity).
4. Analyze your architecture characteristics
   - This, also, is really a sub-step of Step 2.
   - This will depend on the architecture characteristics that were identified.
5. Restructure your components
   - Continue to iterate on component structuring throughout development; lower-level software design considerations will regularly inform component structure.
6. Repeat (likely starting at step 3)

### Component granularity

Determining the right level of component granularity is a balancing act. Too fine-grained component design requires too much communication between components to achieve results. Too coarse-grained component design encourage high-internal coupling, which leads to difficulties in deployability and testability, as well as modularity-related negative side-effects.

### Componet Design

Lots of design techniques exist, all with various trade-offs.

##### Discovering Components

Avoid the "Entity Trap" anti-pattern, which primarily involves a simplistic design with one-to-one mapping of components to domain concepts (at least, as I understood it; this section could use a re-read).

**Actor-Actions Approach**: A popular way to map requirements to components; it involves identifying the actors who perform activities with the application, along with the actions those actors will perform. This can work well when the domain problem nicely fits the actor-action pattern. However, this approach is often best suited to a waterfall development strategy.

**Event Storming**: starts with the assumption that messages/events will be used to communicate between the various components. The team tries to identify which events occurr in the system based on requirements and the identified roles, and then to build components around those events and message handlers. This works especially well in distributed architectures like microservices.

**Workflow Approach**: An alternative to even storming that's a bit more generic. The approach
1. Identifies the key roles (frequently, this includes the "system").
2. Determines the kind of workflows in which these roles engage.
3. Builds components around the identified activities.

# Monoliths vs Distributed Systems

This decision isn't always so obvious as you might think, but can actually be answered by first answering a single question:

Think about whether the entire system will share the same desired architecture characteristics, or whether different sets of characteristics (quanta) can apply separately to different components of the system (i.e. does a particular role require more fault-tolerance than another, or less elasticity than another?). If the answer is the former--the system as a whole should support a shared set of architecture characteristics--the _monolith_ approach actually offers multiple advantages. Of course, in the latter case, favor a distributed architecture.

# Architecte Styles
