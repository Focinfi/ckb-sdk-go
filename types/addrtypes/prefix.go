package addrtypes

import (
	"github.com/Focinfi/ckb-sdk-go/types"
)

const (
	PrefixTestNet = "ckt"
	PrefixMainNet = "ckb"
)

var prefixList = []string{PrefixTestNet, PrefixMainNet}

func IsAllowedPrefix(prefix string) bool {
	for _, p := range prefixList {
		if prefix == p {
			return true
		}
	}
	return false
}

func PrefixForMode(mode types.Mode) string {
	switch mode {
	case types.ModeMainNet:
		return PrefixMainNet
	case types.ModeTestNet:
		return PrefixTestNet
	}
	return ""
}
