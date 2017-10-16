package settings

import (
	"os"
	"log"
	"bufio"
	"strings"
)

type Settings struct {
	settings    map[string]interface{}
}

// method prepares config item to key and value.
// item is a string pattern {key} = {value} and it should be returned as key (string) and value (interface).
func (settings *Settings) prepareConfigItem(item string) (string, string) {
	data := strings.Split(item, "=")

	key := strings.Trim(data[0], " ")
	value := strings.Trim(data[1], " ")

	return key, value
}


// method is reading config and returns go lang object in pattern map[string]interface{}.
func (settings *Settings) Read(path string) map[string]string {
	result := make(map[string]string)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		item := scanner.Text()

		if strings.HasPrefix(item, "//") == false {
			key, value := settings.prepareConfigItem(scanner.Text())
			result[key] = value
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
