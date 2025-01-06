![Basic physics 2D notes](basic-physics-2d-notes.png)

# CharacterBody2D node

Orange is where the physics engine takes over

![alt text](character-body-2d.png)

Nodes that deal with physics typically name themselves with the word "Body" rather than the "Area" used in a lot of non-physics nodes (i.e., physics collisions use `BodyEntered` signal as opposed to non-physics `AreaEntered`).