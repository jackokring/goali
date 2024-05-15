// This package is a cell based game framework.
//
// Written in a normal form, and later expanded for speed memory trade-offs.
// The speed form may be cache serialized also.
package zone

//=================================
//******** Game Framework *********
//=================================

type Game interface {
	// support multiple game contexts
}

type Inventory interface {
}

type InventoryLocation struct {
	// normal form, item at (has) location
	// fast form location contains item(s)?
}

type Item interface {
}

type LocatedItem struct {
	// join
	Item
	Location
}

type Location interface {
}

type Map interface {
}

type MapLocation struct {
}

type Note interface {
}

type NoteLocation struct {
}

type Table interface {
	// support database interface
}

type Tile interface {
}

type TileGroup struct {
	// sparse occupancy
}
