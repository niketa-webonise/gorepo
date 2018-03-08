package main

import   
		("flag"
		 "os"
 		"fmt"
 		"encoding/csv"
 		"strings"
 		"time")


func main() {
	
	getcsvFile := flag.String("csv","csvfile.csv","csv file that represents a question with respective answer")
	timeLimit := flag.Int("limit",5,"The limit of the quiz in seconds")
	flag.Parse()
	
	file, err := os.Open(*getcsvFile)
	if err != nil {
		fmt.Printf("cannot open csv file:%s\n",*getcsvFile)
		os.Exit(1)

	}
	readCSV := csv.NewReader(file)
	lines,err := readCSV.ReadAll()
	if err != nil{
		fmt.Println("failed to read lines of csv file")
		os.Exit(1)

	} 
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	
	correct := 0

	problemloop:
	for i,p := range problems {
		fmt.Printf("Problem #%d : %s = \n",i+1,p.prob)
		answerCh := make(chan string)

		go func(){
		var answer string 
		fmt.Scanf("%s\n",&answer)
		answerCh <- answer
		}()

		select{
		case <- timer.C:
			fmt.Println()
			break problemloop
        case answer := <-answerCh:
		if answer == p.ans{
			correct ++ 
		         }
	         }
      }

	fmt.Printf("You scored %d out of %d \n",correct,len(problems))
}


func parseLines(lines [][]string) []problem {

	ret := make([]problem,len(lines))
	fmt.Println("ret:",ret)
	for i,line := range lines {
		ret[i] = problem{
			prob : line[0],
			ans : strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type problem struct {
	prob string
	ans string
}