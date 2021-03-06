package main

import (
	"testing"

	"github.com/jangler/oracles-randomizer/graph"
	"github.com/jangler/oracles-randomizer/logic"
	"github.com/jangler/oracles-randomizer/rom"
)

func TestGraph(t *testing.T) {
	testSeasonsGraph(t)
	// testAgesGraph(t)
}

// check that graph logic is working as expected
func testSeasonsGraph(t *testing.T) {
	rom.Init(rom.GameSeasons)
	r := NewRoute(rom.GameSeasons)
	g := r.Graph

	checkReach(t, g,
		map[string]string{
			"feather 1": "d0 sword chest",
		}, "maku tree", false, false)

	checkReach(t, g,
		map[string]string{
			"sword 1": "d0 sword chest",
		}, "maku tree", false, true)

	checkReach(t, g,
		map[string]string{
			"sword 1":          "d0 sword chest",
			"ember tree seeds": "horon village seed tree",
			"satchel 1":        "maku tree",
			"member's card":    "d0 rupee chest",
		}, "member's shop 1", false, true)

	checkReach(t, g,
		map[string]string{
			"sword 1":  "d0 sword chest",
			"bracelet": "maku tree",
		}, "floodgate keeper's house", false, false)

	checkReach(t, g,
		map[string]string{
			"bracelet":         "d0 sword chest",
			"spring":           "horon village SW chest",
			"flippers":         "maku tree",
			"satchel 1":        "natzu region, across water",
			"ember tree seeds": "horon village seed tree",

			"woods of winter default summer": "",
			"woods of winter default winter": "start",
		}, "holly's house", false, false)

	// check normal vs hard
	checkReach(t, g,
		map[string]string{
			"sword 1":            "d0 sword chest",
			"satchel 1":          "d0 rupee chest",
			"feather 1":          "maku tree",
			"feather 2":          "shop, 150 rupees",
			"pegasus tree seeds": "horon village seed tree",

			"north horon default winter": "",
			"north horon default summer": "start",
		}, "village portal", false, false)
	checkReach(t, g,
		map[string]string{
			"sword 1":            "d0 sword chest",
			"satchel 1":          "d0 rupee chest",
			"feather 1":          "maku tree",
			"feather 2":          "shop, 150 rupees",
			"pegasus tree seeds": "horon village seed tree",

			"north horon default winter": "",
			"north horon default summer": "start",
		}, "village portal", true, true)

	// make sure that all slots in the game are reachable, given vanilla
	// progression.
	for slotName, _ := range rom.ItemSlots {
		r := NewRoute(rom.GameSeasons)
		g := r.Graph
		checkReach(t, g, map[string]string{
			"sword 1":            "d0 sword chest",
			"gnarled key":        "maku tree",
			"bombs, 10":          "d1 railway chest",
			"satchel 1":          "d1 basement",
			"ember tree seeds":   "horon village seed tree",
			"boomerang 1":        "subrosian dance hall",
			"winter":             "tower of winter",
			"shovel":             "holly's house",
			"mystery tree seeds": "woods of winter seed tree",
			"bracelet":           "d2 moblin chest",
			"ricky's flute":      "shop, 150 rupees",
			"pegasus tree seeds": "spool swamp seed tree",
			"floodgate key":      "floodgate keeper's house",
			"square jewel":       "spool swamp cave",
			"star ore":           "subrosia seaside",
			"ribbon":             "subrosia market, 1st item",
			"summer":             "tower of summer",
			"feather 1":          "d3 moldorm chest", // vanilla unsafe
			"master's plaque":    "master diver's challenge",
			"flippers":           "master diver's reward",
			"spring":             "tower of spring",
			"spring banana":      "spring banana tree",
			"dragon key":         "goron mountain, across pits",
			"pyramid jewel":      "diving spot outside D4",
			"slingshot 1":        "d4 north of entrance", // vanilla unsafe
			"autumn":             "tower of autumn",
			"magnet gloves":      "d5 spiral chest", // vanilla unsafe
			"round jewel":        "old man in treehouse",
			"x-shaped jewel":     "black beast's chest",
			"boomerang 2":        "d6 armos hall",
			"rusty bell":         "samasa desert pit",
			"feather 2":          "d7 spike chest",
			"slingshot 2":        "d8 armos chest",

			// vanilla non-required
			"member's card": "subrosia market, 5th item",
			"red ore":       "subrosia village chest",
			"blue ore":      "subrosian wilds chest",
			"hard ore":      "great furnace",
		}, slotName, false, true)
	}
}

