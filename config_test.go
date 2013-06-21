package allegro

import "testing"

func TestConfigLifecycle(t *testing.T) {
	c := CreateConfig()
	defer c.Destroy()
	if c == nil {
		t.FailNow()
	}
}

func TestIterSections(t *testing.T) {
	c := LoadConfig("./test_data/iter.ini")
	defer c.Destroy()

	sectionNames := [3]string{"foo", "bar", "foobar"}
	for _, name := range sectionNames {
		c.AddSection(name)
	}
	found := make([]string, 0, len(sectionNames))

	for secname := range c.IterSections() {
		found = append(found, secname)
	}

	if len(sectionNames) != len(found) {
		t.Errorf("SectionNames and found had different lengths: %v != %v",
			len(sectionNames), len(found))
	}
	for _, secname := range found {
		validName := false
		for _, realName := range sectionNames {
			validName = validName || realName == secname
		}
		if !validName {
			t.Errorf("Name returned but not in sectionNames: %v", secname)
		}
	}
	for _, secname := range sectionNames {
		validName := false
		for _, realName := range found {
			validName = validName || realName == secname
		}
		if !validName {
			t.Errorf("Name in sectionNames but not in found: %v", secname)
		}
	}
}

func TestIterKeys(t *testing.T) {
	c := LoadConfig("./test_data/iter.ini")
	defer c.Destroy()

	sectionName := "foobar"
	keyNames := []string{"foo", "bar", "asd"}

	found := make([]string, 0, len(keyNames))

	for keyName := range c.IterKeys(sectionName) {
		found = append(found, keyName)
	}

	if len(keyNames) != len(found) {
		t.Errorf("KeyNames and found had different lengths: %v != %v",
			len(keyNames), len(found))
	}
	for _, keyName := range found {
		validName := false
		for _, realName := range keyNames {
			validName = validName || realName == keyName
		}
		if !validName {
			t.Errorf("Name returned but not in keyNames: %v", keyName)
		}
	}
	for _, keyName := range keyNames {
		validName := false
		for _, realName := range found {
			validName = validName || keyName == realName
		}
		if !validName {
			t.Errorf("Name in keyNames but not in found: %v", keyName)
		}
	}
}
