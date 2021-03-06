package commands

import "testing"

// TestgetSource - Tests getSource function
func TestGetSource(t *testing.T) {
	input := `vpc::https://raw.githubusercontent.com/daidokoro/qaz/master/examples/vpc/config.yml`
	_, _, err := getSource(input)
	if err != nil {
		t.Error(err.Error())
	}
}

// TestAwsSession - tests this awsSession function
func TestAwsSession(t *testing.T) {

	if _, err := manager.GetSess("default"); err != nil {
		t.Error(err.Error())
	}
}

// TestInvoke - test lambda invoke Functions
func TestInvoke(t *testing.T) {
	f := awsLambda{
		name:    "hello",
		payload: []byte(`{"name":"qaz"}`),
	}

	sess, err := manager.GetSess("default")
	if err != nil {
		t.Error(err.Error())
	}

	if err := f.Invoke(sess); err != nil {
		t.Errorf(err.Error())
	}
}

// TestExports - test Excport function
func TestExports(t *testing.T) {
	sess, err := manager.GetSess("default")
	if err != nil {
		t.Error(err.Error())
	}

	if err := Exports(sess); err != nil {
		t.Error(err.Error())
	}
}
