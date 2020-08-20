package genvutils

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestIsProduction(t *testing.T) {
	os.Setenv("ENVIRONMENT", "PROD")
	got := IsProduction()
	if got != true {
		t.Errorf("IsProduction() = %v; want true", got)
	}

	os.Setenv("ENVIRONMENT", "PRODUCTION")
	got = IsProduction()
	if got != true {
		t.Errorf("IsProduction() = %v; want true", got)
	}

	os.Setenv("ENVIRONMENT", "BLA")
	got = IsProduction()
	if got != false {
		t.Errorf("IsProduction() = %v; want false", got)
	}

	os.Setenv("APP_ENV", "PROD")
	got = IsProduction()
	if got != true {
		t.Errorf("IsProduction() = %v; want true", got)
	}

	os.Setenv("APP_ENV", "PRODUCTION")
	got = IsProduction()
	if got != true {
		t.Errorf("IsProduction() = %v; want true", got)
	}

	os.Setenv("APP_ENV", "BLA")
	got = IsProduction()
	if got != false {
		t.Errorf("IsProduction() = %v; want false", got)
	}

	os.Unsetenv("ENVIRONMENT")
	os.Unsetenv("APP_ENV")
}

func TestIsDevelopment(t *testing.T) {
	os.Setenv("ENVIRONMENT", "DEV")
	got := IsDevelopment()
	if got != true {
		t.Errorf("IsDevelopment() = %v; want true", got)
	}

	os.Setenv("ENVIRONMENT", "DEVELOPMENT")
	got = IsDevelopment()
	if got != true {
		t.Errorf("IsDevelopment() = %v; want true", got)
	}

	os.Setenv("ENVIRONMENT", "BLA")
	got = IsDevelopment()
	if got != false {
		t.Errorf("IsDevelopment() = %v; want false", got)
	}

	os.Setenv("APP_ENV", "DEV")
	got = IsDevelopment()
	if got != true {
		t.Errorf("IsDevelopment() = %v; want true", got)
	}

	os.Setenv("APP_ENV", "DEVELOPMENT")
	got = IsDevelopment()
	if got != true {
		t.Errorf("IsDevelopment() = %v; want true", got)
	}

	os.Setenv("APP_ENV", "BLA")
	got = IsDevelopment()
	if got != false {
		t.Errorf("IsDevelopment() = %v; want false", got)
	}

	os.Unsetenv("ENVIRONMENT")
	os.Unsetenv("APP_ENV")
}

func TestIsTesting(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TEST")
	got := IsTesting()
	if got != true {
		t.Errorf("IsTesting() = %v; want true", got)
	}

	os.Setenv("ENVIRONMENT", "TESTING")
	got = IsTesting()
	if got != true {
		t.Errorf("IsTesting() = %v; want true", got)
	}

	os.Setenv("ENVIRONMENT", "BLA")
	got = IsTesting()
	if got != false {
		t.Errorf("IsTesting() = %v; want false", got)
	}

	os.Setenv("APP_ENV", "TEST")
	got = IsTesting()
	if got != true {
		t.Errorf("IsTesting() = %v; want true", got)
	}

	os.Setenv("APP_ENV", "TESTING")
	got = IsTesting()
	if got != true {
		t.Errorf("IsTesting() = %v; want true", got)
	}

	os.Setenv("APP_ENV", "BLA")
	got = IsTesting()
	if got != false {
		t.Errorf("IsTesting() = %v; want false", got)
	}

	os.Unsetenv("ENVIRONMENT")
	os.Unsetenv("APP_ENV")
}

func TestGetEnv(t *testing.T) {
	os.Setenv("SERVER_PORT", "8080")
	got := GetEnv("SERVER_PORT", "8080")
	if got != "8080" {
		t.Errorf("GetEnv(\"SERVER_PORT\", \"8080\") = %v; want 8080", got)
	}

	got = GetEnv("REDIS_PORT", "6379")
	if got != "6379" {
		t.Errorf("GetEnv(\"REDIS_PORT\", \"6379\") = %v; want 6379", got)
	}

	os.Unsetenv("SERVER_PORT")
}

