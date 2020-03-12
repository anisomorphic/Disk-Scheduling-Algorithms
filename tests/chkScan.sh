rm *.out #remove previous student output files
rm ex # remove previous student executable
go build -o ex $1 # build new student executable

echo "run all SCAN samples"
./ex scan01.txt>scan01.out
diff scan01.base scan01.out
./ex scan20.txt>scan20.out
diff scan20.base scan20.out
./ex scanPA.txt>scanPA.out
diff scanPA.base scanPA.out
