# Gift Appreciator

This app is a small CLI app (and possibly web) that will let you input data about gifts, givers and personalised messages. It will then generate emails from a template then send the emails to the relevant people.
```
Available Commands:
  add         Adds an entry
  clear       Removes all entries
  list        List all entries
  gen         Generates emails of appreciation and prints to STDOUT
  send        Generates then sends email using SMTP defined in the config
  web         Begin a web server to allow input from an internet browser
```

# Spinup Instructions

- Install Go & Glide

```
brew install go
brew install glide
```

- Get the repo

```
go get github.com/nii236/appreciator
```

- Setup Repo - navigate to directory and run:

```
cd $GOPATH/src/github.com/nii236/appreciator
glide install
gulp watch
open localhost:3000
```
