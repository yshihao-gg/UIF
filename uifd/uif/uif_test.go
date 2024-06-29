package uif

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/stretchr/testify/assert"
)

func TestUnusedPort(t *testing.T) {
	_, err := GetUnusedPort()
	assert.Nil(t, err)
}

func TestGetKey(t *testing.T) {
	defer os.Remove(GetWorkSpace() + "/uif_key.txt")

	key1 := GetKey()
	assert.Equal(t, GetKey(), key1)
}

func TestCloseCore(t *testing.T) {
}

func TestUIFConfig(t *testing.T) {
	defer os.Remove(GetWorkSpace() + "/uif.json")

	config := ReadUIFConfig()
	assert.Equal(t, config, "{}")
	assert.Equal(t, ReadUIFConfig(), config)

	testConfig := "{\"abc\": 1}"
	SaveUIFConfig(testConfig)
	config = ReadUIFConfig()
	assert.Equal(t, testConfig, config)
}

func TestAppInfo(t *testing.T) {
	var m ConnectInfo
	json.Unmarshal([]byte(GetInfo()), &m)

	log.Println(string(GetInfo()))
	assert.NotEmpty(t, m.Path)
	assert.NotEmpty(t, m.StartTime)
}

func TestPing(t *testing.T) {
	fmt.Println(Ping("www.google.com"))
}

func TestGetDefaultInter(t *testing.T) {

	fmt.Println(GetDefaultInterface())
}

func TestGetIPAddress(t *testing.T) {
	fmt.Println(GetOutboundIP())
	assert.NotEqual(t, GetOutboundIP(), "")

	fmt.Println(GetPublicIP())
	assert.NotEqual(t, GetPublicIP(), "")
}

func TestParseVersion(t *testing.T) {
	assert.Equal(t, 1, ParseVersion("0.0.1"))
	assert.Equal(t, 88, ParseVersion("0.0.88"))
	assert.Equal(t, 111, ParseVersion("0.0.111"))

	assert.Equal(t, 201, ParseVersion("0.2.1"))
	assert.Equal(t, 30201, ParseVersion("3.2.1"))
	assert.Equal(t, 778899, ParseVersion("77.88.99"))
	assert.Equal(t, 523, ParseVersion("0.4.123"))
	assert.Equal(t, 523, ParseVersion("v0.4.123"))
	assert.Equal(t, 303, ParseVersion("v0.3.3"))
	assert.Equal(t, 302, ParseVersion("v0.3.2"))

	assert.Equal(t, 0, ParseVersion("v.0.4.123"))
}

func TestDownloadNewestUIF(t *testing.T) {
	DownloadNewestUIF()
}

func TestCheckUpdate(t *testing.T) {
	log.Println(GetNewestCoreVersion())
	log.Println(GetNewestUIFVersion())
}

func TestGetAddress(t *testing.T) {
	assert.Equal(t, "127.0.0.1:9413", GetAPIAddress())
	assert.Equal(t, "127.0.0.1:9527", GetWebAddress())

	port, _ := GetWebAddressPort()
	assert.Equal(t, "9527", port)

	port, _ = GetAPIAddressPort()
	assert.Equal(t, "9413", port)
}

func TestGetSystemInfo(t *testing.T) {
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
	}

	c, err := cpu.Info()
	if err != nil {
		fmt.Println(err)
	}
	p, err := cpu.Percent(time.Second, true)
	fmt.Println(p)

	fmt.Println(c)
	fmt.Println("")

	// convert to JSON. String() is also implemented
	fmt.Println(v)
}
