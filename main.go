package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

/*Method Summary
static void	backgroundMode(int mode)
          This specifies how the background is drawn.
 long	backward(double distance)	🟩🟩🟩
          Moves the turtle backward by the distance.
static void	bgcolor(java.awt.Color color)
          Sets the background color.
static void	bgcolor(java.lang.String color)
          Sets the background color.
static void	bgpic(java.lang.String filename)
          Set the background image.
static double	canvasX(double screenX)
          Converts screen coordinates to canvas coordinates.
static double	canvasY(double screenY)
          Converts screen coordinates to canvas coordinates.
 void	clear()		🟩🟩🟩
          Clears all the drawing that a turtle has done but all the turtle settings remain the same.
 Turtle	clone()
          This creates a cloned copy of a turtle.
 boolean	contains(int x, int y)
          Determines if a turtle is covering a screen position
 double	distance(double x, double y)
          Gets the distance to another position.
 long	dot()	🟩🟩🟩
          Put a dot 3 times the size of the penWidth on the canvas.
 long	dot(java.awt.Color color)	🟩🟩🟩
          Put a dot 3 times the size of the penWidth on the canvas.
 long	dot(java.awt.Color color, double dotSize)	🟩🟩🟩
          Put a dot on the canvas.
 long	dot(java.lang.String color)		🟩🟩🟩
          Put a dot 3 times the size of the penWidth on the canvas.
 long	dot(java.lang.String color, double dotSize)		🟩🟩🟩
          Put a dot on the canvas.
 long	down()		🟩🟩🟩
          Puts the turtle's tail down so it will draw on the screen as it moves.
static void	exit()
 long	face(double x, double y)	🟩🟩🟩
          Sets the direction in such a way that it faces (x,y)
 long	forward(double distance)	🟩🟩🟩
          Moves the turtle forward by the distance.
 double	GetDirection()
          Gets the direction the turtle is facing neglecting tilt.
 double	GetSpeed()
          Gets the speed of the animation.
 double	GetTilt()
          Gets the rotation of the turtle's shape away from the turtle's direction.
 double	GetX()
          Gets the x coordinate of the turtle.
 double	GetY()
          Gets the y coordinate of the turtle.
 long	hide()
          Hides the turtle but it can still draw.
 long	home()		🟩🟩🟩
          Moves the turtle to (0,0) and facing east.
 long	left(double angle)  	🟩🟩🟩
          Turns the turtle left by the number of indicated degrees.
 long	outlineColor(java.awt.Color outlineColor)
          Sets the turtle's outlineColor color.
 long	outlineColor(java.lang.String outlineColor)
          Sets the turtle's outlineColor color.
 long	outlineWidth(double width)
          Sets the width of the turtle's outline.
 long	penColor(java.awt.Color penColor)		🟩🟩🟩
          Sets the turtle's path color.
 long	penColor(java.lang.String penColor)		🟩🟩🟩
          Sets the turtle's path color.
 void	redo()
          Redo turtle state changes.
 void	redo(int steps)
          Redo turtle state changes.
static void	refreshMode(int mode)
          This specifies when the screen Gets refreshed.
 long	right(double angle)		🟩🟩🟩
          Turns the turtle right by the number of indicated degrees.
static void	save(java.lang.String filename)		🟩🟩🟩
          Saves the visible canvas to an image.
static void	setCanvasSize(int width, int height)
          Changes the size of the canvas effectively changing the size of the window.
 long	setDirection(double direction)		🟩🟩🟩
          Sets the direction the turtle is facing neglecting tilt.
 long	setPosition(double x, double y)		🟩🟩🟩
          Sets the position and direction of a turtle.
 longFONT>	setPosition(double x, double y, double direction)	🟩🟩🟩
          Sets the position and direction of a turtle.
 long	setTilt(double angle)
          Sets the angle to rotate the turtle's shape when rendering.
 long	shape(java.lang.String shape)
          Sets the shape of the turtle using the built in shapes (turtle,square, rectangle,triangle,arrow,circle) or to a image.
 long	shapeSize(int width, int height)

 long	show()
          Makes the turtle visible.
 long	speed(double delay)
          Sets the speed of the animation.
 long	stamp()
          Put a copy of the current turtle shape on the canvas.
static void	startApplet(javax.swing.JApplet applet)
          This is an experimental method that should allow you to make turtle applets in the future.
 long	tilt(double angle)
          Adds an additional angle to rotation of the turtle's shape when rendering.
 double	towards(double x, double y)		🟩🟩🟩
          Gets direction towards (x,y)
 void	undo()
          Undo the last turtle state change.
 void	undo(int steps)
          Undo turtle state changes.
 long	up()	🟩🟩🟩
          Picks the turtle's tail up so it won't draw on the screen as it moves.
static void	update()
          Force redraw when the refreshMode is set to on demand.
 long	width(double penWidth)		🟩🟩🟩
          Sets the width of the turtles path
 long	write(java.lang.String text, java.lang.String fontName, int fontSize, int justification, double xOffset, double yOffset)

static void	zoom(double minx, double miny, double maxx, double maxy)
          Fits the indicated box in the center of the screen as large as possible.
static void	zoomFit()
          Fits everything on the screen.
*/

