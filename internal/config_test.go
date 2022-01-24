package internal

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestReadConfig(t *testing.T) {
	cf := Config{}
	err := cf.ReadConfig()
	if err != nil {
		t.Errorf("except file exist %v", err)
	}
	t.Logf("read count %v", cf.Count)
}

func TestWriteConfig(t *testing.T) {
	cf := Config{Count: 12}
	cf.WriteConfig()
	err := cf.ReadConfig()
	if err != nil {
		t.Errorf("except file exist %v", err)
	}
	if cf.Count != 12 {
		t.Fatal("c must be 12")
	}
}