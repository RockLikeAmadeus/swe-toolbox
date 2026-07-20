# Keeping Secrets
    
Private keys and API keys should not be stored even in private remote repositories. Emacs a built-in network security manager that reads login secrets from a standard text file called `~/.authinfo`.

If it doesn't exist, create a local file named `~/.authinfo` (or `~/.authinfo.gpg` to encrypt it) on your machine. Add contents of the form:

```
machine api.google.com login apikey password YOUR_COPIED_GEMINI_API_KEY
```

See config file for example of how to use the value in your doom config.
