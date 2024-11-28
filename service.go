package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
	"math"
	"net/http"
	"os"
	"strconv"
	// "encoding/json"
	"strings"
)

type Response struct {
	Code int
	Json string
}

var log = logrus.New()
var port = ":9999"
var level = logrus.InfoLevel
var slice []int = readInputSlice("input.txt")

func init() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	if err := viper.ReadInConfig(); err != nil {
		log.Infof("The configuration file was not read: %v", err)
	}

	portconf := viper.GetInt("PORT")
	if portconf <= 0 || portconf > 65535 {
		log.Infof("Invalid port %v, will use 9999 instead", portconf)
	} else {
		port = fmt.Sprintf(":%d", portconf)
	}

	switch viper.GetString("LOG_LEVEL") {
	case "Info":
		level = logrus.InfoLevel
	case "Debug":
		level = logrus.DebugLevel
	case "Error":
		level = logrus.ErrorLevel
	default:
		log.Info("Incorrect logging mode, Info mode will be used")
	}

	log.SetLevel(level)

	file, err := os.OpenFile("service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to open log file, using standard output")
	}

}

func main() {

	r := gin.Default()
	r.Use(gin.LoggerWithWriter(log.Writer()))

	r.GET("/endpoint/:value", func(c *gin.Context) {
		value := c.Param("value")
		intValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Conversion error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error message": "Bad input data"})
		} else {
			result := indexSearch(slice, intValue)
			c.Data(result.Code, "application/json", []byte(result.Json))
		}
	})

	r.Run(port)
}

func readInputSlice(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			fmt.Println("Conversion error:", err)
		} else {
			numbers = append(numbers, int(number))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return numbers
}

func indexSearch(arr []int, target int) Response {
	if target > 1000000 || target < 0 {
		return prepareResponse("error", "Out of range", -1)
	}

	idx, found := slices.BinarySearch(arr, target)
	if found {
		log.Debugf("FOUND EXACT MATCH: %d", target)
		return prepareResponse("success", "Exact match", idx)
	} else {
		var lowValue = int(math.Ceil(float64(target) * 0.9))
		var highValue = int(math.Floor(float64(target) * 1.1))
		log.Debugf("LOW VALUE: %d", lowValue)
		log.Debugf("HIGH VALUE: %d", highValue)
		log.Debugf("INDEX: %d", idx)
		log.Debugf("FOUND VALUE: %d", arr[idx])

		if arr[idx] <= highValue && arr[idx] >= lowValue {
			log.Debugf("FOUND TEN PERCENT LEVEL VALUE: %d", arr[idx])
			return prepareResponse("success", "Conformation is at `10% level`", idx)
		}

	}

	return prepareResponse("error", "Index is not found", -2)
}

func prepareResponse(status string, message string, index int) Response {

	var builder strings.Builder
	var response Response

	builder.WriteString("{")
	if status == "success" {
		response.Code = http.StatusOK
		builder.WriteString(fmt.Sprintf(`"index": "%d"`, index))
	} else {
		response.Code = http.StatusBadRequest
		builder.WriteString(fmt.Sprintf(`"index": "Not found"`))
		builder.WriteString(fmt.Sprintf(`, "error_message": "%s"`, message))
	}
	builder.WriteString("}")
	response.Json = builder.String()
	log.Debugf("RESPONSE JSON STRING: '%s'", response.Json)

	return response
}
