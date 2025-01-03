import { useGlobalState } from "../hooks/globalState";

const AuditLog = () => {
    const { auditLog } = useGlobalState();

    return (
        <div className="audit-log">
            {auditLog && auditLog.map((entry, index) => (
                <div key={index}>{entry}</div>
            ))}
        </div>
    );
}

export default AuditLog;