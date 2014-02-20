package c4

import (
  "math/rand"
)

type Game struct {
  turn int
  board [6][7]int
  moveCount int
  winner int
}

func NewGame() (Game) {
  game := Game{}
  game.turn = rand.Intn(2) + 1
  return game
}

func (self *Game) Turn() int { return self.turn }
func (self *Game) Winner() int { return self.winner }
func (self *Game) Board() [6][7]int { return self.board }

func (self Game) String() string {
  str := ""
  for i := 0 ; i<6 ; i++ {
    str += "| "
    for j := 0 ; j<7 ; j++ {
      switch(self.board[i][j]){
      case 1:
        str += "#"
      case 2:
        str += "@"
      default:
        str += "."
      }
      str += " "
    }
    str += "|\n"
  }
  str += "| 0 1 2 3 4 5 6 |\n"
  if (self.turn != 0){
    if (self.turn == 1){
      str += "# to move"
    } else {
      str += "@ to move"
    }
  } else {
    if (self.winner == 1){
      str += "Game Over: # wins"
    } else if (self.winner == 2){
      str += "Game Over: @ wins"
    } else {
      str += "Game Over: tie"
    }
  }
  return str
}

func (self *Game) Move(col int) bool{
  if (col < 0 || col > 6 || self.board[0][col] != 0 || self.turn == 0){
    return false
  }
  for i := 5 ; i >= 0 ; i-- {
    if (self.board[i][col] == 0){
      self.board[i][col] = self.turn
      self.turn = 3 - self.turn
      self.moveCount++
      return self.checkWin(i,col)
    }
  }
  return false
}

type coord struct{
  row int
  col int
}

func (self *Game) checkWin(row, col int) bool{
  if (self.moveCount) < 6 {
    return false
  }
  if (self.moveCount == 42){
    return self.gameover()
  }
  mover := 3 - self.turn

  var linesToTry [4][]coord
  for i := 0 ; i<4 ; i++ {
    linesToTry[i] = make([]coord, 0, 7)
  }
  c := 0
  r := 0
  for i := 0; i<4; i++ {
    r = row + i
    if (r < 6){
      linesToTry[0] = append(linesToTry[0], coord{r, col})
    }
  }
  for i:=-3; i<4; i++ {
    c = col + i
    r = row + i
    if ( c >= 0 && c <= 6){
      linesToTry[1] = append(linesToTry[1], coord{row, c})
      if (r >= 0 && r <= 5){
        linesToTry[2] = append(linesToTry[2], coord{r, c})
      }
    }

    c = col - i
    if ( c >= 0 && c <= 6 && r >= 0 && r <= 5){
      linesToTry[3] = append(linesToTry[3], coord{r, c})
    }
  }

  count := 0
  for _,lineToTry := range linesToTry {
    count = 0
    for _,coord := range lineToTry {
      if (self.board[coord.row][coord.col] == mover){
        count += 1
        if (count == 4){
          return self.gameover()
        }
      } else {
        count = 0
      }
    }
  }
  
  return false
}

func (self *Game) gameover () bool {
  if (self.moveCount != 42){
    self.winner = 3 - self.turn
  }
  self.turn = 0
  return true
}

func (self *Game) Moves() []int{
  moves := make([]int,0,7)
  for i := 0; i<7; i++ {
    if (self.board[0][i] == 0){
      moves = append(moves, i)
    }
  }
  return moves
}

func (self *Game) Copy()(Game){
  board := self.board
  return Game{self.turn, board, self.moveCount, self.winner}
}

