package main

import (
	"slices"
	"testing"
)

var (
	hero1 = &Hero{
		"hero1",
		"dps",
		true,
	}
	hero2 = &Hero{
		"hero2",
		"dps",
		false,
	}
	hero3 = &Hero{
		"hero3",
		"support",
		true,
	}
	hero4 = &Hero{
		"hero4",
		"support",
		false,
	}
	hero5 = &Hero{
		"hero5",
		"tank",
		true,
	}
	hero6 = &Hero{
		"hero6",
		"tank",
		false,
	}
	testHeroes = []*Hero{
		hero1,
		hero2,
		hero3,
		hero4,
		hero5,
		hero6,
	}
)

func TestGetValidHeroes(t *testing.T) {
	tests := []struct {
		name      string
		heroData  *[]*Hero
		role      string
		isStadium bool
		expected  []*Hero
	}{
		{"all roles", &testHeroes, "all", false, testHeroes},
		{"role specific", &testHeroes, "tank", false, []*Hero{hero5, hero6}},
		{"stadium", &testHeroes, "all", true, []*Hero{hero1, hero3, hero5}},
	}
	for _, test := range tests {
		result := getValidHeroes(test.heroData, test.role, test.isStadium)
		if !slices.Equal(result, test.expected) {
			t.Errorf("Get valid heroes for %v: Got %v, expected %v", test.name, result, test.expected)
		}
	}
}