type Turtle struct {
	x          float64
	y          float64
	sizeX      int
	sizeY      int
	isDown     bool
	showTurtle bool
	direction  float64 // in radians
	width      float64
	color      *color.RGBA64
	bgColor    *color.RGBA64
	cxt        *gg.Context
	bg         *gg.Context
}

func NewTurtle() *Turtle {
	cnvs := gg.NewContext(1000, 1000)
	bg := gg.NewContext(1000, 1000)
	bg.SetRGBA(1, 1, 1, 1)
	bg.LineTo(0, 1000)
	bg.LineTo(1000, 1000)
	bg.LineTo(1000, 0)
	bg.LineTo(0, 0)
	bg.Fill()
	cnvs.MoveTo(500, 500)
	cnvs.SetLineWidth(5)
	cnvs.SetRGBA(0, 0, 0, 0)
	cnvs.Clear()
	cnvs.SetRGBA(0, 0, 0, 1)
	t := &Turtle{x: 500, y: 500, sizeX: 1000, sizeY: 1000, cxt: cnvs, bg: bg, isDown: true, showTurtle: true, direction: -gg.Radians(90), width: 5, color: &color.RGBA64{R: 0, G: 0, B: 0, A: 1}, bgColor: &color.RGBA64{1, 1, 1, 1}}
	return t
}

func (t *Turtle) Save(path string) {
	tmp := gg.NewContext(t.sizeX, t.sizeY)
	tmp.DrawImage(t.bg.Image(), 0, 0)
	tmp.DrawImage(t.cxt.Image(), 0, 0)

	if t.showTurtle {

		// loads Turtle Marker
		f, err := os.Open("marker.png")
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		defer f.Close()

		// Gens Temp Turtle Context
		marker, _, _ := image.Decode(f)
		imgTmp := gg.NewContext(64, 64)
		imgTmp.RotateAbout(t.direction+(math.Pi/2), 32, 32)
		imgTmp.DrawImageAnchored(marker, 32, 32, 0.5, 0.5)

		tmp.DrawImageAnchored(imgTmp.Image(), int(t.x), int(t.y), 0.5, 0.5)
	}
	tmp.SavePNG("./static/turtle.png")
}

func (t *Turtle) Clear() {
	t.cxt.MoveTo(500, 500)
	t.cxt.SetLineWidth(5)
	t.cxt.SetRGBA(0, 0, 0, 0)
	t.cxt.Clear()
	t.cxt.SetRGBA(0, 0, 0, 1)
}

func (t *Turtle) BgColor(color *image.RGBA) {}

func (t *Turtle) GetX() float64 {
	return t.x
}

func (t *Turtle) GetY() float64 {
	return t.y
}

func (t *Turtle) Distance(x float64, y float64) float64 {
	return math.Sqrt(math.Pow((x-t.x), 2) + math.Pow((y-t.y), 2))
}

func (t *Turtle) Contains() bool {
	return false
}

func (t *Turtle) Move(dist float64, dir float64) {
	x_ := math.Cos(dir)*dist + t.x
	y_ := math.Sin(dir)*dist + t.y
	if t.isDown {
		t.SetWidth(t.width)
		fmt.Println("Line Drawn from", t.x-500, 500-t.y, "to", x_-500, 500-y_)
		t.cxt.DrawLine(t.x, t.y, x_, y_)
		t.cxt.Stroke()
	} else {
		t.cxt.MoveTo(x_, y_)
	}
	t.x = x_
	t.y = y_
}

func (t *Turtle) Forward(dist float64) {
	t.Move(dist, t.direction)
}

func (t *Turtle) Backward(dist float64) {
	t.Move(dist, t.direction-math.Pi)
}

func (t *Turtle) Up() {
	t.isDown = false
}

func (t *Turtle) Down() {
	t.isDown = true
}

func (t *Turtle) Dot(color *color.RGBA64, dotSize float64) {
	currColor := t.color
	t.cxt.SetColor(color)
	t.cxt.DrawPoint(t.x, t.y, dotSize/2)
	t.cxt.SetColor(currColor)
}

func (t *Turtle) Towards(x float64, y float64) float64 {
	return math.Atan2(y-t.y, x-t.x)
}

func (t *Turtle) Face(x float64, y float64) {
	t.direction = t.Towards(x, y)
}

