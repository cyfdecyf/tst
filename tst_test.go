package tst

import (
	"reflect"
	"testing"
)

func TestTST(t *testing.T) {
	testData := []struct {
		key string
		val string
	}{
		{"hello", "world"},
		{"he", "llo"},
		{"h", "ello"},
		{"hel", "lo"},
		{"foo", "bar"},
		{"helloworld", "great"},
		{"hey", "yes"},
		{"mind", "one's"},
		{"p's", "and q's"},
	}

	trie := &Trie{}
	for _, td := range testData {
		trie.Put(td.key, td.val)
		if trie.Get(td.key) == nil {
			t.Errorf("get nil after inserting (%s, %s)\n", td.key, td.val)
		}
	}

	for _, td := range testData {
		val := trie.Get(td.key).(string)
		if val != td.val {
			t.Errorf("(%s, %s) get wrong value: %s", td.key, td.val, val)
		}
	}

	trie.Put("", nil)
	if trie.Get("") != nil {
		t.Error("get empty string should return nil")
	}

	noSuchKey := []string{"aaa", "bbb", "heee", "ddd"}
	for _, s := range noSuchKey {
		if trie.Get(s) != nil {
			t.Errorf("get %s should return nil\n", s)
		}
	}
}

func TestGetShortestPrefix(t *testing.T) {
	testData := []struct {
		key string
		val string
	}{
		{"hello", "world"},
		{"h", "e"},
		{"foo", "bar"},
		{"com", "com"},
		{"com.example", "example"},
	}

	trie := &Trie{}
	for _, td := range testData {
		trie.Put(td.key, td.val)
	}

	searchData := []struct {
		key string
		val interface{}
	}{
		{"hey", "e"},
		{"hello", "e"},
		{"com", "com"},
		{"come", "com"},
		{"communication", "com"},
		{"coo", nil},
		{"bar", nil},
	}

	for _, td := range searchData {
		val := trie.GetShortestPrefix(td.key)
		if !reflect.DeepEqual(val, td.val) {
			t.Error("key:", td.key, "expected:", td.val, "got:", val)
		}
	}
}
