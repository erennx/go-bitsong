package types

import (
	"fmt"
	btsg "github.com/bitsongofficial/go-bitsong/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
	"time"
)

type Content struct {
	Name       string `json:"name" yaml:"name"`
	Uri        string `json:"uri" yaml:"uri"`
	Metadata   string `json:"metadata" yaml:"metadata"`       // JSON.stringify()
	ContentUri string `json:"content_uri" yaml:"content_uri"` // /ipfs/QM.......

	StreamPrice   sdk.Coin `json:"stream_price" yaml:"stream_price"`
	DownloadPrice sdk.Coin `json:"download_price" yaml:"download_price"`

	RightsHolders RightsHolders `json:"rights_holders" yaml:"rights_holders"`

	Denom       string   `json:"denom" yaml:"denom"`
	Volume      sdk.Coin `json:"volume" yaml:"volume"`
	TotalSupply sdk.Coin `json:"total_supply" yaml:"total_supply"`

	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	CreatedAt time.Time      `json:"created_at" yaml:"created_at"`
}

func NewContent(name, uri, metadata, contentUri, denom string, streamPrice, downloadPrice sdk.Coin, creator sdk.AccAddress, rhs RightsHolders) Content {
	return Content{
		Name:          name,
		Uri:           uri,
		Metadata:      metadata,
		ContentUri:    contentUri,
		Denom:         denom,
		StreamPrice:   streamPrice,
		DownloadPrice: downloadPrice,
		Creator:       creator,
		RightsHolders: rhs,
		Volume:        sdk.NewCoin(btsg.BondDenom, sdk.ZeroInt()),
		TotalSupply:   sdk.NewCoin(denom, sdk.ZeroInt()),
	}
}

func (c Content) String() string {
	return fmt.Sprintf(`Name: %s
Uri: %s
Metadata: %s
ContentUri: %s
Denom: %s
Stream Price: %s
Download Price: %s
CreatedAt: %s
Creator: %s
Rights Hoders: %s`,
		c.Name, c.Uri, c.Metadata, c.ContentUri, c.Denom, c.StreamPrice, c.DownloadPrice, c.CreatedAt, c.Creator, c.RightsHolders,
	)
}

func (c Content) Equals(content Content) bool {
	return c.Name == content.Name && c.Uri == content.Uri && c.Metadata == content.Metadata &&
		c.ContentUri == content.ContentUri && c.Denom == content.Denom && c.StreamPrice == content.StreamPrice && c.DownloadPrice == content.DownloadPrice &&
		c.RightsHolders.Equals(content.RightsHolders)
}

func (c Content) Validate() error {
	if len(strings.TrimSpace(c.Name)) == 0 {
		return fmt.Errorf("name cannot be empty")
	}

	if len(c.Name) > MaxNameLength {
		return fmt.Errorf("name cannot be longer than %d characters", MaxUriLength)
	}

	if len(strings.TrimSpace(c.Uri)) == 0 {
		return fmt.Errorf("uri cannot be empty")
	}

	if len(c.Uri) > MaxUriLength {
		return fmt.Errorf("uri cannot be longer than %d characters", MaxUriLength)
	}

	if len(c.Metadata) > MaxMetadataLength {
		return fmt.Errorf("metadata cannot be longer than %d characters", MaxMetadataLength)
	}

	if len(strings.TrimSpace(c.ContentUri)) == 0 {
		return fmt.Errorf("content-uri cannot be empty")
	}

	if len(c.ContentUri) > MaxUriLength {
		return fmt.Errorf("content-uri cannot be longer than %d characters", MaxUriLength)
	}

	if err := sdk.ValidateDenom(c.Denom); err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	if c.Creator == nil {
		return fmt.Errorf("invalid creator: %s", c.Creator)
	}

	if err := c.RightsHolders.Validate(); err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	return nil
}
