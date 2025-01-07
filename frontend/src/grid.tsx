interface GridProps {
    children: React.ReactNode | React.ReactNode[];
}

const Grid = ({children}: GridProps) => {
    return (
        <div className="grid">
            {children}
        </div>
    )
}

export default Grid;