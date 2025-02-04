# Gator: a command-line RSS Reader!

## Requirements:
- Go
- PostgreSQL

## Installation:
To install the file, use:

```sh
go install https://github.com/bzelaznicki/gator
```

Schemas can be installed using Goose and are found in the /sql/schema/ folder.

The application also requires a .gatorconfig.json file in your home directory.

## Commands
* `register <username>`: Lets you register as an user.
* `login <username>`: Lets you log into an user account.
* `users`: Lists the current users.
* `feeds`: Lists the current feeds.
* `addfeed <name> <url>`: Lets you add a new feed.
* `agg <time>`: Aggregates RSS data per specified period.
* `following`: Lets you see the feeds you follow.
* `follow <url>`: Lets you follow a feed.
* `unfollow <url>`: Lets you unfollow a feed.
* `browse`: Lets you view the feeds you follow. Optional argument: <limit>, letting you display the last n posts.