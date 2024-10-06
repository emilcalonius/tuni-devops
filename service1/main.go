package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"strings"
	"time"

	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type ServiceInfo struct {
	IP                string `json:"ip"`
	RunningProcesses  string `json:"runningProcesses"`
	DiskSpace         string `json:"diskSpace"`
	TimeSinceLastBoot string `json:"timeSinceLastBoot"`
}

type Response struct {
	Service1 ServiceInfo `json:"service1"`
	Service2 ServiceInfo `json:"service2"`
}

func main() {
	router := gin.Default()
	router.GET("/", getSystemInformation)
	router.Run("0.0.0.0:8199")
}

// Respond with container information
func getSystemInformation(c *gin.Context) {
	service2Ip := os.Getenv("SERVICE2_IP")
	service2Port := os.Getenv("SERVICE2_PORT")

	url := "http://" + service2Ip + ":" + service2Port

	// Get information of service2
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	service1Info := ServiceInfo{getIPInformation(), executeCommand("ps", "-ax"), executeCommand("df"), getTimeSinceLastBoot()}
	var service2Info ServiceInfo
	if err := json.Unmarshal(body, &service2Info); err != nil {
		fmt.Println("Can not parse JSON of service 2 response")
	}

	response := Response{service1Info, service2Info}

	c.IndentedJSON(http.StatusOK, response)
}

// Returns the primary IP address of the system
func getIPInformation() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "Unable to get IP address information"
	}

	defer conn.Close()
	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP.String()
}

// Execute a command and return the output
func executeCommand(commands ...string) string {
	out, err := exec.Command(commands[0], commands[1:]...).Output()
	if err != nil {
		return "Unable to execute command '" + strings.Join(commands, " ") + "'"
	}
	str := string(out)
	return str
}

// Get time sinco last boot by checking first process start time
func getTimeSinceLastBoot() string {
	out := executeCommand("stat", "/proc/1")
	date := strings.Split(out, "Change: ")[1]
	date = strings.Split(date, ".")[0]
	bootTime, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return "Unable to parse last boot date"
	}
	currentTime := time.Now().Local()
	difference := currentTime.Sub(bootTime)
	hours := int(math.Floor(difference.Hours()))
	minutes := int(math.Floor(difference.Minutes() - float64(hours*60)))
	seconds := int(math.Floor(difference.Seconds() - float64(hours*60*60) - float64(minutes*60)))
	return fmt.Sprintf("%d", hours) + " hours, " + fmt.Sprintf("%d", minutes) + " minutes, " + fmt.Sprintf("%d", seconds) + " seconds"
}
