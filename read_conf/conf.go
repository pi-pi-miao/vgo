package read_conf

import (
	"bufio"
	"errors"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	conf map[string]url.Values
}

func (c *Config) String(tag string) string {
	spl := strings.Split(tag, ".")
	key := strings.Join(spl[1:], "_")
	if len(spl) < 2 || spl[1] == "" {
		return ""
	}

	return c.conf[spl[0]].Get(key)
}

func (c *Config) Int(tag string) (int, error) {
	return strconv.Atoi(c.String(tag))
}

func (c *Config) Int64(tag string) (int64, error) {
	return strconv.ParseInt(c.String(tag), 10, 64)
}

func (c *Config) Float64(tag string) (float64, error) {
	return strconv.ParseFloat(c.String(tag), 64)
}

func ReadFile_Conf(filePath string) (*Config, error) {

	cf := &Config{
		conf: make(map[string]url.Values, 10),
	}

	f, err := NewFileReader(filePath)
	if err != nil {
		return nil, errors.New("Error:can not read file \"" + filePath + "\"")
	}
	defer f.Close()

	tag := ""
	buf := bufio.NewReader(f)
	replacer := strings.NewReplacer(" ", "")

	for {
		lstr, err := buf.ReadString('\n')
		if err != nil && err != errors.New("EOF") {
			break
		}

		if lstr == "" {
			break
		}

		lstr = strings.TrimSpace(lstr)
		if lstr == "" {
			continue
		}

		if idx := strings.Index(lstr, "["); idx != -1 {
			if lstr[len(lstr)-1:] != "]" {
				return nil, errors.New("Error:field to parse this symbol style:\"" + lstr + "\"")
			}
			tag = lstr[1 : len(lstr)-1]
			cf.conf[tag] = url.Values{}
		} else {
			lstr = replacer.Replace(lstr)
			spl := strings.Split(lstr, "=")

			if lstr[0:1] == ";" {
				continue
			}

			if len(spl) < 2 {
				return nil, errors.New("error:" + lstr)
			}
			cf.conf[tag].Set(strings.Replace(spl[0], ".", "_", -1), spl[1])
		}
	}

	return cf, nil
}

func NewFileReader(filePath string) (*os.File, error) {
	if !PathExists(filePath) {
		return nil, errors.New("Error:File not exists:" + filePath)
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}
