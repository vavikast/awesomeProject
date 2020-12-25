package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main()  {
	csvFilename := flag.String("csv", "problems.csv", "a csv   file in the format of 'question,anser'")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file : %s", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided csv file")
	}
	problems := paraseLine(lines)
	correct := 0
	go func() {
		<-time.After(5*time.Second)
		os.Exit(0)
	}()

	for i,p := range  problems{
		fmt.Printf("Problem #%d: %s\n",i+1,p.q)
		var answer string
		fmt.Scanf("%s\n",&answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n",correct,len(problems))
}
func exit( msg string)  {
	fmt.Println(msg)
	os.Exit(1)

}

type problem struct {
	q string
	a string
}

func paraseLine(lines [][]string) []problem  {
	ret := make([]problem,len(lines))
	for i,line := range lines{
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return  ret
}