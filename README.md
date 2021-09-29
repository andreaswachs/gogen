# gogen

*Go* *Generator*: a template generator written in *Go*.

Developed for the purpose of easily generate new text files, based on predefined templates.

This was developed to remove slightly repetitive work when doing new homework assignments at university by having to copy the same template LaTeX files around.

## Installation

In order to install this program you need to have the go tools installed. You can look into downloading and installing those [here](https://golang.org/doc/install).

```bash
make install
```

This installs the compiled binary in your GOBIN folder. This folder needs to be in your $PATH environmental variable, and it needs to be set with the command `go env -w GOBIN=<path>`

## Usage

When installed, the `gogen` binary should be available to use.

TODO: Write a nice section later...

### For developers

You need the *GNU* tool *make* in order to use the predefined features of the Makefile.

#### Build the binary

```bash
make build
```

Produces the gogen binary file in the root of the source folder.

#### Run without compiling

```bash
make run [ARGS]
```

This uses the `go run` command and passes the args to the program.

#### Remove the binary from the root folder in the source directory

```bash
make clean
```

This removes the binary file that was created with `make build`
