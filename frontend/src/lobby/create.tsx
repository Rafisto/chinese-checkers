const Create = () => {
    return (
        <div>
            <h2>Create a Game</h2>
            <label>Number of Players</label>
            <select>
                <option value="2">2 Players</option>
                <option value="3">3 Players</option>
                <option value="4">4 Players</option>
                <option value="6">6 Players</option>
            </select>
            <br/>
            <button>Create</button>
        </div>
    )

}

export default Create;