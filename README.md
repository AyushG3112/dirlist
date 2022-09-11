# DirList [![Go Reference](https://pkg.go.dev/badge/github.com/ayushg3112/dirlist.svg)](https://pkg.go.dev/github.com/ayushg3112/dirlist)

DirList allows you to quickly spin up an HTTP server to view and browse directory listing of any of your folders, and browse them on your browser. 

## Installing

```
go install github.com/ayushg3112/dirlist/cmd/dirlist@latest
```

## Usages

### View Help

```
$ dirlist --help 
Usage: 
  -c, --cached              Run in cached mode. Cached mode generates the structure once and always shows that even if the underlying structure has changed
  -h, --help                View help
  -p, --port int            Port on which to start the listing server (default 8000)
  -d, --root-dir string     Root directory to start directory listing. Defaults to $PWD (default ".")
  -f, --sort-field string   Field to sort by (default "modifiedAt")
      --sort-order string   Sorting order. ASC/DESC. (default "ASC")
```

### Start a Server

To start a server in the current directory with default options:
```
$ dirlist
2022/09/03 16:13:10 starting the server at port 8000
```

To start a server in another directory:
```
$ dirlist --root-dir /path/to/folder --sort-order ASC --sort-field modifiedAt
2022/09/03 16:15:24 starting the server at port 8000
```

### To run locally after cloning
```
$ go run ./cmd/dirlist/... --root-dir /path/to/folder --sort-order ASC --sort-field modifiedAt
2022/09/03 16:15:24 starting the server at port 8000
```
