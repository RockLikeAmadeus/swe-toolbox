> You have two main tools to process the player's input in Godot:
> 1. The built-in input callbacks, mainly `_unhandled_input()`. Like `_process()`, it's a built-in virtual function that Godot calls every time the player presses a key. **It's the tool you want to use to react to events that don't happen every frame, like pressing Space to jump.** To learn more about input callbacks, see Using InputEvent.
> 2. The Input singleton. A singleton is a globally accessible object. Godot provides access to several in scripts. **It's the right tool to check for input every frame.**

The singleton approach looks like:

```cs
if (Input.IsActionPressed("move_right"))
{
    velocity.X += 1;
}
```

The string that is passed to `IsActionPressed()` is defined in Project > Project Settings > Input Map. You can add new inputs with whatever name you like here.