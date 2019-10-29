TARGET = find-readme

$(TARGET): clean main.go
	go build -o $(TARGET) .

BENCH_CMD_1 = './find-readme ~/ghq' 
BENCH_CMD_2 = 'find ~/ghq -iname "readme.md" -not -path "*/node_modules/*"'
BENCH_CMD_3 = 'ag -l "" ~/ghq | ag -i "readme.md"' 

bench.md: $(TARGET) clear-screen
	hyperfine --warmup 3 \
	  $(BENCH_CMD_1) \
	  $(BENCH_CMD_2) \
	  $(BENCH_CMD_3) \
	  --export-markdown bench.md

.PHONY: clean
clean:
	rm -f $(TARGET)

.PHONY: clear-screen
clear-screen:
	clear
