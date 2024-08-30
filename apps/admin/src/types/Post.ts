import {Site} from "./Site.ts";

export type Post = {
    Slug: string;
    Up: number;
    Down: number;
    Read: number;
    SiteID: number;
    Site: Site;
    Comments: string;
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
}

export type PostListData = {
    Slug: string;
    Up: number;
    Down: number;
    Read: number;
}

export type PostListDataResponse = {
    data: PostListData[];
    message: string;
    code: number;
}