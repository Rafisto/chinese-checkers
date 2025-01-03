import { useEffect, useRef, useState } from "react";
import { useGlobalState } from "./hooks/globalState";
import Connect from "./lobby/connect";
import Create from "./lobby/create"
import Join from "./lobby/join"
import Stats from "./lobby/stats";

const Menu = () => {
    const [connected, setConnected] = useState<boolean>(false);
    const [joined, setJoined] = useState<boolean>(false);
    const { serverAddress, auditLog, setAuditLog, gameID, playerID } = useGlobalState();

    const ws = useRef<WebSocket | null>(null);

    useEffect(() => {
        if (!joined) {
            return;
        }

        ws.current = new WebSocket(`${serverAddress.replace('http://', 'ws://')}/ws?gameID=${gameID}&playerID=${playerID}`);
        ws.current.onopen = () => {
            setAuditLog([...auditLog, "Connected to websocket"]);
        }
        ws.current.onclose = () => {
            setAuditLog([...auditLog, "Disconnected from websocket"]);
        }
        ws.current.onmessage = (event) => {
            const message = JSON.parse(event.data);
            setAuditLog([...auditLog, `Received message: ${message.message}`]);
        }

        const wsCurrent = ws.current;

        return () => {
            wsCurrent.close();
        }
    }, [joined]);


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
                </>
            }
        </div>
    )
}

export default Menu;