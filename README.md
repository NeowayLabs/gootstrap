# Gootstrap

Gootstrap stands for Go Bootstrap. It aims to provide bootstraping when
starting Go projects, helping you to go from 0 to 100 and with some
opinionated good practices/tooling like:

* Versioned builds
* Static Analysis
* Coverage
* Continuous Integration
* Git hooks
* Reproducible builds

The idea is to give a head start when starting a project not
contemplate every detail that a project may need, so after
generating the files you probably will need to add more
things as necessary.

# Install

Just run:

```
go install github.com/NeowayLabs/gootstrap/cmd/gootstrap
```

# Usage

Create the repository for your project, clone the
empty project and inside it run:

```
gootstrap -module "whatever.com/group/project" -image "group/project"
```

And it will generate all the files so you can start coding.

The **-module** parameter is the name of the Go module exported by your
project. Even if you are not developing a library this is necessary
for imports inside your own project since it defines the full import
path that you will use to import different packages inside your
own project.

Previously to Go modules the import path of packages where defined
by where they where in the file system, now the module declaration
controls how packages are imported (their path). For more
details on how Go modules works check [this](https://blog.golang.org/using-go-modules).

If you want to explicitly pass the directory where the files
are going to be generated you can use **--output-dir**, it default
to the working directory where the command is being executed.

Any files that already exist on the given directory will
be left untouched, gootstrap is only constructive, never destructive.

Files will be generated assuming that your project builds a
executable, like a command or a service/daemon. Support for libraries
is yet to be built.
