# gopt

Gootstrap stands for Go Bootstrap. It aims to provide bootstraping when
starting Go projects, helping you to go from 0 to 100 and with some
opinionated good practices/tooling like:

* Static Analysis
* Coverage
* Continuous Integration
* Formatting git hooks
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

Create the repository for your project, clone it and inside it run:

```
gootstrap --docker-image <image-name> --project <project name>
```

And it will generate all the files so you can start coding.

If you want to explicitly pass the directory where the files
are going to be generated you can use **--output-dir**.

Any files that already exist on the given directory will
be left untouched, gootstrap is only constructive, never destructive
(at least it strives to be).

Files will be generated assuming that your project builds a
executable, like a command or a service/daemon. Support for libraries
is yet to be built.
