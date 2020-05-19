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
<img src="https://github.com/ItsAnas/git-commit-configurator/raw/master/img/demo-commit.gif"/>

## Custom rules

```json
{
    "Name": "My commit config",
    "Description": "Use this field to add some metadata",
    "Rules": [
        {
            "Prefix": "feat",
            "Description": "feat: Implement new feature"
        },
        {
            "Prefix": "doc",
            "Description": "doc: writing doc"
        },
        {
            "Prefix": "fix",
            "Description": "fix: fix bug"
        },
        {
            "Prefix": "style",
            "Description": "style: change something without modyfing feature"
        },
        {
            "Prefix": "ci",
            "Description": "ci: improve ci"
        }
    ]
}
```
