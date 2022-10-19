llvm:
	echo "teste"
	llc -filetype=obj source.ll -o main.o
	clang main.o -o main
	./main
