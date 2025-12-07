import { useState } from "react";
import { findUser } from "../api/users";

export default function FindUser() {
    const [email, setEmail] = useState("");
    const [response, setResponse] = useState("");

    const handleFind = async () => {
        try {
            const user = await findUser(email);
            setResponse(JSON.stringify(user, null, 2));
        } catch (err: any) {
            setResponse("Error: " + err.message);
        }
    };

    return (
        <div>
            <h2>Find User</h2>

            <input
                placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
            /><br /><br />

            <button onClick={handleFind}>Find User</button>

            <pre>{response}</pre>
        </div>
    );
}
