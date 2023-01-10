# GO Todo

CLI Todolist written in GO with a nice TUI. Adopted by https://github.com/joefazee/go-toto-app.



### How to use

Build the executable using ```go build```
```
$ go build ./cmd/todo 
````

#### Commands
1. Display the list
```
$ ./todo -list 
````
2. Add a task
```
$ ./todo -add <TASK_NAME> 
````
3. Delete a task
```
$ ./todo -delete=<TASK_INDEX>
````
4. Complete a task
```
$ ./todo -complete=<TASK_INDEX>
````
