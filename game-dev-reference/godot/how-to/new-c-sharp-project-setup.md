1. Project >> Project Settings >> General >> Editor >> Naming >> Scene Name Casing: PascalCase (Must turn on Advanced Settings in the top right of the settings window first). Do the same for Script Name Casing.
2. Editor >> Editor Settings >> Dotnet >> Editor >> External Editor >> Visual Studio Code
3. Create new task and project launch configuration to launch the game from VS Code.
   1. Make a new folder at the root of the project called `.vscode`.
   2. Add `launch.json` to this folder with the following contents (this relies on having the GODOT4 environment variable set).
   ```
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
   3. Add `tasks.json` with the following contents.
   ```
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