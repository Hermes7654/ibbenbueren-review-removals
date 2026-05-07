package mapsreview

const (
	OutputDir     = "output"
	ResultsJSON   = "output/places.json"
	ResultsCSV    = "output/places.csv"
	DiscoveryJSON = "output/discovery.json"
	MetadataJSON  = "output/metadata.json"
)

var IbbenbuerenPostcodes = []string{
	"48429", "48431", "48432", "48477", "48480", "48493", "48496", "48499",
	"49074", "49076", "49078", "49080", "49082", "49084", "49086", "49088",
	"49090", "49124", "49134", "49143", "49152", "49163", "49170", "49176",
	"49179", "49186", "49191", "49196", "49201", "49205", "49214", "49219",
	"49324", "49325", "49326", "49327", "49328", "49448", "49451", "49453",
	"49456", "49457", "49477", "49479", "49492", "49497", "49504", "49509",
	"49525", "49536", "49545", "49549", "49565", "49577", "49584", "49586",
	"49593", "49594", "49596", "49597", "49599", "49610", "49624", "49626",
	"49632", "49635", "49637", "49638", "49661", "49681", "49685", "49688",
	"49692", "49696", "49716", "49733", "49740", "49744", "49751", "49757",
	"49762", "49767", "49770", "49774", "49777", "49779", "49809", "49811",
	"49824", "49828", "49832", "49835", "49838", "49843", "49844", "49846",
	"49847", "48143", "48145", "48147", "48149", "48151", "48153", "48155", "48157", "48159", "48161", "48163", "48165", "48167", "48268", "48291",
}

var DefaultQueries = []string{
	// Gastro (original)
	"restaurant", "café", "imbiss", "pizzeria", "bäckerei",
	"döner", "burger", "sushi", "schnitzel", "frühstück", "brunch",
	// Bars & Nightlife
	"bar", "kneipe", "pub", "biergarten", "brauerei",
	"cocktail bar", "lounge", "weinstube",
	"club", "nachtclub", "diskothek",
	// Hotels
	"hotel",
	// Beauty & Wellness
	"friseur", "barbier", "barbershop",
	"fitnessstudio", "fitness",
	// Shopping & Daily
	"supermarkt", "metzgerei",
	"apotheke",
	// Services
	"tankstelle",
}

var IbbenbuerenPostcodeSet = func() map[string]bool {
	set := make(map[string]bool, len(IbbenbuerenPostcodes))
	for _, postcode := range IbbenbuerenPostcodes {
		set[postcode] = true
	}
	return set
}()
