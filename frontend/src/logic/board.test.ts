import { LogicBoard } from './board';

describe('LogicBoard', () => {
    it('should have 17 rows', () => {
        expect(LogicBoard.length).toBe(17);
    });

    it('should have 25 columns in each row', () => {
        LogicBoard.forEach(row => {
            expect(row.length).toBe(25);
        });
    });

    it('should have -1 in all invalid positions', () => {
        LogicBoard.forEach(row => {
            row.forEach(cell => {
                expect([0, 1, 2, 3, 4, 5, 6, -1]).toContain(cell);
            });
        });
    });
});