# Functional Options in Go

This example Go program makes use of the "Functional Options" pattern to create
and initialize a structure with options. So instead of:
```
myStuct := &MyStruct{
  Field1: default1
  Field2: default2
}
//user options
myString.Field1 := option1
myString.Field2 := option2
```
we can follow this pattern:
```
myStruct := NewMyStruct(
    WithOption(option1),
    WithOption(option2),
    )
```
## Links
  * https://www.sohamkamani.com/golang/options-pattern/#:~:text=%20Functional%20Options%20in%20Go%3A%20Implementing%20the%20Options,Functional%20Options%20to%20our%20Constructor.%20The...%20More%20

## git setup
  * create github project
    * https://github.com/illianov/go-funcopts
  * create top level dev directory locally
    * mkdir /d/dev/github.com/illianov/go-funcopts
    * cd /d/dev/github.com/illianov/go-funcopts
  * initialize git
    * git init
    * git branch -M main
    * git remote add origin git@github.com:illianov/go-funcopts.git
  * add readme.md
    * git add README.md
    * git commit -m "adding readme.md"
    * git push -u origin main
  * update readme.md
    * git commit -m "updating readme.md"
    * git push origin main

