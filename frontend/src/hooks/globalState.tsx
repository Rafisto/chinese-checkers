import React, { useContext } from 'react';

interface GlobalStateContextType {
    serverAddress: string;
    setServerAddress: (address: string) => void;
    playerID: string;
    setPlayerID: (playerID: string) => void;
    playerName: string;
    setPlayerName: (playerName: string) => void;
    gameID: string;
    setGameID: (gameID: string) => void;
}

const GlobalStateContext = React.createContext<GlobalStateContextType | undefined>(undefined);

interface GlobalStateProviderProps {
    children: React.ReactNode;
}

export const GlobalStateProvider = ({ children }: GlobalStateProviderProps) => {
    const [serverAddress, setServerAddress] = React.useState('http://localhost:8080/');
    const [playerID, setPlayerID] = React.useState('');
    const [playerName, setPlayerName] = React.useState('');
    const [gameID, setGameID] = React.useState('');

    const value = {
        serverAddress,
        setServerAddress,
        playerID,
        setPlayerID,
        playerName,
        setPlayerName,
        gameID,
        setGameID,
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