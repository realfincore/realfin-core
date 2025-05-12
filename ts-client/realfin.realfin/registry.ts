import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryParamsRequest } from "./types/realfin/realfin/query";
import { QueryParamsResponse } from "./types/realfin/realfin/query";
import { GenesisState } from "./types/realfin/realfin/genesis";
import { MsgUpdateParamsResponse } from "./types/realfin/realfin/tx";
import { Params } from "./types/realfin/realfin/params";
import { MsgUpdateParams } from "./types/realfin/realfin/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/realfin.realfin.QueryParamsRequest", QueryParamsRequest],
    ["/realfin.realfin.QueryParamsResponse", QueryParamsResponse],
    ["/realfin.realfin.GenesisState", GenesisState],
    ["/realfin.realfin.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/realfin.realfin.Params", Params],
    ["/realfin.realfin.MsgUpdateParams", MsgUpdateParams],
    
];

export { msgTypes }