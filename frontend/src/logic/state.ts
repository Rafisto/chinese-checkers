// chinese checkers state 0d (20rows x 17cols) 

import { Point } from "./types";

const LogicState = Array.from({ length: 18 }, () => Array(25).fill(0));

const singleMoves = [
    { row: 0, col: 2 },
    { row: 0, col: -2 },
    { row: 1, col: 1 },
    { row: 1, col: -1 },
    { row: -1, col: 1 },
    { row: -1, col: -1 },
]

const jumpMoves = [
    { row: 0, col: 4 },
    { row: 0, col: -4 },
    { row: 2, col: 2 },
    { row: 2, col: -2 },
    { row: -2, col: 2 },
    { row: -2, col: -2 },
]

const PerformMove = (state: number[][], start: Point, end: Point): number[][] => {
    const newState = state.map((row) => row.slice());

    if (state[start.row][start.col] === 0) {
        return state;
    }

    newState[end.row][end.col] = newState[start.row][start.col];
    newState[start.row][start.col] = 0;

    return newState;
}

const AdjacencyMap: Record<number, number> = {
    1: 2,
    2: 1,
    3: 4,
    4: 3,
    5: 6,
    6: 5
}

const NoExitCheck = (board: number[][], state: number[][], start: Point, end: Point, inverted: boolean): boolean => {
    if (inverted) {
        const adjacent = AdjacencyMap[state[start.row][start.col]];
        return board[start.row][start.col] != adjacent || board[end.row][end.col] == adjacent;
    }
    else {
        const normal = state[start.row][start.col];
        return board[start.row][start.col] != normal || board[end.row][end.col] == normal;
    }
}

const FindMoves = (
    board: number[][],
    state: number[][],
    position: Point,
    passMap: boolean[][],
    moves: Point[] = [],
    inverted: boolean = true
): Point[] => {
    passMap[position.row][position.col] = true;

    // Check single moves
    if (moves.length === 0) {
        for (const move of singleMoves) {
            const newRow = position.row + move.row;
            const newCol = position.col + move.col;

            if (
                newRow >= 0 && newRow < state.length &&
                newCol >= 0 && newCol < state[0].length &&
                board[newRow][newCol] !== -1 &&
                state[newRow][newCol] === 0 &&
                !passMap[newRow][newCol]
            ) {
                const newPosition = { row: newRow, col: newCol };
                moves.push(newPosition);
            }
        }
    }

    // Check jump moves
    for (const move of jumpMoves) {
        const midRow = position.row + move.row / 2;
        const midCol = position.col + move.col / 2;
        const newRow = position.row + move.row;
        const newCol = position.col + move.col;

        if (
            midRow >= 0 && midRow < state.length &&
            midCol >= 0 && midCol < state[0].length &&
            newRow >= 0 && newRow < state.length &&
            newCol >= 0 && newCol < state[0].length &&
            board[newRow][newCol] !== -1 &&
            board[midRow][midCol] !== -1 &&
            state[midRow][midCol] !== 0 &&
            state[newRow][newCol] === 0 &&
            !passMap[newRow][newCol]
        ) {
            const newPosition = { row: newRow, col: newCol };
            moves.push(newPosition);
            FindMoves(board, state, newPosition, passMap, moves, inverted);
        }
    }

    // Check wheter any of the moves doesn't pass no adjacent exit rule
    return moves.filter((move) => NoExitCheck(board, state, position, move, inverted));
};


export { LogicState };
export { FindMoves, PerformMove };