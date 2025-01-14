1. Set the `$ ECHO $GODOT_BIN` to check if the environment variable is set to the path to Godot. If not, [set it](/game-dev-reference/godot/how-to/installation-and-setup.md#set-system-wide-environment-variable) (use `GODOT_BIN="/opt/Godot_43/Godot_43"`).
2. [Install and activate](https://mikeschulze.github.io/gdUnit4/first_steps/install/) via the AssetLib. This gets installed on a per-project basis.
3. Complete the [C# Testing setup](https://mikeschulze.github.io/gdUnit4/csharp_project_setup/csharp-setup/).
4. [Test your setup in the editor](https://mikeschulze.github.io/gdUnit4/csharp_project_setup/csharp-setup/#test-you-c-build-settings-in-the-godot-editor)
5. [Configure your IDE for GDUnit4](https://mikeschulze.github.io/gdUnit4/csharp_project_setup/vstest-adapter/).
   1. My settings.json looks like this: ```{
    "settings": {
        "dotnet.unitTests.runSettingsPath": "./test/.runsettings"
    }
}```

Note: My CSProject file contains this:

```xml
<ItemGroup>
    <PackageReference Include="Microsoft.NET.Test.Sdk" Version="17.9.0" />
    <PackageReference Include="gdUnit4.api" Version="4.3.*" />
    <PackageReference Include="gdUnit4.api" Version="4.3.*" />
    <PackageReference Include="gdUnit4.test.adapter" Version="2.*" />
</ItemGroup>
```

For an example of a working setup, see my Godot Practice Projects repo, the Tappy project