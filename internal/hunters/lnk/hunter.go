package lnk

import (
	"fmt"

	"github.com/MustafaNafizDurukan/Guard/pkg/types"
)

type Hunter struct {
}

func (i *Hunter) Init() {}

func (i *Hunter) Descriptor() *types.Descriptor {
	return &types.Descriptor{
		Name:        "LNK Hunter",
		Description: "Hunts suspicious LNK files",
		ShortFlag:   "-lnk",
		LongFlag:    "--lnkfiles",
	}
}

func (i *Hunter) Hunt() error {
	fmt.Println("Process hunted successfully!")
	return nil
}
