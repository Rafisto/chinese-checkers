import { useState } from "react";
import { LogicBoard } from "../logic/board";
import { LogicState } from "../logic/state";

import { GlobalStateContext } from "./globalStateContext";

interface LobbyState {
    gameVariant: string;
    playerName: string;
    playerID: number;
    gameID: number;
}

interface GameState {
    color: number;
    players: number[];
    current: number;
    turn: number;
    board: number[][];
    state: number[][];
}

interface GlobalStateProviderProps {
    children: React.ReactNode;
}

export const GlobalStateProvider = ({ children }: GlobalStateProviderProps) => {
    const [serverAddress, setServerAddress] = useState<string>('http://localhost:8080');
    const [auditLog, setAuditLog] = useState<string[]>(["Chinese-Checkers log"]);
    const [ws, setWS] = useState<WebSocket | null>(null);
    const [lobbyState, setLobbyState] = useState<LobbyState>({
        gameVariant: 'classic',
        playerName: 'player',
        playerID: -1,
        gameID: -1,
    })
    const [gameState, setGameState] = useState<GameState>({
        color: -1,
        current: -1,
        players: [],
        turn: -1,
        board: LogicBoard,
        state: LogicState,
    });

    return (
        <GlobalStateContext.Provider value={{
            serverAddress,
            setServerAddress,
            lobbyState,
            setLobbyState,
            gameState,
            setGameState,
            auditLog,
            setAuditLog,
            ws,
            setWS,
        }}>
            {children}
        </GlobalStateContext.Provider>
    );
};

export type { LobbyState, GameState };