import { useState } from "react";
import { APICreateGame } from "../api/lobby";
import { useGlobalState } from "../hooks/useGlobalState";
import StyledButton from "../components/styledButton";
import {
    DEFAULT_NUM_PLAYERS,
    DEFAULT_GAME_VARIANT,
    PLAYER_OPTIONS,
    VARIANT_OPTIONS
} from "../logic/create";


const Create = () => {
    const { serverAddress, auditLog, setAuditLog } = useGlobalState();

    const [numPlayers, setNumPlayers] = useState<number>(DEFAULT_NUM_PLAYERS);
    const [gameVariant, setGameVariant] = useState<string>(DEFAULT_GAME_VARIANT);
    const [loading, setLoading] = useState<boolean>(false);

    const handleCreateGame = async () => {
        setLoading(true);
        try {
            await APICreateGame(serverAddress, numPlayers, gameVariant);
            setAuditLog([...auditLog, "Game Created"]);
        } catch (error) {
            setAuditLog([...auditLog, "Failed to create game"]);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div>
            <h2>Create a Game</h2>

            <label>Number of Players</label>
            <select value={numPlayers} onChange={(e) => setNumPlayers(parseInt(e.currentTarget.value))}>
                {PLAYER_OPTIONS.map((opt) => (
                    <option key={opt.value} value={opt.value}>
                        {opt.label}
                    </option>
                ))}
            </select>

            <label>Game Variant</label>
            <select value={gameVariant} onChange={(e) => setGameVariant(e.currentTarget.value)}>
                {VARIANT_OPTIONS.map((opt) => (
                    <option key={opt.value} value={opt.value}>
                        {opt.label}
                    </option>
                ))}
            </select>

            <br />
            <StyledButton
                text={"Create Game"}
                handleClick={handleCreateGame}
                loading={loading}
                loadingText={"Creating Game..."}
                className={"wide"}
            />
        </div>
    );
};

export default Create;
