import { useState } from "react";
import { APICreateGame } from "../api/lobby";
import StyledButton from "../components/styledButton";
import { useGlobalState } from "../hooks/globalState";

const Create = () => {
    const { serverAddress, auditLog, setAuditLog } = useGlobalState();

    const [numPlayers, setNumPlayers] = useState<number>(2);
    const [gameVariant, setGameVariant] = useState<string>("classic");

    const [loading, setLoading] = useState<boolean>(false);

    const handleCreateGame = async () => {
        setLoading(true);
        try {
            await APICreateGame(serverAddress, numPlayers, gameVariant);
            setAuditLog([...auditLog, "Game Created"]);
        } catch (error) {
            setAuditLog([...auditLog, "Failed to create game"]);
        }
        setLoading(false);

        setTimeout(() => {
        }, 1000);
    }


    return (
        <div>
            <h2>Create a Game</h2>
            <label>Number of Players</label>
            <select value={numPlayers} onChange={(e) => setNumPlayers(parseInt(e.currentTarget.value))}>
                <option value="2">2 Players</option>
                <option value="3">3 Players</option>
                <option value="4">4 Players</option>
                <option value="6">6 Players</option>
            </select>
            <select value={gameVariant} onChange={(e) => setGameVariant(e.currentTarget.value)}>
                <option value="classic">Classic</option>
                <option value="chaos">Chaos</option>
            </select>
            <br />
            <StyledButton text={"Create Game"} handleClick={handleCreateGame} loading={loading} loadingText={"Creating Game..."} className={"wide"} />
        </div>
    )

}

export default Create;