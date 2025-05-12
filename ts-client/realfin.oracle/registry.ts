import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryParamsResponse } from "./types/realfin/oracle/query";
import { MsgUpdateParamsResponse } from "./types/realfin/oracle/tx";
import { GenesisState } from "./types/realfin/oracle/genesis";
import { Params } from "./types/realfin/oracle/params";
import { MsgUpdateParams } from "./types/realfin/oracle/tx";
import { QueryParamsRequest } from "./types/realfin/oracle/query";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/realfin.oracle.QueryParamsResponse", QueryParamsResponse],
    ["/realfin.oracle.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/realfin.oracle.GenesisState", GenesisState],
    ["/realfin.oracle.Params", Params],
    ["/realfin.oracle.MsgUpdateParams", MsgUpdateParams],
    ["/realfin.oracle.QueryParamsRequest", QueryParamsRequest],
    
];

export { msgTypes }