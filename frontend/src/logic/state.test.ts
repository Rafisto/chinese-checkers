import { FindMoves, LogicState } from './state';
import { Point } from './types';

describe('FindMoves', () => {
    const board = Array.from({ length: 18 }, () => Array(25).fill(0));
    const passMap = Array.from({ length: 18 }, () => Array(25).fill(false));

    it('should find single moves correctly', () => {
        const state = LogicState.map((row) => row.slice());
        state[9][13] = 1;
        const position: Point = { row: 9, col: 13 };
        const moves = FindMoves(board, state, position, passMap);

        expect(moves).toEqual([
            { row: 9, col: 15 },
            { row: 9, col: 11 },
            { row: 10, col: 14 },
            { row: 10, col: 12 },
            { row: 8, col: 14 },
            { row: 8, col: 12 }
        ]);
    });

    it('should find jump moves correctly', () => {
        const state = LogicState.map((row) => row.slice());
        state[9][13] = 1; // Place a piece at (9, 13)
        state[9][15] = 2; // Place a piece at (9, 15) to jump over
        const position: Point = { row: 9, col: 13 };
        const moves = FindMoves(board, state, position, passMap);

        expect(moves).toEqual([
            { row: 9, col: 11 },
            { row: 10, col: 14 },
            { row: 10, col: 12 },
            { row: 8, col: 14 },
            { row: 8, col: 12 },
            { row: 9, col: 17 }
        ])
    });

    it('should not find moves blocked by other pieces', () => {
        const state = LogicState.map((row) => row.slice());
        state[9][13] = 1; // Place a piece at (8, 12)
        state[9][15] = 2; // Place a piece at (8, 14) to jump over
        state[9][17] = 3; // Place a piece at (8, 16) to block the jump
        const position: Point = { row: 9, col: 13 };
        const moves = FindMoves(board, state, position, passMap);

        expect(moves).toEqual([
            { row: 9, col: 11 },
            { row: 10, col: 14 },
            { row: 10, col: 12 },
            { row: 8, col: 14 },
            { row: 8, col: 12 },
        ])
    });

    it('should not find moves outside the board', () => {
        const state = LogicState.map((row) => row.slice());
        state[0][12] = 1;
        const position: Point = { row: 0, col: 12 };
        const moves = FindMoves(board, state, position, passMap);

        expect(moves).toEqual([
            { col: 14, row: 0 },
            { col: 10, row: 0 },
            { col: 13, row: 1 },
            { col: 11, row: 1 }
        ]);
    });
});