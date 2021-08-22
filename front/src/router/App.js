import { Route, Switch } from "react-router";
import Home from "./Home";
import Room from "./Room";
import Rooms from "./Rooms";

function App() {
    return (
        <main>
            <Switch>
                <Route exact path="/" component={Home}/>
                <Route exact path="/rooms" component={Rooms}/>
                <Route path="/rooms/:roomId" component={Room}/>
            </Switch>
        </main>
    )
}

export default App;