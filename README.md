# simpledog

A simple process watchdog. Ensures child process won't outlive parent.

- Spawns arguments as a subprocess.
- Reads stdin forever looking for EOF (e.g. parent exit)
- Terminates subprocess if parent exit detected

Implemented with Go 1.4

Date: January 2015

## Authors

- Matt Kangas <kangas@gmail.com>
