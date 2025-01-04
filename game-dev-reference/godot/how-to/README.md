#### Determine the boundaries of the screen

`GetViewportRect();`

#### Remove the current node/scene from the game tree

```cs
private void OnAreaEntered(Area2D area) {
    QueueFree(); // <- removes from tree at the end of the frame
}
```

#### Reference an existing node in the scene tree

Don't do this:

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

#### Reference multiple nodes in the scene tree

There are multiple ways to do this.

##### Call `GetChildren()` from within a script attached to the root node

##### Assign nodes to a particular group and get nodes in that group

#### Spawn an object in the game

```cs
[Export] private PackedScene _enemyScene;
```

After rebuilding in the IDE, drag the .tscn file from the file system view into this exported property.

```cs
public void SpawnEnemy() {
    Enemy enemy = (Enemy)_enemyScene.Instantiate();
    AddChild(enemy); // Adds new node as child of `this`
    enemy.Position = new Vector2(0,0);
	enemy.OnHit += OnHit; // is this a memory leak?
}
```

#### Perform an action at a regular interval

Add a **Timer** node as a child of your scene and set the **Wait Time** interval. The **One Shot** flag will only run the timer a single time. Toggle **Autostart** if necessary. This Node will simply emit the **timeout** event every cycle, so you can subscribe to this to add behavior.

#### Stop execution of the Process function on a node
```cs
SetProcess(false);
```

