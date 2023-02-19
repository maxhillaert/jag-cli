package generate

import (
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"strconv"
)

func JsonPayload(size int) ([]byte, error) {
	var payload []byte
	var err error

	// Generate a JSON payload until we reach the desired size
	for len(payload) < size {
		// Generate fake data using the faker library
		data := struct {
			Field1  string `faker:"word"`
			Field2  string `faker:"word"`
			Field3  string `faker:"word"`
			Field4  string `faker:"word"`
			Field5  string `faker:"word"`
			Field6  string `faker:"word"`
			Field7  string `faker:"word"`
			Field8  string `faker:"word"`
			Field9  string `faker:"word"`
			Field10 string `faker:"word"`
		}{}

		err = faker.FakeData(&data)
		if err != nil {
			return nil, err
		}

		// Convert the data to a JSON string and append it to the payload
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		payload = append(payload, jsonBytes...)
		payload = append(payload, []byte(",")...)
	}

	// Remove the trailing comma from the final JSON object
	if len(payload) > 0 {
		payload[len(payload)-1] = byte('}')
	}

	return payload, nil
}

func PayloadFromSizeString(sizeStr string) ([]byte, error) {
	// Parse the size string to determine the target size in bytes
	size := 0
	lastChar := sizeStr[len(sizeStr)-1]
	if lastChar == 'm' || lastChar == 'M' {
		sizeStr = sizeStr[:len(sizeStr)-1]
		sizeInt, err := strconv.Atoi(sizeStr)
		if err != nil {
			return nil, err
		}
		size = sizeInt * 1024 * 1024
	} else if lastChar == 'k' || lastChar == 'K' {
		sizeStr = sizeStr[:len(sizeStr)-1]
		sizeInt, err := strconv.Atoi(sizeStr)
		if err != nil {
			return nil, err
		}
		size = sizeInt * 1024
	} else {
		sizeInt, err := strconv.Atoi(sizeStr)
		if err != nil {
			return nil, err
		}
		size = sizeInt
	}

	// Generate a JSON payload of the specified size using the generateJsonPayload function
	payload, err := JsonPayload(size)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
