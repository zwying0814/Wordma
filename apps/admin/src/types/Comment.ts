import type {Post} from "./Post";
import type {User} from "./User";

export type CommentDataType = {
    Content: string;
    UA: string;
    IP: string;
    Region: string;
    Type: string;
    Up: number;
    Down: number;
    PostID: number;
    PostTitle: string;
    UserID: number;
    Username: string;
    Parent: number;
    ID: number;
    CreatedAt: string;
}

export type CommentResponseType = {
    Content: string;
    UA: string;
    IP: string;
    Region: string;
    Type: string;
    Up: number;
    Down: number;
    PostID: number;
    Post: Post;
    UserID: number;
    User: User;
    Parent: number;
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
}

export type CommentResponse = {
    data: CommentDataType[];
    message: string;
    code: number;
}