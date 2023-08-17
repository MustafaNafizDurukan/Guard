package hunters

import (
	"fmt"

	"github.com/MustafaNafizDurukan/Guard/internal/hunters/lnk"
	"github.com/MustafaNafizDurukan/Guard/pkg/types"
)

var allHunterList = []types.IHunter{
	&lnk.Hunter{},
}

func List() []types.IHunter {
	return allHunterList
}

func Select(shortFlag string, longFlag string) types.IHunter {
	for _, hunter := range allHunterList {
		if hunter.Descriptor().ShortFlag == shortFlag {
			return hunter
		}

		if hunter.Descriptor().LongFlag == longFlag {
			return hunter
		}
	}

	return nil
}

func Init(hunter types.IHunter) {
	hunter.Init()
}

func Hunt(hunter types.IHunter) bool {
	err := hunter.Hunt()
	if err != nil {
		fmt.Printf("An error ocurred while hunting with %s error: %v", hunter.Descriptor().Name, err)
		return false
	}

	return true
}