// Returns the Direction in Degrees
func (t *Turtle) GetDirection() float64 {
	return gg.Degrees(t.direction)
}

func (t *Turtle) SetDirection(angle float64) {
	t.direction = gg.Radians(angle)
}

func (t *Turtle) Home() {
	if t.isDown {
		t.cxt.DrawLine(t.x, t.y, 0, 0)
	} else {
		t.cxt.MoveTo(0, 0)
	}
}

func (t *Turtle) SetPenColor(color *color.RGBA64) {
	t.cxt.SetColor(t.color)
}
func (t *Turtle) SetBgColor(color contextColor) {
	t.bg.SetRGBA255(color.R, color.G, color.B, color.A)
	t.bg.LineTo(0, 1000)
	t.bg.LineTo(1000, 1000)
	t.bg.LineTo(1000, 0)
	t.bg.LineTo(0, 0)
	t.bg.Fill()
}

func (t *Turtle) Left(angle float64) {
	t.SetDirection(gg.Degrees(t.direction) - angle)
}

func (t *Turtle) Right(angle float64) {
	t.SetDirection(gg.Degrees(t.direction) + angle)
}

func (t *Turtle) SetPosition(x float64, y float64, direction float64) {
	t.x = x
	t.y = y
	t.cxt.MoveTo(x, y)
	t.SetDirection(direction)
}

func (t *Turtle) SetWidth(width float64) {
	t.width = width
	t.cxt.SetLineWidth(width)
}

func cmdInput(t *Turtle, h *History, r *bufio.Reader, inpStr string) bool {
	cmdList := strings.Fields(inpStr)
	switch strings.ToLower(cmdList[0]) {
	case "fd":
		dist, _ := strconv.Atoi(cmdList[1])
		t.Forward(float64(dist))
	case "bk":
		dist, _ := strconv.Atoi(cmdList[1])
		t.Backward(float64(dist))
	case "rt":
		angle, _ := strconv.Atoi(cmdList[1])
		t.Right(float64(angle))
	case "lt":
		angle, _ := strconv.Atoi(cmdList[1])
		t.Left(float64(angle))
	case "pu":
		t.Up()
	case "pd":
		t.Down()
	case "ht":
		t.showTurtle = false
	case "st":
		t.showTurtle = true
	case "bg":
		if len(cmdList) > 1 {
			tmp := strings.Trim(cmdList[1], " \n\t\r")
			if tmp == "" {
				t.SetBgColor(GetRGBAInput(r))
			} else {
				t.SetBgColor(GetRGBA(tmp))
			}
		} else {
			t.SetBgColor(GetRGBAInput(r))
		}
	case "cs":
		t.Clear()
	case "repeat":
		cmds, loopCnt := getCommandsToLoop(strings.ToLower(inpStr))
		tmp := parseCmds(strings.Fields(cmds))
		fmt.Printf("Repeating {%s} %d times :\n", cmds, loopCnt)
		for i := 0; i < loopCnt; i++ {
			for _, v := range tmp {
				cmdInput(t, h, r, v)
			}
		}
	case "undo":
		h.Undo()
		img := h.Undo()
		if img != nil {
			t.cxt.DrawImage(img, 0, 0)
			buf := new(bytes.Buffer)
			png.Encode(buf, img)
			bt := buf.Bytes()
			ioutil.WriteFile("test.png", bt, 'w')
		}
		return false
	case "redo":
		h.Redo()
		t.cxt.DrawImage(h.UndoPeek(), 0, 0)
		return false
	case "exit":
		return true
	default:
		fmt.Println(`Error: Unknown Command "` + cmdList[0] + `", Try Something Else.`)
		return false
	}
	// buf := new(bytes.Buffer)
	// png.Encode(buf, t.cxt.Image())
	// bt := buf.Bytes()
	// ioutil.WriteFile(strconv.Itoa(h.undoLength)+".png", bt, 'w')
	h.Push(t.cxt.Image())
	t.Save("static/turtle.png")
	return false
}

func main() {

	d, _ := os.Getwd()
	fmt.Println(d)

	t := NewTurtle()
	h := NewHistory()
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Would You Like to Open LoGo in Network Mode (Y/N): ")
	choice, _ := reader.ReadString('\n')
	choice = strings.Trim(choice, " \n\t\r")
	if strings.ToLower(choice) == "y" {
		go serveTurtleNet(t, h, reader)
	}

	t.Save("static/turtle.png")
	h.Push(t.cxt.Image())

	for {
		fmt.Printf(">>> ")
		cmd, _ := reader.ReadString('\n')
		exit := cmdInput(t, h, reader, cmd)
		fmt.Println("Saving...")
		// fmt.Println("T:", &t)
		if exit {
			break
		}
	}
}
