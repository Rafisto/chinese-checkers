import { useEffect, useState } from "react";
import { useGlobalState } from "./hooks/useGlobalState";
import { createWebSocketConnection } from "./api/wsconnect";
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

        const ws = createWebSocketConnection(
            serverAddress,
            lobbyState.gameID,
            lobbyState.playerID,
            setAuditLog,
            setWS
        );

        return () => {
            ws.close();
            setWS(null);
        };
    }, [joined, serverAddress, lobbyState, setAuditLog, setWS]);

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