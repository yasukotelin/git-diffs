# git-diffs

git-diffs is the git subcommand that is diff files selector.

## Description

You can use the `git-diffs` as a git subcommand like `git diffs`.<br>
This shows diff file names and show diff when you select one.

## Install

```
go get -u github.com/yasukotelin/git-diffs
```

## Usage

```
$ git diffs
[1] .gitignore                      
[2] LICENSE                         
[3] README.md                       
[4] git.go                          
[5] go.mod                          
[6] go.sum                          
[7] main.go                         
[8] test/sample.txt                 
                                    
Select number (empty is cancel) => 5
```

## Licence

MIT

## Author

yasukotelin