func TestParse(t *testing.T) {

	type serverConfig struct {
		// simple test
		ServerPort string `genv:"SERVER_PORT,8080"`
		MongoUrl   string `genv:"MONGO_URL,mongodb://localhost:27017"`

		// string join test
		MongoClusterUrl string `genv:"MONGO_URL,mongodb://mongodb,mongodb1,mongodb2/?replicaSet=rs0"`

		// edge cases
		RedisUrl   string `genv:""`
		RedisPort  int    `genv:"REDIS_PORT, 6371"`
		Compress   bool   `genv:"COMPRES, true"`

		// empty
		Bla        bool
	}

	var srvConf serverConfig
	err := Parse(&srvConf)
	if err != nil {
		t.Errorf("Parse(&srvConf) | return %v;", err)
	}
	if srvConf.ServerPort != "8080" {
		t.Errorf("Parse(&srvConf) | ServerPort = %v; want 8080", srvConf.ServerPort)
	}
	if srvConf.MongoUrl != "mongodb://localhost:27017" {
		t.Errorf("Parse(&srvConf) | MongoUrl = %v; want mongodb://localhost:27017", srvConf.MongoUrl)
	}
	if srvConf.RedisPort != 6371 {
		t.Errorf("Parse(&srvConf) | RedisPort = %v; want 6371 as int", srvConf.RedisPort)
	}
	if srvConf.Compress != true {
		t.Errorf("Parse(&srvConf) | Compress = %v; want true", srvConf.Compress)
	}

	os.Setenv("SERVER_PORT", "8181")
	os.Setenv("MONGO_URL", "mongodb://localhost:76623")

	err = Parse(&srvConf)
	if err != nil {
		t.Errorf("Parse(&srvConf) | return %v;", err)
	}
	if srvConf.ServerPort != "8181" {
		t.Errorf("Parse(&srvConf) | ServerPort = %v; want 8181", srvConf.ServerPort)
	}
	if srvConf.MongoUrl != "mongodb://localhost:76623" {
		t.Errorf("Parse(&srvConf) | MongoUrl = %v; want mongodb://localhost:76623", srvConf.MongoUrl)
	}

	os.Unsetenv("SERVER_PORT")
	os.Unsetenv("MONGO_URL")
}

func TestLoad(t *testing.T) {
	err := Load()
	if err != nil && err != ErrDotenvNotFound {
			t.Errorf("Load() | error %v;", err)
	}

	err = ioutil.WriteFile(".env", []byte(`
BLA_BLA=42
# some comment`), 0755)
	err = Load(".env")
	if err != nil {
		t.Errorf("Load() | error %v;", err)
	}
	got := os.Getenv("BLA_BLA")
	if got != "42" {
		t.Errorf("Load() | BLA_BLA = %v; want 42", got)
	}
	err = os.Unsetenv("BLA_BLA")
	if err != nil {
		t.Errorf("os.Unsetenv(BLA_BLA) | error %v;", err)
	}

	err = ioutil.WriteFile(".env.development", []byte(`
# some comment
BLO_BLO=42
# another one`), 0755)
	err = Load(".env.development")
	if err != nil {
		t.Errorf("Load(.env.development) | error %v;", err)
	}
	got = os.Getenv("BLO_BLO")
	if got != "42" {
		t.Errorf("Load(.env.development) | BLO_BLO = %v; want 42", got)
	}
	err = os.Unsetenv("BLO_BLO")
	if err != nil {
		t.Errorf("os.Unsetenv(BLO_BLO) | error %v;", err)
	}

	err = ioutil.WriteFile(".env.test.local", []byte(`
# some comment
BLU_BLU=42
# another one`), 0755)
	err = Load(".env.test.local")
	if err != nil {
		t.Errorf("Load(.env.test.local) | error %v;", err)
	}
	got = os.Getenv("BLU_BLU")
	if got != "42" {
		t.Errorf("Load(.env.test.local) | BLU_BLU = %v; want 42", got)
	}
	err = os.Unsetenv("BLU_BLU")
	if err != nil {
		t.Errorf("os.Unsetenv(BLU_BLU) | error %v;", err)
	}

	err = os.Remove(".env")
	err = os.Remove(".env.development")
	err = os.Remove(".env.test.local")
	if err != nil {
		t.Errorf("os.Remove(.env) | error %v;", err)
	}
}

func TestFileExists(t *testing.T) {
	got := fileExists("README.md")
	if got != true {
		t.Errorf("fileExists(README.md) = %v; want true", got)
	}
	got = fileExists("god.hs")
	if got != false {
		t.Errorf("fileExists(god.hs) = %v; want false", got)
	}
}

func TestGetFromPriorityList(t *testing.T) {
	got, err := getFromPriorityList()
	if err != ErrDotenvNotFound {
		t.Errorf("getFromPriorityList() = %v; want ErrDotenvNotFound", got)
	}

	err = ioutil.WriteFile(".env", []byte(``), 0755)
	err = ioutil.WriteFile(".env.local", []byte(``), 0755)
	err = ioutil.WriteFile(".env.production", []byte(``), 0755)
	if err != nil {
		t.Errorf("ioutil.WriteFile | error %v;", err)
	}
	got, err = getFromPriorityList()
	if got != ".env.production" {
		t.Errorf("getFromPriorityList() = %v; want .env.production", got)
	}

	err = os.Remove(".env")
	err = os.Remove(".env.local")
	err = os.Remove(".env.production")
}
