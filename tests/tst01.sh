rm *.out #remove previous student output files
rm ex # remove previous student executable
go build -o ex $1 # build new student executable

./ex fcfs01.txt>check
./ex sstf01.txt>>check
./ex scan01.txt>>check
./ex c-scan01.txt>>check
./ex look01.txt>>check
./ex c-look01.txt>>check

rm check
