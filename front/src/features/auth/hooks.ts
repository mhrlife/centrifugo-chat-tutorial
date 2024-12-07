import {useAuthStore} from "./store.ts";
import {useMutation} from "react-query";
import {register} from "./service.ts";
import {RegisterRequest, UserWithToken} from "../../types/serializer.ts";
import {ApiError} from "../../api/post.ts";

export const useRegister = () => {
    const {loggedIn} = useAuthStore()

    return useMutation<UserWithToken, ApiError, RegisterRequest>(register, {
        onSuccess: loggedIn,
    })
}