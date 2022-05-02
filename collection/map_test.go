package collection

import "testing"

func TestShouldExtractKeysFromMap(t *testing.T) {

	//given a map
	m := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
		"key4": 4,
	}

	//when extracting its keys
	keys := MapKeys(m)

	//then we should get the keys
	if len(keys) != len(m) {
		t.Error("Invalid return size")
	}

outer:
	for k := range m {
		for _, i := range keys {
			if k == i {
				continue outer
			}
		}
		t.Errorf("Unable to find the key %s", k)
	}
}

func TestShouldExtractValuesFromMap(t *testing.T) {

	//given a map
	m := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
		"key4": 4,
	}

	//when extracting its keys
	values := MapValues(m)

	//then we should get the keys
	if len(values) != len(m) {
		t.Error("Invalid return size")
	}

outer:
	for _, v := range m {
		for _, i := range values {
			if v == i {
				continue outer
			}
		}
		t.Errorf("Unable to find the key %d", v)
	}
}
