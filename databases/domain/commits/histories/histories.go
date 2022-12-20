package histories

type histories struct {
	mp   map[string]History
	list []History
}

func createHistories(
	mp map[string]History,
	list []History,
) Histories {
	out := histories{
		mp:   mp,
		list: list,
	}

	return &out
}

// List returns the histories
func (obj *histories) List() []History {
	return obj.list
}

// Compare compares histories and return the ones with the better score
func (obj *histories) Compare(ins Histories) ([]History, error) {
	output := []History{}
	paramList := ins.List()
	for idx, oneParamHistory := range paramList {
		keyname := oneParamHistory.Commit().String()
		if currentHistory, ok := obj.mp[keyname]; ok {
			if oneParamHistory.Score() > currentHistory.Score() {
				output = append(output, oneParamHistory)
				continue
			}

			continue
		}

		if len(obj.list)-1 < idx {
			output = append(output, oneParamHistory)
			continue
		}

		if obj.list[idx].Score() < oneParamHistory.Score() {
			output = append(output, oneParamHistory)
			continue
		}
	}

	return output, nil
}
