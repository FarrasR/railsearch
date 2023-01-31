# Railsearch
An \*.osm.pbf parser to count things around other things based on their tags written in Go. Originally written to search buildings around railway, hence railsearch.

## Approach
It will get tags from a json file, the program then will search objects on \*.osm.pbf that have that tags and output the file as json. THEN it scans again for things around those objects.

Currently only support nodes on osm, somethign like railway(a way) not supported

## Instruction
If you run/build this on windows (KEK), tough luck, theres a lot of dependencies you need to install which I won't cover (I tried running this on windows to no avail), the recommended way to run/build this on windows is via wsl, and these instructions assumed you already familiar with linux/wsl

1. Install golang https://go.dev/doc/install
2. copy `config.json` from `config.json.sample`
3. go run ./cmd/main.go

## Notes
Thanks to DennyDark for helping me with the formula

## License
licensed under GNU General Public License version 3