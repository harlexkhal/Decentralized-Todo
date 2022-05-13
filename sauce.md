I installed solidity compiler building directly from sourc using cmake (followed solidity docs)

I installed protoc which is really not required but just to prevent any errors from its github repo go to /releases and wget or curl the linux version and then unzipped it get into the directory and ```mv bin/protoc $GOBIN```

I had go ethereum aleady, so i navigated to the $GO directory went to go ethereum directory entered ethereum@version directory and ran the command ```make devtool``` which installs ```abigen```

$ solc --bin --abi contract/todo.sol -o build => (builds two files in the build directory Todo.abi and Todo.bin
$ abigen --bin=build/Todo.bin --abi=build/Todo.abi --pkg=todo --out=gen/todo.go => (used to build the todo.go)