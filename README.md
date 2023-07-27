# Drone & Turret Field Project

## Building & Running

A [`Makefile`](./Makefile) has been created with the following targets:

### Go

* `go-get`: Downloads go dependencies needed for project.
* `go-run`: Directly runs go code found in `./go`.
  * _Note: Requires go to be installed_
* `go-test`: Runs all go tests with a race checker enabled
* `go-test-no-cache`: Same as `go-test` except disables test caching. e(Go caches tests if it detects that no changes have been made since the last test execution)
* `go-build`: Creates an executable named `drone` in `./bin` that can be run with `run` (See next bullet)
* `run`: Attempts the `drone` executable in `./bin`.
  * _Note: Go does not need to be installed and the executable was built targeting macOS_
