# KKHCLI

**admin tool for managing the KKHC Serever**

### Installation

You need to have Docker installed in order to build kkhcli for yourself

Edit the **build.sh** script to change target OS or Architecture, see more in the script

Run
```Shell
./build.sh
```

If build is successful you got an executable binary called **kkhcli**

### kkhcli.yaml

Edit this file to set the API URL and Admin Password of the KKHC Server\
Note that these variables can be set with kkhcli.yaml < env < flag

### Usage

```Shell
./kkhcli [command]
```

Available Commands:\
**avatar**      Root command of Avatars. Subcommands: list add seed\
**collection**  Root command of Collections. Subcommands: list flush seed\
**help**        Help about any command\
**user**        Root command of Users. Subcommands: list add reset

Flags:\
  -p, --adminPass string   Admin Password\
  -u, --apiURL string      API URL\
  -h, --help               help for kkhcli

Use "kkhcli \[command\] --help" for more information about a command.
