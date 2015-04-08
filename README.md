# timebench

Command line tool to calculate min/max/avg/stddev of a list of durations.

## Usage

    ./timebench [-p PRECISION]
    
    Arguments:
    -p=ms	Precision of the output. Options: h, m, s, ms, us, ns.

## Example

    ./timebench <<'EOF'
        1m30s
        0m23.41s
        49.2s
        2m12s
        2m20.4s
    EOF    

    pts	5
    avg	1m27.002s
    min	23.41s
    max	2m20.4s

## Authors

- Ahmet Alp Balkan

