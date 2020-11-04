# PDDIKTI API Example

Access dikti API and get lecturer data from all registered university in indonesian
use by https://pddikti.kemdikbud.go.id and then print it to csv file

<br>

## Problem

duplicate data in csv file

<br>



## Disclaimer

i do not try to do anything criminal activity, i accidentally just discover access to api used by https://pddikti.kemdikbud.go.id in their own js file, all this is just for educational purposes

<br>


## How to run via local fo golang dev

now lets try get basic list of dosen from university
by searching their name containing letter, example : A

```

go build && ./pddikti-api-example --mode=basic --univ={UNIVERSITY-ID} --ls={LECTURER-SEARCH}

```

now lets try get list of detail dosen from university
by searching their name containing letter, example : A

```

go build && ./pddikti-api-example --mode=advance --univ={UNIVERSITY-ID} --ls={LECTURER-SEARCH}

```