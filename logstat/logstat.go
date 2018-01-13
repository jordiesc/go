package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

//mapping to file
type Tomcat struct {
	IP       string
	User     string
	Time     time.Time
	Method   string
	Code     string
	Duration int
}

//sorting interface Interface to apliacr sort.Sort
// type ByDuration []Tomcat

// func (a ByDuration) Len() int           { return len(a) }
// func (a ByDuration) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByDuration) Less(i, j int) bool { return a[i].Duration < a[j] }

func main() {
	// read data from CSV file

	csvFile, err := os.Open("./log.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.Comma = ' ' // Use tab-delimited instead of comma <---- here!
	reader.LazyQuotes = true

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var allRecords []Tomcat
	//map of categories in this case urls to group
	//and get information by url or API end point
	mapMethodRectods := make(map[string][]*Tomcat)

	for _, each := range csvData {
		var oneRecord Tomcat

		oneRecord.IP = each[0]
		oneRecord.User = each[1]

		t := each[4][:len(each[4])-1]
		oneRecord.Time = loadDate(each[3][1:] + t)
		oneRecord.Method = each[len(each)-3]
		duration := each[len(each)-1]
		i, err := strconv.Atoi(duration)

		if err != nil {
			fmt.Println(err)
		}
		oneRecord.Duration = i
		allRecords = append(allRecords, oneRecord)

		dicctionary, ok := mapMethodRectods[oneRecord.Method]
		if ok {

			dicctionary = append(dicctionary, &oneRecord)
			mapMethodRectods[oneRecord.Method] = dicctionary
		} else {
			//dicctionary  []*Tomcat
			dicctionary = append(dicctionary, &oneRecord)
			mapMethodRectods[oneRecord.Method] = dicctionary
		}

		//errores(err)
	}

	PrintMostDurableTomcatRecords(allRecords)

	for k, v := range mapMethodRectods {

		fmt.Printf(" clave -> %s\n", k)
		PrintMostDurablePartial(v)
	}

}

//print the most expensive records
func PrintMostDurableTomcatRecords(allRecords []Tomcat) {
	//ordena el slice por el campo Duration
	sort.Slice(allRecords, func(i, j int) bool { return allRecords[i].Duration < allRecords[j].Duration })
	var sublista []Tomcat
	if len(allRecords) > 10 {
		sublista = allRecords[len(allRecords)-11 : len(allRecords)-1]
	}

	for _, record := range sublista {
		fmt.Printf("tomcat valores timpoe %d metodo %s \n", record.Duration, record.Method)
	}
}

//print the most expensive records
func PrintMostDurablePartial(allRecords []*Tomcat) {
	//ordena el slice por el campo Duration
	sort.Slice(allRecords, func(i, j int) bool { return allRecords[i].Duration < allRecords[j].Duration })
	var sublista []*Tomcat
	if len(allRecords) > 10 {
		sublista = allRecords[len(allRecords)-11 : len(allRecords)-1]
	}

	for _, record := range sublista {
		fmt.Printf("tomcat valores timpoe %d metodo %s \n", record.Duration, record.Method)
	}
}

func loadDate(tiempo string) time.Time {
	layout := "02/Jan/2006:15:04:05 -0700"
	t, _ := time.Parse(layout, tiempo)
	return t
}

func errores(err error) {
	if err != nil {
		panic(err)
	}
}

type statisticsTomcat struct {
	mean float32
	max  float32
	sum  int16
}
