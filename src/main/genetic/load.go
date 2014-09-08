package genetic

/*
	singleton struct for loaded variable info, such as the size of the map
*/
import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//might need to return pointers here.
	return lines, scanner.Err()
}

func LoadFromFile(path string) Mapper {
	//for each line in the file, add a new Location to the map until the file is over

	lines, _ := ReadLines(path)
	eucMap := EuclideanMap{}
	for _, element := range lines {

		i := strings.Split(element, " ")
		if _, err := strconv.Atoi(i[0]); err == nil {
			xi, err1 := strconv.ParseFloat(i[1], 64)
			yi, err2 := strconv.ParseFloat(i[2], 64)

			if err1 != nil || err2 != nil {
				//SOMETHING BREAKS
				log.Println("eating line: ", element)
				break
			}
			eucMap.position = append(eucMap.position, Location{xi, yi})

		}
	}

	//return a pointer here, so we arent constructing, deconstructing
	var m Mapper
	m = &eucMap
	return m
}
