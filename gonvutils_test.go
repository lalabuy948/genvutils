package genvutils

import (
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
		MongoURL   string `genv:"MONGO_URL, mongodb://localhost:27017"`

		// string join test
		MongoClusterURL string `genv:"MONGO_URL,mongodb://mongodb,mongodb1,mongodb2/?replicaSet=rs0"`

		// edge cases
		RedisURL  string `genv:""`
		RedisPort int    `genv:"REDIS_PORT, 6371"`
		Compress  bool   `genv:"COMPRES, true"`

		// empty
		Bla bool
	}

	var srvConf serverConfig
	err := Parse(&srvConf)
	if err != nil {
		t.Errorf("Parse(&srvConf) | return %v;", err)
	}
	if srvConf.ServerPort != "8080" {
		t.Errorf("Parse(&srvConf) | ServerPort = %v; want 8080", srvConf.ServerPort)
	}
	if srvConf.MongoURL != "mongodb://localhost:27017" {
		t.Errorf("Parse(&srvConf) | MongoURL = %v; want mongodb://localhost:27017", srvConf.MongoURL)
	}
	if srvConf.MongoClusterURL != "mongodb://mongodb,mongodb1,mongodb2/?replicaSet=rs0" {
		t.Errorf("Parse(&srvConf) | MongoClusterURL = %v; want mongodb://mongodb,mongodb1,mongodb2/?replicaSet=rs0", srvConf.MongoClusterURL)
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
	if srvConf.MongoURL != "mongodb://localhost:76623" {
		t.Errorf("Parse(&srvConf) | MongoURL = %v; want mongodb://localhost:76623", srvConf.MongoURL)
	}

	os.Unsetenv("SERVER_PORT")
	os.Unsetenv("MONGO_URL")
}
