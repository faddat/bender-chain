// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdatePost } from "./types/benderchain/tx";
import { MsgCreatePost } from "./types/benderchain/tx";
import { MsgDeletePost } from "./types/benderchain/tx";
const types = [
    ["/faddat.benderchain.benderchain.MsgUpdatePost", MsgUpdatePost],
    ["/faddat.benderchain.benderchain.MsgCreatePost", MsgCreatePost],
    ["/faddat.benderchain.benderchain.MsgDeletePost", MsgDeletePost],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgUpdatePost: (data) => ({ typeUrl: "/faddat.benderchain.benderchain.MsgUpdatePost", value: data }),
        msgCreatePost: (data) => ({ typeUrl: "/faddat.benderchain.benderchain.MsgCreatePost", value: data }),
        msgDeletePost: (data) => ({ typeUrl: "/faddat.benderchain.benderchain.MsgDeletePost", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
