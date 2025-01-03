import { useGlobalState } from "../hooks/globalState";

const Connect = () => {
    const { serverAddress, setServerAddress } = useGlobalState();

    return (
        <div>
            <h2>Connect to the Server</h2>
            <input type="text" value={serverAddress} onChange={(e) => setServerAddress(e.currentTarget.value)}></input>
            <br />
            <button>Connect</button>
        </div>
    )
}

export default Connect;