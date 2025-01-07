import { GameState } from "../hooks/globalStateProvider";
import { PerformMove } from "../logic/state";
import { Point } from "../logic/types";

const handleWebSocketMessages = (ws: WebSocket | null, setAuditLog: React.Dispatch<React.SetStateAction<string[]>>, setGameState: React.Dispatch<React.SetStateAction<GameState>>) => {
    if (!ws) return;

    ws.onopen = () => {
        console.log("WebSocket connection established.");
        sendBoardRequest(ws, setAuditLog);
        sendPawnsRequest(ws, setAuditLog);
        sendStateRequest(ws, setAuditLog);
    };

    ws.onmessage = (event) => {
        const data = JSON.parse(event.data);
        const message = JSON.parse(data.message);
        console.log(message);
        setAuditLog((prevAuditLog) => [...prevAuditLog, `RX ${JSON.stringify(message)}`]);

        if (message.type === "server" && message.board !== undefined) {
            if (message.board == null || message.board.length === 0) {
                console.log("Received empty board state.");
            }
            setGameState((prevGameState: GameState) => ({ ...prevGameState, board: message.board }));
        }

        if (message.type === "server" && message.pawns !== undefined) {
            if (message.pawns == null || message.pawns.length === 0) {
                console.log("Received empty pawns state.");
            }
            setGameState((prevGameState: GameState) => ({ ...prevGameState, state: message.pawns }));
        }

        if (message.type === "server" && message.action === "state") {
            setGameState((prevGameState: GameState) => ({ ...prevGameState, players: message.players, turn: message.turn, current: message.current, color: message.color, ended: message.ended }));
        }

        if (message.type === "server" && message.action === "move") {
            setGameState((prevGameState: GameState) => ({ ...prevGameState, state: PerformMove(prevGameState.state, message.start, message.end) }));
            sendStateRequest(ws, setAuditLog);
        }

        if (message.message == "Skipped Turn" && message.type === undefined) {
            sendStateRequest(ws, setAuditLog);
        }
    }
};

const sendStateRequest = (ws: WebSocket | null, setAuditLog: React.Dispatch<React.SetStateAction<string[]>>) => {
    if (!ws || ws.readyState !== WebSocket.OPEN) {
        console.warn("WebSocket is not ready. Cannot send state request.");
        return;
    }

    const request = JSON.stringify({
        type: "player",
        action: "state",
    });

    setAuditLog((prevAuditLog) => [...prevAuditLog, `TX ${request}`]);
    ws.send(request);
};

const sendBoardRequest = (ws: WebSocket | null, setAuditLog: React.Dispatch<React.SetStateAction<string[]>>) => {
    if (!ws || ws.readyState !== WebSocket.OPEN) {
        console.warn("WebSocket is not ready. Cannot send board request.");
        return;
    }

    const request = JSON.stringify({
        type: "player",
        action: "board",
    });

    setAuditLog((prevAuditLog) => [...prevAuditLog, `TX ${request}`]);
    ws.send(request);
};

const sendPawnsRequest = (ws: WebSocket | null, setAuditLog: React.Dispatch<React.SetStateAction<string[]>>) => {
    if (!ws) return;

    const request = JSON.stringify({
        type: "player",
        action: "pawns",
    });

    setAuditLog((prevAuditLog) => [...prevAuditLog, `TX ${request}`]);
    ws.send(request);
};

const sendNewMove = (ws: WebSocket | null, setAuditLog: React.Dispatch<React.SetStateAction<string[]>>, playerID: number, start: Point, end: Point) => {
    if (!ws) return;

    const request = JSON.stringify({
        type: "player",
        action: "move",
        player_id: playerID,
        start,
        end,
    });

    setAuditLog((prevAuditLog) => [...prevAuditLog, `TX ${request}`]);
    ws.send(request);
};


const handleSkipTurn = (ws: WebSocket | null, setAuditLog: React.Dispatch<React.SetStateAction<string[]>>, playerID: number) => {
    if (!ws) return;

    const request = JSON.stringify({
        type: "player",
        action: "move",
        player_id: playerID,
    });

    setAuditLog((prevAuditLog) => [...prevAuditLog, `TX ${request}`]);
    ws.send(request);
};

export { handleWebSocketMessages, sendStateRequest, sendBoardRequest, sendPawnsRequest, sendNewMove, handleSkipTurn };