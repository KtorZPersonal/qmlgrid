// qmlgrid aims at giving a wrapper around qml to easily define a grid with various set of tiles
// In that way, it is possible to quickly display a grid and operate small changes on several tiles
// to visualize the effect of an algorithm.
package qmlgrid

import "gopkg.in/qml.v1"

type Grid struct {
	width    int
	height   int
	tilesize int
	tiles    [][]qml.Object
	window   *qml.Window
}

const (
	empty = iota
	blocked
	visited
	active
	goal
)

func (grid *Grid) Draw(parent qml.Object) {
	for _, line := range grid.tiles {
		for _, tile := range line {
			tile.Set("parent", parent)
		}
	}
}

func create(width, height, tilesize int) (*Grid, error) {
    grid := &Grid{
        width,
        height,
        tilesize,
        nil,
        nil,
    }

	engine := qml.NewEngine()

	tileComponent, err := engine.LoadFile("tile.qml")
	if err != nil {
		return nil, err
	}

	grid.tiles = make([][]qml.Object, height)
	for i := range grid.tiles {
		grid.tiles[i] = make([]qml.Object, width)
		for j := range grid.tiles[i] {
			tile := tileComponent.Create(nil)
			tile.Set("x", j*tilesize)
			tile.Set("y", i*tilesize)
			tile.Set("width", tilesize)
			tile.Set("height", tilesize)
			tile.Set("kind", empty)
			grid.tiles[i][j] = tile
		}
	}

	context := engine.Context()
	context.SetVar("grid", grid)

	gridComponent, err := engine.LoadFile("grid.qml")
	if err != nil {
		return nil, err
	}

	grid.window = gridComponent.CreateWindow(context)
	grid.window.Set("width", width*tilesize)
	grid.window.Set("height", height*tilesize)
	return grid, err
}

// SetVisited transform an existing tile of the grid into a visited one
// i.e. color will change to gray
func (grid *Grid) SetVisited(i int, j int) {
	grid.tiles[i][j].Set("kind", visited)
}

// SetActive transform an existing tile of the grid into an one
// i.e. color will change to green
func (grid *Grid) SetActive(i int, j int) {
	grid.tiles[i][j].Set("kind", active)
}

// SetBlocked transform an existing tile of the grid into a blocked one
// i.e. color will change to dark blue. Also, this will have an impact
// on the IsWalkable method which will now return false for the given tile
func (grid *Grid) SetBlocked(i int, j int) {
	grid.tiles[i][j].Set("kind", blocked)
}

// SetGoal transform an existing tile of the grid into a goal
// i.e. color will change to red. Also, thid will have an impact
// on the IsGoal method which will now return true for the given tile
func (grid *Grid) SetGoal(i int, j int) {
	grid.tiles[i][j].Set("kind", goal)
}

// SetEmpty transform an existing tile of the grid into an empty one
// i.e. color will change to light gray
func (grid *Grid) SetEmpty(i int, j int) {
	grid.tiles[i][j].Set("kind", empty)
}

// IsWalkable return false if the given coordinates are out of the grid or,
// if they target a Blocked tile. True otherwise.
func (grid *Grid) IsWalkable(i, j int) bool {
    if i < 0 || i >= grid.height || j < 0 || j >= grid.width {
        return false
    }
    return grid.tiles[i][j].Property("kind") != blocked
}

// IsGoal return true if the target is a Goal tile
func (grid *Grid) IsGoal(i, j int) bool {
    return grid.tiles[i][j].Property("kind") == goal
}

// New create a new tile and execute the running function passed as argument
func New(w int, h int, s int, f func(g *Grid) error) error {
	return qml.Run(func () error {
        grid, err := create(w, h, s)
        if err != nil {
            return err
        }
        grid.window.Show()
        err = f(grid)
        if err != nil {
            return err
        }
        grid.window.Wait()
        return nil
    })
}
