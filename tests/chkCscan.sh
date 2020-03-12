rm *.out #remove previous student output files
rm ex # remove previous student executable
go build -o ex $1 # build new student executable

echo "run all C-SCAN samples"
./ex c-scan01.txt>c-scan01.out
diff c-scan01.base c-scan01.out
./ex c-scan20.txt>c-scan20.out
diff c-scan20.base c-scan20.out
./ex c-scanPA.txt>c-scanPA.out
diff c-scanPA.base c-scanPA.out
