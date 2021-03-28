import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "faddat.benderchain.benderchain";
export interface Add {
    creator: string;
    id: number;
    post: string;
    title: string;
    body: string;
    ipfs: string;
}
export declare const Add: {
    encode(message: Add, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Add;
    fromJSON(object: any): Add;
    toJSON(message: Add): unknown;
    fromPartial(object: DeepPartial<Add>): Add;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
