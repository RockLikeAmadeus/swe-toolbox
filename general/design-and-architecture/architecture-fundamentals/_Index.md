Sources: Fundamentals of Software Architecture; Richards & Ford

Architecture, like every other part of software development, should be developed iteratively.


# [Software Architecture Characteristics](architecture-characteristics)

The first step in defining the architecture of a new application or system is to decide on which architecture characteristics are most important to support. Supporting architecture characteristics is a matter of making trade-offs, which means all systems cannot support all characteristics. As a good starting point, attempt to determine the top 3 most important characteristics that the system must support in order for the project to be successful.

# Component Identification

After prioritizing architecture characteristics, one of the first concerns is how to divide the system up into discrete components (libraries, services, etc.). This should always be an iterative process, with the following five steps in a general case (though modifications to this process can apply for certain specialized domains, such as those with complex security requirements):

1. Identify your initial components
   - Base this on top-level partitioning, i.e.
     - Technical partitioning, as in layered architectures like MVC, or
     - Domain partitioning, like a conceptually fractured structure for a microservices approach.
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

