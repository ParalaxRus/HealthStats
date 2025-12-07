import { useState } from "react";
import { createUser } from "../api/users";

export default function CreateUser() {
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [response, setResponse] = useState("");

    const handleCreate = async () => {
        try {
            const userId = await createUser({ name, email, password });
            setResponse(JSON.stringify(userId, null, 2));
        } catch (err: any) {
            setResponse("Error: " + err.message);
        }
    };

    return (
        <div>
            <h1>Create User</h1>

            <input placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} /><br /><br />
            <input placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} /><br /><br />
            <input placeholder="Password" type="password" value={password} onChange={(e) => setPassword(e.target.value)} /><br /><br />

            <button onClick={handleCreate}>Create User</button>

            <pre>{response}</pre>
        </div>
    );
}