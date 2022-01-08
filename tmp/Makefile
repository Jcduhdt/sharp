all: install

install: clean gen_output build
	mkdir -p output/bin
	mv sharp output/bin/

gen_output:
	mkdir -p output/log
	mkdir -p output/conf
	mkdir -p output/data/export
	cp -rf conf/* output/conf
	cp -rf control.sh output/

clean:
	rm -rf output/bin
	rm -rf output/conf
	rm -rf output/control.sh
	rm -rf output/data

build:
	go build -o sharp

run:
	./output/bin/sharp # 前台运行，方便退出