# Basic Structure of Event Handlers, Callbacks
An event is a message sent by an object. The object that raises the event is called the *event sender*. The sender doesn't know which object or method will handle the event it raises. The event is typically a member of the event sender i.e. the `Click` event is raised by a `Button` in Button.Click().

The pieces of the puzzle look like this:

- EventSender (i.e. Button)
	- event (the definition of the message, i.e. Click)
	- raise event method (OnEventName, i.e. OnClick)
	
- EventReceiver
	- explicit subscription (i.e. button.Click += HandleClick)
	- Handle[Event] (defines the *callback*, i.e. HandleClick)
		- Takes as arguments the sender (object) and the event args (EventArgs type or some derivation of it).


# Testing Event-Driven Systems
Two parts: 1. verify that an events your class fires are fired when appropriate conditions are met, and 2. test that any handlers for the events perform the expected actions.

1. Is simple: set up the sender object in a state that you would expect to fire the event, and then verify that "an appropriately populated event is fired and any other state changes take place". (from https://stackoverflow.com/questions/29545777/testing-event-driven-behavior) If you're only implementing the event receiver, this test is outside of the scope of your work.

2. Is also simple: event handlers are just methods. Your tests should just call them directly, as if they were the event sender, and then verify the behavior.