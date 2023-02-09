# Railsearch
An \*.osm.pbf parser to count things around other things based on their tags written in Go. Originally written to search buildings around railway, hence railsearch.

## Approach
It will index all node and way in a file with the desired tag. then convert all those waynodes into nodes that can be easily calculated. 

## Limitations
Because osm uses latitude/longitude as its coordinate, the program will convert this into a mercator projection, which affect accuracy somewhat.

The radius used is also square instead of circle to accelerate computation, as using circle will take much more time to compute(also hard maths).

The way for target also used the first node in the way instead of the way centroid, because calculating those is also expensive on database side.

## Instruction
If you run/build this on windows (KEK), tough luck, theres a lot of dependencies you need to install which I won't cover (I tried running this on windows to no avail), the recommended way to run/build this on windows is via wsl, and these instructions assumed you already familiar with linux/wsl

1. Install golang https://go.dev/doc/install
2. Install mysql 
3. Copy `config.json` from `config.json.sample`
4. change the config accordingly
5. DO THIS IN ORDER
6. `make migrate` to migrate necessary table
7. `make sanity` to check the nodes and way count (useful for estimates)
8. `make index` to index all nodes
9. `make processway` to process the way into indexed nodes
10. `make target` to search 

Once you're done the target_nodes table should be filled with the node/way data that nears the searched node/way by the configured radius (guaranteed around 78% accuracy)

## Notes
Thanks to Denny Rengganis for helping me with the formula