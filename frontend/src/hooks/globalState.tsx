import React, { useContext, useState } from 'react';
import { LogicState } from '../logic/state';
import { LogicBoard } from '../logic/board';

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

interface GlobalStateContextType {
    serverAddress: string;
    setServerAddress: (address: string) => void;
    lobbyState: LobbyState;
    setLobbyState: (state: LobbyState) => void;
    gameState: GameState;
    setGameState: (state: ((prevGameState: GameState) => GameState) | GameState) => void;
    auditLog: string[];
    setAuditLog: (auditLog: ((prevAuditLog: string[]) => string[]) | string[]) => void;
    ws: WebSocket | null;
    setWS: (ws: WebSocket | null) => void;
}

const GlobalStateContext = React.createContext<GlobalStateContextType | undefined>(undefined);

interface GlobalStateProviderProps {
    children: React.ReactNode;
}

export const GlobalStateProvider = ({ children }: GlobalStateProviderProps) => {
    const [serverAddress, setServerAddress] = useState<string>('http://localhost:8080');
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

    const [auditLog, setAuditLog] = useState<string[]>(["Chinese-Checkers log"]);
    const [ws, setWS] = useState<WebSocket | null>(null);

    const value = {
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
    };

    return (
        <GlobalStateContext.Provider value={value}>
            {children}
        </GlobalStateContext.Provider>
    );
};

export const useGlobalState = () => {
    const context = useContext(GlobalStateContext);
    if (!context) {
        throw new Error('useGlobalState must be used within a GlobalStateProvider');
    }
    return context;
};
