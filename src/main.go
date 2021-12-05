package main

import (
	"fmt"
	"log"
	"milrabliars/milrabliars"
	"os"
	"runtime"
)

func main() {
	cores := runtime.NumCPU()
	snapshots := [...]int{
		100,
		128,
		256,
		500,
		512,
		1000,
		1024,
		2048,
		4096,
		8192,
		10000,
		16384,
		20000,
		32768,
		50000,
		65536,
		100000,
		131072,
		200000,
		262144,
		400000,
		500000,
		524288,
	}

	runningTotal := milrabliars.NewRunningTotalLiars()

	// make the ../out directory
	err := os.Mkdir("../out", 0777)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	// open the running log file
	runningLog, err := os.OpenFile("../out/running.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	for _, snapshot := range snapshots {
		changeChannel := make(chan milrabliars.Top10ChangesAt)
		go milrabliars.ComputeUntil(runningTotal, snapshot, cores, changeChannel)

		for changes := range changeChannel {
			log.Printf("When N=%d, the top 10 changes:", changes.N)
			PrintTop10Movement(changes.Changes, log.Printf)
			_, err = runningLog.WriteString(fmt.Sprintf("When N=%d, the top 10 changes:\n", changes.N))
			if err != nil {
				log.Fatal(err)
			}

			PrintTop10Movement(changes.Changes, func(format string, v ...interface{}) {
				_, err := fmt.Fprintf(runningLog, format, v...)
				if err != nil {
					log.Fatal(err)
				}
				_, err = runningLog.WriteString("\n")
				if err != nil {
					log.Fatal(err)
				}
			})
			_, err = runningLog.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}

		PrintTop10(runningTotal, log.Printf)
		PrintTop10(runningTotal, func(format string, v ...interface{}) {
			_, err = fmt.Fprintf(runningLog, format, v...)
			if err != nil {
				log.Fatal(err)
			}
			_, err = runningLog.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		})

		snapshotFile := fmt.Sprintf("../out/snapshot-%d.log", snapshot)
		snapshotLog, err := os.OpenFile(snapshotFile, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		PrintTop10(runningTotal, func(format string, v ...interface{}) {
			_, err = fmt.Fprintf(snapshotLog, format, v...)
			if err != nil {
				log.Fatal(err)
			}
			_, err = snapshotLog.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		})
		err = snapshotLog.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	err = runningLog.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func PrintTop10Movement(changes []milrabliars.Top10Movement, logger func(string, ...interface{})) {
	for _, change := range changes {
		if change.OldIndex == -1 {
			logger("%d enters the top 10 at spot %d with %d lies", change.Witness, change.NewIndex+1, change.WitnessNumLies)
		} else if change.NewIndex == -1 {
			logger("%d exits the top 10 with only %d lies", change.Witness, change.WitnessNumLies)
		} else {
			logger("%d is now in position %d with %d lies", change.Witness, change.NewIndex+1, change.WitnessNumLies)
		}
	}
}

func PrintTop10(runningTotal *milrabliars.RunningTotalLiars, logger func(string, ...interface{})) {
	logger("The top 10 least reliable witnesses for values up to %d are:", runningTotal.HighestNumberChecked)
	for i, witness := range runningTotal.Top10Liars {
		if witness == 0 {
			break
		}
		spacer := " "
		if i == 9 {
			spacer = ""
		}
		logger("%d.%s %d with %d lies", i+1, spacer, witness, runningTotal.TimesLied[witness])
	}
}
