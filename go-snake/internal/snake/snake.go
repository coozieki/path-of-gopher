package snake

import (
	"go-snake/internal/config"
)

type Snake interface {
	Move()
	PushBlock()
	GetBlocks() []Block
	ChangeDir(direction Direction)
	GetNextHeadCoords() (x, y float64)
	BlockExistsAt(x, y float64) bool
}

type snake struct {
	blocks     []Block
	currentDir Direction
	blocksMap  map[float64]map[float64]bool
}

func NewSnake() *snake {
	return &snake{
		blocks:     []Block{{X: 0, Y: 0, Direction: DirectionRight}},
		currentDir: DirectionRight,
		blocksMap: map[float64]map[float64]bool{
			0: {0: true},
		},
	}
}

func (s *snake) PushBlock() {
	firstBlock := s.blocks[len(s.blocks)-1]

	var x, y = firstBlock.X, firstBlock.Y
	switch s.currentDir {
	case DirectionRight:
		x++
	case DirectionDown:
		y++
	case DirectionUp:
		y--
	case DirectionLeft:
		x--
	default:
		return
	}
	s.blocks = append(s.blocks, Block{X: x, Y: y, Direction: s.currentDir})
	_, ok := s.blocksMap[x]
	if !ok {
		s.blocksMap[x] = map[float64]bool{}
	}
	s.blocksMap[x][y] = true
}

func (s *snake) GetBlocks() []Block {
	return s.blocks
}

func (s *snake) Move() {
	temp := s.blocks[len(s.blocks)-1]
	temp.Direction = s.currentDir
	temp.X, temp.Y = s.GetNextHeadCoords()
	if s.BlockExistsAt(temp.X, temp.Y) {
		return
	}
	firstBlock := s.blocks[0]
	s.blocksMap[firstBlock.X][firstBlock.Y] = false
	s.blocks = s.blocks[1:]
	s.blocks = append(s.blocks, temp)
	_, ok := s.blocksMap[temp.X]
	if !ok {
		s.blocksMap[temp.X] = map[float64]bool{}
	}
	s.blocksMap[temp.X][temp.Y] = true
}

func (s *snake) ChangeDir(direction Direction) {
	switch direction {
	case DirectionRight:
		if s.currentDir == DirectionLeft {
			return
		}
	case DirectionDown:
		if s.currentDir == DirectionUp {
			return
		}
	case DirectionLeft:
		if s.currentDir == DirectionRight {
			return
		}
	case DirectionUp:
		if s.currentDir == DirectionDown {
			return
		}
	}
	s.currentDir = direction
}

func (s *snake) GetNextHeadCoords() (x float64, y float64) {
	block := s.blocks[len(s.blocks)-1]
	switch s.currentDir {
	case DirectionRight:
		block.X++
		if block.X > config.FieldWidthInBlocks-1 {
			block.X = 0
		}
	case DirectionDown:
		block.Y++
		if block.Y > config.FieldWidthInBlocks-1 {
			block.Y = 0
		}
	case DirectionLeft:
		block.X--
		if block.X < 0 {
			block.X = config.FieldWidthInBlocks - 1
		}
	case DirectionUp:
		block.Y--
		if block.Y < 0 {
			block.Y = config.FieldWidthInBlocks - 1
		}
	}
	return block.X, block.Y
}

func (s *snake) BlockExistsAt(x, y float64) bool {
	return s.blocksMap[x][y]
}