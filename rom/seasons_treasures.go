package rom

func seasonsTreasure(id, subID byte, offset uint16,
	mode, param, text, sprite byte) *Treasure {
	return &Treasure{id, subID, Addr{0x15, offset}, mode, param, text, sprite}
}

var SeasonsTreasures = map[string]*Treasure{
	// equip items
	"wooden shield":   seasonsTreasure(0x01, 0x00, 0x52bd, 0x0a, 0x01, 0x1f, 0x13),
	"shield L-2":      seasonsTreasure(0x01, 0x01, 0x52c1, 0x0a, 0x02, 0x20, 0x14),
	"bombs, 10":       seasonsTreasure(0x03, 0x00, 0x52c9, 0x38, 0x10, 0x4d, 0x05),
	"sword 1":         seasonsTreasure(0x05, 0x00, 0x52d9, 0x00, 0x01, 0x1c, 0x10),
	"sword 2":         seasonsTreasure(0x05, 0x01, 0x52dd, 0x09, 0x01, 0x1c, 0x10),
	"boomerang 1":     seasonsTreasure(0x06, 0x00, 0x52f1, 0x0a, 0x01, 0x22, 0x1c),
	"boomerang 2":     seasonsTreasure(0x06, 0x01, 0x52f5, 0x38, 0x01, 0x22, 0x1c),
	"rod":             seasonsTreasure(0x07, 0x00, 0x52f9, 0x38, 0x07, 0x0a, 0x1e),
	"spring":          seasonsTreasure(0x07, 0x02, 0x5301, 0x09, 0x00, 0x0d, 0x1e),
	"summer":          seasonsTreasure(0x07, 0x03, 0x5305, 0x09, 0x01, 0x0b, 0x1e),
	"autumn":          seasonsTreasure(0x07, 0x04, 0x5309, 0x09, 0x02, 0x0c, 0x1e),
	"winter":          seasonsTreasure(0x07, 0x05, 0x530d, 0x09, 0x03, 0x0a, 0x1e),
	"magnet gloves":   seasonsTreasure(0x08, 0x00, 0x5149, 0x38, 0x00, 0x30, 0x18),
	"bombchus":        seasonsTreasure(0x0d, 0x00, 0x531d, 0x0a, 0x10, 0x32, 0x24),
	"moosh's flute":   seasonsTreasure(0x0e, 0x00, 0x5161, 0x0a, 0x0d, 0x3a, 0x4d),
	"dimitri's flute": seasonsTreasure(0x0e, 0x00, 0x5161, 0x0a, 0x0c, 0x39, 0x4c),
	"strange flute":   seasonsTreasure(0x0e, 0x00, 0x5161, 0x0a, 0x0d, 0x3b, 0x23),
	"ricky's flute":   seasonsTreasure(0x0e, 0x00, 0x5161, 0x0a, 0x0b, 0x38, 0x4b),
	"slingshot 1":     seasonsTreasure(0x13, 0x00, 0x5325, 0x38, 0x01, 0x2e, 0x21),
	"slingshot 2":     seasonsTreasure(0x13, 0x01, 0x5329, 0x38, 0x01, 0x2e, 0x21),
	"shovel":          seasonsTreasure(0x15, 0x00, 0x517d, 0x0a, 0x00, 0x25, 0x1b),
	"bracelet":        seasonsTreasure(0x16, 0x00, 0x5181, 0x38, 0x00, 0x26, 0x19),
	"feather 1":       seasonsTreasure(0x17, 0x00, 0x532d, 0x38, 0x01, 0x27, 0x16),
	"feather 2":       seasonsTreasure(0x17, 0x01, 0x5331, 0x38, 0x01, 0x27, 0x16),
	"satchel 1":       seasonsTreasure(0x19, 0x00, 0x52b5, 0x0a, 0x01, 0x2d, 0x20),
	"satchel 2":       seasonsTreasure(0x19, 0x01, 0x52b9, 0x01, 0x01, 0x2d, 0x20),
	"fool's ore":      seasonsTreasure(0x1e, 0x00, 0x51a1, 0x00, 0x00, 0x36, 0x4a),

	// not used because of progressive item upgrades
	// "sword L-2":       &Treasure{0x05, 0x01, 0x52dd, 0x09, 0x02, 0x1d, 0x11},
	// "boomerang L-2":   &Treasure{0x06, 0x01, 0x52f5, 0x38, 0x02, 0x23, 0x1d},
	// "slingshot L-2":   &Treasure{0x13, 0x01, 0x5329, 0x38, 0x02, 0x2f, 0x22},
	// "feather L-2":     &Treasure{0x17, 0x01, 0x5331, 0x38, 0x02, 0x28, 0x17},
	// "satchel 2":       &Treasure{0x19, 0x01, 0x52b9, 0x01, 0x00, 0x46, 0x20},

	// non-inventory items
	"rupees, 1":        seasonsTreasure(0x28, 0x00, 0x5355, 0x38, 0x01, 0x01, 0x28),
	"rupees, 5":        seasonsTreasure(0x28, 0x01, 0x5359, 0x38, 0x03, 0x02, 0x29),
	"rupees, 10":       seasonsTreasure(0x28, 0x02, 0x535d, 0x38, 0x04, 0x03, 0x2a),
	"rupees, 20":       seasonsTreasure(0x28, 0x03, 0x5361, 0x38, 0x05, 0x04, 0x2b),
	"rupees, 30":       seasonsTreasure(0x28, 0x04, 0x5365, 0x38, 0x07, 0x05, 0x2b),
	"rupees, 50":       seasonsTreasure(0x28, 0x05, 0x5369, 0x38, 0x0b, 0x06, 0x2c),
	"rupees, 100":      seasonsTreasure(0x28, 0x06, 0x536d, 0x38, 0x0c, 0x07, 0x2d),
	"heart container":  seasonsTreasure(0x2a, 0x00, 0x5399, 0x1a, 0x04, 0x16, 0x3b),
	"piece of heart":   seasonsTreasure(0x2b, 0x01, 0x5391, 0x38, 0x01, 0x17, 0x3a),
	"rare peach stone": seasonsTreasure(0x2b, 0x02, 0x5395, 0x02, 0x01, 0x17, 0x4e),

	// rings
	"green joy ring": seasonsTreasure(0x2d, 0x01, 0x53bd, 0x29, 0x27, 0x54, 0x0e),
	"discovery ring": seasonsTreasure(0x2d, 0x04, 0x53c9, 0x38, 0x28, 0x54, 0x0e),
	"moblin ring":    seasonsTreasure(0x2d, 0x05, 0x53cd, 0x38, 0x2b, 0x54, 0x0e),
	"steadfast ring": seasonsTreasure(0x2d, 0x06, 0x53d1, 0x38, 0x10, 0x54, 0x0e),
	"blast ring":     seasonsTreasure(0x2d, 0x07, 0x53d5, 0x38, 0x0c, 0x54, 0x0e),
	"rang ring L-1":  seasonsTreasure(0x2d, 0x08, 0x53d9, 0x38, 0x0d, 0x54, 0x0e),
	"octo ring":      seasonsTreasure(0x2d, 0x09, 0x53dd, 0x38, 0x2a, 0x54, 0x0e),
	"quicksand ring": seasonsTreasure(0x2d, 0x0a, 0x53e1, 0x38, 0x23, 0x54, 0x0e),
	"armor ring L-2": seasonsTreasure(0x2d, 0x0b, 0x53e5, 0x38, 0x05, 0x54, 0x0e),
	"power ring L-1": seasonsTreasure(0x2d, 0x0e, 0x53f1, 0x38, 0x01, 0x54, 0x0e),
	"subrosian ring": seasonsTreasure(0x2d, 0x10, 0x53f9, 0x38, 0x2d, 0x54, 0x0e),

	// dungeon items
	"small key":   seasonsTreasure(0x30, 0x03, 0x5409, 0x38, 0x01, 0x1a, 0x42),
	"boss key":    seasonsTreasure(0x31, 0x03, 0x5419, 0x38, 0x00, 0x1b, 0x43),
	"d5 boss key": seasonsTreasure(0x31, 0x00, 0x540d, 0x19, 0x00, 0x1b, 0x43),
	"d4 boss key": seasonsTreasure(0x31, 0x02, 0x5415, 0x49, 0x00, 0x1b, 0x43),
	"d1 boss key": seasonsTreasure(0x31, 0x03, 0x5419, 0x38, 0x00, 0x1b, 0x43),
	"d2 boss key": seasonsTreasure(0x31, 0x03, 0x5419, 0x38, 0x00, 0x1b, 0x43),
	"d3 boss key": seasonsTreasure(0x31, 0x03, 0x5419, 0x38, 0x00, 0x1b, 0x43),
	"d6 boss key": seasonsTreasure(0x31, 0x03, 0x5419, 0x38, 0x00, 0x1b, 0x43),
	"d7 boss key": seasonsTreasure(0x31, 0x03, 0x5419, 0x38, 0x00, 0x1b, 0x43),
	"d8 boss key": seasonsTreasure(0x31, 0x03, 0x5419, 0x38, 0x00, 0x1b, 0x43),
	"compass":     seasonsTreasure(0x32, 0x02, 0x5425, 0x68, 0x00, 0x19, 0x41),
	"dungeon map": seasonsTreasure(0x33, 0x02, 0x5431, 0x68, 0x00, 0x18, 0x40),

	// collection items
	"ring box L-1":    seasonsTreasure(0x2c, 0x00, 0x53a5, 0x02, 0x01, 0x57, 0x33),
	"ring box L-2":    seasonsTreasure(0x2c, 0x01, 0x53a9, 0x02, 0x02, 0x34, 0x34),
	"flippers":        seasonsTreasure(0x2e, 0x00, 0x51e1, 0x02, 0x00, 0x31, 0x31),
	"gasha seed":      seasonsTreasure(0x34, 0x01, 0x5341, 0x38, 0x01, 0x4b, 0x0d),
	"gnarled key":     seasonsTreasure(0x42, 0x00, 0x5465, 0x29, 0x00, 0x42, 0x44),
	"floodgate key":   seasonsTreasure(0x43, 0x00, 0x5235, 0x09, 0x00, 0x43, 0x45),
	"dragon key":      seasonsTreasure(0x44, 0x00, 0x5239, 0x09, 0x00, 0x44, 0x46),
	"star ore":        seasonsTreasure(0x45, 0x00, 0x523d, 0x5a, 0x00, 0x40, 0x57),
	"ribbon":          seasonsTreasure(0x46, 0x00, 0x5241, 0x0a, 0x00, 0x41, 0x4f),
	"spring banana":   seasonsTreasure(0x47, 0x00, 0x5245, 0x0a, 0x00, 0x66, 0x54),
	"ricky's gloves":  seasonsTreasure(0x48, 0x00, 0x5249, 0x09, 0x01, 0x67, 0x55),
	"rusty bell":      seasonsTreasure(0x4a, 0x00, 0x546d, 0x0a, 0x00, 0x55, 0x5b),
	"treasure map":    seasonsTreasure(0x4b, 0x00, 0x5255, 0x0a, 0x00, 0x6c, 0x49),
	"round jewel":     seasonsTreasure(0x4c, 0x00, 0x5259, 0x0a, 0x00, 0x47, 0x36),
	"pyramid jewel":   seasonsTreasure(0x4d, 0x00, 0x5479, 0x08, 0x00, 0x4a, 0x37),
	"square jewel":    seasonsTreasure(0x4e, 0x00, 0x5261, 0x38, 0x00, 0x48, 0x38),
	"x-shaped jewel":  seasonsTreasure(0x4f, 0x00, 0x5265, 0x38, 0x00, 0x49, 0x39),
	"red ore":         seasonsTreasure(0x50, 0x00, 0x5269, 0x38, 0x00, 0x3f, 0x59),
	"blue ore":        seasonsTreasure(0x51, 0x00, 0x526d, 0x38, 0x00, 0x3e, 0x58),
	"hard ore":        seasonsTreasure(0x52, 0x00, 0x5271, 0x0a, 0x00, 0x3d, 0x5a),
	"member's card":   seasonsTreasure(0x53, 0x00, 0x5275, 0x0a, 0x00, 0x45, 0x48),
	"master's plaque": seasonsTreasure(0x54, 0x00, 0x5279, 0x38, 0x00, 0x70, 0x26),

	// not real treasures, just placeholders for seeds in trees
	"ember tree seeds":   &Treasure{id: 0x00},
	"scent tree seeds":   &Treasure{id: 0x01},
	"pegasus tree seeds": &Treasure{id: 0x02},
	"gale tree seeds":    &Treasure{id: 0x03},
	"mystery tree seeds": &Treasure{id: 0x04},
}
