import Connect from "./lobby/connect";
import Create from "./lobby/create"
import Join from "./lobby/join"

const Menu = () => {
    return (
        <div className="menu">
            <h1 style={{textAlign:"center"}}>Chinese Checkers</h1>
            <hr/>
            <Connect />
            <hr/>
            <Create />
            <hr/>
            <Join />
            <hr/>
        </div>
    )
}

export default Menu;