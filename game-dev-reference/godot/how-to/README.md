###### Determine the boundaries of the screen

`GetViewportRect();`

###### Remove the current node/scene from the game tree

```cs
private void OnAreaEntered(Area2D area) {
    QueueFree(); // <- removes from tree at the end of the frame
}
```

###### Reference an existing node in the scene tree

```cs
Gem gem = GetNode<Gem>("gem"); // Not ideal
```

Instead, export a property that you can assign to in the IDE via drag and drop:

```cs
[Export] private Game _gem;
```

Also possible:

```cs
[Export] private NodePath _gemPath; // Gem
private Gem _gem;

public override void Ready() {
    _gem = GetNode<Gem>(_gemPath);
}
```