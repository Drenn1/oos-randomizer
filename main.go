package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/jangler/oos-randomizer/graph"
	"github.com/jangler/oos-randomizer/rom"
)

// fatals if the command got the wrong number of arguments
func checkNumArgs(op string, expected int) {
	if flag.NArg() != expected {
		log.Printf("%s takes %d argument(s); got %d",
			op, expected, flag.NArg())
		os.Exit(2)
	}
}

func main() {
	// init flags
	flagGoal := flag.String("goal", "done",
		"comma-separated list of nodes that must be reachable")
	flagForbid := flag.String("forbid", "",
		"comma-separated list of nodes that must not be reachable")
	flagMaxlen := flag.Int("maxlen", -1,
		"if >= 0, maximum number of slotted items in the route")
	flagSeed := flag.String("seed", "",
		"specific random seed to use (32-bit hex number)")
	flagUpdate := flag.Bool(
		"update", false, "update randomized ROM to this version")
	flagVerbose := flag.Bool(
		"verbose", false, "print more detailed output to terminal")
	flag.Parse()

	checkNumArgs("randomizer", 2)

	// load rom
	romData, err := readFileBytes(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	// split node params
	goal := parseDelimitedArg(*flagGoal, ",")
	forbid := []string{}
	if *flagForbid != "" {
		forbid = parseDelimitedArg(*flagForbid, ",")
	}

	// randomize according to params, unless we're just updating
	if *flagUpdate {
		_, err := rom.Update(romData)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		seed := setRandomSeed(*flagSeed)

		summary, summaryDone := getSummaryChannel()
		summary <- fmt.Sprintf("seed: %08x", seed)

		if errs := randomize(romData, flag.Arg(1),
			[]string{"horon village"}, goal, forbid,
			*flagMaxlen, summary, *flagVerbose); errs != nil {
			for _, err := range errs {
				log.Print(err)
			}
			os.Exit(1)
		}

		close(summary)
		<-summaryDone
	}

	// write to file
	f, err := os.Create(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.Write(romData); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote new ROM to %s", flag.Arg(1))
}

// parses a delimited (e.g. with comma) command-line argument, stripping spaces
// around each entry.
func parseDelimitedArg(arg, delimiter string) []string {
	a := make([]string, 0)

	for _, s := range strings.Split(arg, delimiter) {
		a = append(a, strings.TrimSpace(s))
	}

	return a
}

// sets a 32-bit unsigned random seed based on a hexstring, if non-empty, or
// else the current time, and returns that seed.
func setRandomSeed(hexString string) uint32 {
	seed := uint32(time.Now().UnixNano())
	if hexString != "" {
		v, err := strconv.ParseUint(
			strings.Replace(hexString, "0x", "", 1), 16, 32)
		if err != nil {
			log.Fatalf(`fatal: invalid seed "%s"`, hexString)
		}
		seed = uint32(v)
	}
	rand.Seed(int64(seed))

	return seed
}

// return the contents of the names file as a slice of bytes
func readFileBytes(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

// messes up rom data and writes it to a file. this also calls rom.Verify().
func randomize(romData []byte, outFilename string, start, goal,
	forbid []string, maxlen int, summary chan string, verbose bool) []error {
	// make sure rom data is a match first
	if errs := rom.Verify(romData); errs != nil {
		return errs
	}

	// check args against graph
	r := NewRoute(start)
	for _, name := range goal {
		if _, ok := r.Graph[name]; !ok {
			log.Fatal("fatal: unknown goal node: ", name)
		}
	}
	for _, name := range forbid {
		if _, ok := r.Graph[name]; !ok {
			log.Fatal("fatal: unknown forbid node: ", name)
		}
	}

	// give each routine its own random source, so that the same -seed will
	// yield the same result.
	sources := make([]rand.Source, runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		sources[i] = rand.NewSource(rand.Int63())
	}

	// search for route, parallelized
	routeChan := make(chan *RouteLists)
	logChan := make(chan string)
	doneChan := make(chan int)
	log.Printf("using %d threads", runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go searchAsync(rand.New(sources[i]), start, goal, forbid, maxlen,
			verbose, logChan, routeChan, doneChan)
	}

	// log messages from all threads
	go func() {
		for {
			select {
			case msg := <-logChan:
				log.Print(msg)
			case <-doneChan:
				return
			}
		}
	}()

	// get return values
	var rl *RouteLists
	for i := 0; i < runtime.NumCPU(); i++ {
		rl = <-routeChan
		if rl != nil {
			break
		}
	}

	// tell all the other routines to stop
	go func() {
		for {
			doneChan <- 1
		}
	}()

	// didn't find any route
	if rl == nil {
		log.Fatal("fatal: no route found")
	}

	// place selected treasures in slots
	usedLines := make([]string, 0, rl.UsedSlots.Len())
	for rl.UsedSlots.Len() > 0 {
		slotName :=
			rl.UsedSlots.Remove(rl.UsedSlots.Front()).(*graph.Node).Name
		treasureName :=
			rl.UsedItems.Remove(rl.UsedItems.Front()).(*graph.Node).Name
		rom.ItemSlots[slotName].Treasure = rom.Treasures[treasureName]

		usedLines =
			append(usedLines, fmt.Sprintf("%s <- %s", slotName, treasureName))
	}

	// do it! (but don't write anything)
	checksum, err := rom.Mutate(romData)
	if err != nil {
		return []error{err}
	}

	// write info to summary file
	summary <- fmt.Sprintf("sha-1 sum: %x", checksum)
	summary <- ""
	summary <- "used items, in order:"
	summary <- ""
	for _, usedLine := range usedLines {
		summary <- usedLine
	}
	summary <- ""
	summary <- "unused items:"
	summary <- ""
	for e := rl.UnusedItems.Front(); e != nil; e = e.Next() {
		summary <- e.Value.(*graph.Node).Name
	}

	summary <- ""
	summary <- "default seasons:"
	summary <- ""
	for name, area := range rom.Seasons {
		summary <- fmt.Sprintf("%s - %s", name, seasonsByID[int(area.New[0])])
	}

	return nil
}

// searches for a route and logs and returns a route on the given channels.
func searchAsync(src *rand.Rand, start, goal, forbid []string,
	maxlen int, verbose bool, logChan chan string, retChan chan *RouteLists,
	doneChan chan int) {
	// find a viable random route
	r := NewRoute(start)
	retChan <- findRoute(src, r, start, goal, forbid, maxlen, verbose,
		logChan, doneChan)
}
