package tst

import (
	"reflect"
	"testing"
)

func reverseString(s string) string {
	b := []byte(s)
	n := len(b)
	for i := 0; i < n/2; i++ {
		b[i], b[n-1-i] = b[n-1-i], b[i]
	}
	return string(b)
}

func TestTST(t *testing.T) {
	testData := []struct {
		key string
		val string
		rk  bool
	}{
		{"hello", "world", false},
		{"he", "llo", true},
		{"h", "ello", false},
		{"hel", "lo", true},
		{"foo", "bar", false},
		{"helloworld", "great", true},
		{"hey", "yes", false},
		{"mind", "one's", true},
		{"p's", "and q's", false},
		{"4321", "1234", true},
		{"5678", "5678", false},
	}

	trie := &Trie{}
	for _, td := range testData {
		put := (*Trie).Put
		get := (*Trie).Get
		if td.rk {
			put = (*Trie).PutRK
			get = (*Trie).GetRK
		}
		put(trie, td.key, td.val)
		if get(trie, td.key) == nil {
			t.Errorf("get nil after inserting (%s, %s)\n", td.key, td.val)
		}
	}

	for _, td := range testData {
		get := (*Trie).Get
		get2 := (*Trie).GetRK
		if td.rk {
			get = (*Trie).GetRK
			get2 = (*Trie).Get
		}

		v := get(trie, td.key)
		v2 := get2(trie, reverseString(td.key))
		if v != v2 {
			t.Errorf("(%s, %s) reversed search for reversed key does NOT match\n",
				td.key, td.val)
		}
		if v == nil {
			t.Fatalf("(%s, %s) value not found\n", td.key, td.val)
		}
		val := v.(string)
		if val != td.val {
			t.Errorf("(%s, %s) get wrong value: %s\n", td.key, td.val, val)
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
		val2 := trie.GetShortestPrefixRK(reverseString(td.key))
		if val != val2 {
			t.Errorf("(%s, %s) reversed search for reversed key does NOT match\n",
				td.key, td.val)
		}
		if !reflect.DeepEqual(val, td.val) {
			t.Error("key:", td.key, "expected:", td.val, "got:", val)
		}
	}
}
