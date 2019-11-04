package utils

import "github.com/Focinfi/ckb-sdk-go/types"

func EmptyWitnessesByLen(len int) []string {
	witnesses := make([]string, 0, len)
	for i := 0; i < len; i++ {
		witnesses = append(witnesses, types.HexStrPrefix)
	}
	return witnesses
}
