interface StyledButtonProps {
    text: string;
    handleClick: () => void;
    loading: boolean;
    loadingText: string;
    className?: string;
}

const StyledButton = ({ text, handleClick, loading, loadingText, className }: StyledButtonProps) => {
    return <div>
        <button
            className={className}
            onClick={handleClick}
            disabled={loading}
            style={{
                color: 'white',
                fontWeight: 'bold',
                padding: '0.5rem 1rem',
                borderRadius: '0.25rem',
                opacity: loading ? 0.5 : 1,
                cursor: loading ? 'not-allowed' : 'pointer',
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center'
            }}
        >
            {loading ? (
                <span style={{ display: 'flex', alignItems: 'center' }}>
                    <svg
                        style={{ animation: 'spin 1s linear infinite', height: '1.25rem', width: '1.25rem', marginRight: '0.75rem', color: 'white' }}
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                    >
                        <circle style={{ opacity: 0.25 }} cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" />
                        <path style={{ opacity: 0.75 }} fill="currentColor" d="M4 12a8 8 0 018-8v8H4z" />
                    </svg>
                    {loadingText}
                </span>
            ) : (
                <span>
                    {text}
                </span>
            )}
        </button>
    </div>
}

export default StyledButton;
