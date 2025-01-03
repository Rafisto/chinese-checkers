import { useState } from "react";
import Connect from "./lobby/connect";
import Create from "./lobby/create"
import Join from "./lobby/join"

const Menu = () => {
    const [connected, setConnected] = useState<boolean>(false);
    const [joined, setJoined] = useState<boolean>(false);

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
        </div>
    )
}

export default Menu;