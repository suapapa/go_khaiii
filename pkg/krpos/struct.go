// Copyright 2020 Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krpos

import "fmt"

// POS represents Korean Part Of Speech; 품사
type POS struct {
	Pos         string
	Category    string
	Description string
}

func (p POS) String() string {
	return fmt.Sprintf("%s(%s-%s)", p.Pos, p.Category, p.Description)
}
