import { APIListGames } from "../api/lobby";
import { useGlobalState } from "../hooks/globalState";

type ConnectProps = {
    setConnected: (connected: boolean) => void;
}

const Connect = ({ setConnected }: ConnectProps) => {
    const { serverAddress, setServerAddress, playerName, setPlayerName, auditLog, setAuditLog } = useGlobalState();

    const connect = async () => {
        try {
            await APIListGames(serverAddress);
            setAuditLog([...auditLog, "Connected to server."]);
            setConnected(true);
        }
        catch (error) {
            console.error(error);
            setAuditLog([...auditLog, "Unable to connect to the server."]);
            setConnected(false);
        }
    }

    return (
        <div>
            <p>Server API</p>
            <input type="text" value={serverAddress} onChange={(e) => setServerAddress(e.currentTarget.value)}></input>
            <br />
            <p>Enter your name</p>
            <input type="text" value={playerName} onChange={(e) => setPlayerName(e.currentTarget.value)}></input>
            <button onClick={() => connect()} className={"wide"}>Connect</button>
        </div>
    )
}

export default Connect;