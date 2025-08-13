BPF_CLANG=clang
BPF_LLVM=llvm-strip

all:
	$(BPF_CLANG) -g -O2 -target bpf -c bpf/drop.c -o bpf/drop.o
	$(BPF_LLVM) bpf/drop.o
	bpf2go DropObj bpf/drop.o --target native --strip
