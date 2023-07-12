.PHONY: build install clean watchStyles watchPages watch

build:
	./bin/alvu
	./bin/tailwindcss -i ./main.css -o ./dist/style.css --minify

install:
	mkdir -p ./bin
	# Downloading alvu
	# https://github.com/barelyhuman/alvu
	curl -sf https://goblin.run/github.com/barelyhuman/alvu | PREFIX=./bin sh
	chmod +x ./bin/alvu
	# Downloading Tailwind CSS CLI for macOS arm64
	# https://github.com/tailwindlabs/tailwindcss/releases/latest
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
	chmod +x ./tailwindcss-macos-arm64
	mv tailwindcss-macos-arm64 ./bin/tailwindcss

clean:
	rm ./bin/alvu ./bin/tailwindcss

watchStyles:
	./bin/tailwindcss -i ./main.css -o ./dist/style.css --watch

watchPages:
	ls ./pages/*.{html,md} | entr -cr ./bin/alvu -serve

watch:
	${MAKE} -j4 watchPages watchStyles 
