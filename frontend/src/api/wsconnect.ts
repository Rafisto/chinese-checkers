export const createWebSocketConnection = (serverAddress: string,
    gameID: number,
    playerID: number, setAuditLog: (auditLog: ((prevAuditLog: string[]) => string[]) | string[]) => void,
    setWS: (ws: WebSocket | null) => void
) => {
    const ws = new WebSocket(`${serverAddress.replace('http://', 'ws://')}/ws?gameID=${gameID}&playerID=${playerID}`);

    ws.onopen = () => {
        setAuditLog((prevAuditLog: string[]) => [...prevAuditLog, "Connected to websocket"]);
    };

    ws.onclose = () => {
        setAuditLog((prevAuditLog: string[]) => [...prevAuditLog, "Disconnected from websocket"]);
    };

    ws.onmessage = (event) => {
        const data = JSON.parse(event.data);
        setAuditLog((prevAuditLog: string[]) => [...prevAuditLog, `Received message: ${data}`]);
    };

    setWS(ws);

    return ws;
};
