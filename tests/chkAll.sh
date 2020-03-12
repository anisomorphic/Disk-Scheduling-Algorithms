rm *.out #remove previous student output files
rm ex # remove previous student executable
go build -o ex $1 # build new student executable

echo "run all FCFS cases"
./ex fcfs01.txt>fcfs01.out
diff fcfs01.base fcfs01.out
./ex fcfs20.txt>fcfs20.out
diff fcfs20.base fcfs20.out
./ex fcfsPA.txt>fcfsPA.out
diff fcfsPA.base fcfsPA.out

echo "run all SSTF cases"
./ex sstf01.txt>sstf01.out
diff sstf01.base sstf01.out
./ex sstf20.txt>sstf20.out
diff sstf20.base sstf20.out
./ex sstfPA.txt>sstfPA.out
diff sstfPA.base sstfPA.out

echo "run all SCAN cases"
./ex scan01.txt>scan01.out
diff scan01.base scan01.out
./ex scan20.txt>scan20.out
diff scan20.base scan20.out
./ex scanPA.txt>scanPA.out
diff scanPA.base scanPA.out

echo "run all C-SCAN cases"
./ex c-scan01.txt>c-scan.out
diff c-scan.out c-scan01.base
./ex c-scan20.txt>c-scan20.out
diff c-scan20.out c-scan20.base
./ex c-scanPA.txt>c-scanPA.out
diff c-scanPA.out c-scanPA.base

echo "run all LOOK cases"
./ex look01.txt>look01.out
diff look01.base look01.out
./ex look20.txt>look20.out
diff look20.base look20.out
./ex lookPA.txt>lookPA.out
diff lookPA.base lookPA.out

echo "run all C-LOOK cases"
./ex c-look01.txt>c-look01.out
diff c-look01.out c-look01.base
./ex c-look20.txt>c-look20.out
diff c-look20.out c-look20.base
./ex c-lookPA.txt>c-lookPA.out
diff c-lookPA.out c-lookPA.base
