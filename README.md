# Railsearch
An \*.osm.pbf parser to count things around other things based on their tags written in Go. Originally written to search buildings around railway, hence railsearch.

## Approach
It will get tags from a json file, the program then will search objects on osmpbf that have that tags and output the file as json. THEN it scans again for things around those objects.

## Instruction
If you run/build this on windows (KEK), tough luck, theres a lot of dependencies you need to install which I won't cover (I tried running this on windows to no avail), the recommended way to run/build this on windows is via wsl, and these instructions assumed you already familiar with linux/wsl

1. Install golang https://go.dev/doc/install
2. go run ./cmd/main.go

## Notes
licensed under GNU General Public License version 3