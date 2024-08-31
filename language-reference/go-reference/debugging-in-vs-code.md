In your project root create .vscode/launch.json

```json
{
  // Use IntelliSense to learn about possible attributes.
  // Hover to view descriptions of existing attributes.
  // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/cmd"
    },
    {
      "name": "Test",
      "type": "go",
      "request": "launch",
      "mode": "test",
      "program": "${workspaceFolder}/cmd"
    }
  ]
}
```

The second configuration is to make testing work--`"mode": "test"` is the important part. The "program" field needs to be set to the entry point path. Click Run >> Start Debugging to debug the code, or right click on the green arrows next to tests and select Debug Test to debug an individual test.