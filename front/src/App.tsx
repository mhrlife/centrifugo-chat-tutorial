import './App.css'
import {useRegister} from "./features/auth/hooks.ts";
import {useAuthStore} from "./features/auth/store.ts";
import {useState} from "react";

function App() {
    const {mutate: register, isLoading, isError, error} = useRegister()
    const {isLoggedIn, loggedOut} = useAuthStore()
    const [username, setUsername] = useState("")

    if (isLoading) {
        return <div>Loading...</div>
    }

    if (isError) {
        return <div>Error: {error?.message}</div>
    }

    if (isLoggedIn()) {
        return <div>
            <h3>Logout</h3>
            <button onClick={() => loggedOut()}>Logout</button>
        </div>
    }

    return (
        <div>
            <h1>Register</h1>
            <input placeholder="Username" onChange={e => setUsername(e.target.value)}/>
            <button onClick={() => register({display_name: username})}>Register as {username}</button>
        </div>
    )

}

export default App
