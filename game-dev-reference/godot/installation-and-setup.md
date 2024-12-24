# Debian Linux


## .NET

For the .NET version for C# coding, you need the .NET SDK installed first (change the SDK version according to the latest stable):

```
$ sudo apt install -y dotnet-sdk-8.0
```

Verify the installation with

```
$ dotnet
$ dotnet --list-sdks
```

## Godot

Download Godot for Linux from the Godot website. Rather than being an application that needs to be installed, Godot is simply a standalone executable. I simply extracted the zip file and renamed the resulting folder to `Godot_43` (or whatever version you installed), which I then moved to the `/opt/` directory (this directory is intended for third-party software, but you may need to change the ownership of the folder before you'll be allowed to place things there).

You should also shorten the name of the Godot executable in the folder (call it `Godot_XX`).