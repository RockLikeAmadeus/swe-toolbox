#### [Set up a new C# Project](new-c-sharp-project-setup.md)
#### [Decide what type of root node my new scene should have]()

#### Determine the boundaries of the screen

`GetViewportRect();`

#### Be notified when an object becomes (or becomes not) visible on the screen

Use the VisibleOnScreenNotifier2D Node. (There is currently a bug where the related Signals are only emitted when the game window is in focus)

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

#### Stop execution of child node processing (i.e. pause)

```cs
_pipesSpawnTimer.Stop();
foreach (Pipes pipes in _pipesContainer.GetChildren()) {
    pipes.SetProcess(false);
}
```

#### Share variables, constants, shared logic, etc. between scenes

Use Godot's "Autoloads", which are the engine's version of Singletons (take a look at the documentation). Autoloads still inherit from Node, so you still have access to all of its parents methods and properties, like `_Ready()`.

Create a new folder in the file system explorer in the Godot IDE, and call it Globals or Autoloads. Right click the directory >> Create new >> Script. Typically you might call it `_____Manager` or, `SignalBus`.

Finally, register the new script as an Autoload: Project >> Project Settings >> Globals >> Autoload and add the new script. The order of the autoloads listed here determines the order that Godot will instantiate them, which is relevant when autoloads depend on one another.

For GDScript autoloads, Godot will automatically set the script up as a singleton. For C#, however, you need to do it manually.

```cs
public partial class GameManager : Node
{
	public static GameManager Instance { get; private set; }

	private PackedScene _gameScene =
		GD.Load<PackedScene>("res://Scenes/Game/Game.tscn");
	private PackedScene _mainScene =
		GD.Load<PackedScene>("res://Scenes/Main/Main.tscn");

	// Called when the node enters the scene tree for the first time.
	public override void _Ready()
	{
		Instance = this;
	}

	private static void LoadMain() { // Making this static makes calling easier
		Instance.GetTree().ChangeSceneToPacked(_mainScene);
	}

	private static void LoadGame() { // Making this static makes calling easier
		Instance.GetTree().ChangeSceneToPacked(_gameScene);
	}
}
```