export type LoginUser = {
    username: string;
    password: string;
}

export type User = {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
    Name: string;
    Email: string;
    Url: string;
    Password: string;
    LastIP: string;
    LastUA: string;
    Role: string;
    Notice: boolean;
    Comments: string;
}

export type UserLoginResponseInfo = {
    user: User;
    token: string;
}
export type UserLoginResponse = {
    data: UserLoginResponseInfo;
    message: string;
    code: number;
}