import {create} from "zustand";
import {persist} from "zustand/middleware";
import {UserWithToken} from "../../types/serializer.ts";


export interface AuthState {
    user: UserWithToken | null;

    loggedIn(user: UserWithToken): void;

    loggedOut(): void;

    isLoggedIn(): boolean;
}

export const useAuthStore = create<AuthState>()(persist((setState, getState) => ({
    user: null,
    requestState: 'IDLE',

    loggedIn(user: UserWithToken) {
        setState({user});
    },
    loggedOut() {
        setState({user: null});
    },
    isLoggedIn(): boolean {
        return !!getState().user;
    }
}), {
    name: "auth-storage",
}))