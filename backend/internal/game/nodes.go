package game

type MatrixNode struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"` // TELEPORT or ENTRY
	Cost        int    `json:"cost"`
}

var EngineeringMatrix = []MatrixNode{
	// Teleport Path
	{
		ID:          "SIGNAL_BOOSTER_1",
		Name:        "Signal Booster I",
		Description: "+1 Encounter visibility",
		Path:        "TELEPORT",
		Cost:        100,
	},
	{
		ID:          "DIMENSIONAL_ANCHOR",
		Name:        "Dimensional Anchor",
		Description: "-10% Fuel consumption during Teleport",
		Path:        "TELEPORT",
		Cost:        250,
	},
	{
		ID:          "VOID_STABILIZER",
		Name:        "Void Stabilizer",
		Description: "Reduces Accident chance by 5%",
		Path:        "TELEPORT",
		Cost:        500,
	},
	{
		ID:          "FREQUENCY_TUNER",
		Name:        "Frequency Tuner",
		Description: "+5% chance for Refined signals",
		Path:        "TELEPORT",
		Cost:        750,
	},
	{
		ID:          "WORMHOLE_NAVIGATOR",
		Name:        "Wormhole Navigator",
		Description: "Unlocks Short-range Jump",
		Path:        "TELEPORT",
		Cost:        1000,
	},

	// Entry Path
	{
		ID:          "HEAT_SHIELDING_1",
		Name:        "Heat Shielding I",
		Description: "-15% O2 consumption during entry",
		Path:        "ENTRY",
		Cost:        100,
	},
	{
		ID:          "REINFORCED_HULL",
		Name:        "Reinforced Hull",
		Description: "+10% Mech HP after landing",
		Path:        "ENTRY",
		Cost:        250,
	},
	{
		ID:          "SHOCK_ABSORBERS",
		Name:        "Shock Absorbers",
		Description: "Reduces Critical Damage chance by 10%",
		Path:        "ENTRY",
		Cost:        500,
	},
	{
		ID:          "CARGO_STABILIZER",
		Name:        "Cargo Stabilizer",
		Description: "+5% chance to keep items on Accident",
		Path:        "ENTRY",
		Cost:        750,
	},
	{
		ID:          "DESCENT_THRUSTERS",
		Name:        "Descent Thrusters",
		Description: "Unlocks Precision Landing",
		Path:        "ENTRY",
		Cost:        1000,
	},
}
