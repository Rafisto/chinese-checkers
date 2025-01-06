import React from "react";
import { GameState, LobbyState } from "./globalStateProvider";

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

export { GlobalStateContext };