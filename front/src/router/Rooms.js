import { Component } from "react"
import axios from "axios"

class Rooms extends Component {
    constructor (props) {
        super(props)
        this.state = {
            rooms:[]
        }
    }
    componentDidMount() {
        axios.defaults.headers = {
            'Access-Control-Allow-Origin': '*',
        }
        axios.get(`http://localhost:8080/api/rooms?token=${localStorage.getItem("token")}`)
            .then((response) => {
                console.log(response)
            })
            .catch((error, err) => {
                console.log(error, err)
            })
    }
    render() {
        return (
            <div>Rooms</div>
        )
    }
}

export default Rooms;