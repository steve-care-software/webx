package applications

import (
	application_affiliate "github.com/steve-care-software/affiliates/applications"
	application_blockchain "github.com/steve-care-software/fungible-unit-pow-blockchains/applications"
)

// Application represents an application
type Application interface {
	Blockchain() application_blockchain.Application
	Affiliate() application_affiliate.Application
}
