import { PLAYER_OPTIONS } from './create';

describe('PLAYER_OPTIONS', () => {
    it('should have correct label for 2 players', () => {
        expect(PLAYER_OPTIONS.find(option => option.value === 2)?.label).toBe("2 Players");
    });

    it('should have correct label for 3 players', () => {
        expect(PLAYER_OPTIONS.find(option => option.value === 3)?.label).toBe("3 Players");
    });

    it('should have correct label for 4 players', () => {
        expect(PLAYER_OPTIONS.find(option => option.value === 4)?.label).toBe("4 Players");
    });

    it('should have correct label for 6 players', () => {
        expect(PLAYER_OPTIONS.find(option => option.value === 6)?.label).toBe("6 Players");
    });

    it('should have correct number of player options', () => {
        expect(PLAYER_OPTIONS.length).toBe(4);
    });
});