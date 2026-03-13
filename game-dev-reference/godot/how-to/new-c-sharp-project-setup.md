1. Create a new project in the Godot editor; "Mobile" is a pretty balanced choice for the renderer (even if you're not making a mobile game).
2. Import any existing assets by dragging them into the Godot editor from the file explorer.
3. Project >> Project Settings >> General >> Display >> Window and set the Viewport width and height per your target.  Set the viewport to stretch mode - canvas items, and keep the aspect ratio.
4. Turn on Advanced Settings. Project >> Project Settings >> General >> Editor >> Naming >> Scene Name Casing: PascalCase. Do the same for Script Name Casing.
5. Editor >> Editor Settings >> Dotnet >> Editor >> External Editor >> Visual Studio Code
6. Create new task and project launch configuration to launch the game from VS Code.
   1. Open the project folder in VS Code.
   2. Make a new folder at the root of the project called `.vscode`.
   3. Add `launch.json` to this folder with the following contents (this relies on having the GODOT4 environment variable set).
   ```json
   {
        "version": "0.2.0",
        "configurations": [
            {
                "name": "Debug Game",
                "type": "coreclr",
                "request": "launch",
                "preLaunchTask": "build",
                "program": "${env:GODOT4}",
                "args": [],
                "cwd": "${workspaceFolder}",
                "stopAtEntry": false,
                "console": "internalConsole"
            },
            {
                "name": "Debug Current Scene",
                "type": "coreclr",
                "request": "launch",
                "preLaunchTask": "build",
                "program": "${env:GODOT4}",
                "args": [
                    "${fileDirname}/${fileBasenameNoExtension}.tscn"
                ],
                "cwd": "${workspaceFolder}",
                "stopAtEntry": false,
                "console": "internalConsole"
            }
        ]
    }
   ```
   4. Add `tasks.json` with the following contents.
   ```json
   {
        "version": "2.0.0",
        "tasks": [
            {
                "label": "build",
                "command": "dotnet",
                "type": "process",
                "args": [
                    "build"
                ],
                "problemMatcher": "$msCompile",
                "group": {
                    "kind": "build",
                    "isDefault": true
                }
            }
        ]
    }
   ```

7. Create a Scenes, Classes, Globals, Tests, and Resources folder at the project root (rethink this for what works for you, consider lower-casing these?)
8. Create a new 2D Scene called Game, just a Node2D should be fine. Save it in a folder called `Game` under the `Scenes` folder.
9. Create a new C# script in the `classes` folder (you can just call it `Game.cs` for now) in order to force Godot to build the project and solution files.
10. [Set up testing](../testing/install-and-setup-gdunit4.md).
11. Build and run the game to ensure there are no errors. Set Game as your main Scene.