package main

import (
	"github.com/nhooyr/color/log"
	"github.com/pelletier/go-toml"
)

func necessary(tree *toml.TomlTree, key string) string {
	v := tree.Get(key)
	if v == nil {
		log.Fatalf("%s: missing %q key", pos(tree, ""), key)
	}
	s, ok := v.(string)
	if !ok {
		log.Fatalf("%s: type of %q is incorrect, should be string", pos(tree, key), key)
	}
	return s
}

func optional(tree *toml.TomlTree, key string) string {
	s, ok := tree.GetDefault(key, "").(string)
	if !ok {
		log.Fatalf("%s: type of %q is incorrect, should be string", pos(tree, key), key)
	}
	return s
}

// TODO unnecessary
func pos(tree *toml.TomlTree, key string) string {
	p := tree.GetPosition(key)
	return p.String()
}
