package main

const Player1 = 1
const Player2 = 2

type GameImp struct {
	Board
	activePlayer int
	gameOver     bool
}

func NewGameState() *GameImp {
	newGame := &GameImp{Board: NewBoard()}
	newGame.activePlayer = Player1
	newGame.setupPieces()
	return newGame
}

func (game *GameImp) setupPieces() {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			player := Player1
			if i == 0 || i == 1 {
				player = Player2
			}
			posn := IPosn{i, j}
			info := PieceInfo{player: player, IPosn: posn}

			var piece Piece
			if i == 0 || i == 7 {
				switch j {
				case 0, 7:
					piece = NewPiece(rook, info)
				case 1, 6:
					piece = NewPiece(knight, info)
				case 2, 5:
					piece = NewPiece(bishop, info)
				case 3:
					piece = NewPiece(queen, info)
				case 4:
					piece = NewPiece(king, info)
				}
			} else if i == 1 || i == 6 {
				piece = NewPiece(pawn, info)
			}
			game.setSquare(posn, piece)
		}
	}
}

func (game *GameImp) validMoves() (validMoves map[string][]string) {
	validMoves = make(map[string][]string)
	movesLeft := false

	for _, piece := range game.pieces() {
		if piece.pieceInfo().player == game.activePlayer {
			iMoves := piece.validDests(game.Board)

			// convert to StdPosn
			stdMoves := make([]string, len(iMoves))
			for i, move := range iMoves {
				stdMoves[i] = move.toStdPosn().String()
			}

			validMoves[piece.pieceInfo().IPosn.toStdPosn().String()] = stdMoves
			movesLeft = movesLeft || len(iMoves) > 0
		}
	}

	game.gameOver = !movesLeft
	return validMoves
}

func (game *GameImp) move(move Move) {
	src := move.src.toIPosn()
	dest := move.dest.toIPosn()

	// update board
	piece := game.at(src)
	game.setSquare(dest, piece)
	piece.updatePosn(dest)
	game.setSquare(src, nil)

	game.activePlayer = 3 - game.activePlayer // switch players
}

func (game GameImp) ActivePlayer() int {
	return game.activePlayer
}

func (game *GameImp) isOver() bool {
	return game.gameOver
}

func (game *GameImp) fen() (s string) {
	return game.Board.fen()
}