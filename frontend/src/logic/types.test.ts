import { Point } from './types';

describe('Point type', () => {
    it('should create a Point with correct row and col values', () => {
        const point: Point = { row: 1, col: 2 };
        expect(point.row).toBe(1);
        expect(point.col).toBe(2);
    });

    it('should allow updating row and col values', () => {
        const point: Point = { row: 1, col: 2 };
        point.row = 3;
        point.col = 4;
        expect(point.row).toBe(3);
        expect(point.col).toBe(4);
    });

    it('should not allow extra properties', () => {
        const point: Point = { row: 1, col: 2 } as Point & { extra?: string };
        expect(point.row).toBe(1);
        expect(point.col).toBe(2);
        expect((point as unknown as {extra: string}).extra).toBeUndefined();
    });
});