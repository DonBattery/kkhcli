# kkhcli

Simple admin CLI tool for KKHC APP

### Usage: kkhcli \[option\] \[username\]

```SHELL
Options:
  -a, --add string
    	Add user
  -c, --collections
    	List collections
  -f, --flush string
    	Flush collection
  -l, --list
    	List users
  -r, --reset string
    	Reset password
  -s, --seed
    	Seed database
  -u, --url string
    	WEB URL (default "https://kkhc.eu/admin")
```

You can flush multiple collection atonce by separating their names with commas (,)
```SHELL
kkhcli -f CommentFlow,Folder,Image
```

