# find readme

`find-readme` - Find README.md under path.

## DON'T USE THIS

This App is a sample implementation for using `path/filepath/Walk`.  
Use [the_silver_searcher](https://github.com/ggreer/the_silver_searcher) is more faster and powerful 😌

```sh
$ function find-readme { ag -l "" ${1:?} | ag -i 'readme.md'; }
$ find-readme ~/go
```

see [benchmark](./bench.md)

## Usage

```
$ find-readme [global options] command [command options] path ...
```

### COMMANDS:
```
help, h  Shows a list of commands or help for one command
```

### GLOBAL OPTIONS:
```
--ignore name, -i name  Directory name to ignore.
--help, -h              show help
--version, -v           print the version
```

## Installation
```
$ go get github.com/micheam/find-readme
```

## Example

Show All readme.md or README.md under _$HOME/go_:

    $ find-readme $HOME/go

Same as above but ignore some directory [^1]:

    $ find-readme --ignore pkg --ignore golang.org $HOME/go

[^1]: `node_modules` will always ignored.

## License
MIT

## Author
Michito Maeda <https://github.com/micheam>
