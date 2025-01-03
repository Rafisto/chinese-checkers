import { useState } from "react";
import { APICreateGame } from "../api/lobby";
import StyledButton from "../components/styledButton";

const Create = () => {
    const [numPlayers, setNumPlayers] = useState<number>(2);
    
    const [loading, setLoading] = useState<boolean>(false);
    const [response, setResponse] = useState<string>("");
    const [error, setError] = useState<string>("");
    
    const handleCreateGame = async () => {
        setLoading(true);
        setResponse("");
        setError("");
        try {
            await APICreateGame(numPlayers);
            setResponse("Game Created");
        } catch (error) {
            console.error(error);
            setError("Failed to create game");
        }
        setLoading(false);

        setTimeout(() => {
            setResponse("");
            setError("");
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
            <br />
            <StyledButton text={"Create Game"} handleClick={handleCreateGame} loading={loading} loadingText={"Creating Game..."} className={"wide"}  />
            {response && <span style={{color:"green"}}>{response}</span>}
            {error && <span style={{color:"red"}}>{error}</span>}
        </div>
    )

}

export default Create;