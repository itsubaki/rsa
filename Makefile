SHELL := /bin/bash

run: clean compile
	scala Hello

compile: clean
	scalac *.scala -deprecation

clean:
	-rm *.class