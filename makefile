llvm:
	echo "generate code object"
	mv source.tmp source.c
	clang++ -S -emit-llvm source.c -o source.ll
	llc -filetype=obj source.ll -o main.o
	rm source.c
	rm source.ll
	clang++ main.o -o main
	rm main.o
	./main
