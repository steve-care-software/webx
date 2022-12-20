package histories

import "github.com/steve-care-software/webx/databases/domain/commits"

type adapter struct {
	builder        Builder
	historyBuilder HistoryBuilder
}

func createAdapter(
	builder Builder,
	historyBuilder HistoryBuilder,
) Adapter {
	out := adapter{
		builder:        builder,
		historyBuilder: historyBuilder,
	}

	return &out
}

// FromCommitsToHistories converts commits to histories
func (app *adapter) FromCommitsToHistories(commits []commits.Commit) (Histories, error) {
	list := []History{}
	for _, oneCommit := range commits {
		score := uint(0)
		if oneCommit.HasMine() {
			score = oneCommit.Mine().Score()
		}

		hash := oneCommit.Hash()
		history, err := app.historyBuilder.Create().WithScore(score).WithCommit(hash).Now()
		if err != nil {
			return nil, err
		}

		list = append(list, history)
	}

	return app.builder.Create().WithList(list).Now()
}

// ToContent converts histories to content
func (app *adapter) ToContent(ins Histories) ([]byte, error) {
	return nil, nil
}

// ToHistories converts content to histories
func (app *adapter) ToHistories(content []byte) (Histories, error) {
	return nil, nil
}
