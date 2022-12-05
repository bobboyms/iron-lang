llvm:
	echo "generate code object"
	mv source.tmp source.cpp
	clang++ -S -emit-llvm source.cpp -o source.ll -std=c++20 -stdlib=libc++
	llc -filetype=obj source.ll -o main.o
	rm source.ll
	rm source.cpp
	clang++ main.o -o main
	#./main
	rm main.o
	rm main


