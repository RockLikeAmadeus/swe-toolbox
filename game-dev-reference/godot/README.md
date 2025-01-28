# Project Reference

See [Dodge The Creeps](https://github.com/RockLikeAmadeus/dodge-the-creeps).

# Key Concepts

A Godot game is a _tree_ of _nodes_ that you organize together into _scenes_. Nodes can communicate between each other using _signals_.

## Scenes

Godot **Scenes** serve a flexible dual purpose, acting as the Scene equivalent in an engine like Unity, while simultaneously behaving similarly to Unity's prefabs. Godot scenes are nestable, and are always composed of one or more _nodes_.

"A Godot scene could be a Weapon, a Character, an Item, a Door, a Level, part of a level… anything you'd like. It works like a class in pure code, except you're free to design it by using the editor, using only the code, or mixing and matching the two. It's different from prefabs you find in several 3D engines, as you can then inherit from and extend those scenes."

"On top of acting like nodes, scenes have the following characteristics:

1. They always have one root node, like the "Player" in our example.
2. You can save them to your local drive and load them later.
3. You can create as many instances of a scene as you'd like. You could have five or ten characters in your game, created from your Character scene."

## Nodes

**Nodes** are your game's smallest building blocks that you arrange into trees. A tree of nodes can be saved as a _scene_, much like Unity allows you to define a game object composed of components that can be saved as prefabs.

"Nodes are part of a tree and always inherit from their parents up to the Node class. Although the engine does feature some nodes like collision shapes that a parent physics body will use, most nodes work independently from one another. In other words, Godot's nodes do not work like components in some other game engines."

"All nodes have the following characteristics:

- A name.
- Editable properties.
- They receive callbacks to update every frame.
- You can extend them with new properties and functions.
- You can add them to another node as a child.
"

## The Scene Tree

Come back to this.

## Signals

Nodes emit **Signals** when certain events occur. They are Godot's version of the [_Observer_](https://gameprogrammingpatterns.com/observer.html) pattern.

# Creating a "Game Object Prefab" (Scene)

The first step is to decide on a root node for the scene. "As a general rule, a scene's root node should reflect the object's desired functionality - what the object is." `Area2D` is a useful one.

Once you've decided on a root node type and selected it (after clicking "Other Node" in the Scene dock), you can double click the name of the root node to rename the root node to be the name of the Scene you're creating (i.e. change `Area2D` to `Player`).

Now that you've set the scene's root node, you can add child nodes to give it more functionality.

A useful sequence for creating a basic game object is to create an `Area2D` root node, then add a `Sprite2D` child (or `AnimatedSprite2D`), as well as a collider (`CollisionShape2D`). If the object needs custom logic, right click the node and select `Attach Script`.