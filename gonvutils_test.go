package gonvutils

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
