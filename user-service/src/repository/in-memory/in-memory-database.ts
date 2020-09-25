import { IUser } from "../../models/user";
import { IUserRepository } from "../user.repository";

interface IUserDb {
    [userName: string]: IUser;
}

export class InMemoryDatabase implements IUserRepository {

    private readonly db: IUserDb = {};

    addUser(user: IUser): IUser {
        if (this.db[user.username]) {
            throw new Error('Username taken');
        }

        this.db[user.username] = user;
        return user;
    }

    getUser(username: string): IUser {
        return this.db[username] || undefined;
    }

}