// check that graph logic is working as expected
func testAgesGraph(t *testing.T) {
	rom.Init(rom.GameAges)
	r := NewRoute(rom.GameAges)
	g := r.Graph

	checkReach(t, g, map[string]string{
		"sword 1":          "starting chest",
		"shovel":           "black tower worker",
		"satchel 1":        "maku tree",
		"ember tree seeds": "south lynna tree",
		"graveyard key":    "grave under tree",
	}, "enter d1", false, true)

	checkReach(t, g, map[string]string{
		"harp 1":     "starting chest",
		"harp 2":     "nayru's house",
		"bracelet 1": "black tower worker",
	}, "enter d2", false, true)

	checkReach(t, g, map[string]string{
		"dimitri's flute": "starting chest",
	}, "enter d3", false, true)

	checkReach(t, g, map[string]string{
		"harp 1":     "starting chest",
		"harp 2":     "nayru's house",
		"harp 3":     "black tower worker",
		"flippers 1": "lynna city chest",
		"sword 1":    "fairies' woods chest",
		"tuni nut":   "tokkey's composition",
	}, "symmetry past", false, true)

	checkReach(t, g, map[string]string{
		"sword 1":            "starting chest",
		"satchel 1":          "black tower worker",
		"ember tree seeds":   "south lynna tree",
		"graveyard key":      "grave under tree",
		"switch hook 1":      "lynna city chest",
		"feather":            "nayru's house",
		"bomb flower":        "d1 east terrace",
		"bracelet 1":         "d1 crystal room",
		"flippers 1":         "d1 west terrace",
		"harp 1":             "d1 pot chest",
		"harp 2":             "d1 crossroads",
		"pegasus tree seeds": "rolling ridge west tree",
		"crown key":          "under moblin keep",
	}, "enter d5", false, true)

	checkReach(t, g, map[string]string{
		"harp 1":      "starting chest",
		"harp 2":      "nayru's house",
		"flippers 1":  "black tower worker",
		"flippers 2":  "fairies' woods chest",
		"feather":     "lynna city chest",
		"mermaid key": "hidden tokay cave",
	}, "enter d6 past", false, true)

	checkReach(t, g, map[string]string{
		"harp 1":          "starting chest",
		"harp 2":          "nayru's house",
		"flippers 1":      "black tower worker",
		"flippers 2":      "fairies' woods chest",
		"feather":         "lynna city chest",
		"old mermaid key": "hidden tokay cave",
	}, "enter d6 present", false, true)

	checkReach(t, g, map[string]string{
		"harp 1":           "starting chest",
		"harp 2":           "nayru's house",
		"harp 3":           "black tower worker",
		"flippers 1":       "fairies' woods chest",
		"flippers 2":       "lynna city chest",
		"switch hook 1":    "hidden tokay cave",
		"sword 1":          "zora village present",
		"satchel 1":        "zora palace chest",
		"ember tree seeds": "zora village tree",
		"fairy powder":     "grave under tree",
		"graveyard key":    "under crescent island",
	}, "enter d7", false, true)

	checkReach(t, g, map[string]string{
		"sword 1":       "starting chest",
		"flippers 1":    "nayru's house",
		"flippers 2":    "black tower worker",
		"tokay eyeball": "hidden tokay cave",
		"feather":       "under crescent island",
		"bombs, 10":     "tokay crystal cave",
		"bracelet 1":    "ambi's palace chest",
		"cane":          "tokay bomb cave",
	}, "enter d8", false, true)

	// make sure that all slots in the game are reachable, given vanilla
	// progression.
	for slotName, _ := range rom.ItemSlots {
		r := NewRoute(rom.GameAges)
		g := r.Graph
		checkReach(t, g, map[string]string{
			"sword 1":            "starting chest",
			"shovel":             "black tower worker",
			"satchel 1":          "maku tree",
			"ember tree seeds":   "south lynna tree",
			"graveyard key":      "grave under tree",
			"bracelet 1":         "d1 basement",
			"harp 1":             "nayru's house",
			"mystery tree seeds": "deku forest tree",
			"bombs, 10":          "deku forest soldier",
			"feather":            "d2 thwomp tunnel",
			"flippers 1":         "cheval's test",
			"cheval rope":        "cheval's invention",
			"ricky's gloves":     "south shore dirt",
			"island chart":       "balloon guy's gift",
			"scent seedling":     "wild tokay game",
			"scent tree seeds":   "crescent island tree",
			"seed shooter":       "d3 pols voice chest",
			"moosh's flute":      "shop, 150 rupees",
			"tuni nut":           "symmetry city brother",
			"harp 2":             "tokkey's composition",
			"switch hook 1":      "d4 small floor puzzle",
			"pegasus tree seeds": "rolling ridge west tree",
			"bomb flower":        "defeat great moblin",
			"crown key":          "goron elder",
			"cane":               "d5 blue peg chest", // vanilla unsafe
			"brother emblem":     "goron dance present",
			"rock brisket":       "target carts 1",
			"goron vase":         "trade rock brisket",
			"goronade":           "trade goron vase",
			"old mermaid key":    "big bang game",
			"lava juice":         "goron shooting gallery",
			"goron letter":       "trade lava juice",
			"mermaid key":        "goron dance, with letter",
			"flippers 2":         "d6 present vire chest",
			"harp 3":             "rescue nayru",
			"library key":        "king zora",
			"book of seals":      "library present",
			"fairy powder":       "library past",
			"switch hook 2":      "d7 miniboss chest",
			"d7 boss key":        "d7 post-hallway chest",
			"zora scale":         "zora's reward",
			"tokay eyeball":      "piratian captain",
			"bracelet 2":         "d8 floor puzzle",
		}, slotName, false, true)
	}
}

func BenchmarkGraphExplore(b *testing.B) {
	// init graph
	r := NewRoute(rom.GameSeasons)
	b.ResetTimer()

	// explore all items from the d0 sword chest
	for name := range logic.SeasonsExtraItems() {
		r.Graph.Explore(make(map[*graph.Node]bool), false, r.Graph[name])
	}
}

// helper function for testing whether a node is reachable given a certain
// slotting
func checkReach(t *testing.T, g graph.Graph, parents map[string]string,
	target string, hard, expect bool) {
	t.Helper()

	// add parents at the start of the function, and remove them at the end. if
	// a parent is blank, remove it at the start and add it at the end (only
	// useful for default seasons).
	for child, parent := range parents {
		if parent == "" {
			g[child].ClearParents()
		} else {
			g[child].AddParents(g[parent])
		}
	}
	defer func() {
		for child, parent := range parents {
			if parent == "" {
				g[child].AddParents(g["start"])
			} else {
				g[child].ClearParents()
			}
		}
	}()
	g.ExploreFromStart(hard)

	if (g[target].GetMark(g[target], hard) == graph.MarkTrue) != expect {
		if expect {
			t.Errorf("expected to reach %s, but could not", target)
		} else {
			t.Errorf("expected not to reach %s, but could", target)
		}
	}
}
