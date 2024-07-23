package applications

import (
	stencil_applications "github.com/steve-care-software/webx/engine/stencils/applications"
	"github.com/steve-care-software/webx/engine/vms/applications"
	instructions "github.com/steve-care-software/webx/engine/vms/applications/instructions"
	assignments "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments"
	assignables "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables"
	assignables_bytes "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/compilers"
	assignables_constants "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/constants"
	cryptography "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography"
	cryptography_decrypts "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/decrypts"
	cryptography_encrypts "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/encrypts"
	cryptography_keys "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/keys"
	cryptography_keys_encryptions "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/keys/encryptions"
	cryptography_keys_encryptions_decrypts "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	cryptography_keys_encryptions_encrypts "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	cryptography_keys_signatures "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/keys/signatures"
	cryptography_keys_signatures_signs "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	cryptography_keys_signatures_signs_creates "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	cryptography_keys_signatures_signs_validates "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	cryptography_keys_signatures_votes "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	cryptography_keys_signatures_votes_creates "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	cryptography_keys_signatures_votes_validates "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	executables "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/excutables"
	assignables_executions "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/executions"
	assignables_executions_executes "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/executions/executes"
	assignables_executions_inits "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/executions/inits"
	assignables_executions_retrieves "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/executions/retrieves"
	assignables_lists "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/lists"
	assignables_lists_fetches "github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/lists/fetches"
	instructions_executions "github.com/steve-care-software/webx/engine/vms/applications/instructions/executions"
	instructions_executions_merges "github.com/steve-care-software/webx/engine/vms/applications/instructions/executions/merges"
	instructions_lists "github.com/steve-care-software/webx/engine/vms/applications/instructions/lists"
	instructions_lists_deletes "github.com/steve-care-software/webx/engine/vms/applications/instructions/lists/deletes"
	instructions_lists_inserts "github.com/steve-care-software/webx/engine/vms/applications/instructions/lists/inserts"
	"github.com/steve-care-software/webx/engine/vms/infrastructure/edwards25519"
	"github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances"
)

type factory struct {
	localApplicationBuilder  stencil_applications.LocalBuilder
	remoteApplicationBuilder stencil_applications.RemoteBuilder
}

func createFactory(
	localApplicationBuilder stencil_applications.LocalBuilder,
	remoteApplicationBuilder stencil_applications.RemoteBuilder,
) applications.Factory {
	out := factory{
		localApplicationBuilder:  localApplicationBuilder,
		remoteApplicationBuilder: remoteApplicationBuilder,
	}

	return &out
}

// Create creates a new application
func (app *factory) Create() applications.Application {
	instanceAdapter := instances.NewAdapter()
	encryptor := edwards25519.NewEncryptor()
	return applications.NewApplication(
		instructions.NewApplication(
			assignments.NewApplication(
				assignables.NewApplication(
					compilers.NewApplication(
						instanceAdapter,
					),
					assignables_executions.NewApplication(
						assignables_executions_executes.NewApplication(),
						assignables_executions_inits.NewApplication(),
						assignables_executions_retrieves.NewApplication(),
					),
					assignables_bytes.NewApplication(),
					assignables_constants.NewApplication(),
					cryptography.NewApplication(
						cryptography_decrypts.NewApplication(
							encryptor,
						),
						cryptography_encrypts.NewApplication(
							encryptor,
						),
						cryptography_keys.NewApplication(
							cryptography_keys_encryptions.NewApplication(
								cryptography_keys_encryptions_decrypts.NewApplication(),
								cryptography_keys_encryptions_encrypts.NewApplication(),
								keyEncryptionBitrate,
							),
							cryptography_keys_signatures.NewApplication(
								cryptography_keys_signatures_votes.NewApplication(
									cryptography_keys_signatures_votes_creates.NewApplication(),
									cryptography_keys_signatures_votes_validates.NewApplication(),
								),
								cryptography_keys_signatures_signs.NewApplication(
									cryptography_keys_signatures_signs_creates.NewApplication(),
									cryptography_keys_signatures_signs_validates.NewApplication(),
								),
							),
						),
					),
					assignables_lists.NewApplication(
						assignables_lists_fetches.NewApplication(),
					),
					executables.NewApplication(
						app.localApplicationBuilder,
						app.remoteApplicationBuilder,
					),
				),
			),
			instructions_lists.NewApplication(
				instructions_lists_inserts.NewApplication(),
				instructions_lists_deletes.NewApplication(),
			),
			instructions_executions.NewApplication(
				instructions_executions_merges.NewApplication(),
			),
		),
		NewLayerBinaryApplication(),
	)
}
