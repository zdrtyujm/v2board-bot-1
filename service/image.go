package service

import (
	"image"
	"image/color"
	"os"
	"unicode"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

var (
	Aliceblue            = color.RGBA{0xf0, 0xf8, 0xff, 0xff} // rgb(240, 248, 255)
	Antiquewhite         = color.RGBA{0xfa, 0xeb, 0xd7, 0xff} // rgb(250, 235, 215)
	Aqua                 = color.RGBA{0x00, 0xff, 0xff, 0xff} // rgb(0, 255, 255)
	Aquamarine           = color.RGBA{0x7f, 0xff, 0xd4, 0xff} // rgb(127, 255, 212)
	Azure                = color.RGBA{0xf0, 0xff, 0xff, 0xff} // rgb(240, 255, 255)
	Beige                = color.RGBA{0xf5, 0xf5, 0xdc, 0xff} // rgb(245, 245, 220)
	Bisque               = color.RGBA{0xff, 0xe4, 0xc4, 0xff} // rgb(255, 228, 196)
	Black                = color.RGBA{0x00, 0x00, 0x00, 0xff} // rgb(0, 0, 0)
	Blanchedalmond       = color.RGBA{0xff, 0xeb, 0xcd, 0xff} // rgb(255, 235, 205)
	Blue                 = color.RGBA{0x00, 0x00, 0xff, 0xff} // rgb(0, 0, 255)
	Blueviolet           = color.RGBA{0x8a, 0x2b, 0xe2, 0xff} // rgb(138, 43, 226)
	Brown                = color.RGBA{0xa5, 0x2a, 0x2a, 0xff} // rgb(165, 42, 42)
	Burlywood            = color.RGBA{0xde, 0xb8, 0x87, 0xff} // rgb(222, 184, 135)
	Cadetblue            = color.RGBA{0x5f, 0x9e, 0xa0, 0xff} // rgb(95, 158, 160)
	Chartreuse           = color.RGBA{0x7f, 0xff, 0x00, 0xff} // rgb(127, 255, 0)
	Chocolate            = color.RGBA{0xd2, 0x69, 0x1e, 0xff} // rgb(210, 105, 30)
	Coral                = color.RGBA{0xff, 0x7f, 0x50, 0xff} // rgb(255, 127, 80)
	Cornflowerblue       = color.RGBA{0x64, 0x95, 0xed, 0xff} // rgb(100, 149, 237)
	Cornsilk             = color.RGBA{0xff, 0xf8, 0xdc, 0xff} // rgb(255, 248, 220)
	Crimson              = color.RGBA{0xdc, 0x14, 0x3c, 0xff} // rgb(220, 20, 60)
	Cyan                 = color.RGBA{0x00, 0xff, 0xff, 0xff} // rgb(0, 255, 255)
	Darkblue             = color.RGBA{0x00, 0x00, 0x8b, 0xff} // rgb(0, 0, 139)
	Darkcyan             = color.RGBA{0x00, 0x8b, 0x8b, 0xff} // rgb(0, 139, 139)
	Darkgoldenrod        = color.RGBA{0xb8, 0x86, 0x0b, 0xff} // rgb(184, 134, 11)
	Darkgray             = color.RGBA{0xa9, 0xa9, 0xa9, 0xff} // rgb(169, 169, 169)
	Darkgreen            = color.RGBA{0x00, 0x64, 0x00, 0xff} // rgb(0, 100, 0)
	Darkgrey             = color.RGBA{0xa9, 0xa9, 0xa9, 0xff} // rgb(169, 169, 169)
	Darkkhaki            = color.RGBA{0xbd, 0xb7, 0x6b, 0xff} // rgb(189, 183, 107)
	Darkmagenta          = color.RGBA{0x8b, 0x00, 0x8b, 0xff} // rgb(139, 0, 139)
	Darkolivegreen       = color.RGBA{0x55, 0x6b, 0x2f, 0xff} // rgb(85, 107, 47)
	Darkorange           = color.RGBA{0xff, 0x8c, 0x00, 0xff} // rgb(255, 140, 0)
	Darkorchid           = color.RGBA{0x99, 0x32, 0xcc, 0xff} // rgb(153, 50, 204)
	Darkred              = color.RGBA{0x8b, 0x00, 0x00, 0xff} // rgb(139, 0, 0)
	Darksalmon           = color.RGBA{0xe9, 0x96, 0x7a, 0xff} // rgb(233, 150, 122)
	Darkseagreen         = color.RGBA{0x8f, 0xbc, 0x8f, 0xff} // rgb(143, 188, 143)
	Darkslateblue        = color.RGBA{0x48, 0x3d, 0x8b, 0xff} // rgb(72, 61, 139)
	Darkslategray        = color.RGBA{0x2f, 0x4f, 0x4f, 0xff} // rgb(47, 79, 79)
	Darkslategrey        = color.RGBA{0x2f, 0x4f, 0x4f, 0xff} // rgb(47, 79, 79)
	Darkturquoise        = color.RGBA{0x00, 0xce, 0xd1, 0xff} // rgb(0, 206, 209)
	Darkviolet           = color.RGBA{0x94, 0x00, 0xd3, 0xff} // rgb(148, 0, 211)
	Deeppink             = color.RGBA{0xff, 0x14, 0x93, 0xff} // rgb(255, 20, 147)
	Deepskyblue          = color.RGBA{0x00, 0xbf, 0xff, 0xff} // rgb(0, 191, 255)
	Dimgray              = color.RGBA{0x69, 0x69, 0x69, 0xff} // rgb(105, 105, 105)
	Dimgrey              = color.RGBA{0x69, 0x69, 0x69, 0xff} // rgb(105, 105, 105)
	Dodgerblue           = color.RGBA{0x1e, 0x90, 0xff, 0xff} // rgb(30, 144, 255)
	Firebrick            = color.RGBA{0xb2, 0x22, 0x22, 0xff} // rgb(178, 34, 34)
	Floralwhite          = color.RGBA{0xff, 0xfa, 0xf0, 0xff} // rgb(255, 250, 240)
	Forestgreen          = color.RGBA{0x22, 0x8b, 0x22, 0xff} // rgb(34, 139, 34)
	Fuchsia              = color.RGBA{0xff, 0x00, 0xff, 0xff} // rgb(255, 0, 255)
	Gainsboro            = color.RGBA{0xdc, 0xdc, 0xdc, 0xff} // rgb(220, 220, 220)
	Ghostwhite           = color.RGBA{0xf8, 0xf8, 0xff, 0xff} // rgb(248, 248, 255)
	Gold                 = color.RGBA{0xff, 0xd7, 0x00, 0xff} // rgb(255, 215, 0)
	Goldenrod            = color.RGBA{0xda, 0xa5, 0x20, 0xff} // rgb(218, 165, 32)
	Gray                 = color.RGBA{0x80, 0x80, 0x80, 0xff} // rgb(128, 128, 128)
	Green                = color.RGBA{0x00, 0x80, 0x00, 0xff} // rgb(0, 128, 0)
	Greenyellow          = color.RGBA{0xad, 0xff, 0x2f, 0xff} // rgb(173, 255, 47)
	Grey                 = color.RGBA{0x80, 0x80, 0x80, 0xff} // rgb(128, 128, 128)
	Honeydew             = color.RGBA{0xf0, 0xff, 0xf0, 0xff} // rgb(240, 255, 240)
	Hotpink              = color.RGBA{0xff, 0x69, 0xb4, 0xff} // rgb(255, 105, 180)
	Indianred            = color.RGBA{0xcd, 0x5c, 0x5c, 0xff} // rgb(205, 92, 92)
	Indigo               = color.RGBA{0x4b, 0x00, 0x82, 0xff} // rgb(75, 0, 130)
	Ivory                = color.RGBA{0xff, 0xff, 0xf0, 0xff} // rgb(255, 255, 240)
	Khaki                = color.RGBA{0xf0, 0xe6, 0x8c, 0xff} // rgb(240, 230, 140)
	Lavender             = color.RGBA{0xe6, 0xe6, 0xfa, 0xff} // rgb(230, 230, 250)
	Lavenderblush        = color.RGBA{0xff, 0xf0, 0xf5, 0xff} // rgb(255, 240, 245)
	Lawngreen            = color.RGBA{0x7c, 0xfc, 0x00, 0xff} // rgb(124, 252, 0)
	Lemonchiffon         = color.RGBA{0xff, 0xfa, 0xcd, 0xff} // rgb(255, 250, 205)
	Lightblue            = color.RGBA{0xad, 0xd8, 0xe6, 0xff} // rgb(173, 216, 230)
	Lightcoral           = color.RGBA{0xf0, 0x80, 0x80, 0xff} // rgb(240, 128, 128)
	Lightcyan            = color.RGBA{0xe0, 0xff, 0xff, 0xff} // rgb(224, 255, 255)
	Lightgoldenrodyellow = color.RGBA{0xfa, 0xfa, 0xd2, 0xff} // rgb(250, 250, 210)
	Lightgray            = color.RGBA{0xd3, 0xd3, 0xd3, 0xff} // rgb(211, 211, 211)
	Lightgreen           = color.RGBA{0x90, 0xee, 0x90, 0xff} // rgb(144, 238, 144)
	Lightgrey            = color.RGBA{0xd3, 0xd3, 0xd3, 0xff} // rgb(211, 211, 211)
	Lightpink            = color.RGBA{0xff, 0xb6, 0xc1, 0xff} // rgb(255, 182, 193)
	Lightsalmon          = color.RGBA{0xff, 0xa0, 0x7a, 0xff} // rgb(255, 160, 122)
	Lightseagreen        = color.RGBA{0x20, 0xb2, 0xaa, 0xff} // rgb(32, 178, 170)
	Lightskyblue         = color.RGBA{0x87, 0xce, 0xfa, 0xff} // rgb(135, 206, 250)
	Lightslategray       = color.RGBA{0x77, 0x88, 0x99, 0xff} // rgb(119, 136, 153)
	Lightslategrey       = color.RGBA{0x77, 0x88, 0x99, 0xff} // rgb(119, 136, 153)
	Lightsteelblue       = color.RGBA{0xb0, 0xc4, 0xde, 0xff} // rgb(176, 196, 222)
	Lightyellow          = color.RGBA{0xff, 0xff, 0xe0, 0xff} // rgb(255, 255, 224)
	Lime                 = color.RGBA{0x00, 0xff, 0x00, 0xff} // rgb(0, 255, 0)
	Limegreen            = color.RGBA{0x32, 0xcd, 0x32, 0xff} // rgb(50, 205, 50)
	Linen                = color.RGBA{0xfa, 0xf0, 0xe6, 0xff} // rgb(250, 240, 230)
	Magenta              = color.RGBA{0xff, 0x00, 0xff, 0xff} // rgb(255, 0, 255)
	Maroon               = color.RGBA{0x80, 0x00, 0x00, 0xff} // rgb(128, 0, 0)
	Mediumaquamarine     = color.RGBA{0x66, 0xcd, 0xaa, 0xff} // rgb(102, 205, 170)
	Mediumblue           = color.RGBA{0x00, 0x00, 0xcd, 0xff} // rgb(0, 0, 205)
	Mediumorchid         = color.RGBA{0xba, 0x55, 0xd3, 0xff} // rgb(186, 85, 211)
	Mediumpurple         = color.RGBA{0x93, 0x70, 0xdb, 0xff} // rgb(147, 112, 219)
	Mediumseagreen       = color.RGBA{0x3c, 0xb3, 0x71, 0xff} // rgb(60, 179, 113)
	Mediumslateblue      = color.RGBA{0x7b, 0x68, 0xee, 0xff} // rgb(123, 104, 238)
	Mediumspringgreen    = color.RGBA{0x00, 0xfa, 0x9a, 0xff} // rgb(0, 250, 154)
	Mediumturquoise      = color.RGBA{0x48, 0xd1, 0xcc, 0xff} // rgb(72, 209, 204)
	Mediumvioletred      = color.RGBA{0xc7, 0x15, 0x85, 0xff} // rgb(199, 21, 133)
	Midnightblue         = color.RGBA{0x19, 0x19, 0x70, 0xff} // rgb(25, 25, 112)
	Mintcream            = color.RGBA{0xf5, 0xff, 0xfa, 0xff} // rgb(245, 255, 250)
	Mistyrose            = color.RGBA{0xff, 0xe4, 0xe1, 0xff} // rgb(255, 228, 225)
	Moccasin             = color.RGBA{0xff, 0xe4, 0xb5, 0xff} // rgb(255, 228, 181)
	Navajowhite          = color.RGBA{0xff, 0xde, 0xad, 0xff} // rgb(255, 222, 173)
	Navy                 = color.RGBA{0x00, 0x00, 0x80, 0xff} // rgb(0, 0, 128)
	Oldlace              = color.RGBA{0xfd, 0xf5, 0xe6, 0xff} // rgb(253, 245, 230)
	Olive                = color.RGBA{0x80, 0x80, 0x00, 0xff} // rgb(128, 128, 0)
	Olivedrab            = color.RGBA{0x6b, 0x8e, 0x23, 0xff} // rgb(107, 142, 35)
	Orange               = color.RGBA{0xff, 0xa5, 0x00, 0xff} // rgb(255, 165, 0)
	Orangered            = color.RGBA{0xff, 0x45, 0x00, 0xff} // rgb(255, 69, 0)
	Orchid               = color.RGBA{0xda, 0x70, 0xd6, 0xff} // rgb(218, 112, 214)
	Palegoldenrod        = color.RGBA{0xee, 0xe8, 0xaa, 0xff} // rgb(238, 232, 170)
	Palegreen            = color.RGBA{0x98, 0xfb, 0x98, 0xff} // rgb(152, 251, 152)
	Paleturquoise        = color.RGBA{0xaf, 0xee, 0xee, 0xff} // rgb(175, 238, 238)
	Palevioletred        = color.RGBA{0xdb, 0x70, 0x93, 0xff} // rgb(219, 112, 147)
	Papayawhip           = color.RGBA{0xff, 0xef, 0xd5, 0xff} // rgb(255, 239, 213)
	Peachpuff            = color.RGBA{0xff, 0xda, 0xb9, 0xff} // rgb(255, 218, 185)
	Peru                 = color.RGBA{0xcd, 0x85, 0x3f, 0xff} // rgb(205, 133, 63)
	Pink                 = color.RGBA{0xff, 0xc0, 0xcb, 0xff} // rgb(255, 192, 203)
	Plum                 = color.RGBA{0xdd, 0xa0, 0xdd, 0xff} // rgb(221, 160, 221)
	Powderblue           = color.RGBA{0xb0, 0xe0, 0xe6, 0xff} // rgb(176, 224, 230)
	Purple               = color.RGBA{0x80, 0x00, 0x80, 0xff} // rgb(128, 0, 128)
	Red                  = color.RGBA{0xff, 0x00, 0x00, 0xff} // rgb(255, 0, 0)
	Rosybrown            = color.RGBA{0xbc, 0x8f, 0x8f, 0xff} // rgb(188, 143, 143)
	Royalblue            = color.RGBA{0x41, 0x69, 0xe1, 0xff} // rgb(65, 105, 225)
	Saddlebrown          = color.RGBA{0x8b, 0x45, 0x13, 0xff} // rgb(139, 69, 19)
	Salmon               = color.RGBA{0xfa, 0x80, 0x72, 0xff} // rgb(250, 128, 114)
	Sandybrown           = color.RGBA{0xf4, 0xa4, 0x60, 0xff} // rgb(244, 164, 96)
	Seagreen             = color.RGBA{0x2e, 0x8b, 0x57, 0xff} // rgb(46, 139, 87)
	Seashell             = color.RGBA{0xff, 0xf5, 0xee, 0xff} // rgb(255, 245, 238)
	Sienna               = color.RGBA{0xa0, 0x52, 0x2d, 0xff} // rgb(160, 82, 45)
	Silver               = color.RGBA{0xc0, 0xc0, 0xc0, 0xff} // rgb(192, 192, 192)
	Skyblue              = color.RGBA{0x87, 0xce, 0xeb, 0xff} // rgb(135, 206, 235)
	Slateblue            = color.RGBA{0x6a, 0x5a, 0xcd, 0xff} // rgb(106, 90, 205)
	Slategray            = color.RGBA{0x70, 0x80, 0x90, 0xff} // rgb(112, 128, 144)
	Slategrey            = color.RGBA{0x70, 0x80, 0x90, 0xff} // rgb(112, 128, 144)
	Snow                 = color.RGBA{0xff, 0xfa, 0xfa, 0xff} // rgb(255, 250, 250)
	Springgreen          = color.RGBA{0x00, 0xff, 0x7f, 0xff} // rgb(0, 255, 127)
	Steelblue            = color.RGBA{0x46, 0x82, 0xb4, 0xff} // rgb(70, 130, 180)
	Tan                  = color.RGBA{0xd2, 0xb4, 0x8c, 0xff} // rgb(210, 180, 140)
	Teal                 = color.RGBA{0x00, 0x80, 0x80, 0xff} // rgb(0, 128, 128)
	Thistle              = color.RGBA{0xd8, 0xbf, 0xd8, 0xff} // rgb(216, 191, 216)
	Tomato               = color.RGBA{0xff, 0x63, 0x47, 0xff} // rgb(255, 99, 71)
	Turquoise            = color.RGBA{0x40, 0xe0, 0xd0, 0xff} // rgb(64, 224, 208)
	Violet               = color.RGBA{0xee, 0x82, 0xee, 0xff} // rgb(238, 130, 238)
	Wheat                = color.RGBA{0xf5, 0xde, 0xb3, 0xff} // rgb(245, 222, 179)
	White                = color.RGBA{0xff, 0xff, 0xff, 0xff} // rgb(255, 255, 255)
	Whitesmoke           = color.RGBA{0xf5, 0xf5, 0xf5, 0xff} // rgb(245, 245, 245)
	Yellow               = color.RGBA{0xff, 0xff, 0x00, 0xff} // rgb(255, 255, 0)
	Yellowgreen          = color.RGBA{0x9a, 0xcd, 0x32, 0xff} // rgb(154, 205, 50)

	BlueGreyDarken_4 = color.RGBA{0x26, 0x32, 0x38, 0xff} // rgb(38, 50, 56)
	BlueGreyDarken_1 = color.RGBA{0x54, 0x76, 0x7a, 0xff} // rgb(84, 110, 122)
)

type Image struct {
	i *image.NRGBA
	f *freetype.Context
	d *font.Drawer

	fontSize float64
	t        table
}

type table struct {
	h     int
	v     int
	hsize int
	vsize int
}

type tableBlock struct {
	fromx, fromy int
	tox, toy     int
}

// 获取第几列的所有单元格
func (t *table) hblock(h int) []tableBlock {
	r := make([]tableBlock, 0)
	if h > t.h {
		return r
	}

	for i := 0; i < t.v*t.vsize; i += t.vsize {
		r = append(r, tableBlock{
			fromx: h*t.hsize - t.hsize,
			fromy: i,
			tox:   h * t.hsize,
			toy:   i + t.vsize,
		})
	}
	return r
}

func (t *table) vblock(v int) []tableBlock {
	r := make([]tableBlock, 0)
	if v > t.v {
		return r
	}

	for i := 0; i < t.h*t.hsize; i += t.hsize {
		r = append(r, tableBlock{
			fromx: i,
			fromy: v*t.vsize - t.vsize,
			tox:   i + t.hsize,
			toy:   v * t.vsize,
		})
	}
	return r
}

func IsChinese(str string) bool {
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}

func (img *Image) SetTableText(ss [][]string) {
	for i := 1; i <= img.t.h; i++ {
		ths := img.t.hblock(i)
		for j := 0; j < len(ss[i-1]); j++ {
			img.DrawText(ss[i-1][j], ths[j].fromx+10, ths[j].fromy+(img.t.vsize/2+10))
		}
	}
}

type TableData struct {
	HSize int
	VSize int
	L     [][]string

	MinV int

	FontPath  string
	FontSize  float64
	FontDPI   float64
	FontColor color.Color
}

func NewDefaultTable(ss [][]string, font string) (*Image, error) {
	t := TableData{
		HSize:     300,
		VSize:     50,
		MinV:      6,
		FontPath:  font,
		FontSize:  10,
		FontDPI:   200,
		FontColor: Burlywood,
		L:         ss,
	}

	return Table(t)
}

func Table(t TableData) (*Image, error) {
	h := len(t.L)
	v := len(t.L[0])
	if v < t.MinV {
		v = t.MinV
	}
	img := NewTable(h, v, t.HSize, t.VSize)
	img.SetFontColor(t.FontColor)
	err := img.SetFont(t.FontPath, t.FontSize, t.FontDPI)
	if err != nil {
		return nil, err
	}
	img.SetTableText(t.L)
	return img, nil
}

// h: 多少列
//
// v: 多少行
//
// size: 一列/行多宽
func NewTable(h, v, hsize, vsize int) *Image {
	img := &Image{
		i: image.NewNRGBA(image.Rect(0, 0, h*hsize, v*vsize)),
		f: freetype.NewContext(),
		d: &font.Drawer{},
		t: table{
			h:     h,
			v:     v,
			hsize: hsize,
			vsize: vsize,
		},
	}
	var c color.Color
	var n int

	for j := 0; j <= v*vsize; j += vsize {
		if n%2 == 0 {
			c = Whitesmoke
		} else {
			c = White
		}
		if n == 0 {
			c = BlueGreyDarken_1
		}
		n++
		img.SetBlock(c, 0, img.i.Rect.Dx(), j, j+vsize)
	}
	for i := 0; i <= h*hsize; i += hsize {
		if i != 0 {
			img.SetVerLine(BlueGreyDarken_4, 0, vsize, i)
			img.SetVerLine(Whitesmoke, vsize, img.i.Rect.Dy(), i)
		}
	}
	return img
}

func (img *Image) GetImage() *image.NRGBA {
	return img.i
}

func (img *Image) SetFontColor(color color.Color) {
	img.f.SetSrc(image.NewUniform(color))
}

func (img *Image) DrawText(s string, x, y int) error {
	_, err := img.f.DrawString(s, freetype.Pt(x, y))
	return err
}

func (img *Image) SetFont(fontPath string, size, dpi float64) error {
	b, err := os.ReadFile(fontPath)
	if err != nil {
		return err
	}
	font, err := freetype.ParseFont(b)
	if err != nil {
		return err
	}

	//设置字体
	img.f.SetFont(font)
	// 设置字体大小
	img.f.SetFontSize(size)
	img.fontSize = size

	//设置分辨率
	img.f.SetDPI(dpi)
	// 设置绘制矩阵大小
	img.f.SetClip(img.i.Bounds())

	//设置输出的图片
	img.f.SetDst(img.i)

	return nil
}

func (img *Image) SetBlock(color color.Color, fromx, tox, fromy, toy int) {
	for i := fromx; i <= tox; i++ {
		for j := fromy; j <= toy; j++ {
			img.i.Set(i, j, color)
		}
	}
}

func (img *Image) SetHorizLine(color color.Color, fromx, tox, y int) {
	for i := fromx; i <= tox; i++ {
		img.i.Set(i, y, color)
	}
}

func (img *Image) SetVerLine(color color.Color, fromy, toy, x int) {
	for i := fromy; i <= toy; i++ {
		img.i.Set(x, i, color)
	}
}

func (img *Image) SetBackgroudColor(color color.Color) {
	for i := 0; i <= img.i.Rect.Dx(); i++ {
		for j := 0; j <= img.i.Rect.Dy(); j++ {
			img.i.Set(i, j, color)
		}
	}
}
