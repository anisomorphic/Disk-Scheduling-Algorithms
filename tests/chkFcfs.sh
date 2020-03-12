rm *.out #remove previous student output files
rm ex # remove previous student executable
go build -o ex $1 # build new student executable

echo "run all FCFS samples"
./ex fcfs01.txt>fcfs01.out
diff fcfs01.base fcfs01.out
./ex fcfs20.txt>fcfs20.out
diff fcfs20.base fcfs20.out
./ex fcfsPA.txt>fcfsPA.out
diff fcfsPA.base fcfsPA.out
