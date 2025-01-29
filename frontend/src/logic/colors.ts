const BoardColors: Record<string, string> = {
    '-1': 'transparent',
    '0': 'gray',
    '1': 'red',
    '2': 'blue',
    '3': 'lime',
    '4': 'cyan',
    '5': 'magenta',
    '6': 'yellow'
}

const PlayerColors: Record<string, string> = {
    '-1': 'red',
    '0': 'transparent',
    '1': 'red',
    '2': 'blue',
    '3': 'lime',
    '4': 'cyan',
    '5': 'magenta',
    '6': 'yellow'
}

const ThreePlayerColors: Record<string, string> = {
    '-1': 'red',
    '0': 'transparent',
    '1': 'red',
    '2': 'lime',
    '3': 'magenta',
}

const getPlayerColor = (playerID: number, playerCount: number) => {
    if (playerCount == 3) {
        return ThreePlayerColors[playerID + 1];
    }
    return PlayerColors[playerID + 1];
}

const getPlayerTurnColor = (playerID: number, playerCount: number) => {
    if (playerCount == 3) {
        return ThreePlayerColors[playerID];
    }
    return PlayerColors[playerID];
}


export { BoardColors, PlayerColors, ThreePlayerColors };
export { getPlayerColor, getPlayerTurnColor };