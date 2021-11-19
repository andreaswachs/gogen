# Credit: https://stackoverflow.com/a/14061796
# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

.PHONY: build
build:
	go build ./src; mv ./src/gogen ..

.PHONY: install
install:
	go install ./src

.PHONY: run
run:
	go run ./ $(RUN_ARGS)

.PHONY: clean
clean:
	rm gogen