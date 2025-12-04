import { UserServiceClientImpl, GrpcWebImpl } from "../gen/service";
import { RegisterResponse, RegisterRequest } from "../gen/service";
import * as grpcWeb from "grpc-web";

export interface IUserId {
    id: number;
}

export class UserService {
    private grpcWeb = new GrpcWebImpl("http://localhost:5005", {
        debug: true,
        metadata: { Authorization: "Bearer abc123" },
    });

    private client = new UserServiceClientImpl(this.grpcWeb);

    private mapUserId(resp: RegisterResponse): IUserId {
        return { id: resp.id ?? -1 }
    }

    private mapError(error: grpcWeb.RpcError): Error {
        return new Error(`StorageService error: ${error.message}`);
    }

    async register(): Promise<IUserId> {
        const request = RegisterRequest.create({
            name: "test1",
            email: "test1@gmail.com",
            password: "testpass1"
        });

        return new Promise((resolve, reject) => {
            this.client.Register(request, {}, (err, resp) => {
                if (err) {
                    reject(this.mapError(err));
                    return;
                }

                if (!resp) {
                    reject(new Error("Nil response"));
                    return;
                }

                resolve(this.mapUserId(resp));
            });
        });
    }

    async list(): Promise<void> {
        alert("List users")
    }
}