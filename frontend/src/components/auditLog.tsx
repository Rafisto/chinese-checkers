import { useState } from "react";
import { useGlobalState } from "../hooks/globalState";

const AuditLog = () => {
    const { auditLog } = useGlobalState();
    const [showLog, setShowLog] = useState<boolean>(false);

    return (
        <div>
            <button onClick={() => setShowLog(!showLog)} title={showLog ? "Hide Log" : "Show Log"} className={"audit-log-btn"}>
                {showLog ? "x" : "o"}
            </button>
            {showLog && (
                <div className={"audit-log"}>
                    {auditLog && auditLog.slice().reverse().map((entry, index) => (
                        <div 
                            key={index} 
                            className="audit-log-entry" 
                            style={{ color: entry.includes("TX") ? "orange" : entry.includes("RX") ? "lime" : "inherit" }}
                        >
                            {entry}
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
}

export default AuditLog;