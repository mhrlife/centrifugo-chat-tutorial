import {RegisterRequest, UserWithToken} from "../../types/serializer.ts";
import {post} from "../../api/post.ts";

export const register = async (data: RegisterRequest): Promise<UserWithToken> => {
    return await post<RegisterRequest, UserWithToken>("/auth/register", data)
}