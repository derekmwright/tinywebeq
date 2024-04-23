package db

import "strings"

// ClassesFromMask returns a string of classes from a bitmask
func ClassesFromMask(in int) string {
	out := ""
	if in == 65535 {
		return "ALL"
	}
	if in&16384 != 0 {
		out += "BST "
	}
	if in&8192 != 0 {
		out += "ENC "
	}
	if in&4096 != 0 {
		out += "MAG "
	}
	if in&2048 != 0 {
		out += "WIZ "
	}
	if in&1024 != 0 {
		out += "NEC "
	}
	if in&512 != 0 {
		out += "SHM "
	}
	if in&256 != 0 {
		out += "ROG "
	}
	if in&128 != 0 {
		out += "BRD "
	}
	if in&64 != 0 {
		out += "MNK "
	}
	if in&32 != 0 {
		out += "DRU "
	}
	if in&16 != 0 {
		out += "SHD "
	}
	if in&8 != 0 {
		out += "RNG "
	}
	if in&4 != 0 {
		out += "PAL "
	}
	if in&2 != 0 {
		out += "CLR "
	}
	if in&1 != 0 {
		out += "WAR "
	}

	out = strings.TrimSuffix(out, " ")
	return out
}
