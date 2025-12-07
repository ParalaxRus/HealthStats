export interface CreateUserRequest {
    name: string;
    email: string;
    password: string;
}

export interface UserId {
    id: number;
}

export interface User {
    name: string;
    email: string;
}

const ENDPOINT = "http://localhost:8080";

export async function createUser(request: CreateUserRequest): Promise<UserId> {
    const res = await fetch(`${ENDPOINT}/users`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(request),
    });

    if (!res.ok) {
        throw new Error(`Failed to create user: ${res.statusText}`);
    }

    return await res.json();
}

export async function findUser(email: string): Promise<User> {
    const res = await fetch(`${ENDPOINT}/users?email=${email}`);

    if (!res.ok) {
        throw new Error(`Failed to find user by email: ${res.statusText}`);
    }

    return await res.json();
}
