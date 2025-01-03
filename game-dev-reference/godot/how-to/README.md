###### Determine the boundaries of the screen

`GetViewportRect();`

###### Remove the current node/scene from the game tree

```cs
private void OnAreaEntered(Area2D area) {
    QueueFree(); // <- removes from tree at the end of the frame
}
```