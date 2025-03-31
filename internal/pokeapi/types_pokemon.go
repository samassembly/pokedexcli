package pokeapi

//Pokemon -
type PokemonInfo struct {
	Abilities 			[]Ability_model 	`json:"abilities"`
	Base_experience 	int 		`json:"base_experience"`
	Cries 				struct {
		Latest 			string 		`json:"latest"`
		Legacy 			string 		`json:"legacy"`
	} 								`json:"cries"`
	Forms 				[]Form 		`json:"forms"`
	Game_indices 		[]any	 	`json:"game_indices"`
	Height 				int			`json:"height"`
	Held_items 			[]any 		`json:"held_items"`
	Id 					int 		`json:"id"`
	Is_default			bool		`json:"is_default"`
	Location_area_encounters string `json:"location_area_encounter"`
	Moves				[]MoveDetail	`json:"moves"`
	Name				string		`json:"name"`
	Order				int			`json:"order"`
	Past_abilities		[]any		`json:"past_abilities"`
	Past_types			[]any		`json:"past_types"`
	Species				struct {
		Name			string		`json:"name"`
		URL				string		`json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault       *string `json:"back_default"`
		BackFemale        *string `json:"back_female"`
		BackShiny         *string `json:"back_shiny"`
		BackShinyFemale   *string `json:"back_shiny_female"`
		FrontDefault      *string `json:"front_default"`
		FrontFemale       *string `json:"front_female"`
		FrontShiny        *string `json:"front_shiny"`
		FrontShinyFemale  *string `json:"front_shiny_female"`
		Other             Other   `json:"other"`
		Versions          Versions `json:"versions"`
	}`json:"sprites"`
	Stats               []Stat_model   	`json:"stats"`
	Types				[]Type_model	`json:"types"`
	Weight				int			`json:"weight"`
}

type Form struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type Move struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type MoveLearnMethod struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type VersionGroup struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type VersionGroupDetail struct {
    LevelLearnedAt   int             `json:"level_learned_at"`
    MoveLearnMethod  MoveLearnMethod `json:"move_learn_method"`
    Order            *int            `json:"order"`
    VersionGroup     VersionGroup    `json:"version_group"`
}

type MoveDetail struct {
    Move                Move                `json:"move"`
    VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type Type_model struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"type"`
}

type Stat_model struct {
	Base_stat int `json:"base_stat"`
	Effort int `json:"effort"`
	Stat struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"stat"`
}

type Ability_model struct {
	Ability struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"ability"`
	Is_hidden bool `json:"is_hidden"`
	Slot      int  `json:"slot"`
}

type Other struct {
    DreamWorld       DreamWorld `json:"dream_world"`
    Home             Home       `json:"home"`
    OfficialArtwork  OfficialArtwork `json:"official-artwork"`
    Showdown         Showdown   `json:"showdown"`
}

type DreamWorld struct {
    FrontDefault *string `json:"front_default"`
    FrontFemale  *string `json:"front_female"`
}

type Home struct {
    FrontDefault      *string `json:"front_default"`
    FrontFemale       *string `json:"front_female"`
    FrontShiny        *string `json:"front_shiny"`
    FrontShinyFemale  *string `json:"front_shiny_female"`
}

type OfficialArtwork struct {
    FrontDefault *string `json:"front_default"`
    FrontShiny   *string `json:"front_shiny"`
}

type Showdown struct {
    BackDefault       *string `json:"back_default"`
    BackFemale        *string `json:"back_female"`
    BackShiny         *string `json:"back_shiny"`
    BackShinyFemale   *string `json:"back_shiny_female"`
    FrontDefault      *string `json:"front_default"`
    FrontFemale       *string `json:"front_female"`
    FrontShiny        *string `json:"front_shiny"`
    FrontShinyFemale  *string `json:"front_shiny_female"`
}

type Versions struct {
    GenerationI   GenerationI   `json:"generation-i"`
    GenerationII  GenerationII  `json:"generation-ii"`
    GenerationIII GenerationIII `json:"generation-iii"`
    GenerationIV  GenerationIV  `json:"generation-iv"`
    GenerationV   GenerationV   `json:"generation-v"`
    GenerationVI  GenerationVI  `json:"generation-vi"`
    GenerationVII GenerationVII `json:"generation-vii"`
    GenerationVIII GenerationVIII `json:"generation-viii"`
}

type GenerationI struct {
    RedBlue RedBlue `json:"red-blue"`
    Yellow  Yellow  `json:"yellow"`
}

type RedBlue struct {
    BackDefault      *string `json:"back_default"`
    BackGray         *string `json:"back_gray"`
    BackTransparent  *string `json:"back_transparent"`
    FrontDefault     *string `json:"front_default"`
    FrontGray        *string `json:"front_gray"`
    FrontTransparent *string `json:"front_transparent"`
}

type Yellow struct {
    BackDefault      *string `json:"back_default"`
    BackGray         *string `json:"back_gray"`
    BackTransparent  *string `json:"back_transparent"`
    FrontDefault     *string `json:"front_default"`
    FrontGray        *string `json:"front_gray"`
    FrontTransparent *string `json:"front_transparent"`
}

type GenerationII struct {
    Crystal Crystal `json:"crystal"`
    Gold    Gold    `json:"gold"`
    Silver  Silver  `json:"silver"`
}

type Crystal struct {
    BackDefault           *string `json:"back_default"`
    BackShiny             *string `json:"back_shiny"`
    BackShinyTransparent  *string `json:"back_shiny_transparent"`
    BackTransparent       *string `json:"back_transparent"`
    FrontDefault          *string `json:"front_default"`
    FrontShiny            *string `json:"front_shiny"`
    FrontShinyTransparent *string `json:"front_shiny_transparent"`
    FrontTransparent      *string `json:"front_transparent"`
}

type Gold struct {
    BackDefault      *string `json:"back_default"`
    BackShiny        *string `json:"back_shiny"`
    FrontDefault     *string `json:"front_default"`
    FrontShiny       *string `json:"front_shiny"`
    FrontTransparent *string `json:"front_transparent"`
}

type Silver struct {
    BackDefault      *string `json:"back_default"`
    BackShiny        *string `json:"back_shiny"`
    FrontDefault     *string `json:"front_default"`
    FrontShiny       *string `json:"front_shiny"`
    FrontTransparent *string `json:"front_transparent"`
}

type GenerationIII struct {
    Emerald          Emerald          `json:"emerald"`
    FireredLeafgreen FireredLeafgreen `json:"firered-leafgreen"`
    RubySapphire     RubySapphire     `json:"ruby-sapphire"`
}

type Emerald struct {
    FrontDefault *string `json:"front_default"`
    FrontShiny   *string `json:"front_shiny"`
}

type FireredLeafgreen struct {
    BackDefault  *string `json:"back_default"`
    BackShiny    *string `json:"back_shiny"`
    FrontDefault *string `json:"front_default"`
    FrontShiny   *string `json:"front_shiny"`
}

type RubySapphire struct {
    BackDefault  *string `json:"back_default"`
    BackShiny    *string `json:"back_shiny"`
    FrontDefault *string `json:"front_default"`
    FrontShiny   *string `json:"front_shiny"`
}

type GenerationIV struct {
    DiamondPearl        DiamondPearl        `json:"diamond-pearl"`
    HeartgoldSoulsilver HeartgoldSoulsilver `json:"heartgold-soulsilver"`
    Platinum            Platinum            `json:"platinum"`
}

type DiamondPearl struct {
    BackDefault       *string `json:"back_default"`
    BackFemale        *string `json:"back_female"`
    BackShiny         *string `json:"back_shiny"`
    BackShinyFemale   *string `json:"back_shiny_female"`
    FrontDefault      *string `json:"front_default"`
    FrontFemale       *string `json:"front_female"`
    FrontShiny        *string `json:"front_shiny"`
    FrontShinyFemale  *string `json:"front_shiny_female"`
}

type HeartgoldSoulsilver struct {
    BackDefault       *string `json:"back_default"`
    BackFemale        *string `json:"back_female"`
    BackShiny         *string `json:"back_shiny"`
    BackShinyFemale   *string `json:"back_shiny_female"`
    FrontDefault      *string `json:"front_default"`
    FrontFemale       *string `json:"front_female"`
    FrontShiny        *string `json:"front_shiny"`
    FrontShinyFemale  *string `json:"front_shiny_female"`
}

type Platinum struct {
    BackDefault       *string `json:"back_default"`
    BackFemale        *string `json:"back_female"`
    BackShiny         *string `json:"back_shiny"`
    BackShinyFemale   *string `json:"back_shiny_female"`
    FrontDefault      *string `json:"front_default"`
    FrontFemale       *string `json:"front_female"`
    FrontShiny        *string `json:"front_shiny"`
    FrontShinyFemale  *string `json:"front_shiny_female"`
}

type GenerationV struct {
    BlackWhite BlackWhite `json:"black-white"`
}

type BlackWhite struct {
    Animated         Animated `json:"animated"`
    BackDefault      *string  `json:"back_default"`
    BackFemale       *string  `json:"back_female"`
    BackShiny        *string  `json:"back_shiny"`
    BackShinyFemale  *string  `json:"back_shiny_female"`
    FrontDefault     *string  `json:"front_default"`
    FrontFemale      *string  `json:"front_female"`
    FrontShiny       *string  `json:"front_shiny"`
    FrontShinyFemale *string  `json:"front_shiny_female"`
}

type Animated struct {
    BackDefault      *string `json:"back_default"`
    BackFemale       *string `json:"back_female"`
    BackShiny        *string `json:"back_shiny"`
    BackShinyFemale  *string `json:"back_shiny_female"`
    FrontDefault     *string `json:"front_default"`
    FrontFemale      *string `json:"front_female"`
    FrontShiny       *string `json:"front_shiny"`
    FrontShinyFemale *string `json:"front_shiny_female"`
}

type GenerationVI struct {
    OmegarubyAlphasapphire OmegarubyAlphasapphire `json:"omegaruby-alphasapphire"`
    XY                    XY                    `json:"x-y"`
}

type OmegarubyAlphasapphire struct {
    FrontDefault     *string `json:"front_default"`
    FrontFemale      *string `json:"front_female"`
    FrontShiny       *string `json:"front_shiny"`
    FrontShinyFemale *string `json:"front_shiny_female"`
}

type XY struct {
    FrontDefault     *string `json:"front_default"`
    FrontFemale      *string `json:"front_female"`
    FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female`
}

type GenerationVII struct {
	Icons	Icons	`json:"icons"`
	UltraSunUltraMoon UltraSunUltraMoon `json:"ultra-sun-ultra-moon"`
}

type Icons struct {
	FrontDefault *string `json:"front_default`
	FrontFemale *string `json:"front_female"`
}

type UltraSunUltraMoon struct {
	FrontDefault *string `json:"front_default"`
	FrontFemale *string `json:"front_female"`
	FrontShiny *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type GenerationVIII struct {
	Icons Icons `json:"icons"`
}