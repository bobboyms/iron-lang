llvm:
	echo "generate code object"
	mv source.tmp source.c
	clang++ -S -emit-llvm source.c -o source.ll -std=c++20 -stdlib=libc++
	llc -filetype=obj source.ll -o main.o
	rm source.c
	rm source.ll
	#rm main.o
	clang++ main.o -o main
	rm main.o


