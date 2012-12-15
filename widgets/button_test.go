/*
   Copyright 2012 the go.uik authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package widgets

import (
	"code.google.com/p/draw2d/draw2d"
	"image"
	"testing"
)

func BenchmarkButtonDraw(b *testing.B) {
	b.StopTimer()
	button := NewButton("test")
	ctx := draw2d.NewGraphicContext(image.NewRGBA(image.Rect(0, 0, 100, 100)))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		button.DoPaint(ctx)
	}
}
