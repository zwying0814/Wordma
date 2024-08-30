export type Site = {
    Url: string;
    Name: string;
    Posts: string;
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
}

export type SiteResponse = {
    data: Site[];
    message: string;
    code: number;
}