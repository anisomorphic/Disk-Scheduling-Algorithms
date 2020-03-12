rm *.out #remove previous student output files
rm ex # remove previous student executable
go build -o ex $1 # build new student executable

./ex fcfs20.txt>check
./ex sstf20.txt>>check
./ex scan20.txt>>check
./ex c-scan20.txt>>check
./ex look20.txt>>check
./ex c-look20.txt>>check

rm check
