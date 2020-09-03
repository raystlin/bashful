# Bashful
A RESTfull server to trigger scripts on the host. I use it with home-assistant and its RESTful switch so as to change the inner state of my servers.

# Compilation
I usually compile it using go docker, but this will disable sqlite support because the base image does not have gcc. If you want to be able to store the configuration using sqlite you must enable CGO_ENABLED (everything is well explained at github.com/mattn/go-sqlite).

# Execution
The cli has the following help. 

```bash
USAGE:
   bashful [global options] command [command options] [arguments...]

COMMANDS:
   init       Create a sample command file
   add, a     Adds a command and its statuses to the database. If you need a template call init subcommand
   delete, d  Delete a command from the database
   set        Sets the current status of a command, triggering its action
   list, l    List all commands and their current status
   server     Starts a bashful server
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d    Debug level for logs (default: false)
   --store value  Store type to use: sqlite3, bolt (default: "bolt")
   --help, -h     show help (default: false)
```

## Template generation
First of all it is recommended to generate a template to configure the commands and its states

```sh
USAGE:
   bashful init [command options] [arguments...]

OPTIONS:
   --command value  Command file (default: "command.json")
   --help, -h       show help (default: false)
```

After a blank execution we will have an example template named command.json

## Insert new command
Once we have filled the template we can add it using add subcommand

```bash
USAGE:
   bashful add [command options] [arguments...]

OPTIONS:
   --command value             Command file (default: "command.json")
   --database value, -d value  Database with the command definitions (default: "data.db")
   --help, -h                  show help (default: false)
```

## Run the server
We are now ready to run our http server using the following subcommand

```bash
USAGE:
   bashful server [command options] [arguments...]

OPTIONS:
   --address value, -a value   Address and port to listen (default: ":8083")
   --cert value                Server certificate path for https (PEM)
   --key value                 Server key path for https (PEM)
   --database value, -d value  Database with the command definitions (default: "data.db")
   --help, -h                  show help (default: false)
```

_--cert_ and _--key_ parameters are not mandatory. If given the server will listen to https requests, but if not a regular http server will be launched.
Once the server is launched we will be able to retrieve the current status of each command using the following entrypoint
```
http://<server>:<port>/api/v1/cmd/<commandName>
```
POST requests to these entrypoints will be used to set new states and run the linked commands. 
The entrypoint _/api/v1/cmd_ will give you a list of <command_name>:<status>

## Listing known commands
To list the commands registered on our database we can use the following subcommand

```bash
USAGE:
   bashful list [command options] [arguments...]

OPTIONS:
   --database value, -d value  Database with the command definitions (default: "data.db")
   --help, -h                  show help (default: false)
```

## Changing the current status via cli
_set_ subcommand will change the status and trigger the linked command, as if we did it via web.

```bash
USAGE:
   bashful list [command options] [arguments...]

OPTIONS:
   --database value, -d value  Database with the command definitions (default: "data.db")
   --help, -h                  show help (default: false)
```

## Deleting commands
CLI also provides a delete subcommand

```bash
USAGl delete [command options] [arguments...]

OPTIONS:
   --name value, -n value      
   --database value, -d value  Database with the command definitions (default: "data.db")
   --help, -h                  show help (default: false)
```
