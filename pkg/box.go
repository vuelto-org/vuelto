/*
 * Copyright (C) 2025 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1.1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package vuelto

type Box2D struct {
	Window      *Window
	Renderer    *Renderer2D
	Children    []*Box2D
	IsChild     bool
	HasChildren bool
}

// Creates a new box, with only a window
func (w *Window) NewBox() *Box2D {
	return &Box2D{
		Window:      w,
		Renderer:    w.NewRenderer2D(),
		IsChild:     false,
		HasChildren: false,
	}
}

// Create a new box with an existing renderer
func (w *Window) NewBoxER(ren *Renderer2D) *Box2D {
	return &Box2D{
		Window:      w,
		Renderer:    ren,
		IsChild:     false,
		HasChildren: false,
	}
}
