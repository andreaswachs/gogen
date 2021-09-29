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
	cd src && go build && mv gogen ..

.PHONY: install
install:
	cd src && go install

.PHONY: run
run:
	cd src && go run . $(RUN_ARGS)

.PHONY: clean
clean:
	rm gogen