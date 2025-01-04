import { useEffect, useState } from "react";
import { useGlobalState } from "./hooks/globalState";
import Connect from "./lobby/connect";
import Create from "./lobby/create"
import Join from "./lobby/join"
import Stats from "./lobby/stats";

const Menu = () => {
    const [connected, setConnected] = useState<boolean>(false);
    const [joined, setJoined] = useState<boolean>(false);
    const { serverAddress, setAuditLog, lobbyState, setWS } = useGlobalState();

    useEffect(() => {
        if (!joined) {
            return;
        }

        let ws = new WebSocket(`${serverAddress.replace('http://', 'ws://')}/ws?gameID=${lobbyState.gameID}&playerID=${lobbyState.playerID}`);
        ws.onopen = () => {
            setAuditLog((prevAuditLog: string[]) => [...prevAuditLog, "Connected to websocket"]);
        }
        ws.onclose = () => {
            setAuditLog((prevAuditLog: string[]) => [...prevAuditLog, "Disconnected from websocket"]);
        }
        ws.onmessage = (event) => {
            const data = JSON.parse(event.data);
            setAuditLog((prevAuditLog: string[]) => [...prevAuditLog, `Received message: ${data}`]);
        }

        setWS(ws);

        const wsCurrent = ws;

        return () => {
            wsCurrent.close();
            setWS(null);
        }

    }, [joined, serverAddress, lobbyState, setAuditLog]);

    return (
        <div className="menu">
            <h1 style={{ textAlign: "center" }}>Chinese Checkers</h1>

            {!connected &&
                <>
                    <hr />
                    <Connect setConnected={setConnected} />
                </>
            }

            {connected && !joined &&
                <>
                    <hr />
                    <Create />
                    <hr />
                    <Join joined={joined} setJoined={setJoined} />
                    <hr />
                </>
            }

            {joined &&
                <>
                    <hr />
                    <Stats />
                </>
            }
        </div>
    )
}

export default Menu;