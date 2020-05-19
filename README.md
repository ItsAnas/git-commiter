# git-commit-configurator

This program allows you to define and use rules for your git commits. The goal is to have a repository with structured commit but also explain to people who would like to contribute how commits should be written.

## Prerequesite

- [Golang](https://golang.org/)

## Install

```shell
cd git-commit-configurator
go install .
```

## Quick Test

```shell
# Supposed you are not in a git repository
 mkdir test
 cd test
 git init
 echo 'Bonjour' >> foo
 git add foo
 git-commit-configurator
```

## Custom rules

Soon 🔜:
> At the end of the project you should be able to use a file .commit.json that we will read by the program to let you commit correctly with a description of each header