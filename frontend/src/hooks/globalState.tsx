import React, { useContext } from 'react';
import { TwoTeamState } from '../logic/state';

interface GlobalStateContextType {
    serverAddress: string;
    setServerAddress: (address: string) => void;
    playerName: string;
    setPlayerName: (playerName: string) => void;
    playerID: number;
    setPlayerID: (playerID: number) => void;
    gameID: number;
    setGameID: (gameID: number) => void;
    boardState: number[][];
    setBoardState: (gamestate: ((prevGamStage: number[][]) => number[][]) | number[][]) => void;
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
    const [serverAddress, setServerAddress] = React.useState<string>('http://localhost:8080');
    const [playerName, setPlayerName] = React.useState<string>('player');
    const [playerID, setPlayerID] = React.useState<number>(-1);
    const [gameID, setGameID] = React.useState<number>(-1);
    const [boardState, setBoardState] = React.useState<number[][]>(TwoTeamState);
    const [auditLog, setAuditLog] = React.useState<string[]>(["Chinese-Checkers log"]);
    const [ws, setWS] = React.useState<WebSocket | null>(null);

    const value = {
        serverAddress,
        setServerAddress,
        playerID,
        setPlayerID,
        playerName,
        setPlayerName,
        gameID,
        setGameID,
        boardState,
        setBoardState,
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