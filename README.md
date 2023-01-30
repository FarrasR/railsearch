# Railsearch

An openstreetmap parser to search relevant things around railways written in Go.
I was going to include cobra for this app but the scope of this app is smaller than i thought so wcyd

## Approach
it will get lists of nodes that have tags "train" in it and scan again to see if a building is withing radius 

## Instruction

If you run/build this on windows (KEK), tough luck, theres a lot of mendo dependency you need to install which i won't cover (trust me i tried running this on windows to no avail), the best way to run/build this on windows is via wsl, and these instructions assumed you already familiar with linux/wsl

1. Install golang https://go.dev/doc/install
2. go run ./cmd/main.go

## Notes
licensed under GNU General Public License version 3