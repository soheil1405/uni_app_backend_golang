package env

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

// GetBasePath ...
func GetBasePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	return exPath
}

var c *Config

func Get(key string) interface{} { return c.Get(key) }

func GetString(key string) string { return c.GetString(key) }

func GetInt(key string) int { return c.GetInt(key) }

func GetFloat64(key string) float64 { return c.GetFloat64(key) }

func GetBool(key string) bool { return c.GetBool(key) }

func GetStringMap(key string) map[string]interface{} { return c.GetStringMap(key) }

func GetStringMapString(key string) map[string]string { return c.GetStringMapString(key) }

func GetStringSlice(key string) []string { return c.GetStringSlice(key) }

func Sub(key string) *Config {
	sv := c.Sub(key)
	if sv == nil {
		return nil
	}

	return &Config{sv}
}

func GetViper() *viper.Viper { return c.Viper }

func UnmarshalKey(key string, rawVal interface{}) error { return c.UnmarshalKey(key, rawVal) }

func IsSet(key string) bool { return c.IsSet(key) }

func (v *Config) Init(filename string) {
	v.Viper = viper.GetViper()

	v.SetConfigType("json")

	if filename != "" {
		v.SetConfigFile(filename)
	} else {
		v.AddConfigPath(GetBasePath())
		v.SetConfigName("config")
	}

	// v.SetEnvPrefix(`go-clean`)
	// v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println("Using config file:", v.ConfigFileUsed())
}

// NewViperConfig ...
func NewViperConfig(path string) *Config {
	c = new(Config)
	c.Init(path)
	return c
}

// func FillPlaceHolder(texts ...*string) (err error) { return c.FillPlaceHolder(texts...) }
func FillPlaceHolder(ctx echo.Context, texts ...*string) (err error) {
	return c.FillPlaceHolder(ctx, texts...)
}

// func (v *Config) FillPlaceHolder(texts ...*string) (err error) {
func (v *Config) FillPlaceHolder(ctx echo.Context, texts ...*string) (err error) {
	var (
		re *regexp.Regexp
		m  = make(map[string]interface{})
	)

	if re, err = regexp.Compile(`{{@([\w.]+)}}`); err != nil {
		return
	}

	for _, text := range texts {
		placeHolders := re.FindAllStringSubmatch(*text, -1)
		for _, placeHolder := range placeHolders {
			if len(placeHolder) < 2 {
				continue
			}

			match := placeHolder[0]
			submatch := placeHolder[1]
			// *text = strings.ReplaceAll(*text, match, v.GetString(submatch))

			m[match] = v.GetString(submatch)
		}
	}

	for _, text := range texts {
		*text = helpers.ReplaceMap(*text, m)
	}

	return nil
}
