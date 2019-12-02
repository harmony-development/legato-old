export interface IPacket {
    type: string;
    data: unknown;
}

export interface IGuildData {
    guildid: string;
    guildname: string;
    picture: string;
}
