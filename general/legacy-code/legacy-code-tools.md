- Some thoughts from myself
	- To start to understand an unfamiliar codebase:
		1. Build a dependency diagram. This can help you determine which modules exist mainly to serve other modules. The modules with the most outgoing dependencies (that is, this module depends on many other modules) and the fewest incoming dependencies (that is, many other modules depend on this module) are likely to be a good place to start investigating.
		2. Build some high-level UML diagrams (add link here). Lower-level diagrams like detailed class diagrams are less likely to be useful early on in your understanding.
		3. Find, then step through and review the code for service/application entry points or startup (main) functions. The central application `main()` function is probably the best starting point for this, in order to get a foundational grasp of what the software is doing.
	- Build a project/solution structure diagram
	- It's easiest to follow the logic of a file/class when the public methods and properties are clearly separated from the private ones--this make it such that the interface (policy) is explicit and distinct from the implementation (details)
		- Change methods that are currently public but not used externally to private
	- Draw diagrams detailing who asks whom for what, as well as what info is provided in the inquiry
	- Take ownership of the codebase and re-organize, re-factor, and re-name to make everything easier to deal with
		- Also re-organize and add `#regions`
		- If a name isn't descriptive enough, change it
	- If you need to add unit testing to a class that has dependencies with behavior, you have two options:
		1. If the dependency is within your own control, you can extract an interface from the dependency and then depend on the interface. This is possible even for languages where the dependency is not in your control for languages where interface implemtation is implicit, such as Go. If you're using something like C# or Java and need to extract the dependency you can:
		2. If the dependency is outside of your control and you can't modify it to change it's implemented interfaces, you can replace the dependency with a *wrapper* that you do control, and then depend again on the interface. Then you can inject an instance of the wrapper in production. The wrapper probably shouldn't have unit testing, because by definition of being a wrapper for a concrete implementation, its implementation is going to be tightly-coupled to that dependency.
- Tools from Michael Feathers's Book
	- Feathers's algorithm for making changes to legacy code
		1. Identify change points
			1. This is not always simple, see Chapters 16 and 17 for help
		2. Find test points
			1. Also sometimes difficult, see Chapter 11 and 12
		3. Break dependencies
			1. Chapter 23 is important for doing this safely
			2. Afterword, Chapters 9 and 10 can help get past common problems
			3. See Dependency Breaking Techniques below (not yet created)
			4. See Chapter 22 and Chapter 7 if running into further problems
		4. Write tests
			1. This is often different than writing tests for new code; see Chapter 13
		5. Make changes and refactor
			1. Chapter 8 should be a helpful start
			2. Chapters 20, 21, and 22 can provide specific techniques

	- Mock objects are not the same thing as Fake objects
		- Fake objects work for most cases
	- Changing legacy code workflow (merge this into above algorithm)
		- If you need to make a change to a legacy class, the fist step is to attempt to instantiate the class in a test harness.
		- If time constraints disallow breaking dependencies
			- Techniques:
				- Sprout Method
			- Remember that if you don't cover the code that calls these new functions, you aren't testing their use