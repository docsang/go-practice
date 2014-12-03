go-practice
===========

only for my go study

## go install 

	mkdir $HOME/go
	
	export GOPATH=$HOME/go
	export PATH=$PATH:$GOPATH/bin

## run cli demo

	go run run.go 
	
- 学习如何执行go代码
- 学习如何调用shell

## log demo

https://github.com/docsang/go-cli-log

执行过程

	cd examples
	go get github.com/visionmedia/go-cli-log   
	go run simple.go
	
非常简单，

- 学习go get命令


## c-go-cgo ?

http://golang.org/doc/install/gccgo

http://blog.golang.org/c-go-cgo


执行

	go run test_import.go 


说明：

	/*
	#include <stdlib.h>
	*/
	import "C"

此句不能有空行，否则编译无法通过


```
	➜  go-practice git:(master) ✗ go run test_import.go
	# _/Users/sang/workspace/github/go-practice/rand
	37: error: use of undeclared identifier 'random'
	37: error: use of undeclared identifier 'srandom'
```

参考

- http://tonybai.com/2012/09/26/interoperability-between-go-and-c/

## mem

http://golang.org/ref/mem