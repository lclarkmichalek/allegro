package allegro

import "testing"
import "strconv"

func TestConfigLifecycle(t *testing.T) {
	c := CreateConfig()
	defer c.Destroy()
	if c == nil {
		t.FailNow()
	}
}

func TestIterSections(t *testing.T) {
	c := CreateConfig()
	defer c.Destroy()

	sectionNames := [3]string{"asd", "bsd", "csd"}
	for _, name := range sectionNames {
		c.AddSection(name)
	}
	for secname := range c.IterSections() {
		found := false
		for _, name := range sectionNames {
			if name == secname {
				found = true
			}
		}
		if !found {
			t.Error("Could not find name " + secname)
		}
	}
}

func TestIterKeys(t *testing.T) {
	c := CreateConfig()
	defer c.Destroy()

	sectionName := "foo"
	keyNames := [3]string{"asd", "bsd", "csd"}
	for i, key := range keyNames {
		c.Set(sectionName, key, strconv.Itoa(i))
	}
	for keyName := range c.IterKeys(sectionName) {
		found := false
		for _, name := range keyNames {
			if keyName == name {
				found = true
			}
		}
		if !found {
			t.Error("Could not find key " + keyName)
		}
	}
}
