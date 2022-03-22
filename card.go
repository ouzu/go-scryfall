package scryfall

import (
	"time"

	"github.com/google/uuid"
	"ouzu.tech/go-scryfall.git/BorderColor"
	"ouzu.tech/go-scryfall.git/Color"
	"ouzu.tech/go-scryfall.git/Component"
	"ouzu.tech/go-scryfall.git/Finish"
	"ouzu.tech/go-scryfall.git/Game"
	"ouzu.tech/go-scryfall.git/ImageStatus"
	"ouzu.tech/go-scryfall.git/Layout"
	"ouzu.tech/go-scryfall.git/Legality"
	"ouzu.tech/go-scryfall.git/Rarity"
	"ouzu.tech/go-scryfall.git/SecurityStamp"
)

type Date time.Time

const dateLayout = `"2006-01-02"`

func (d *Date) UnmarshalJSON(bytes []byte) error {
	date, err := time.Parse(dateLayout, string(bytes))
	*d = Date(date)
	return err
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(*d).Format(dateLayout)), nil
}

type ImageURIs struct {
	Small      string `json:"small"`
	Normal     string `json:"normal"`
	Large      string `json:"large"`
	PNG        string `json:"png"`
	ArtCrop    string `json:"art_crop"`
	BorderCrop string `json:"border_crop"`
}

// https://scryfall.com/docs/api/cards#card-face-objects

type CardFace struct {
	Artist          string       `json:"artist"`
	CMC             float64      `json:"cmc"`
	ColorIndicator  Color.Colors `json:"color_indicator"`
	Colors          Color.Colors `json:"colors"`
	FlavorText      string       `json:"flavor_text"`
	IllustrationID  uuid.UUID    `json:"illustration_id"`
	ImageURIs       ImageURIs    `json:"image_uris"`
	Layout          string       `json:"layout"`
	Loyalty         string       `json:"loyalty"`
	ManaCost        string       `json:"mana_cost"`
	Name            string       `json:"name"`
	Object          string       `json:"object"`
	OracleID        uuid.UUID    `json:"oracle_id"`
	OracleText      string       `json:"oracle_text"`
	Power           string       `json:"power"`
	PrintedName     string       `json:"printed_name"`
	PrintedText     string       `json:"printed_text"`
	PrintedTypeLine string       `json:"printed_type_line"`
	Toughness       string       `json:"toughness"`
	TypeLine        string       `json:"type_line"`
	Watermark       string       `json:"watermark"`
}

// https://scryfall.com/docs/api/cards#related-card-objects

type RelatedCard struct {
	ID        uuid.UUID           `json:"id"`
	Object    string              `json:"object"`
	Component Component.Component `json:"component"`
	Name      string              `json:"name"`
	TypeLine  string              `json:"type_line"`
	URI       string              `json:"uri"`
}

// https://scryfall.com/docs/api/cards

type Card struct {
	// Core Card Fields
	ArenaID           int       `json:"arena_id"`
	ID                uuid.UUID `json:"id"`
	Lang              string    `json:"lang"`
	MtgoID            int       `json:"mtgo_id"`
	MtgoFoilID        int       `json:"mtgo_foil_id"`
	MultiverseIDs     []int     `json:"multiverse_ids"`
	TcgPlayerID       int       `json:"tcgplayer_id"`
	TcgPlayerEtchedID int       `json:"tcgplayer_etched_id"`
	CardmarketID      int       `json:"cardmarket_id"`
	Object            string    `json:"object"`
	OracleID          uuid.UUID `json:"oracle_id"`
	PrintsSearchURI   string    `json:"prints_search_uri"`
	RulingsURI        string    `json:"rulings_uri"`
	ScryfallURI       string    `json:"scryfall_uri"`
	URI               string    `json:"uri"`

	// Gameplay Fields
	AllParts       []interface{}       `json:"all_parts"`
	CardFaces      []CardFace          `json:"card_faces"`
	CMC            float64             `json:"cmc"`
	ColorIdentity  Color.Colors        `json:"color_identity"`
	ColorIndicator Color.Colors        `json:"color_indicator"`
	Colors         Color.Colors        `json:"colors"`
	EdhrecRank     int                 `json:"edhrec_rank"`
	HandModifier   string              `json:"hand_modifier"`
	Keywords       []string            `json:"keywords"`
	Layout         Layout.Layout       `json:"layout"`
	Legalities     Legality.Legalities `json:"legalities"`
	LifeModifier   string              `json:"life_modifier"`
	Loyalty        string              `json:"loyalty"`
	ManaCost       string              `json:"mana_cost"`
	Name           string              `json:"name"`
	OracleText     string              `json:"oracle_text"`
	Oversized      bool                `json:"oversized"`
	Power          string              `json:"power"`
	OroducedMana   Color.Colors        `json:"produced_mana"`
	Reversed       bool                `json:"reversed"`
	Toughness      string              `json:"toughness"`
	TypeLine       string              `json:"type_line"`

	// Print Fields
	Artist          string                      `json:"artist"`
	Booster         bool                        `json:"booster"`
	BorderColor     BorderColor.BorderColor     `json:"border_color"`
	CardBackID      uuid.UUID                   `json:"card_back_id"`
	CollectorNumber string                      `json:"collector_number"`
	ContentWarning  bool                        `json:"content_warning"`
	Digital         bool                        `json:"digital"`
	Finishes        []Finish.Finish             `json:"finishes"`
	FlavorName      string                      `json:"flavor_name"`
	FlavorText      string                      `json:"flavor_text"`
	FrameEffects    interface{}                 `json:"frame_effects"`
	Frame           Layout.Layout               `json:"frame"`
	FullArt         bool                        `json:"full_art"`
	Games           []Game.Game                 `json:"games"`
	HighresImage    bool                        `json:"highres_image"`
	IllustrationID  uuid.UUID                   `json:"illustration_id"`
	ImageStatus     ImageStatus.ImageStatus     `json:"image_status"`
	ImageURIs       ImageURIs                   `json:"image_uris"`
	Prices          interface{}                 `json:"prices"`
	PrintedName     string                      `json:"printed_name"`
	PrintedText     string                      `json:"printed_text"`
	PrintedTypeLine string                      `json:"printed_type_line"`
	Promo           bool                        `json:"promo"`
	PromoTypes      []string                    `json:"promo_types"`
	PurchaseURIs    interface{}                 `json:"purchase_uris"`
	Rarity          Rarity.Rarity               `json:"rarity"`
	RelatedURIs     interface{}                 `json:"related_uris"`
	ReleasedAt      Date                        `json:"released_at"`
	Repring         bool                        `json:"reprint"`
	ScryfallSetURI  string                      `json:"scryfall_set_uri"`
	SetName         string                      `json:"set_name"`
	SetSearchURI    string                      `json:"set_search_uri"`
	SetType         string                      `json:"set_type"`
	SetURI          string                      `json:"set_uri"`
	Set             string                      `json:"set"`
	SetID           uuid.UUID                   `json:"set_id"`
	StorySpotlight  bool                        `json:"story_spotlight"`
	Textless        bool                        `json:"textless"`
	Variation       bool                        `json:"variation"`
	VariationOf     uuid.UUID                   `json:"variation_of"`
	SecurityStamp   SecurityStamp.SecurityStamp `json:"security_stamp"`
	Watermark       string                      `json:"watermark"`
}
