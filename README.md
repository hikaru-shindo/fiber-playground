# fiber-playground

This is my playground to check out golang, fiber and some of its features.
It has no purpose beyond testing features and the framework but might help some
other people checking out golang or fiber too.

## Getting Started

This project is setup in a way that you only need [mise](https://mise.jdx.dev/)
to get started.

You may list the available commands with `mise task ls` and run tasks with `mise run <task_name>`.

Building an image requires either Podman (recommended) or Docker to be installed on
you machine.

Some useful commands:

```bash
# running go fmt
mise exec -- go fmt ./...
# building the project
mise run build
# watching and automatically restarting the application
mise run watch
# running tests / lint
mise run test
mise run coverage # same as test but with coverage output
mise run lint
```
