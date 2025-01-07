import { BoardColors } from './colors';

describe('BoardColors', () => {
    it('should have correct color for -1', () => {
        expect(BoardColors['-1']).toBe('transparent');
    });

    it('should have correct color for 0', () => {
        expect(BoardColors['0']).toBe('gray');
    });

    it('should have correct color for 1', () => {
        expect(BoardColors['1']).toBe('red');
    });

    it('should have correct color for 2', () => {
        expect(BoardColors['2']).toBe('blue');
    });

    it('should have correct color for 3', () => {
        expect(BoardColors['3']).toBe('lime');
    });

    it('should have correct color for 4', () => {
        expect(BoardColors['4']).toBe('cyan');
    });

    it('should have correct color for 5', () => {
        expect(BoardColors['5']).toBe('magenta');
    });

    it('should have correct color for 6', () => {
        expect(BoardColors['6']).toBe('yellow');
    });
});