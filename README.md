# Gator App

Purpose of leaning go on boot.dev

## Installation

### Requirement goose and sqlc

- [Pressly/Goose](https://github.com/pressly/goose)
- [Sqlc](https://docs.sqlc.dev/en/latest/overview/install.html)
- Go version 1.23+
- Postgresql 16

## Step 1

Go to root run

```bash
sqlc generate
```

then go to sql/schema run

```bash
goose postgres "postgres://username:yourpassword@pass@host:port/db" up
```

finally create file `~/.gatorconfig.json` in your home directory

## Step 2

Install gator by run `go build or go install`

## Finial

command list

`gator register yourname` register user

`gator login username` login user

`gator users` list all users register


`gator addfeed "name" "https://yoururl.com"` store your feed

`gator feeds` listing all feeds

`gator follow "https://yoururl.com"` follow feed

`gator unfollow "https://yoururl.com"` follow feed

`gator following "https://yoururl.com"` list all your folow feeds


`gator agg` scan feed url to get posts

`gator browse (limit)` browse posts (limit is optional number, default 2)


> Dangerous

`gator reset` delete all data in database


