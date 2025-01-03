import { useEffect, useState } from "react";
import { useGlobalState } from "./hooks/globalState";
import Connect from "./lobby/connect";
import Create from "./lobby/create"
import Join from "./lobby/join"
import Stats from "./lobby/stats";

const Menu = () => {
    const [connected, setConnected] = useState<boolean>(false);
    const [joined, setJoined] = useState<boolean>(false);
    const { serverAddress, setAuditLog, gameID, playerID, ws, setWS } = useGlobalState();

    // const ws = useRef<WebSocket | null>(null);

    useEffect(() => {
        if (!joined) {
            return;
        }

        let ws = new WebSocket(`${serverAddress.replace('http://', 'ws://')}/ws?gameID=${gameID}&playerID=${playerID}`);
        ws.onopen = () => {
            setAuditLog((prevAuditLog: string[]) => [...prevAuditLog, "Connected to websocket"]);
        }
        ws.onclose = () => {
            setAuditLog((prevAuditLog: string[]) => [...prevAuditLog, "Disconnected from websocket"]);
        }

        setWS(ws);

        const wsCurrent = ws;

        return () => {
            wsCurrent.close();
            setWS(null);
        }

    }, [joined, serverAddress, gameID, playerID, setAuditLog]);

    const sendJSON = () => {
        if (!ws) {
            return;
        }

        ws.send(JSON.stringify({ message: "Hello, server!" }));
    }

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
                    <Join setJoined={setJoined} />
                    <hr />
                </>
            }

            {joined &&
                <>
                    <hr />
                    <Stats />
                    <button onClick={() => sendJSON()}>Send JSON</button>
                </>
            }
        </div>
    )
}

export default Menu;