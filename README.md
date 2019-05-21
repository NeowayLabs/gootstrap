# Gootstrap

Gootstrap stands for Go Bootstrap. It aims to provide bootstraping when
starting Go projects, helping you to go from 0 to 100 and with some
opinionated good practices/tooling like:

* Reproducible builds
* Versioned builds
* Static Analysis
* Coverage Analysis
* Code Formatting
* Continuous Integration
* Git hooks
* Tagged releases

The idea is to give a head start when starting a project not
contemplate every detail that a project may need, so after
generating the files you probably will need to add more
things as necessary.

The generated project relies on [Docker](https://www.docker.com/)
as a development environment build tool and releasing tool.
But the generated code will also work with using Go directly
from the command line in your host.

# Install

To generate a project you need to install gootstrap, if you
have Go installed in your host it is as easy as:

```
go install github.com/NeowayLabs/gootstrap/cmd/gootstrap
```

To use the generated project you will need [Docker](https://www.docker.com/).

# Usage

To quickly checkout how a project is generated run:

```
gootstrap -module "whatever.com/group/project" -image "group/project" -output-dir /tmp/test
```

And you can check your brand new project at **/tmp/test**,
If no **-output-dir** is passed it defaults
to the working directory where the command is being executed.

So if you already have a git repository created for your project
(preferably empty) just clone the project and inside it run:

```
gootstrap -module "whatever.com/group/project" -image "group/project"
```

And it will generate all the files so you can start coding
(you need to add the generated files to git manually).

The **-module** parameter is the name of the Go module exported by your
project. Even if you are not developing a library this is necessary
for imports inside your own project since it defines the full import
path that you will use to import different packages inside your
own project.

Previously to Go modules the import path of packages where defined
by where they where in the file system, now the module declaration
controls how packages are imported (their path). For more
details on how Go modules works check [this](https://blog.golang.org/using-go-modules).

Any files that already exist on the given directory will
be left untouched, gootstrap don't change any files or delete them.

Files will be generated assuming that your project builds a
executable, like a command or a service/daemon. Support for libraries
is yet to be built (although it is not hard to just delete the command
related targets and files).
