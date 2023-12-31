go-run:
	@cd Go && go run main.go

java-run:
	@cd Java && javac Main.java && java Main

js-run:
	@cd JavaScript && node main.js

cpp-run:
	@cd C++ && g++ loop.cpp -o loop && ./loop

swift-run:
	@cd Swift && swift loop.swift

pre-commit:
	@pre-commit run --all-file
