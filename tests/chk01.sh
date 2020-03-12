rm *.out #remove previous student output files
rm ex # remove previous student executable
go build -o ex $1 # build new student executable

./ex fcfs01.txt>fcfs01.out
diff fcfs01.base fcfs01.out
./ex sstf01.txt>sstf01.out
diff sstf01.base sstf01.out
./ex scan01.txt>scan01.out
diff scan01.base scan01.out
./ex c-scan01.txt>c-scan.out
diff c-scan.out c-scan01.base
./ex look01.txt>look01.out
diff look01.base look01.out
./ex c-look01.txt>c-look01.base
