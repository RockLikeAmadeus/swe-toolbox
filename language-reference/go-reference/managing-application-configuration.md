# Using Viper for Configuration Management

Viper is a config management solution for Go applications which allows you to specify configuration options in several ways, including using configuration files, environment variables, and command-line flags.

## Viper with Cobra

If your application is a CLI tool and you've used Cobra to build it, Viper can be enabled automatically. Cobra does this by running the function initConfig in `cmd/root.go` when initializing this application.

If the user of your CLI tool includes the `--config` flag to specify a config file, Viper sets it as the configuration file for the application. If not, it sets the configuration as the file `$HOME/.myApp.yaml`. Then it uses the function `viper.AutomaticEnv` to read the configuration from environment variables that match any expected configuration keys. Finally, if the config file exists, Viper reads in the configuration from it.

**If your Cobra application already sets its options using flags, you can create Viper configuration keys by binding them to those flags.**